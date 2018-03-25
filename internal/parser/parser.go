package parser

import (
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/opsidian/parsley/ast"
	"github.com/opsidian/parsley/ast/builder"
	"github.com/opsidian/parsley/combinator"
	"github.com/opsidian/parsley/parser"
	"github.com/opsidian/parsley/parsley"
	"github.com/opsidian/parsley/reader"
	"github.com/opsidian/parsley/text"
	"github.com/opsidian/parsley/text/terminal"
	"github.com/seriousben/net-script/internal/types"
)

type commandHeader struct {
	Method string
	URL    string
}

var newlineToken = terminal.Regexp("newline", "\n", true, 0, "NEWLINE")
var spaceToken = terminal.Rune(' ', "SPACE")

func commentsParser() parser.Func {
	commentLine := combinator.Seq(
		builder.Select(0),
		terminal.Regexp("comment", "^#(.+)", false, 1, "COMMENT"),
		combinator.Optional(newlineToken),
	)

	return combinator.Many1(
		builder.All(
			"COMMENTS",
			ast.InterpreterFunc(func(ctx interface{}, nodes []ast.Node) (interface{}, reader.Error) {
				comments := []string{}
				for _, node := range nodes {
					val, _ := node.Value(ctx)
					comment, _ := val.(string)
					commentContent := strings.Trim(comment, " #")
					if commentContent != "" {
						comments = append(comments, commentContent)
					}
				}
				return comments, nil
			}),
		),
		commentLine,
	)
}

func commandHeaderParser() parser.Func {
	return combinator.Seq(
		builder.All(
			"COMMAND_HEADER",
			ast.InterpreterFunc(func(ctx interface{}, nodes []ast.Node) (interface{}, reader.Error) {
				methodRaw, _ := nodes[0].Value(ctx)
				urlRaw, _ := nodes[2].Value(ctx)

				return commandHeader{
					Method: methodRaw.(string),
					URL:    urlRaw.(string),
				}, nil
			}),
		),
		// GET http://example.com
		terminal.Regexp("method", "^[A-Z]+", false, 0, "METHOD"),
		spaceToken,
		terminal.Regexp("url", ".*", false, 0, "URL"),
	)
}

func headerParser() parser.Func {
	return combinator.Seq(
		builder.All(
			"HEADER",
			ast.InterpreterFunc(func(ctx interface{}, nodes []ast.Node) (interface{}, reader.Error) {
				value0, _ := nodes[0].Value(ctx)
				value3, _ := nodes[3].Value(ctx)

				return []string{value0.(string), value3.(string)}, nil
			}),
		),
		terminal.Regexp("header name", "[^:]*", false, 0, "HEADER_NAME"),
		terminal.Rune(':', "COLON"),
		spaceToken,
		terminal.Regexp("any char", ".*", false, 0, "HEADER_VALUE"),
	)
}

func headersParser() parser.Func {
	headerLine := combinator.Seq(
		builder.Select(0),
		headerParser(),
		combinator.Optional(newlineToken),
	)

	return combinator.Many(
		builder.All(
			"HEADERS",
			ast.InterpreterFunc(func(ctx interface{}, nodes []ast.Node) (interface{}, reader.Error) {
				headers := make(http.Header)
				for _, node := range nodes {
					val, _ := node.Value(ctx)
					headerVals, _ := val.([]string)
					headers.Add(headerVals[0], headerVals[1])
				}
				return headers, nil
			}),
		),
		headerLine,
	)
}

func bodyParser() parser.Func {
	bodyLine := combinator.Seq(
		builder.Select(0),
		terminal.Regexp("body line", ".+", false, 0, "BODY_LINE"),
		combinator.Optional(newlineToken),
	)

	return combinator.Many(
		builder.All(
			"BODY",
			ast.InterpreterFunc(func(ctx interface{}, nodes []ast.Node) (interface{}, reader.Error) {
				body := []string{}
				for _, node := range nodes {
					val, _ := node.Value(ctx)
					body = append(body, val.(string))
				}
				return []byte(strings.Join(body, "\n")), nil
			}),
		),
		bodyLine,
	)
}

func commandParser() parser.Func {
	return combinator.Seq(
		builder.All(
			"LINES",
			ast.InterpreterFunc(func(ctx interface{}, nodes []ast.Node) (interface{}, reader.Error) {
				commentsRaw, _ := nodes[0].Value(ctx)
				comments := commentsRaw.([]string)

				cmdHeaderRaw, _ := nodes[1].Value(ctx)
				cmdHeader := cmdHeaderRaw.(commandHeader)

				headers := http.Header{}
				if nodes[2] != nil {
					headersRaw, _ := nodes[2].Value(ctx)
					headers = headersRaw.(http.Header)
				}

				body := []byte{}
				if nodes[3] != nil {
					bodyRaw, _ := nodes[3].Value(ctx)
					body = bodyRaw.([]byte)
				}

				return types.Command{
					Comment: strings.Join(comments, "\n"),
					Method:  cmdHeader.Method,
					URL:     cmdHeader.URL,
					Headers: headers,
					Body:    body,
				}, nil
			}),
		),
		commentsParser(),
		commandHeaderParser(),
		combinator.Optional(combinator.Seq(
			builder.Select(1),
			newlineToken,
			headersParser(),
		)),
		combinator.Optional(combinator.Seq(
			builder.Select(1),
			newlineToken,
			bodyParser(),
		)),
	)
}

func scriptParser() parser.Func {
	command := combinator.Seq(
		builder.Select(1),
		combinator.Many(builder.Nil(), newlineToken),
		commandParser(),
		combinator.Many(builder.Nil(), newlineToken),
	)

	return combinator.Many(
		builder.All(
			"SCRIPT",
			ast.InterpreterFunc(func(ctx interface{}, nodes []ast.Node) (interface{}, reader.Error) {
				cmds := []types.Command{}
				for _, node := range nodes {
					val, _ := node.Value(ctx)
					cmds = append(cmds, val.(types.Command))
				}
				return cmds, nil
			}),
		),
		command,
	)
}

func parse(b []byte, filename string) ([]types.Command, error) {
	s := parsley.NewSentence(scriptParser())
	cmdsRaw, _, err := s.Evaluate(text.NewReader(b, filename, false), nil)
	if err != nil {
		return nil, err
	}
	cmds := cmdsRaw.([]types.Command)

	return cmds, nil
}

// Parse bytes into commands
func Parse(b []byte) ([]types.Command, error) {
	return parse(b, "")
}

// ParseFile compiles a file into commands
func ParseFile(filepath string) ([]types.Command, error) {
	b, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	return parse(b, filepath)
}

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
		terminal.Regexp("header name", "[A-Za-z-_]*", false, 0, "HEADER_NAME"),
		terminal.Rune(':', "COLON"),
		spaceToken,
		terminal.Regexp("any char", ".*", false, 0, "HEADER_VALUE"),
	)
}

func headersParser() parser.Func {
	headersInterp := ast.InterpreterFunc(func(ctx interface{}, nodes []ast.Node) (interface{}, reader.Error) {
		headers := make(http.Header)
		for i := 0; i < len(nodes); i += 2 {
			node := nodes[i]
			val, _ := node.Value(ctx)
			headerVals, _ := val.([]string)
			headers.Add(headerVals[0], headerVals[1])
		}
		return headers, nil
	})

	return combinator.SepBy1(
		"HEADER_LINE",
		headerParser(),
		newlineToken,
		headersInterp,
	).Parse
}

func bodyParser() parser.Func {
	bodyLineInterp := ast.InterpreterFunc(func(ctx interface{}, nodes []ast.Node) (interface{}, reader.Error) {
		body := []string{}
		for _, node := range nodes {
			val, _ := node.Value(ctx)
			body = append(body, val.(string))
		}
		return []byte(strings.Join(body, "")), nil
	})

	return combinator.SepBy1(
		"BODY",
		terminal.Regexp("body line", ".+", false, 0, "BODY_LINE"),
		newlineToken,
		bodyLineInterp,
	).Parse
}

func commandParser() parser.Func {
	return combinator.Seq(
		builder.All(
			"LINES",
			ast.InterpreterFunc(func(ctx interface{}, nodes []ast.Node) (interface{}, reader.Error) {
				commentsRaw, _ := nodes[0].Value(ctx)
				comment, _ := commentsRaw.(string)
				comment = strings.Trim(comment, " #")

				cmdHeaderRaw, _ := nodes[2].Value(ctx)
				cmdHeader := cmdHeaderRaw.(commandHeader)

				headers := http.Header{}
				body := []byte{}

				if nodes[3] != nil {
					var hdrBodyComb []interface{}
					hrdBodyRaw, _ := nodes[3].Value(ctx)
					hdrBodyComb = hrdBodyRaw.([]interface{})

					if hdrBodyComb[0] != nil {
						headers = hdrBodyComb[0].(http.Header)
					}

					if hdrBodyComb[1] != nil {
						body = hdrBodyComb[1].([]byte)
					}
				}

				/*
					headers := http.Header{}
					if nodes[3] != nil {
						headersRaw, _ := nodes[3].Value(ctx)
						headers = headersRaw.(http.Header)
					}

					body := []byte{}
					if nodes[4] != nil {
						bodyRaw, _ := nodes[4].Value(ctx)
						body = bodyRaw.([]byte)
					}
				*/

				return types.Command{
					Comment: comment,
					Method:  cmdHeader.Method,
					URL:     cmdHeader.URL,
					Headers: headers,
					Body:    body,
				}, nil
			}),
		),
		terminal.Regexp("comment", "^#(.+)", false, 1, "COMMENT"),
		newlineToken,
		commandHeaderParser(),
		combinator.Optional(combinator.Choice(
			"BODY-HEADER-COMBINATIONS",
			combinator.Seq(
				builder.All(
					"COMB-BOTH",
					ast.InterpreterFunc(func(ctx interface{}, nodes []ast.Node) (interface{}, reader.Error) {
						hdr, _ := nodes[1].Value(ctx)
						body, _ := nodes[4].Value(ctx)
						data := make([]interface{}, 2)
						data[0] = hdr
						data[1] = body
						return data, nil
					}),
				),
				newlineToken,
				headersParser(),
				newlineToken,
				newlineToken,
				bodyParser(),
			),
			combinator.Seq(
				builder.All(
					"COMB-HDR",
					ast.InterpreterFunc(func(ctx interface{}, nodes []ast.Node) (interface{}, reader.Error) {
						hdr, _ := nodes[1].Value(ctx)
						data := make([]interface{}, 2)
						data[0] = hdr
						data[1] = nil
						return data, nil
					}),
				),
				newlineToken,
				headersParser(),
			),
			combinator.Seq(
				builder.All(
					"COMB-BODY",
					ast.InterpreterFunc(func(ctx interface{}, nodes []ast.Node) (interface{}, reader.Error) {
						body, _ := nodes[2].Value(ctx)
						data := make([]interface{}, 2)
						data[0] = nil
						data[1] = body
						return data, nil
					}),
				),
				newlineToken,
				newlineToken,
				bodyParser(),
			),
		)),
	)
}

func scriptParser() parser.Func {
	commandInterp := ast.InterpreterFunc(func(ctx interface{}, nodes []ast.Node) (interface{}, reader.Error) {
		cmds := []types.Command{}
		for i := 0; i < len(nodes); i += 2 {
			node := nodes[i]
			val, err := node.Value(ctx)
			if err != nil {
				return nil, err
			}
			cmds = append(cmds, val.(types.Command))
		}
		return cmds, nil
	})

	return combinator.SepBy(
		"SCRIPT",
		commandParser(),
		combinator.Many1(
			builder.Nil(),
			terminal.Regexp("newline", "\n", true, 0, "NEWLINE"),
		),
		commandInterp,
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

package builder

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr"

	"github.com/seriousben/net-script/internal/parser"
	"github.com/seriousben/net-script/internal/types"
)

type errorListener struct {
	*antlr.DefaultErrorListener
	errors []string
}

func newErrorListener() *errorListener {
	return new(errorListener)
}

func (c *errorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	c.errors = append(c.errors, fmt.Sprintf("line %d:%d %s", line, column, msg))
}

type netscriptListener struct {
	*parser.BaseNetscriptListener
	errors         []error
	currentComment string
	commands       []types.Command
}

func (s *netscriptListener) EnterComment(ctx *parser.CommentContext) {
	comment := ctx.GetText()
	comment = strings.Trim(comment, " #")
	s.currentComment = comment
}

func (s *netscriptListener) EnterCommand(ctx *parser.CommandContext) {
	method := ctx.Method().GetText()
	_, err := url.Parse(ctx.Url().GetText())
	if err != nil {
		s.errors = append(s.errors, fmt.Errorf("bad URL \"%s\": %s", ctx.Url().GetText(), err))
		return
	}

	s.commands = append(s.commands, types.Command{
		Comment: s.currentComment,
		Method:  method,
		URL:     ctx.Url().GetText(),
	})

	s.currentComment = ""
}

// Build builds the script
func Build(filepath string) ([]types.Command, error) {
	// Setup the input
	is, err := antlr.NewFileStream(filepath)
	if err != nil {
		return nil, err
	}

	el := newErrorListener()

	// Create the Lexer
	lexer := parser.NewNetscriptLexer(is)
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(el)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := parser.NewNetscriptParser(stream)
	p.RemoveErrorListeners()
	p.AddErrorListener(el)

	// Create the tree
	script := p.Script()

	if len(el.errors) != 0 {
		return nil, fmt.Errorf("error parsing\n\t", strings.Join(el.errors, "\n\t"))
	}

	nsListener := new(netscriptListener)

	// Finally parse the expression
	antlr.ParseTreeWalkerDefault.Walk(nsListener, script)

	if len(el.errors) != 0 {
		return nil, fmt.Errorf("error parsing after walking\n\t", strings.Join(el.errors, "\n\t"))
	}
	return nsListener.commands, nil
}

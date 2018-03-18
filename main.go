package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr"

	"github.com/seriousben/net-script/parser"
)

type ErrorListener struct {
	*antlr.DefaultErrorListener
	errors []string
}

func NewErrorListener() *ErrorListener {
	return new(ErrorListener)
}

func (c *ErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	c.errors = append(c.errors, fmt.Sprintf("line %d:%d %s", line, column, msg))
}

type netscriptListener struct {
	*parser.BaseNetscriptListener
}

func (s *netscriptListener) EnterComment(ctx *parser.CommentContext) {
	comment := ctx.GetText()
	comment = strings.Trim(comment, " #")
	log.Printf("comment \"%s\" full=%s", comment, ctx.GetText())
}

func (s *netscriptListener) EnterCommand(ctx *parser.CommandContext) {
	method := ctx.Method().GetText()
	command := ctx.Arguments().GetText()
	log.Printf("executing %s %s", method, command)
}

func main() {
	// Setup the input
	is, err := antlr.NewFileStream("./examples/min-script.ns")
	if err != nil {
		log.Fatalf("error %s", err)
	}

	ell := NewErrorListener()
	// Create the Lexer
	lexer := parser.NewNetscriptLexer(is)
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(ell)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := parser.NewNetscriptParser(stream)
	p.RemoveErrorListeners()
	elp := NewErrorListener()
	p.AddErrorListener(elp)

	// Finally parse the expression
	antlr.ParseTreeWalkerDefault.Walk(&netscriptListener{}, p.Script())

	if len(ell.errors) != 0 {
		log.Fatal("error lexing\n\t", strings.Join(ell.errors, "\n\t"))
	}
	if len(elp.errors) != 0 {
		log.Fatal("error parsing\n\t", strings.Join(elp.errors, "\n\t"))
	}
}

// Code generated from Netscript.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 7, 39, 8,
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 3, 2, 3,
	2, 3, 3, 3, 3, 7, 3, 18, 10, 3, 12, 3, 14, 3, 21, 11, 3, 3, 4, 6, 4, 24,
	10, 4, 13, 4, 14, 4, 25, 3, 5, 6, 5, 29, 10, 5, 13, 5, 14, 5, 30, 3, 6,
	6, 6, 34, 10, 6, 13, 6, 14, 6, 35, 3, 6, 3, 6, 2, 2, 7, 3, 3, 5, 4, 7,
	5, 9, 6, 11, 7, 3, 2, 5, 4, 2, 12, 12, 15, 15, 3, 2, 67, 92, 5, 2, 12,
	12, 15, 15, 34, 34, 2, 42, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3,
	2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 3, 13, 3, 2, 2, 2, 5, 15,
	3, 2, 2, 2, 7, 23, 3, 2, 2, 2, 9, 28, 3, 2, 2, 2, 11, 33, 3, 2, 2, 2, 13,
	14, 7, 34, 2, 2, 14, 4, 3, 2, 2, 2, 15, 19, 7, 37, 2, 2, 16, 18, 10, 2,
	2, 2, 17, 16, 3, 2, 2, 2, 18, 21, 3, 2, 2, 2, 19, 17, 3, 2, 2, 2, 19, 20,
	3, 2, 2, 2, 20, 6, 3, 2, 2, 2, 21, 19, 3, 2, 2, 2, 22, 24, 9, 3, 2, 2,
	23, 22, 3, 2, 2, 2, 24, 25, 3, 2, 2, 2, 25, 23, 3, 2, 2, 2, 25, 26, 3,
	2, 2, 2, 26, 8, 3, 2, 2, 2, 27, 29, 10, 4, 2, 2, 28, 27, 3, 2, 2, 2, 29,
	30, 3, 2, 2, 2, 30, 28, 3, 2, 2, 2, 30, 31, 3, 2, 2, 2, 31, 10, 3, 2, 2,
	2, 32, 34, 9, 2, 2, 2, 33, 32, 3, 2, 2, 2, 34, 35, 3, 2, 2, 2, 35, 33,
	3, 2, 2, 2, 35, 36, 3, 2, 2, 2, 36, 37, 3, 2, 2, 2, 37, 38, 8, 6, 2, 2,
	38, 12, 3, 2, 2, 2, 7, 2, 19, 25, 30, 35, 3, 2, 3, 2,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "' '",
}

var lexerSymbolicNames = []string{
	"", "", "COMMENT", "METHOD", "TEXT", "TERMINATOR",
}

var lexerRuleNames = []string{
	"T__0", "COMMENT", "METHOD", "TEXT", "TERMINATOR",
}

type NetscriptLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var lexerDecisionToDFA = make([]*antlr.DFA, len(lexerAtn.DecisionToState))

func init() {
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

func NewNetscriptLexer(input antlr.CharStream) *NetscriptLexer {

	l := new(NetscriptLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "Netscript.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// NetscriptLexer tokens.
const (
	NetscriptLexerT__0       = 1
	NetscriptLexerCOMMENT    = 2
	NetscriptLexerMETHOD     = 3
	NetscriptLexerTEXT       = 4
	NetscriptLexerTERMINATOR = 5
)

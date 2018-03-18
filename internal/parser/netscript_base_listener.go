// Code generated from Netscript.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // Netscript
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseNetscriptListener is a complete listener for a parse tree produced by NetscriptParser.
type BaseNetscriptListener struct{}

var _ NetscriptListener = &BaseNetscriptListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseNetscriptListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseNetscriptListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseNetscriptListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseNetscriptListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterScript is called when production script is entered.
func (s *BaseNetscriptListener) EnterScript(ctx *ScriptContext) {}

// ExitScript is called when production script is exited.
func (s *BaseNetscriptListener) ExitScript(ctx *ScriptContext) {}

// EnterLine is called when production line is entered.
func (s *BaseNetscriptListener) EnterLine(ctx *LineContext) {}

// ExitLine is called when production line is exited.
func (s *BaseNetscriptListener) ExitLine(ctx *LineContext) {}

// EnterComment is called when production comment is entered.
func (s *BaseNetscriptListener) EnterComment(ctx *CommentContext) {}

// ExitComment is called when production comment is exited.
func (s *BaseNetscriptListener) ExitComment(ctx *CommentContext) {}

// EnterCommand is called when production command is entered.
func (s *BaseNetscriptListener) EnterCommand(ctx *CommandContext) {}

// ExitCommand is called when production command is exited.
func (s *BaseNetscriptListener) ExitCommand(ctx *CommandContext) {}

// EnterMethod is called when production method is entered.
func (s *BaseNetscriptListener) EnterMethod(ctx *MethodContext) {}

// ExitMethod is called when production method is exited.
func (s *BaseNetscriptListener) ExitMethod(ctx *MethodContext) {}

// EnterUrl is called when production url is entered.
func (s *BaseNetscriptListener) EnterUrl(ctx *UrlContext) {}

// ExitUrl is called when production url is exited.
func (s *BaseNetscriptListener) ExitUrl(ctx *UrlContext) {}

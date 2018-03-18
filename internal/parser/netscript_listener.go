// Code generated from Netscript.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // Netscript
import "github.com/antlr/antlr4/runtime/Go/antlr"

// NetscriptListener is a complete listener for a parse tree produced by NetscriptParser.
type NetscriptListener interface {
	antlr.ParseTreeListener

	// EnterScript is called when entering the script production.
	EnterScript(c *ScriptContext)

	// EnterLine is called when entering the line production.
	EnterLine(c *LineContext)

	// EnterComment is called when entering the comment production.
	EnterComment(c *CommentContext)

	// EnterCommand is called when entering the command production.
	EnterCommand(c *CommandContext)

	// EnterMethod is called when entering the method production.
	EnterMethod(c *MethodContext)

	// EnterUrl is called when entering the url production.
	EnterUrl(c *UrlContext)

	// ExitScript is called when exiting the script production.
	ExitScript(c *ScriptContext)

	// ExitLine is called when exiting the line production.
	ExitLine(c *LineContext)

	// ExitComment is called when exiting the comment production.
	ExitComment(c *CommentContext)

	// ExitCommand is called when exiting the command production.
	ExitCommand(c *CommandContext)

	// ExitMethod is called when exiting the method production.
	ExitMethod(c *MethodContext)

	// ExitUrl is called when exiting the url production.
	ExitUrl(c *UrlContext)
}

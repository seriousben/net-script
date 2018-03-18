#!/bin/bash
antlr4 -Dlanguage=Go -o parser -package parser Netscript.g4 && go install . && go run main.go

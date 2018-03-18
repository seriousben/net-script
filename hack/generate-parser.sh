#!/bin/bash

set -eu -o pipefail

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

$DIR/antlr4/antlr4.sh -Dlanguage=Go -o "internal/parser" -package "parser" Netscript.g4

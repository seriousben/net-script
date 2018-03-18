#!/bin/bash

set -eu -o pipefail

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
export CLASSPATH=".:$DIR/antlr-4.7.1-complete.jar:$CLASSPATH"

javac $@

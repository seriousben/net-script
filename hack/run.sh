#!/bin/bash

set -eu -o pipefail

go run main.go run "$1"

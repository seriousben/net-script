#!/bin/bash

set -eu -o pipefail

go run main.go lint "$1"

# is-it-binary

Using git way to check if a file is binary.

[![GoDoc](https://godoc.org/github.com/seriousben/is-it-binary?status.svg)](https://godoc.org/github.com/seriousben/is-it-binary)
[![CircleCI](https://circleci.com/gh/seriousben/is-it-binary/tree/master.svg?style=shield)](https://circleci.com/gh/seriousben/is-it-binary/tree/master)
[![Go Reportcard](https://goreportcard.com/badge/github.com/seriousben/is-it-binary)](https://goreportcard.com/report/github.com/seriousben/is-it-binary)
[![codecov](https://codecov.io/gh/seriousben/is-it-binary/branch/master/graph/badge.svg)](https://codecov.io/gh/seriousben/is-it-binary)

## Example

```go
import (
        "fmt"
        "io/ioutil"

        "github.com/seriousben/is-it-binary/binary"
)

func main() {
        isBinary, err := binary.IsBinaryFile("testdata/file") {
        if err != nil {
                fmt.Println("error", err)
                return
        }
        if isBinary {
                fmt.Println("file is binary")
        } else {
                fmt.Println("file is not binary")
        }
}
```

## Implementation

Using git's heuristics to detect a binary file.
See https://git.kernel.org/pub/scm/git/git.git/tree/xdiff-interface.c?id=7668cbc60578f99a4c048f8f8f38787930b8147b#n201 for details.

## Install

`go get -u github.com/seriousben/is-it-binary`

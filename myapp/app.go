package main

import (
	"github.com/bm4cs/gotime/cake"
)

// nested package
// package alias

func main() {
	//templates.PipelineDemo1()
	cake.StockPrice()
}

/*
To build and install:

    go install github.com/bm4cs/example/myapp

Or with vim-go:

    :GoInstall


To run:
    $GOPATH/bin/myapp

TODO: Investigate makefile.
*/

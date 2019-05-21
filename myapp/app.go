package main

import (
	"fmt"

	"github.com/bm4cs/gotime/cake"
)

// nested package
// package alias

func main() {
	//templates.PipelineDemo1()
	// stocks := []int{10, 7, 5, 8, 11, 9}
	stocks := []int{10, 7, 5, 3, 1}
	maxprofit := cake.StockPrice(stocks)
	fmt.Println(maxprofit)
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

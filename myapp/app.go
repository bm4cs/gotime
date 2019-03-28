package main

import (
	"fmt"
	"github.com/bm4cs/example/numbers"
	"github.com/bm4cs/example/strings"
	"github.com/bm4cs/example/strings/greeting" // nested package
	str "strings"                               // package alias
)

func main() {

	fmt.Println(numbers.IsPrime(19))
	fmt.Println(greeting.WelcomeText)
	fmt.Println(strings.Reverse("benjamin"))
	fmt.Println(str.Count("go is awesome. go is great", "go"))

	bigstr := str.ToUpper("testing 123")
	fmt.Println(bigstr)
}

/*
To build and install:

    go install github.com/bm4cs/example/myapp

To run:

    $GOPATH/bin/myapp

TODO: Investigate makefile.
*/

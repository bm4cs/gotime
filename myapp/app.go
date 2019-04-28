package main

import (
	"fmt"
	"github.com/bm4cs/example/numbers"
	"github.com/bm4cs/example/strings"
	"github.com/bm4cs/example/strings/greeting" // nested package
	"log"
	"net/http"
	"os"
	str "strings" // package alias
)

import "html/template"

func main() {

	fmt.Println("hello world v0.1")

	fmt.Println(numbers.IsPrime(19))
	fmt.Println(greeting.WelcomeText)
	fmt.Println(strings.Reverse("benjamin"))
	fmt.Println(str.Count("go is awesome. go is great", "go"))

	// bindTemplate()

	templates := populateTemplates()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		requestedFile := r.URL.Path[1:]
		t := templates.Lookup(requestedFile + ".htm")
		if t != nil {
			err := t.Execute(w, nil)
			if err != nil {
				log.Println(err)
			}
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	})

	http.Handle("/img/", http.FileServer(http.Dir("web/public")))
	http.Handle("/css/", http.FileServer(http.Dir("web/public")))
	http.ListenAndServe(":1337", nil)
}

type fooHandler struct {
	greeting string
}

func (h *fooHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprintf("%v world", h.greeting)))
}

func populateTemplates() *template.Template {
	result := template.New("templates")
	const basePath = "web/templates"
	template.Must(result.ParseGlob(basePath + "/*.htm"))
	return result
}

// in-memory templates
func bindTemplate() {
	templateString := `hello, rax refers to a register. This is a temporary storage location, which values can be written to, read from, or operated on.`

	t, err := template.New("title").Parse(templateString)
	if err != nil {
		fmt.Println(err)
	}

	err = t.Execute(os.Stdout, nil)
	if err != nil {
		fmt.Println(err)
	}
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

package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	str "strings"

	"github.com/bm4cs/gotime/numbers"
	"github.com/bm4cs/gotime/strings"
	"github.com/bm4cs/gotime/strings/greeting"
	"github.com/bm4cs/gotime/viewmodel"
	// nested package
	// package alias
)

func main() {

	fmt.Println("hello world v0.1")

	fmt.Println(numbers.IsPrime(19))
	fmt.Println(greeting.WelcomeText)
	fmt.Println(strings.Reverse("benjamin"))
	fmt.Println(str.Count("go is awesome. go is great", "go"))

	// bindTemplate()

	templates := populateTemplatesWithLayout()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		requestedFile := r.URL.Path[1:]
		//t := templates.Lookup(requestedFile + ".htm")
		t := templates[requestedFile+".htm"]
		context := viewmodel.NewBase()

		if t != nil {
			err := t.Execute(w, context)
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

func populateTemplatesWithLayout() map[string]*template.Template {
	result := make(map[string]*template.Template)
	const basePath = "web/templates"
	layout := template.Must(template.ParseFiles(basePath + "/_layout.htm"))
	template.Must(layout.ParseFiles(basePath + "/_header.htm"))

	dir, err := os.Open(basePath + "/content")
	if err != nil {
		panic("Can't open content templates dir" + err.Error())
	}

	fis, err := dir.Readdir(-1)
	if err != nil {
		panic("Can't read files within content templates dir" + err.Error())
	}

	for _, fi := range fis {
		f, err := os.Open(basePath + "/content/" + fi.Name())
		if err != nil {
			panic("Can't open template file '" + fi.Name() + "'")
		}

		content, err := ioutil.ReadAll(f)
		if err != nil {
			panic("Couldn't read bytes out of file '" + fi.Name() + "'")
		}

		f.Close()

		// create the template, and store in map
		tmpl := template.Must(layout.Clone())
		_, err = tmpl.Parse(string(content))
		if err != nil {
			panic("Couldn't parse content template '" + fi.Name() + "'")
		}

		result[fi.Name()] = tmpl
	}

	return result
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

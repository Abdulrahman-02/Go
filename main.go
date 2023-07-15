package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type RSVP struct {
	name, email, phone string
	willAttend         bool
}

var responses = make([]*RSVP, 0, 10)
var templates = make(map[string]*template.Template, 3)

func loadTemplates() {
	// TODO: load templates here
	templateNames := [5]string{"welcome", "form", "thanks", "sorry", "list"}
	for index, name := range templateNames {
		t, err := template.ParseFiles("layout.htm", name+".htm")
		if err == nil {
			templates[name] = t
			fmt.Println("Loaded template", index, name)
		} else {
			panic(err)
		}
	}
}

func welcomeHandler(writer http.ResponseWriter, request *http.Request) {
	templates["welcome"].Execute(writer, nil)
}

func listHandler(writer http.ResponseWriter, request *http.Request) {
	templates["list"].Execute(writer, responses)
}

type formData struct {
	*RSVP
	Errors []string
}

func formHandler(writer http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet {
		templates["form"].Execute(writer, formData{
			RSVP: &RSVP{}, Errors: []string{},
		})
	}
}

func main() {
	// fmt.Println("TODO: ADD SOME FEATURES")
	loadTemplates()
	http.HandleFunc("/", welcomeHandler)
	http.HandleFunc("/list", listHandler)
	http.HandleFunc("/form", formHandler)

	err := http.ListenAndServe(":5000", nil)
	if err != nil {
		fmt.Println(err)
	}

	// Stopped at page 21
}

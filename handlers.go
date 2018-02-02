package main

import (
	"net/http"
	"fmt"
	"html/template"
	"log"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	// parse form data
	r.ParseForm()
	firstnames := r.Form["people[][firstname]"]
	surnames := r.Form["people[][surname]"]

	// collect form data for people
	var newPeople []Person
	for i,v := range firstnames {
		newPeople = append(newPeople, Person{v, surnames[i]})
	}
	// overwrite stored people
	people = newPeople

	t, err := template.ParseFiles("tmpl/markup.html")
	if err != nil {
		log.Println("Error parsing markup template:",err)
		fmt.Fprintln(w, "Error getting the requested page")
	}

	t.Execute(w, people)
}

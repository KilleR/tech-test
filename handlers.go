package main

import (
	"net/http"
	"fmt"
	"html/template"
	"log"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	firstnames := r.Form["people[][firstname]"]
	surnames := r.Form["people[][surname]"]

	for i,v := range firstnames {
		fmt.Printf("Person %d: %s %s\r\n", i, v, surnames[i])
	}

	t, err := template.ParseFiles("tmpl/markup.html")
	if err != nil {
		log.Println("Error parsing markup template:",err)
		fmt.Fprintln(w, "Error getting the requested page")
	}

	people := []Person{
		Person{"Jeff", "Stelling"},
	}

	t.Execute(w, people)
}

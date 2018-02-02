package main

import (
	"net/http"
	"fmt"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	fmt.Println("Form:",r.Form)
	fmt.Println("People:", r.Form["people[][firstname]"])
	fmt.Println("PostForm:", r.PostForm)

	firstnames := r.Form["people[][firstname]"]
	surnames := r.Form["people[][surname]"]

	for i,v := range firstnames {
		fmt.Printf("Person %d: %s %s\r\n", i, v, surnames[i])
	}


	http.ServeFile(w, r, "markup.html")
}

package main

import (
	"net/http"
	"fmt"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	fmt.Println("Form:",r.Form)
	fmt.Println("People:", r.Form["people"])
	fmt.Println("PostForm:", r.PostForm)


	http.ServeFile(w, r, "markup.html")
}

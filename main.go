package main

import "net/http"

func main() {


	http.HandleFunc("/", formHandler)
	http.ListenAndServe(":8080", nil)
}

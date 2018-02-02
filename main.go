package main

import "net/http"

var DB *PeopleDb

func init() {
	DB = new(PeopleDb)
}

func main() {
	http.HandleFunc("/", formHandler)
	http.ListenAndServe(":8080", nil)
}

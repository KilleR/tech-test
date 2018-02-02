package main

import "net/http"

var people []Person

func init() {
	people = []Person{
		Person{"Jeff", "Stelling"},
		Person{"Chris", "Kamara"},
		Person{"Alex", "Hammond"},
		Person{"Jim", "White"},
		Person{"Natalie", "Sawyer"},
	}
}

func main() {
	http.HandleFunc("/", formHandler)
	http.ListenAndServe(":8080", nil)
}

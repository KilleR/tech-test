package main

import "net/http"

var DB *PeopleDb

func init() {
	DB = new(PeopleDb)
	DB.people = []Person{
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

package main

type PeopleDb struct {
	people []Person
}

func (db *PeopleDb) GetAll() (p []Person) {
	p = db.people
	return
}

func (db *PeopleDb) SetAll(p []Person) {
	db.people = p
}
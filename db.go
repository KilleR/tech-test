package main

import (
	"encoding/json"
	"log"
	"os"
	"io/ioutil"
)

type PeopleDb struct {
}

func (db *PeopleDb) GetAll() (p []Person) {
	// fetch contents of db file
	dbtext, err := ioutil.ReadFile("db")
	if err != nil {
		log.Println("Error reading from DB file:", err)
		return // TODO: Should have error trapping here
	}

	// parse db json into output
	err = json.Unmarshal(dbtext, &p)
	if err != nil {
		log.Println("Error parsing db content:",err)
		return
	}
	return
}

func (db *PeopleDb) SetAll(p []Person) {
	dbtext, err := json.Marshal(p)
	if err != nil {
		log.Println("Error marshalling DB content:", err)
		return
	}
	// open DB file handler - create if not present, truncate before writing
	file, err := os.OpenFile("db", os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0644)

	// write data to file
	_, err = file.Write(dbtext)
	if err != nil {
		log.Println("Error writing to DB file:", err)
	}
}
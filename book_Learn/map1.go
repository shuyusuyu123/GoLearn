package main

import "fmt"

type PersonInfo struct {
	ID   string
	Name string
	Add  string
}

func main() {
	var PersonDB map[string]PersonInfo
	PersonDB = make(map[string]PersonInfo)

	PersonDB["12345"] = PersonInfo{"12345", "Tom", "Room 2003,..."}
	PersonDB["1"] = PersonInfo{"1", "Helen", "Room 2002,..."}
	person, ok := PersonDB["12345"]
	if ok {
		fmt.Println("Found person", person.Name, "with ID 1234")
	} else {
		fmt.Println("Did not find person with ID 1234")
	}
}

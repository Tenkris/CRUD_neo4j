package main

import (
	"fmt"
	"log"
	"neo4j_tutorial_crud/pkg"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

func main() {
	// create a new Neo4j driver
	driver, err := neo4j.NewDriver("bolt://localhost:7687", neo4j.BasicAuth("neo4j", "test1234", ""))
	if err != nil {
		log.Fatalf("Failed to create Neo4j driver: %v", err)
	}
	defer driver.Close()

	// create a new person

	person, err := pkg.CreatePerson(driver, "TEN", 30) // using CreatePerson from mypackage
	if err != nil {
		log.Fatalf("Failed to create person: %v", err)
	}
	log.Printf("Created person: %+v\n", person)

	// get the person by name
	personByName, err := pkg.GetPersonByName(driver, "Alice") // using GetPersonByName from mypackage
	if err != nil {
		log.Fatalf("Failed to get person by name: %v", err)
	}
	fmt.Println(personByName)
	log.Printf("Found person by name: %+v\n", personByName)

	// get the person by ID
	personByID, err := pkg.GetPersonByID(driver, person.ID)
	if err != nil {
		log.Fatalf("Failed to get person by ID: %v", err)
	}
	log.Printf("Found person by ID: %+v\n", personByID)

	// update the person's age
	updatedPerson, err := pkg.UpdatePersonAge(driver, person.ID, 35)
	if err != nil {
		log.Fatalf("Failed to update person's age: %v", err)
	}
	log.Printf("Updated person: %+v\n", updatedPerson)

	// delete the person
	err = pkg.DeletePerson(driver, person.ID)
	if err != nil {
		log.Fatalf("Failed to delete person: %v", err)
	}
	log.Printf("Deleted person with ID %d\n", person.ID)
	fmt.Println("TEN")
}

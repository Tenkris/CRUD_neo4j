// pkg/mypackage.go
package pkg

import (
	"fmt"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

type Person struct {
	ID   int64
	Name string
	Age  int
}

func CreatePerson(driver neo4j.Driver, name string, age int) (*Person, error) {
	session := driver.NewSession(neo4j.SessionConfig{})
	defer session.Close()

	result, err := session.Run(
		"CREATE (p:Person {name: $name, age: $age}) RETURN id(p)",
		map[string]interface{}{"name": name, "age": age},
	)
	if err != nil {
		return nil, err
	}

	record, err := result.Single()
	if err != nil {
		return nil, err
	}

	id, ok := record.Values[0].(int64)
	if !ok {
		return nil, fmt.Errorf("invalid ID type")
	}

	return &Person{ID: id, Name: name, Age: age}, nil
}
func GetPersonByName(driver neo4j.Driver, name string) (*Person, error) {
	session := driver.NewSession(neo4j.SessionConfig{})
	defer session.Close()

	result, err := session.Run(
		"MATCH (p:Person) WHERE p.name = $name RETURN id(p), p.age LIMIT 1",
		map[string]interface{}{"name": name},
	)
	if err != nil {
		return nil, err
	}

	record, err := result.Single()
	if err != nil {
		return nil, err
	}

	id, ok := record.Values[0].(int64)
	if !ok {
		return nil, fmt.Errorf("invalid ID type")
	}

	age, ok := record.Values[1].(int64)
	if !ok {
		return nil, fmt.Errorf("invalid age type")
	}

	return &Person{ID: id, Name: name, Age: int(age)}, nil
}

func GetPersonByID(driver neo4j.Driver, id int64) (*Person, error) {
	session := driver.NewSession(neo4j.SessionConfig{})
	defer session.Close()

	result, err := session.Run(
		"MATCH (p:Person) WHERE id(p) = $id RETURN p.name, p.age",
		map[string]interface{}{"id": id},
	)
	if err != nil {
		return nil, err
	}

	record, err := result.Single()
	if err != nil {
		return nil, err
	}

	name, ok := record.Values[0].(string)
	if !ok {
		return nil, fmt.Errorf("invalid name type")
	}

	age, ok := record.Values[1].(int64)
	if !ok {
		return nil, fmt.Errorf("invalid age type")
	}

	return &Person{ID: id, Name: name, Age: int(age)}, nil
}
func UpdatePersonAge(driver neo4j.Driver, id int64, age int) (*Person, error) {
	session := driver.NewSession(neo4j.SessionConfig{})
	defer session.Close()

	result, err := session.Run(
		"MATCH (p:Person) WHERE id(p) = $id SET p.age = $age RETURN p.name, p.age",
		map[string]interface{}{"id": id, "age": age},
	)
	if err != nil {
		return nil, err
	}

	record, err := result.Single()
	if err != nil {
		return nil, err
	}

	name, ok := record.Values[0].(string)
	if !ok {
		return nil, fmt.Errorf("invalid name type")
	}

	newAge, ok := record.Values[1].(int64)
	if !ok {
		return nil, fmt.Errorf("invalid age type")
	}

	return &Person{ID: id, Name: name, Age: int(newAge)}, nil
}
func DeletePerson(driver neo4j.Driver, id int64) error {
	session := driver.NewSession(neo4j.SessionConfig{})
	defer session.Close()

	_, err := session.Run(
		"MATCH (p:Person) WHERE id(p) = $id DELETE p",
		map[string]interface{}{"id": id},
	)
	if err != nil {
		return err
	}

	return nil
}

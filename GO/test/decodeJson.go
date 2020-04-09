package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// Person struct
type Person struct {
	ID        int    `json:"id"`
	Branch    Branch `json:"branch"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Age       int    `json:"age"`
	Sex       int    `json:"sex"`
}

// Branch struct
type Branch struct {
	ID         int    `json:"id"`
	BranchName string `json:"branch"`
	Prefecture string `json:"prefecture"`
	Address    string `json:"address"`
}

func main() {
	branch := Branch{
		ID:         1,
		BranchName: "officeA",
		Prefecture: "Tokyo",
		Address:    "Shinjuku 1-2-3, Tokyo",
	}

	person := Person{
		ID:        1,
		Branch:    branch,
		FirstName: "Keisuke",
		LastName:  "AAA",
		Age:       29,
		Sex:       0,
	}

	err := (&person).encode("person.json")
	if err != nil {
		fmt.Println("Error encode json :", err)
	}

	var person1 Person
	if person1, err = unmarshal("person.json"); err != nil {
		fmt.Println("Error unmarshalling json :", err)
		return
	}
	var person2 Person
	if person2, err = decode("person.json"); err != nil && err != io.EOF {
		fmt.Println("Error decoding json :", err)
		return
	}

	fmt.Println(person1)
	fmt.Println(person2)

}

func (p *Person) encode(fileName string) (err error) {

	_, err = os.Create(fileName)
	if err != nil {
		fmt.Println("Error create json file :", err)
		return
	}

	data, err := json.MarshalIndent(&p, "", "\t")

	if err = ioutil.WriteFile(fileName, data, 0644); err != nil {
		fmt.Println("Error write structure :", err)
		return
	}
	return
}

func decode(fileName string) (person Person, err error) {
	jsonFile, err := os.Open(fileName)
	if err != nil {
		return
	}
	decoder := json.NewDecoder(jsonFile)

	for {
		err = decoder.Decode(&person)
		if err == io.EOF {
			break
		}
		if err != nil {
			return
		}
	}
	return
}

func unmarshal(fileName string) (person Person, err error) {
	jsonFile, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening file :", err)
		return
	}
	b, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println("Error reading file :", err)
		return
	}
	if err = json.Unmarshal(b, &person); err != nil {
		fmt.Println("Error unmarshalling json :", err)
		return
	}
	return
}

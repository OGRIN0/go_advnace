package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

type Address struct {
	Street  string `json:"street"`
	City    string `json:"city"`
	State   string `json:"state"`
	ZipCode string `json:"zip_code"`
}

type Contact struct {
	Email     *string `json:"email,omitempty"`    // Optional field
	Phone     *string `json:"phone,omitempty"`    // Optional field
	Preferred string  `json:"preferred,omitempty"` // Optional field
}

type Person struct {
	ID        int       `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	BirthDate time.Time `json:"-"` // Custom handling, not directly marshaled
	Address   Address   `json:"address"`
	Contact   Contact   `json:"contact"`
	Tags      []string  `json:"tags,omitempty"`
	Active    bool      `json:"active"`
}

const dateFormat = "2006-01-02"

func (p Person) MarshalJSON() ([]byte, error) {
	type PersonAlias Person
	
	return json.Marshal(&struct {
		PersonAlias
		BirthDate string `json:"birth_date"`
		FullName  string `json:"full_name"`
	}{
		PersonAlias: PersonAlias(p),
		BirthDate:   p.BirthDate.Format(dateFormat),
		FullName:    p.FirstName + " " + p.LastName,
	})
}

func (p *Person) UnmarshalJSON(data []byte) error {
	type ContactAlias Contact
	type AddressAlias Address
	type PersonAlias Person
	
	aux := &struct {
		*PersonAlias
		BirthDate string `json:"birth_date"`
	}{
		PersonAlias: (*PersonAlias)(p),
	}
	
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	
	if aux.BirthDate != "" {
		date, err := time.Parse(dateFormat, aux.BirthDate)
		if err != nil {
			return errors.New("invalid date format for birth_date: must be YYYY-MM-DD")
		}
		p.BirthDate = date
	}
	
	return nil
}

func main() {
	email := "john.doe@example.com"
	phone := "555-123-4567"
	
	person := Person{
		ID:        1,
		FirstName: "John",
		LastName:  "Doe",
		BirthDate: time.Date(1985, time.March, 15, 0, 0, 0, 0, time.UTC),
		Address: Address{
			Street:  "123 Main St",
			City:    "Anytown",
			State:   "CA",
			ZipCode: "12345",
		},
		Contact: Contact{
			Email:     &email,
			Phone:     &phone,
			Preferred: "email",
		},
		Tags:   []string{"developer", "golang"},
		Active: true,
	}
	
	jsonData, err := json.MarshalIndent(person, "", "  ")
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}
	
	fmt.Println("Encoded Person with custom format:")
	fmt.Println(string(jsonData))
	
	personMinimal := Person{
		ID:        2,
		FirstName: "Jane",
		LastName:  "Smith",
		BirthDate: time.Date(1990, time.July, 20, 0, 0, 0, 0, time.UTC),
		Address: Address{
			Street: "456 Oak Ave",
			City:   "Othertown",
			State:  "NY",
		},
		Active: false,
	}
	
	jsonData2, _ := json.MarshalIndent(personMinimal, "", "  ")
	fmt.Println("\nEncoded Person with minimal information:")
	fmt.Println(string(jsonData2))
	
	var decodedPerson Person
	err = json.Unmarshal(jsonData, &decodedPerson)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}
	
	fmt.Println("\nDecoded Person:")
	fmt.Printf("ID: %d\n", decodedPerson.ID)
	fmt.Printf("Name: %s %s\n", decodedPerson.FirstName, decodedPerson.LastName)
	fmt.Printf("Birth Date: %s\n", decodedPerson.BirthDate.Format(dateFormat))
	fmt.Printf("Address: %s, %s, %s %s\n", 
		decodedPerson.Address.Street, 
		decodedPerson.Address.City,
		decodedPerson.Address.State,
		decodedPerson.Address.ZipCode)
	
	if decodedPerson.Contact.Email != nil {
		fmt.Printf("Email: %s\n", *decodedPerson.Contact.Email)
	}
	
	if decodedPerson.Contact.Phone != nil {
		fmt.Printf("Phone: %s\n", *decodedPerson.Contact.Phone)
	}
	
	invalidJSON := []byte(`{
		"id": 3,
		"first_name": "Invalid",
		"last_name": "Person",
		"birth_date": "not-a-date",
		"address": {"street": "Error St", "city": "Bug City", "state": "XX"},
		"active": true
	}`)
	
	var invalidPerson Person
	err = json.Unmarshal(invalidJSON, &invalidPerson)
	if err != nil {
		fmt.Println("\nExpected error with invalid date:", err)
	}
}

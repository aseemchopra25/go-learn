package main

import (
	"encoding/json"
	"fmt"
)

// not public hence lower case
type pii struct {
	Name     string `json:"my cool name"`
	Age      int
	Address  string
	Password string   `json:"-"`              //this will remove the password
	Hobbies  []string `json:"tags,omitempty"` // no space and it will omit empty fields
}

func main() {
	fmt.Println("Welcome to JSON in golang")
	// Converting data into json
	EncodeJson()
}

func EncodeJson() {

	info := []pii{
		{"Aseem", 5, "Australia", "abc123", []string{"pentesting", "golang"}},
		{"ET", 500, "Alient", "eeettteee", []string{"lolcats", "scaring people"}},
		{"ET", 500, "Alient", "eeettteee", nil},
	}

	// Marshal is an implementation of JSON
	// Interface is sort of an alternative of struct
	finalJson, err := json.MarshalIndent(info, "", "\t")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", finalJson)

}

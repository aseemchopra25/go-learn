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
	// EncodeJson()
	DecodeJson()
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

func DecodeJson() {
	jsonDataFromWeb := []byte(`
	{
		"my cool name": "Aseem",
		"Age": 5,
		"Address": "Australia",
		"tags": ["pentesting","golang"]
	}
	`)

	var mypii pii

	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid {
		fmt.Println("JSON was valid")
		json.Unmarshal(jsonDataFromWeb, &mypii)
		fmt.Printf("%#v\n", mypii)

	} else {
		fmt.Println("JSON wasn't valid")
	}

	// cases where data needs to be added to key value pairs

	var myonlinedata map[string]interface{} //don't know what kind of data is coming up use interface as anything could come up

	json.Unmarshal(jsonDataFromWeb, &myonlinedata)
	fmt.Printf("%#v\n", myonlinedata)

	for key, value := range myonlinedata {
		fmt.Printf("Key is %v and value is %v and the type is %T\n", key, value, value)
	}

}

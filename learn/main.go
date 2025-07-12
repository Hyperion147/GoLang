package main

import (
	"encoding/json"
	"fmt"
)

// const webUrl = "https://todo.suryansu.pro"

type course struct {
	Name     string
	Price    int
	Platform string
	Password string   `json:"-"`
	Tag      []string `json:"tags,omitempty"`
}

func main() {
	topic := "This program is to handle json."
	fmt.Println(topic)

	EncodeJsonData()

	DecodeJsonData()

}

func EncodeJsonData() {
	courses := []course{
		{"React", 0, "Youtube", "123@hyper", []string{"web-dev", "react"}},
		{"GoLang", 0, "Youtube", "345@hyper", []string{"web-dev", "gofer"}},
		{"JWT", 0, "Youtube", "678@hyper", nil},
	}

	courseJson, err := json.MarshalIndent(courses, "", "\t")

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", courseJson)

}

func DecodeJsonData() {
	jsonData := []byte(`
        {
            "Name": "GoLang",
            "Price": 0,
            "Platform": "Youtube",
            "tags": ["web-dev","gofer"]
        }
	`)

	var courses course 

	check := json.Valid(jsonData)

	if check {
		fmt.Println("JSON is valid")
		json.Unmarshal(jsonData, &courses)
		fmt.Printf("%#v\n", courses)
	} else {
		fmt.Println("Json is invalid")
	}

	var jsonAdd map[string] interface{}
	json.Unmarshal(jsonData, &jsonAdd)

	for key, value := range jsonAdd {
		fmt.Printf("%v, %v, %T\n", key, value, key)
	}

}

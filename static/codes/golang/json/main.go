package main

import (
	"encoding/json"
	"log"
)

func main() {
	jsonStr := `
	["one", "two"]
	`
	var array []string
	err := json.Unmarshal([]byte(jsonStr), &array)
	if err != nil {
		log.Println(err)
	} else {
		log.Println(array)
	}
}

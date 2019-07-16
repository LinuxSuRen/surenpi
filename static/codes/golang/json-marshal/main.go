package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

var json_str = `{
	"value": "GOOD",
	"name":"name"
}`

var json_bool = `{
	"value": false,
	"name":"good"
}`

type Form struct {
	Value ComposeValue
	Name  string
}

type ComposeValue struct {
	Value string
}

func (f *ComposeValue) UnmarshalJSON(data []byte) error {
	originStr := strings.Trim(string(data), `"`)
	switch str := strings.ToLower(originStr); str {
	case "true":
		f.Value = "true"
	case "false":
		f.Value = "false"
	default:
		f.Value = originStr
	}
	return nil
}

func (f ComposeValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(f.Value)
}

func (f ComposeValue) String() string {
	return f.Value
}

func main() {
	var data []byte
	var err error

	form := &Form{}
	if err := json.Unmarshal([]byte(json_str), form); err != nil {
		log.Fatal(err)
	}
	if data, err = json.Marshal(form); err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))

	if err := json.Unmarshal([]byte(json_bool), form); err != nil {
		log.Fatal(err)
	}
	data, _ = json.Marshal(form)
	fmt.Println(string(data))
}

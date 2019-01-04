package main

import "fmt"

func main() {
	demo4Array()

	demo4Map()
}

func demo4Map() {
	var strMap map[string]string

	strMap = make(map[string]string)
	strMap["name"] = "link"

	fmt.Println(strMap)
}

func demo4Array() {
	var array []string
	array = make([]string, 1)

	array[0] = "one"
	array = append(array, "two")
	array = append(array, "three", "four")

	others := []string{"1", "2"}
	array = append(array, others...)

	fmt.Println(array)
	fmt.Println(len(array))
}

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	url := url.URL{
		Host:   "linuxsuren.github.io",
		Path:   "/index.json",
		Scheme: "https",
	}

	fmt.Println(url.RequestURI())
	fmt.Println(url.String())

	response, err := http.Get(url.String())
	if err != nil {
		fmt.Printf("get error %v\n", err)
		return
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("read error %v\n", err)
		return
	}

	fmt.Println(string(data))
}

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	test()
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

func test() {
	data := `json={"parameter":[{"name":"name","value":"bob"},{"name":"age","value":"12"}]}`
	inputURL := &url.URL{
		Host:   "jenkins.k8s.surenpi.com",
		Path:   "job/demo/job/demo-input-request-3/6/input/6be09e0ef4dcfd27eb2c985c3cff02ae/proceed",
		Scheme: "http",
	}
	fmt.Println(data)
	fmt.Println(inputURL)
	formData := url.Values{"json": {`{"parameter":[{"name":"name","value":"bob"},{"name":"age","value":"12"}]}`}}
	payload := strings.NewReader(formData.Encode())
	fmt.Println(formData.Encode())
	fmt.Println(payload)
	fmt.Println(inputURL.String())
	request, err := http.NewRequest(http.MethodPost, inputURL.String(), payload)
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	request.SetBasicAuth("user", "passwd")
	if err != nil {
		fmt.Println(err)
	} else {
		client := &http.Client{}
		resp, err := client.Do(request)
		if err != nil {
			fmt.Println(err)
		} else {
			data, _ := ioutil.ReadAll(resp.Body)
			fmt.Println(string(data))
		}
	}
}

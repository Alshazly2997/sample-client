package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var token string

func main() {
	var name, passowrd string
	response, err := http.Get("http://localhost:8080")
	if err != nil {
		fmt.Printf("The http request failed %v", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(data))
	}

	fmt.Println("Please enter your name and then your password:")
	fmt.Scan(&name)
	fmt.Scan(&passowrd)

	jsonData := map[string]string{"name": name, "password": passowrd}
	jsonValue, _ := json.Marshal(jsonData)
	response, err = http.Post("http://localhost:8080/auth", "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		fmt.Printf("The http request failed %v", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(data))
		token = string(data)
	}

	req, err := http.NewRequest("GET", "http://localhost:8080/query", nil)
	req.Header.Add("authorization", "Bearar "+token)

	client := &http.Client{}
	response, err = client.Do(req)
	if err != nil {
		fmt.Printf("The http request failed %v", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(data))
	}

}

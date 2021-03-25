package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

//GetSource(*http.Client,string)
func GetSource(client *http.Client, url string) (*string, error) {
	//Input:Pointer to the http.Client and the url as a string
	//Returns the source code in a string and an error

	request, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, err
	}

	response, err := client.Do(request)

	if err != nil {
		return nil, err
	}

	if response.Status != "200 OK" {
		fmt.Println(response.Status)
		return nil, errors.New("Page not found")
	}

	defer response.Body.Close()

	responseBytes, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	str := string(responseBytes)
	str_ptr := &str

	return str_ptr, nil
}

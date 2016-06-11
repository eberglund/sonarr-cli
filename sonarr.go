package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const baseUrl = "http://localhost:8989/api/"

func main() {
	refreshSeries()
}

func refreshSeries() {
	apiKey := readApiKey()
	endpoint := baseUrl + "command?apikey=" + apiKey
	fmt.Println(endpoint)
	data := []byte(`{"name":"RefreshSeries"}`)
	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(data))

	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("submitted!")
	fmt.Printf("%s", body)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readApiKey() string {
	key, err := ioutil.ReadFile("api_key")
	check(err)
	return strings.TrimSpace(string(key))
}

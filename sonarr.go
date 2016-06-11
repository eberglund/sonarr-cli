package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const baseUrl = "http://localhost:8989/api/"

func main() {
	refreshSeries()
}

type command struct {
	Name string
}

func refreshSeries() {
	apiKey := readApiKey()
	endpoint := baseUrl + "command?apikey=" + apiKey

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(command{Name: "RefreshSeries"})
	resp, err := http.Post(endpoint, "application/json", b)

	check(err)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	check(err)

	fmt.Printf("Response:\n%s", body)
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

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const baseUrl = "http://localhost:8989/api/"

func main() {
	switch os.Args[1] {
	case "refresh":
		refreshSeries()
	case "search":
		search()
	case "list":
		list()
	}
}

type command struct {
	Name string
}

func search() {

}

func list() {
	apiKey := readApiKey()
	resp, err := http.Get(baseUrl + "/series?apikey=" + apiKey)

	check(err)

	fmt.Printf("Response:\n%s", getBody(resp))
}

func refreshSeries() {
	sendCommand("RefreshSeries")
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

func sendCommand(name string) {
	apiKey := readApiKey()
	endpoint := baseUrl + "command?apikey=" + apiKey

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(command{Name: name})
	resp, err := http.Post(endpoint, "application/json", b)

	check(err)

	fmt.Printf("Response:\n%s", getBody(resp))
}

func getBody(resp *http.Response) []byte {
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	check(err)

	return body
}

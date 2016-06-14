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
		search(os.Args[2])
	case "list":
		list()
	}
}

type command struct {
	Name string
}

func search(seriesId string) {

}

func list() {
	resp, err := http.Get(getUrl("series"))

	check(err)

	fmt.Printf("Response:\n%s", getBody(resp))
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func refreshSeries() {
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(command{Name: "RefreshSeries"})
	resp, err := http.Post(getUrl("command"), "application/json", b)

	check(err)

	fmt.Printf("Response:\n%s", getBody(resp))
}

func getBody(resp *http.Response) []byte {
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	check(err)

	return body
}

func getUrl(endpoint string) string {
	apiKey := readApiKey()
	return baseUrl + endpoint + "?apikey=" + apiKey
}

func readApiKey() string {
	key, err := ioutil.ReadFile("api_key")
	check(err)
	return strings.TrimSpace(string(key))
}

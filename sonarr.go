package sonarr

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
	s := SonarrAPI("http://localhost:8989/api/", readApiKey())

	switch os.Args[1] {
	case "refresh":
		s.RefreshSeries()
	case "search":
		s.Search()
	case "list":
		s.List()
	}
}

type Sonarr interface {
	RefreshSeries()
	List()
	Search()
}

type api struct {
	baseUrl string
	apiKey  string
}

type command struct {
	Name string
}

func SonarrAPI(baseUrl string, apiKey string) Sonarr {
	return api{baseUrl: baseUrl, apiKey: apiKey}
}

func (a api) Search() {

}

func (a api) RefreshSeries() {

}

func (a api) List() {
	resp, err := http.Get(a.getUrl("series"))

	check(err)

	fmt.Printf("Response:\n%s", getBody(resp))
}

func (a api) refreshSeries() {
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(command{Name: "RefreshSeries"})
	resp, err := http.Post(a.getUrl("command"), "application/json", b)

	check(err)

	fmt.Printf("Response:\n%s", getBody(resp))
}

func (a api) getUrl(endpoint string) string {
	return a.baseUrl + endpoint + "?apikey=" + a.apiKey
}

func readApiKey() string {
	key, err := ioutil.ReadFile("api_key")
	check(err)
	return strings.TrimSpace(string(key))
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getBody(resp *http.Response) []byte {
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	check(err)

	return body
}

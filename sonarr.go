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

func main() {
	s := SonarrAPI("http://localhost:8989/api/", readApiKey())

	switch os.Args[1] {
	case "refresh":
		s.RefreshSeries()
	case "search":
		s.Search()
	case "list":
		series := s.SeriesList()
		for i := 0; i < len(series); i++ {
			fmt.Printf("%s (%d) - %d / %d\n", series[i].Title, series[i].Id, series[i].EpisodeFileCount, series[i].EpisodeCount)
		}
	}
}

type Sonarr interface {
	RefreshSeries()
	SeriesList() []Series
	Search()
}

type Series struct {
	Title            string
	Id               int
	EpisodeCount     int
	EpisodeFileCount int
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

func (a api) SeriesList() []Series {
	resp, err := http.Get(a.getUrl("series"))

	check(err)

	series := make([]Series, 0)

	json.Unmarshal(getBody(resp), &series)

	return series

}

func (a api) refreshSeries() {
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(command{Name: "RefreshSeries"})
	resp, err := http.Post(a.getUrl("command"), "application/json", b)

	check(err)

	fmt.Printf("ASDF", getBody(resp))
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

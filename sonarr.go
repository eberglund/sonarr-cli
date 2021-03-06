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
		s.SearchAllSeries()
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
	SearchAllSeries()
	Search(id int)
}

type Series struct {
	Title            string
	Id               int
	EpisodeCount     int
	EpisodeFileCount int
}

func SonarrAPI(baseUrl string, apiKey string) Sonarr {
	return api{baseUrl: baseUrl, apiKey: apiKey}
}

type api struct {
	baseUrl string
	apiKey  string
}

func (a api) SearchAllSeries() {
	series := a.SeriesList()
	for i := 0; i < len(series); i++ {
		fmt.Printf("Searching for %s\n", series[i].Title)
		a.Search(series[i].Id)
	}
}

func (a api) Search(id int) {
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(struct {
		Name     string
		SeriesId int
	}{
		"SeriesSearch",
		id,
	})
	_, err := http.Post(a.getUrl("command"), "application/json", b)
	check(err)
}

func (a api) SeriesList() []Series {
	resp, err := http.Get(a.getUrl("series"))

	check(err)

	series := make([]Series, 0)

	json.Unmarshal(getBody(resp), &series)

	return series
}

func (a api) RefreshSeries() {
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(struct{ Name string }{"RefreshSeries"})
	_, err := http.Post(a.getUrl("command"), "application/json", b)

	check(err)
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

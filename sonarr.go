package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	refreshSeries()
}

func refreshSeries() {
	endpoint := "http://localhost:8989/api/command"
	data := url.Values{}
	data.Set("name", "RefreshSeries")
	resp, err := http.PostForm(endpoint, data)

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

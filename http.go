package main

import(
	"net/http"
	"encoding/json"
	"io"
)

type LocationResponse struct {
	Count int
	Next *string
	Previous *string
	Results []Location
}

type Location struct {
	Name string
	Url string
}

func getResponse(url string) (*LocationResponse, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil{
		return nil, err
	}
	var lp LocationResponse
	err = json.Unmarshal(data, &lp)
	if err != nil {
		return nil, err
	}
	return &lp, nil
}
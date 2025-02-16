package main

import (
	"net/http"
	"encoding/json"
	"io"
)

func getLocationAreas(url string) (locationAreas, error) {
	res, err := http.Get(url)
	if err != nil {
		return locationAreas{}, err
	}
	defer res.Body.Close()
	
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return locationAreas{}, err
	}

	la := locationAreas{}
	err = json.Unmarshal(data, &la)
	if err != nil {
		return locationAreas{}, err
	}
	return la, err
}
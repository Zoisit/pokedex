package pokeapi

import (
	"net/http"
	"encoding/json"
	"io"
	"github.com/Zoisit/pokedex/internal/pokecache"
	"time"
	"fmt"
)

var (
	cache *pokecache.Cache
)

const (
	api = "https://pokeapi.co/api/v2/"
)

func init() {
    cache = pokecache.NewCache(2 * time.Minute)
}

func GetLocationAreas(url string) (locationAreas, error) {
	data, ok := cache.Get(url)
	if ok {
		la := locationAreas{}
		err := json.Unmarshal(data, &la)
		if err != nil {
			return locationAreas{}, err
		}
		return la, err
	}

	res, err := http.Get(url)
	if err != nil {
		return locationAreas{}, err
	}
	defer res.Body.Close()
	
	data, err = io.ReadAll(res.Body)
	if err != nil {
		return locationAreas{}, err
	}

	la := locationAreas{}
	err = json.Unmarshal(data, &la)
	if err != nil {
		return locationAreas{}, err
	}

	cache.Add(url, data) 

	return la, err
}

func GetLocationInfo(location string) (LocationInfo, error) {
	url := api + "location-area/" + location
	data, ok := cache.Get(url)
	if ok {
		la := LocationInfo{}
		err := json.Unmarshal(data, &la)
		if err != nil {
			return LocationInfo{}, err
		}
		return la, err
	}

	res, err := http.Get(url)
	if err != nil {
		return LocationInfo{}, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return LocationInfo{}, fmt.Errorf(res.Status)
	}
	
	data, err = io.ReadAll(res.Body)
	if err != nil {
		return LocationInfo{}, err
	}

	la := LocationInfo{}
	err = json.Unmarshal(data, &la)
	if err != nil {
		return LocationInfo{}, err
	}

	cache.Add(url, data) 

	return la, err
}

func GetPokemonInfo(pokemon string) (PokemonInfo, error) {
	url := api + "pokemon/" + pokemon
	
	data, err := getJSON(url)

	if err != nil {
		return PokemonInfo{}, err
	} else {
		la := PokemonInfo{}
		err = json.Unmarshal(data, &la)
		if err != nil {
			return PokemonInfo{}, err
		}

		cache.Add(url, data) 

		return la, err
	}

	
}

func getJSON(url string) ([]byte, error) {
	data, ok := cache.Get(url)
	if ok {
		return data, nil
	}

	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode > 299 {
		return nil, fmt.Errorf(res.Status)
	}
	
	data, err = io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return data, nil
}
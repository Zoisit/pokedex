package pokeapi

import (
	"net/http"
	"encoding/json"
	"io"
	"github.com/Zoisit/pokedex/internal/pokecache"
	"time"
)

var (cache *pokecache.Cache)

func init() {
    cache = pokecache.NewCache(30 * time.Second)
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
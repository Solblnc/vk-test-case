package service

import (
	"Vk-internship/pkg/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const storageTime = 10 * time.Second

var (
	client = &http.Client{}
	cache  = make(map[string]*models.Floor)
)

// GetFloor - make and http request to magic eden api for getting a floor (price) of a collection
func GetFloor(symbol string) float64 {
	// checking
	if _, ok := cache[symbol]; ok && time.Since(cache[symbol].Time) < storageTime {
		return -1
	} else {
		// http client and creating request
		client = &http.Client{}
		req, _ := http.NewRequest("GET", fmt.Sprintf("https://api-mainnet.magiceden.dev/v2/collections/%s/stats", symbol), nil)

		// Do request
		res, _ := client.Do(req)

		defer res.Body.Close()

		body, _ := ioutil.ReadAll(res.Body)

		// Unmarshalling response to struct
		var floorResponse models.Stats
		json.Unmarshal(body, &floorResponse)

		cache[symbol] = &models.Floor{
			Value: floorResponse.FloorPrice / 1e9,
			Time:  time.Now(),
		}
	}
	return cache[symbol].Value
}

// GetVolume - make and http request to magic eden api for getting volume of a collection
func GetVolume(symbol string) float64 {

	cache := make(map[string]*models.Stats)

	if _, ok := cache[symbol]; ok {
		return -1
	} else {
		client = &http.Client{}
		req, _ := http.NewRequest(
			"GET", fmt.Sprintf("https://api-mainnet.magiceden.dev/v2/collections/%s/stats", symbol), nil)

		res, _ := client.Do(req)

		defer res.Body.Close()

		body, _ := ioutil.ReadAll(res.Body)

		var floorResponse models.Stats
		json.Unmarshal(body, &floorResponse)

		cache[symbol] = &models.Stats{
			VolumeAll: floorResponse.VolumeAll / 1e9,
		}
	}
	return cache[symbol].VolumeAll
}

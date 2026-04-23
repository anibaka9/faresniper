package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

const (
	countries_link = "https://api.travelpayouts.com/data/en/countries.json"
	cities_link    = "https://api.travelpayouts.com/data/ru/cities.json"
	airport_link   = "https://api.travelpayouts.com/data/en/airports.json"
	airlines_link  = "https://api.travelpayouts.com/data/en/airlines.json"
)

func downloadFile(url string, filepath string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	file, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	return err
}

func main() {
	err := downloadFile(countries_link, "data/iata/countries.json")
	if err != nil {
		log.Fatal(err)
	}
	err = downloadFile(cities_link, "data/iata/cities.json")
	if err != nil {
		log.Fatal(err)
	}
	err = downloadFile(airport_link, "data/iata/airport.json")
	if err != nil {
		log.Fatal(err)
	}
	err = downloadFile(airlines_link, "data/iata/airlines.json")
	if err != nil {
		log.Fatal(err)
	}
}

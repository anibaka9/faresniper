package main

import (
	"io"
	"log"
	"net/http"
	"os"

	"github.com/anibaka9/faresniper/internal/travelpayouts"
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
	log.Println("Downloading reference data...")

	err := downloadFile(countries_link, "data/iata/countries.json")
	if err != nil {
		log.Fatal(err)
	}
	err = downloadFile(cities_link, "data/iata/cities.json")
	if err != nil {
		log.Fatal(err)
	}
	err = downloadFile(airport_link, "data/iata/airports.json")
	if err != nil {
		log.Fatal(err)
	}
	err = downloadFile(airlines_link, "data/iata/airlines.json")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Saving to database...")

	if err := travelpayouts.SaveCountries(); err != nil {
		log.Fatal(err)
	}
	log.Println("Countries saved")

	if err := travelpayouts.SaveCities(); err != nil {
		log.Fatal(err)
	}
	log.Println("Cities saved")

	if err := travelpayouts.SaveAirports(); err != nil {
		log.Fatal(err)
	}
	log.Println("Airports saved")

	if err := travelpayouts.SaveAirlines(); err != nil {
		log.Fatal(err)
	}
	log.Println("Airlines saved")

	log.Println("Done")
}

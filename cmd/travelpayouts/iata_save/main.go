package main

import (
	"log"

	"github.com/anibaka9/faresniper/internal/travelpayouts"
)

func main() {
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

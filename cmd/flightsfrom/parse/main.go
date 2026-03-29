package main

import (
	"log"
	"os"

	"github.com/anibaka9/faresniper/internal/flightsfrom"
)

func main() {
	hjsonFightsData, err := os.ReadFile("data/flightsfrom.hjson")
	if err != nil {
		log.Fatal("cant open data/flightsfrom.hjson", err)
	}
	flightData, err := flightsfrom.Parse(hjsonFightsData)
	if err != nil {
		log.Fatal("cant parse data: ", err)
	}
	err = flightsfrom.SaveToDB(flightData)
	if err != nil {
		log.Fatal("cant save: ", err)
	}
}

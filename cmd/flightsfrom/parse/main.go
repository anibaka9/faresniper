package main

import (
	"fmt"
	"log"
	"os"

	"github.com/anibaka9/faresniper/internal/flightsfrom"
)

func main() {
	hjsonFightsData, err := os.ReadFile("data/flightsfrom.hjson")
	if err != nil {
		log.Fatal("cant open data/flightsfrom.hjson", err)
	}
	fightsData, err := flightsfrom.Parse(hjsonFightsData)
	if err != nil {
		log.Fatal("cant parse data", err)
	}
	fmt.Println(len(fightsData.AllDestinations))
}

package main

import (
	"log"

	"github.com/anibaka9/faresniper/internal/travelpayouts"
)

const (
	origin = "BSZ"
	date   = "2026-04-28"
)

func main() {
	err := travelpayouts.SavePrices(origin, date)
	if err != nil {
		log.Fatal(err)
	}
}

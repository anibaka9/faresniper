package main

import (
	"log"
	"os"

	"github.com/anibaka9/faresniper/internal/flightsfrom"
)

func main() {
	file, err := os.ReadFile("data/flightsfrom.html")
	if err != nil {
		log.Fatal("couldnt open flightsfrom.html")
	}

	data, err := flightsfrom.Exctact(file)
	if err != nil {
		log.Fatal("error extracting data", err)
	}

	if err := os.WriteFile("data/flightsfrom.hjson", data, 0o666); err != nil {
		log.Fatal("coudnt write data", err)
	}
}

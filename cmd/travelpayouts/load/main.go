package main

import (
	"fmt"
	"log"
	"os"

	"github.com/anibaka9/faresniper/internal/travelpayouts"
)

const (
	origin      = "BSZ"
	destination = "ALA"
	month       = "2026-05-01"
)

func main() {
	data, err := travelpayouts.GetCalendar(origin, destination, month)
	if err != nil {
		log.Fatal(err)
	}

	fileName := fmt.Sprintf("data/travelpayouts/%s-%s-%s.json", origin, destination, month)
	if err := os.WriteFile(fileName, data, 0o666); err != nil {
		log.Fatal("coudnt write data", err)
	}
}

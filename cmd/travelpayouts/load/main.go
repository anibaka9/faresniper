package main

import (
	"fmt"
	"log"
	"os"

	"github.com/anibaka9/faresniper/internal/travelpayouts"
)

const (
	origin = "BSZ"
)

func main() {
	data, err := travelpayouts.GetPrices(origin)
	if err != nil {
		log.Fatal(err)
	}

	fileName := fmt.Sprintf("data/travelpayouts/%s.json", origin)
	if err := os.WriteFile(fileName, data, 0o666); err != nil {
		log.Fatal("coudnt write data", err)
	}
}

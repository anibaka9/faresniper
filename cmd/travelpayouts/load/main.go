package main

import (
	"fmt"
	"log"
	"os"
	"time"

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
	currentDay := time.Now().Local().Format("2006-01-02")
	fileName := fmt.Sprintf("data/travelpayouts/%s_%s.json", origin, currentDay)
	if err := os.WriteFile(fileName, data, 0o666); err != nil {
		log.Fatal("coudnt write data", err)
	}
}

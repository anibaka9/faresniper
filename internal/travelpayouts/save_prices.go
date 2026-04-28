package travelpayouts

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/anibaka9/faresniper/internal/database"
	databaseclient "github.com/anibaka9/faresniper/internal/database_client"
)

func SavePrices(origin string, date string) error {
	ctx := context.Background()
	c, err := databaseclient.NewClient()
	if err != nil {
		return err
	}

	fileName := fmt.Sprintf("data/travelpayouts/%s_%s.json", origin, date)
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	var flightsResponse FlightsResponse
	err = json.NewDecoder(file).Decode(&flightsResponse)
	if err != nil {
		return err
	}
	flights := flightsResponse.Data

	requestDate, err := time.Parse("2006-01-02", date)
	if err != nil {
		return err
	}

	fmt.Printf("saving %v flights\n", len(flights))

	for _, flight := range flights {
		err := c.Queries.CreatePriceSnapshot(ctx, database.CreatePriceSnapshotParams{
			FlightNumber:    flight.FlightNumber,
			OriginIata:      flight.OriginAirport,
			DestinationIata: flight.DestinationAirport,
			DepartureAt:     flight.DepartureAt.Format(time.RFC3339),
			AirlineIata:     flight.Airline,
			Price:           flight.Price,
			ObservedAt:      requestDate.Format(time.RFC3339),
		})
		if err != nil {
			return err
		}
	}

	return nil
}

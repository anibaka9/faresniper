package flightsfrom

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/anibaka9/faresniper/internal/database"
	databaseclient "github.com/anibaka9/faresniper/internal/database_client"
	_ "github.com/mattn/go-sqlite3"
)

func SaveToDB(flightData *FlightData) error {
	c, error := databaseclient.NewClient()
	ctx := context.Background()
	if error != nil {
		return fmt.Errorf("cant open db: %w", error)
	}
	for _, country := range flightData.Countries {
		fmt.Printf("saving country %s", country.Country)
		_, error = c.Queries.CreateCountry(ctx, database.CreateCountryParams{
			CountryCode: country.CountryCode,
			Name:        country.Country,
		})

		if error != nil {
			return fmt.Errorf("couldnt insert new country: %w", error)
		}
	}

	for _, destination := range flightData.AllDestinations {
		fmt.Printf("saving destination %s", destination.Airport.Name)

		_, err := c.Queries.GetAirportByIata(ctx, destination.Airport.IATA)
		if err == sql.ErrNoRows {
		} else if err != nil {
			return fmt.Errorf("couldnt get airport: %w", error)
		} else {
			continue
		}

		city, err := c.Queries.GetCityByCountyCodeAndName(ctx, database.GetCityByCountyCodeAndNameParams{
			Name:        destination.Airport.CityName,
			CountryCode: destination.Airport.CountryCode,
		})
		if err == sql.ErrNoRows {
			country, err := c.Queries.GetCountryByCode(ctx, destination.Airport.CountryCode)

			if err == sql.ErrNoRows {
				country, err = c.Queries.CreateCountry(ctx, database.CreateCountryParams{
					CountryCode: destination.Airport.CountryCode,
					Name:        destination.Airport.Country,
				})
				if err != nil {
					return fmt.Errorf("couldnt create country: %w", error)
				}
			} else if err != nil {
				return fmt.Errorf("couldnt get country: %w", error)
			}
			city, err = c.Queries.CreateCity(context.Background(), database.CreateCityParams{
				Name:      destination.Airport.CityName,
				CountryID: country.ID,
			})
		} else if err != nil {
			return fmt.Errorf("couldnt create city: %w", error)
		}

		airport, err := c.Queries.CreateAirport(ctx, database.CreateAirportParams{
			Iata:   destination.Airport.IATA,
			Name:   destination.Airport.CityName,
			CityID: city.ID,
		})
		if err != nil {
			return fmt.Errorf("couldnt create airport: %w", error)
		}
		fmt.Println(airport)
	}

	return nil
}

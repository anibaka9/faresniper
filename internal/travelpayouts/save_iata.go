package travelpayouts

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"github.com/anibaka9/faresniper/internal/database"
	databaseclient "github.com/anibaka9/faresniper/internal/database_client"
)

func LoadCountries() ([]Country, error) {
	file, err := os.Open("data/iata/countries.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var countries []Country
	err = json.NewDecoder(file).Decode(&countries)
	if err != nil {
		return nil, err
	}

	return countries, nil
}

func LoadCities() ([]City, error) {
	file, err := os.Open("data/iata/cities.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var cities []City
	err = json.NewDecoder(file).Decode(&cities)
	if err != nil {
		return nil, err
	}

	return cities, nil
}

func LoadAirports() ([]Airport, error) {
	file, err := os.Open("data/iata/airports.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var airports []Airport
	err = json.NewDecoder(file).Decode(&airports)
	if err != nil {
		return nil, err
	}

	return airports, nil
}

func LoadAirlines() ([]Airline, error) {
	file, err := os.Open("data/iata/airlines.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var airlines []Airline
	err = json.NewDecoder(file).Decode(&airlines)
	if err != nil {
		return nil, err
	}

	return airlines, nil
}

func SaveCountries() error {
	ctx := context.Background()
	c, err := databaseclient.NewClient()
	if err != nil {
		return err
	}

	count, err := c.Queries.CountCountries(ctx)
	if err != nil {
		return err
	}

	if count != 0 {
		return nil
	}

	countries, err := LoadCountries()
	if err != nil {
		return err
	}

	for _, country := range countries {
		_, err := c.Queries.CreateCountry(ctx, database.CreateCountryParams{
			Code: country.Code,
			Name: country.Name,
		})
		if err != nil {
			log.Printf("skip country %s: %v", country.Code, err)
		}
	}

	return nil
}

func SaveCities() error {
	ctx := context.Background()
	c, err := databaseclient.NewClient()
	if err != nil {
		return err
	}

	count, err := c.Queries.CountCities(ctx)
	if err != nil {
		return err
	}

	if count != 0 {
		return nil
	}

	cities, err := LoadCities()
	if err != nil {
		return err
	}

	for _, city := range cities {
		_, err := c.Queries.CreateCity(ctx, database.CreateCityParams{
			Iata:        city.Code,
			Name:        city.Name,
			CountryCode: city.CountryCode,
		})
		if err != nil {
			log.Printf("skip city %s: %v", city.Code, err)
		}
	}

	return nil
}

func SaveAirports() error {
	ctx := context.Background()
	c, err := databaseclient.NewClient()
	if err != nil {
		return err
	}

	count, err := c.Queries.CountAirports(ctx)
	if err != nil {
		return err
	}

	if count != 0 {
		return nil
	}

	airports, err := LoadAirports()
	if err != nil {
		return err
	}

	for _, airport := range airports {
		_, err := c.Queries.CreateAirport(ctx, database.CreateAirportParams{
			Iata:       airport.Code,
			Name:       airport.Name,
			IataType:   airport.IataType,
			Flightable: airport.Flightable,
			CityIata:   airport.CityCode,
		})
		if err != nil {
			log.Printf("skip airport %s: %v", airport.Code, err)
		}
	}

	return nil
}

func SaveAirlines() error {
	ctx := context.Background()
	c, err := databaseclient.NewClient()
	if err != nil {
		return err
	}

	count, err := c.Queries.CountAirlines(ctx)
	if err != nil {
		return err
	}

	if count != 0 {
		return nil
	}

	airlines, err := LoadAirlines()
	if err != nil {
		return err
	}

	for _, airline := range airlines {
		_, err := c.Queries.CreateAirline(ctx, database.CreateAirlineParams{
			Iata:      airline.Code,
			Name:      airline.Name,
			IsLowcost: airline.IsLowcost,
		})
		if err != nil {
			log.Printf("skip airline %s: %v", airline.Code, err)
		}
	}

	return nil
}

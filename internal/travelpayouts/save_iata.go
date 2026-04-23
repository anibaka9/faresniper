package travelpayouts

import (
	"context"
	"encoding/json"
	"os"

	"github.com/anibaka9/faresniper/internal/database"
	databaseclient "github.com/anibaka9/faresniper/internal/database_client"
)

func LoadCountries() ([]Country, error) {
	file, err := os.Open("data/countries")
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
	file, err := os.Open("data/cities")
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
		c.Queries.CreateCountry(ctx, database.CreateCountryParams{
			CountryCode: country.Code,
			Name:        country.Name,
		})
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
		c.Queries.CreateCountry(ctx, database.CreateCountryParams{
			CountryCode: country.Code,
			Name:        country.Name,
		})
	}

	return nil
}

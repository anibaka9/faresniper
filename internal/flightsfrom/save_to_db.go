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
	if err := SaveCountries(flightData.Countries, ctx, c); err != nil {
		return fmt.Errorf("cant save countries: %w", err)
	}

	for _, destination := range flightData.AllDestinations {

		fmt.Printf("saving destination %s\n", destination.Airport.Name)

		airport, err := getAirport(destination.Airport.IATA, ctx, c)
		if err != nil {
			return fmt.Errorf("cant check if airport exist: %w", err)
		}

		if airport != nil {
			fmt.Printf("airport %s already exist, continue", destination.Airport.Name)
			continue
		}

		city, err := getCity(destination.Airport.CityName, destination.Airport.CountryCode, ctx, c)
		if err != nil {
			return fmt.Errorf("couldnt get city: %w", err)
		}

		if city == nil {
			country, err := getCountry(destination.Airport.CountryCode, ctx, c)
			if err != nil {
				return fmt.Errorf("couldnt get country: %w", err)
			}
			if country == nil {
				country, err = CreateCountry(&Country{
					Country:     destination.Airport.Country,
					CountryCode: destination.Airport.CountryCode,
				}, ctx, c)
				if err != nil {
					return fmt.Errorf("couldnt create country: %w", err)
				}
			}
			city, err = CreateCity(destination.Airport.CityName, country.ID, ctx, c)
			if err != nil {
				return fmt.Errorf("cant create city: %w", err)
			}
		}

		newAirport, err := c.Queries.CreateAirport(ctx, database.CreateAirportParams{
			Iata:   destination.Airport.IATA,
			Name:   destination.Airport.CityName,
			CityID: city.ID,
		})
		if err != nil {
			return fmt.Errorf("couldnt create airport: %w", error)
		}
		fmt.Println(newAirport)

	}

	return nil
}

func getCountry(countryCode string, ctx context.Context, c databaseclient.Client) (*database.Country, error) {
	country, err := c.Queries.GetCountryByCode(ctx, countryCode)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("couldnt get country: %w", err)
	} else {
		return &country, nil
	}
}

func getAirport(iata string, ctx context.Context, c databaseclient.Client) (*database.Airport, error) {
	airport, err := c.Queries.GetAirportByIata(ctx, iata)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("couldnt get airport: %w", err)
	} else {
		return &airport, nil
	}
}

func getCity(cityName string, countryCode string, ctx context.Context, c databaseclient.Client) (*database.City, error) {
	city, err := c.Queries.GetCityByCountyCodeAndName(ctx, database.GetCityByCountyCodeAndNameParams{
		Name:        cityName,
		CountryCode: countryCode,
	})
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("couldnt get airport: %w", err)
	} else {
		return &city, nil
	}
}

func CreateCountry(country *Country, ctx context.Context, c databaseclient.Client) (*database.Country, error) {
	countryFromDB, err := c.Queries.CreateCountry(ctx, database.CreateCountryParams{
		CountryCode: country.CountryCode,
		Name:        country.Country,
	})
	if err != nil {
		return nil, fmt.Errorf("cant create conutry: %w", err)
	}
	return &countryFromDB, nil
}

func CreateCity(cityName string, countryID int64, ctx context.Context, c databaseclient.Client) (*database.City, error) {
	city, err := c.Queries.CreateCity(context.Background(), database.CreateCityParams{
		Name:      cityName,
		CountryID: countryID,
	})
	if err != nil {
		return nil, fmt.Errorf("cant create city: %w", err)
	}
	return &city, nil
}

func SaveCountries(countries []Country, ctx context.Context, c databaseclient.Client) error {
	for _, country := range countries {
		fmt.Printf("saving country %s\n", country.Country)

		countryFromDB, err := getCountry(country.CountryCode, ctx, c)
		if err != nil {
			return fmt.Errorf("cant check if country exist: %w", err)
		}

		if countryFromDB != nil {
			fmt.Printf("country %s alredy exist\n", countryFromDB.Name)
			continue
		}

		_, err = CreateCountry(&country, ctx, c)
		if err != nil {
			return fmt.Errorf("couldnt insert new country: %w", err)
		}
	}
	return nil
}

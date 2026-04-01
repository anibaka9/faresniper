package flightsfrom

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/anibaka9/faresniper/internal/database"
	databaseclient "github.com/anibaka9/faresniper/internal/database_client"
	_ "github.com/mattn/go-sqlite3"
)

const (
	tempCountryCode = "KG"
	tempCity        = "Bishkek"
	tempIata        = "BSZ"
	tempAirportName = "Manas International Airport"
)

func SaveToDB(flightData *FlightData) error {
	c, error := databaseclient.NewClient()
	ctx := context.Background()
	if error != nil {
		return fmt.Errorf("cant open db: %w", error)
	}
	if err := SaveCountries(&flightData.Countries, ctx, c); err != nil {
		return fmt.Errorf("cant save countries: %w", err)
	}

	if err := saveDestanations(&flightData.AllDestinations, ctx, c); err != nil {
		return fmt.Errorf("cant save destanations: %w", err)
	}

	if err := saveTempAirportFrom(ctx, c); err != nil {
		return fmt.Errorf("cant create airport from: %w", err)
	}

	if err := saveRoutes(&flightData.Airlines, ctx, c); err != nil {
		return fmt.Errorf("cant save routes: %w", err)
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

func SaveCountries(countries *[]Country, ctx context.Context, c databaseclient.Client) error {
	for _, country := range *countries {
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

func saveDestanations(destinations *[]Destination, ctx context.Context, c databaseclient.Client) error {
	for _, destination := range *destinations {

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
			return fmt.Errorf("couldnt create airport: %w", err)
		}
		fmt.Println(newAirport)

	}
	return nil
}

func saveTempAirportFrom(ctx context.Context, c databaseclient.Client) error {
	_, err := c.Queries.GetAirportByIata(ctx, tempIata)
	var airportExist bool
	if err == sql.ErrNoRows {
		airportExist = false
	} else if err != nil {
		return fmt.Errorf("cant get airport :%w", err)
	} else {
		airportExist = true
	}
	if airportExist {
		return nil
	}

	country, err := c.Queries.GetCountryByCode(ctx, tempCountryCode)
	if err != nil {
		return fmt.Errorf("cant get country: %w", err)
	}
	city, err := c.Queries.CreateCity(ctx, database.CreateCityParams{
		Name:      tempCity,
		CountryID: country.ID,
	})
	if err != nil {
		return fmt.Errorf("cant create city: %w", err)
	}
	_, err = c.Queries.CreateAirport(ctx, database.CreateAirportParams{
		Iata:   tempIata,
		Name:   tempAirportName,
		CityID: city.ID,
	})
	if err != nil {
		return fmt.Errorf("cant create airport: %w", err)
	}
	return nil
}

func saveRoutes(routes *[]AirlineRouteDetail, ctx context.Context, c databaseclient.Client) error {
	for _, route := range *routes {
		_, err := c.Queries.GetRouteByFlightsFromId(ctx, int64(route.ID))
		var noRoute bool
		if err == sql.ErrNoRows {
			noRoute = true
		} else if err != nil {
			return fmt.Errorf("cant get route: %w", err)
		} else {
			noRoute = false
		}
		if noRoute == false {
			fmt.Printf("route from %s to %s by %s already exits, continue\n", route.IataFrom, route.IataTo, route.Airline.IATA)
			continue
		}

		airline, err := c.Queries.GetAirlineByFlightsFromId(ctx, int64(route.Airline.ID))
		if err == sql.ErrNoRows {
			name := ""
			if route.Airline.Name != nil {
				name = *route.Airline.Name
			}
			active := true
			if route.Airline.Active != nil {
				active = *route.Airline.Active == 1
			}

			airline, err = c.Queries.CreateAirline(ctx, database.CreateAirlineParams{
				Iata:     route.Airline.IATA,
				Name:     name,
				IsActive: active,
			})
		} else if err != nil {
			return fmt.Errorf("cant get airline: %w", err)
		}
		airportFrom, err := c.Queries.GetAirportByIata(ctx, route.IataFrom)
		if err != nil {
			return fmt.Errorf("cant get airport route from %s: %w", route.IataFrom, err)
		}
		airportTo, err := c.Queries.GetAirportByIata(ctx, route.IataTo)
		if err != nil {
			return fmt.Errorf("cant get airport route to %s: %w", route.IataTo, err)
		}
		_, err = c.Queries.CreateRoute(ctx, database.CreateRouteParams{
			FlightsfromID: int64(route.ID),
			AirlineID:     airline.ID,
			AirportFromID: airportFrom.ID,
			AirportToID:   airportTo.ID,
			IsActive:      route.IsActive == 1,
		})
		fmt.Printf("created route from %s to %s with airline %s\n", airportFrom.Name, airportTo.Name, airline.Name)

	}
	return nil
}

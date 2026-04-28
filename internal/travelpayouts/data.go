package travelpayouts

import "time"

// Coordinates вынесена в отдельную структуру для переиспользования в Airport и City
type Coordinates struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

// Airline описывает структуру airlines.json
type Airline struct {
	NameTranslations map[string]string `json:"name_translations"`
	Code             string            `json:"code"`
	Name             string            `json:"name"`
	IsLowcost        bool              `json:"is_lowcost"`
}

// Airport описывает структуру airport.json
type Airport struct {
	NameTranslations map[string]string `json:"name_translations"`
	CityCode         string            `json:"city_code"`
	CountryCode      string            `json:"country_code"`
	TimeZone         string            `json:"time_zone"`
	Code             string            `json:"code"`
	IataType         string            `json:"iata_type"`
	Name             string            `json:"name"`
	Coordinates      Coordinates       `json:"coordinates"`
	Flightable       bool              `json:"flightable"`
}

// City описывает структуру cities.json
type City struct {
	NameTranslations     map[string]string `json:"name_translations"`
	Cases                map[string]string `json:"cases"`
	CountryCode          string            `json:"country_code"`
	Code                 string            `json:"code"`
	TimeZone             string            `json:"time_zone"`
	Name                 string            `json:"name"`
	Coordinates          Coordinates       `json:"coordinates"`
	HasFlightableAirport bool              `json:"has_flightable_airport"`
}

// Country описывает структуру countries.json
type Country struct {
	NameTranslations map[string]string `json:"name_translations"`
	Cases            map[string]string `json:"cases"`
	Code             string            `json:"code"`
	Name             string            `json:"name"`
	Currency         string            `json:"currency"`
}

type FlightInfo struct {
	FlightNumber       string    `json:"flight_number"`
	Link               string    `json:"link"`
	OriginAirport      string    `json:"origin_airport"`
	DestinationAirport string    `json:"destination_airport"`
	DepartureAt        time.Time `json:"departure_at"` // Можно использовать string, если не нужен встроенный парсинг времени
	Airline            string    `json:"airline"`
	Destination        string    `json:"destination"`
	Origin             string    `json:"origin"`
	Price              int       `json:"price"` // Можно заменить на float64, если цена может быть дробной
	Gate               string    `json:"gate"`
	ReturnTransfers    int       `json:"return_transfers"`
	Duration           int       `json:"duration"`
	DurationTo         int       `json:"duration_to"`
	DurationBack       int       `json:"duration_back"`
	Transfers          int       `json:"transfers"`
}

type FlightsResponse struct {
	Data     []FlightInfo `json:"data"`
	Currency string       `json:"currency"`
	Success  bool         `json:"success"`
}

package travelpayouts

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

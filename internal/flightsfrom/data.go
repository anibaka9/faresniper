package flightsfrom

// Root Data Struct
type FlightData struct {
	IsInit                    bool                 `json:"isInit"`
	ShowAll                   bool                 `json:"showAll"`
	DurationFrom              int                  `json:"durationFrom"`
	DurationTo                int                  `json:"durationTo"`
	DestinationSorting        string               `json:"destinationSorting"`
	DurationSorting           string               `json:"durationSorting"`
	DepartureFrom             int                  `json:"departureFrom"`
	ArrivalFrom               int                  `json:"arrivalFrom"`
	ArrivalTo                 int                  `json:"arrivalTo"`
	Lowcost                   int                  `json:"lowcost"`
	Oneworld                  int                  `json:"oneworld"`
	DestCount                 int                  `json:"destCount"`
	Staralliance              int                  `json:"staralliance"`
	IsLoading                 bool                 `json:"isLoading"`
	Skyteam                   int                  `json:"skyteam"`
	Carrier                   map[string]any       `json:"carrier"`
	SortingType               string               `json:"sortingType"`
	SortingDirection          string               `json:"sortingDirection"`
	BusinessOnly              bool                 `json:"businessOnly"`
	DepartureTo               int                  `json:"departureTo"`
	PriceFrom                 string               `json:"priceFrom"`
	Filters                   string               `json:"filters"`
	PriceTo                   string               `json:"priceTo"`
	SchedulesCount            int                  `json:"schedulesCount"`
	SelectedCountries         map[string]any       `json:"selectedCountries"`
	From                      string               `json:"from"`
	RefreshAds                bool                 `json:"refreshAds"`
	DistanceFrom              int                  `json:"distanceFrom"`
	DistanceTo                int                  `json:"distanceTo"`
	SchedulesOffset           int                  `json:"schedulesOffset"`
	SelectedMonth             *string              `json:"selectedMonth"`
	EntityType                string               `json:"entityType"`
	Destinations              []Destination        `json:"destinations"`
	AllDestinations           []Destination        `json:"allDestinations"`
	Date                      string               `json:"date"`
	AircraftsToShowCount      int                  `json:"aircraftsToShowCount"`
	Aircrafts                 []Aircraft           `json:"aircrafts"`
	ActiveAircrafts           []string             `json:"activeaircrafts"`
	AircraftsToShow           int                  `json:"aircraftsToShow"`
	SelectedAircrafts         []string             `json:"selectedaircrafts"`
	Schedules                 []any                `json:"schedules"`
	OriginDestinations        any                  `json:"originDestinations"`
	Airlines                  []AirlineRouteDetail `json:"airlines"`
	AirlinesToShow            int                  `json:"airlinesToShow"`
	AirlinesToShowCount       int                  `json:"airlinesToShowCount"`
	CountriesToShowCount      int                  `json:"countriesToShowCount"`
	Days                      map[string]any       `json:"days"`
	Countries                 []Country            `json:"countries"`
	CountriesToShow           int                  `json:"countriesToShow"`
	SortBy                    string               `json:"sortBy"`
	Classes                   map[string]any       `json:"classes"`
	DestinationCountrySorting string               `json:"destinationCountrySorting"`
	PriceSorting              string               `json:"priceSorting"`
	MostFlightsSorting        string               `json:"mostFlightsSorting"`
	AirportNameSorting        string               `json:"airportNameSorting"`
	AirportSorting            string               `json:"airportSorting"`
	DestinationTo             string               `json:"destinationTo"`
	Calendar                  any                  `json:"calendar"`
	SelectedDate              string               `json:"selectedDate"`
	AirportRequest            any                  `json:"airportRequest"`
	SelectedFlight            any                  `json:"selectedFlight"`
	ShowInfoBox               bool                 `json:"showInfoBox"`
	IframeUrl                 string               `json:"iframeUrl"`
}

type Destination struct {
	ID             int            `json:"id"`
	IataFrom       string         `json:"iata_from"`
	IataTo         string         `json:"iata_to"`
	Price          int            `json:"price"`
	Day1           string         `json:"day1"`
	Day2           string         `json:"day2"`
	Day3           string         `json:"day3"`
	Day4           string         `json:"day4"`
	Day5           string         `json:"day5"`
	Day6           string         `json:"day6"`
	Day7           string         `json:"day7"`
	FlightsPerDay  *string        `json:"flights_per_day"`
	CommonDuration int            `json:"common_duration"`
	FlightsPerWeek *int           `json:"flights_per_week"`
	FirstFlight    *string        `json:"first_flight"`
	LastFlight     string         `json:"last_flight"`
	Airport        Airport        `json:"airport"`
	AirlineRoutes  []AirlineRoute `json:"airlineroutes"`
}

type Airport struct {
	IATA        string `json:"IATA"`
	CountryCode string `json:"country_code"`
	Name        string `json:"name"`
	CityName    string `json:"city_name"`
	CityNameEn  string `json:"city_name_en"`
	Country     string `json:"country"`
	NoRoutes    int    `json:"no_routes"`
}

type AirlineRoute struct {
	ID            int     `json:"id"`
	RouteID       int     `json:"route_id"`
	IsNew         int     `json:"is_new"`
	FirstFlight   *string `json:"first_flight"`
	Carrier       string  `json:"carrier"`
	CarrierName   *string `json:"carrier_name"`
	AircraftCodes string  `json:"aircraft_codes"`
}

type Aircraft struct {
	IATA string `json:"IATA"`
	Name string `json:"name"`
}

type AirlineRouteDetail struct {
	ID               int         `json:"id"`
	RouteID          int         `json:"route_id"`
	Carrier          string      `json:"carrier"`
	CarrierName      *string     `json:"carrier_name"`
	Lcc              int         `json:"lcc"`
	IataFrom         string      `json:"iata_from"`
	IataTo           string      `json:"iata_to"`
	Day1             string      `json:"day1"`
	Day2             string      `json:"day2"`
	Day3             string      `json:"day3"`
	Day4             string      `json:"day4"`
	Day5             string      `json:"day5"`
	Day6             string      `json:"day6"`
	Day7             string      `json:"day7"`
	AircraftCodes    string      `json:"aircraft_codes"`
	FirstFlight      *string     `json:"first_flight"`
	LastFlight       string      `json:"last_flight"`
	ClassFirst       int         `json:"class_first"`
	ClassBusiness    int         `json:"class_business"`
	ClassEconomy     int         `json:"class_economy"`
	CommonDuration   int         `json:"common_duration"`
	MinDuration      int         `json:"min_duration"`
	MaxDuration      int         `json:"max_duration"`
	IsNew            int         `json:"is_new"`
	IsActive         int         `json:"is_active"`
	IsLayover        int         `json:"is_layover"`
	PassengersPerDay int         `json:"passengers_per_day"`
	CreatedAt        string      `json:"created_at"`
	UpdatedAt        string      `json:"updated_at"`
	DeletedAt        *string     `json:"deleted_at"`
	LastFound        string      `json:"last_found"`
	FlightsPerWeek   *int        `json:"flights_per_week"`
	FlightsPerDay    *string     `json:"flights_per_day"`
	Airline          AirlineInfo `json:"airline"`
}

type AirlineInfo struct {
	ID                      int     `json:"id"`
	Callsign                *string `json:"callsign"`
	ICAO                    *string `json:"ICAO"`
	IATA                    string  `json:"IATA"`
	Name                    *string `json:"name"`
	FsID                    *string `json:"fs_id"`
	Shortname               *string `json:"shortname"`
	Fullname                *string `json:"fullname"`
	Country                 *string `json:"country"`
	FlagCarrier             *int    `json:"flag_carrier"`
	LanguagesPrimary        *string `json:"languages_primary"`
	LanguagesSecondary      *string `json:"languages_secondary"`
	FlightsLast24Hours      *int    `json:"flights_last_24_hours"`
	Airbourne               *int    `json:"airbourne"`
	Location                *string `json:"location"`
	Phone                   *string `json:"phone"`
	Url                     *string `json:"url"`
	WikiUrl                 *string `json:"wiki_url"`
	IsScheduledPassenger    *int    `json:"is_scheduled_passenger"`
	IsNonscheduledPassenger *int    `json:"is_nonscheduled_passenger"`
	IsCargo                 *int    `json:"is_cargo"`
	IsRailway               *int    `json:"is_railway"`
	IsLowcost               *int    `json:"is_lowcost"`
	Active                  *int    `json:"active"`
	IsOneworld              *int    `json:"is_oneworld"`
	IsStaralliance          *int    `json:"is_staralliance"`
	IsSkyteam               *int    `json:"is_skyteam"`
	IsAllianceaffiliate     *int    `json:"is_allianceaffiliate"`
	RatingIosapp            *string `json:"rating_iosapp"`
	RatingAndroidapp        *string `json:"rating_androidapp"`
	RatingSkytraxReviews    *string `json:"rating_skytrax_reviews"`
	RatingSkytraxStars      *string `json:"rating_skytrax_stars"`
	RatingTripadvisor       *string `json:"rating_tripadvisor"`
	RatingTrustpilot        *string `json:"rating_trustpilot"`
	RatingFlightradar24     *string `json:"rating_flightradar24"`
	Color1                  *string `json:"color1"`
	Color2                  *string `json:"color2"`
	Textcolor               *string `json:"textcolor"`
	KayakApiLastseen        *string `json:"kayak_api_lastseen"`
	KayakApiIsactive        *int    `json:"kayak_api_isactive"`
	CreatedAt               string  `json:"created_at"`
	UpdatedAt               string  `json:"updated_at"`
}

type Country struct {
	Country     string `json:"country"`
	CountryCode string `json:"country_code"`
}

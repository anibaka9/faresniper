package travelpayouts

import "time"

type RoutePrice struct {
	ShowToAffiliates bool      `json:"show_to_affiliates"`
	TripClass        int       `json:"trip_class"`
	Origin           string    `json:"origin"`
	Destination      string    `json:"destination"`
	DepartDate       string    `json:"depart_date"`
	ReturnDate       string    `json:"return_date"`
	NumberOfChanges  int       `json:"number_of_changes"`
	Value            int       `json:"value"`
	FoundAt          time.Time `json:"found_at"`
	Distance         int       `json:"distance"`
	Actual           bool      `json:"actual"`
}

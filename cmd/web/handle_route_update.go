package main

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/anibaka9/faresniper/internal/database"
	"github.com/go-chi/chi/v5"
)

func (s Server) handleRouteUpdate(w http.ResponseWriter, r *http.Request) {
	routeId, err := strconv.Atoi(chi.URLParam(r, "route_id"))
	if err != nil {
		http.Error(w, "wrong route id", http.StatusBadRequest)
		return
	}

	flightsFromIdString := r.FormValue("flightsfrom_id")
	airlineIdString := r.FormValue("airline_id")
	airportFromIdString := r.FormValue("airport_from_id")
	airportToIdString := r.FormValue("airport_to_id")
	isActive := r.FormValue("is_active") == "on"

	if flightsFromIdString == "" {
		http.Error(w, "flightsFrom id should not be empty", http.StatusBadRequest)
	}

	flightsFromId, err := strconv.Atoi(flightsFromIdString)
	if err != nil {
		http.Error(w, "flightsFromId should be number", http.StatusBadRequest)
	}

	if airlineIdString == "" {
		http.Error(w, "airline id should not be empty", http.StatusBadRequest)
	}
	airlineId, err := strconv.Atoi(airlineIdString)
	if err != nil {
		http.Error(w, "airline id should be number", http.StatusBadRequest)
	}

	if airportFromIdString == "" {
		http.Error(w, "airport from id should not be empty", http.StatusBadRequest)
	}

	airportFromId, err := strconv.Atoi(airportFromIdString)
	if err != nil {
		http.Error(w, "airport from id should be number", http.StatusBadRequest)
	}

	if airportToIdString == "" {
		http.Error(w, "airport to id should not be empty", http.StatusBadRequest)
	}

	airportToId, err := strconv.Atoi(airportToIdString)
	if err != nil {
		http.Error(w, "airport to id should be number", http.StatusBadRequest)
	}

	_, err = s.db.Queries.UpdateRoute(r.Context(), database.UpdateRouteParams{
		FlightsfromID: int64(flightsFromId),
		AirlineID:     int64(airlineId),
		AirportFromID: int64(airportFromId),
		AirportToID:   int64(airportToId),
		IsActive:      isActive,
		ID:            int64(routeId),
	})
	if errors.As(err, &sqliteErr) {
		http.Error(w, sqliteErr.ExtendedCode.Error(), http.StatusInternalServerError)
		return
		// switch sqliteErr.ExtendedCode {
		// 	case sql.ErrNoRows
		// }
	}
	http.Redirect(w, r, "/routes/"+strconv.Itoa(routeId), http.StatusSeeOther)
}

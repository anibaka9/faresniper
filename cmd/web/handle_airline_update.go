package main

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/anibaka9/faresniper/internal/database"
	"github.com/go-chi/chi/v5"
)

func (s Server) handleAirlineUpdate(w http.ResponseWriter, r *http.Request) {
	airlineId, err := strconv.Atoi(chi.URLParam(r, "airline_id"))
	if err != nil {
		http.Error(w, "wrong airline id", http.StatusBadRequest)
		return
	}
	name := r.FormValue("name")
	iata := r.FormValue("iata")
	flightsFromIdString := r.FormValue("flightsfrom_id")
	isActive := r.FormValue("is_active") == "on"

	if name == "" {
		http.Error(w, "name should not be empty", http.StatusBadRequest)
	}
	if iata == "" {
		http.Error(w, "iata should not be empty", http.StatusBadRequest)
	}

	if flightsFromIdString == "" {
		http.Error(w, "flightsFromId should not be empty", http.StatusBadRequest)
	}

	flightsFromId, err := strconv.Atoi(flightsFromIdString)
	if err != nil {
		http.Error(w, "flightsFromId should be number", http.StatusBadRequest)
	}

	_, err = s.db.Queries.UpdateAirline(r.Context(), database.UpdateAirlineParams{
		ID:            int64(airlineId),
		Name:          name,
		FlightsfromID: int64(flightsFromId),
		Iata:          iata,
		IsActive:      isActive,
	})
	if errors.As(err, &sqliteErr) {
		http.Error(w, sqliteErr.ExtendedCode.Error(), http.StatusInternalServerError)
		return
		// switch sqliteErr.ExtendedCode {
		// 	case sql.ErrNoRows
		// }
	}
	http.Redirect(w, r, "/airlines/"+strconv.Itoa(airlineId), http.StatusSeeOther)
}

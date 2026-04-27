package main

import (
	"errors"
	"net/http"

	"github.com/anibaka9/faresniper/internal/database"
	"github.com/go-chi/chi/v5"
)

func (s Server) handleAirportUpdate(w http.ResponseWriter, r *http.Request) {
	iata := chi.URLParam(r, "airport_iata")
	if iata == "" {
		http.Error(w, "wrong airport id", http.StatusBadRequest)
		return
	}

	name := r.FormValue("name")
	iataType := r.FormValue("iata_type")
	flightable := r.FormValue("flightable") == "on"
	cityIata := r.FormValue("city_iata")

	if name == "" {
		http.Error(w, "name should not be empty", http.StatusBadRequest)
		return
	}
	if cityIata == "" {
		http.Error(w, "city should not be empty", http.StatusBadRequest)
		return
	}

	_, err := s.db.Queries.UpdateAirport(r.Context(), database.UpdateAirportParams{
		Iata:       iata,
		Name:       name,
		IataType:   iataType,
		Flightable: flightable,
		CityIata:   cityIata,
	})
	if errors.As(err, &sqliteErr) {
		serverError(w, "update airport "+iata, sqliteErr.ExtendedCode)
		return
	}
	http.Redirect(w, r, "/airports/"+iata, http.StatusSeeOther)
}

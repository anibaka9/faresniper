package main

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/anibaka9/faresniper/internal/database"
	"github.com/go-chi/chi/v5"
)

func (s Server) handleAirportUpdate(w http.ResponseWriter, r *http.Request) {
	countryId, err := strconv.Atoi(chi.URLParam(r, "airport_id"))
	if err != nil {
		http.Error(w, "wrong airport id", http.StatusBadRequest)
		return
	}
	name := r.FormValue("name")
	iata := r.FormValue("iata")
	city_id_string := r.FormValue("city_id")

	if name == "" {
		http.Error(w, "name should not be empty", http.StatusBadRequest)
	}
	if iata == "" {
		http.Error(w, "iata should not be empty", http.StatusBadRequest)
	}

	if city_id_string == "" {
		http.Error(w, "city id should not be empty", http.StatusBadRequest)
	}
	city_id, err := strconv.Atoi(city_id_string)
	if err != nil {
		http.Error(w, "city id should be number", http.StatusBadRequest)
	}

	_, err = s.db.Queries.UpdateAirport(r.Context(), database.UpdateAirportParams{
		Name:   name,
		ID:     int64(countryId),
		Iata:   iata,
		CityID: int64(city_id),
	})
	if errors.As(err, &sqliteErr) {
		http.Error(w, sqliteErr.ExtendedCode.Error(), http.StatusInternalServerError)
		return
		// switch sqliteErr.ExtendedCode {
		// 	case sql.ErrNoRows
		// }
	}
	http.Redirect(w, r, "/airports/"+strconv.Itoa(countryId), http.StatusSeeOther)
}

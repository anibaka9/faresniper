package main

import (
	"errors"
	"net/http"

	"github.com/anibaka9/faresniper/internal/database"
	"github.com/go-chi/chi/v5"
)

func (s Server) handleCityUpdate(w http.ResponseWriter, r *http.Request) {
	iata := chi.URLParam(r, "city_id")
	if iata == "" {
		http.Error(w, "wrong city id", http.StatusBadRequest)
		return
	}

	name := r.FormValue("name")
	countryCode := r.FormValue("country_code")

	if name == "" {
		http.Error(w, "name should not be empty", http.StatusBadRequest)
		return
	}
	if countryCode == "" {
		http.Error(w, "country code should not be empty", http.StatusBadRequest)
		return
	}

	_, err := s.db.Queries.UpdateCity(r.Context(), database.UpdateCityParams{
		Iata:        iata,
		Name:        name,
		CountryCode: countryCode,
	})
	if errors.As(err, &sqliteErr) {
		http.Error(w, sqliteErr.ExtendedCode.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/cities/"+iata, http.StatusSeeOther)
}

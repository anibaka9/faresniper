package main

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/anibaka9/faresniper/internal/database"
	"github.com/go-chi/chi/v5"
)

func (s Server) handleCityUpdate(w http.ResponseWriter, r *http.Request) {
	cityId, err := strconv.Atoi(chi.URLParam(r, "city_id"))
	if err != nil {
		http.Error(w, "wrong city id", http.StatusBadRequest)
		return
	}
	name := r.FormValue("name")
	country_id_string := r.FormValue("country_id")

	if name == "" {
		http.Error(w, "name should not be empty", http.StatusBadRequest)
	}

	if country_id_string == "" {
		http.Error(w, "country id should not be empty", http.StatusBadRequest)
	}
	country_id, err := strconv.Atoi(country_id_string)
	if err != nil {
		http.Error(w, "country id should be number", http.StatusBadRequest)
	}

	_, err = s.db.Queries.UpdateCity(r.Context(), database.UpdateCityParams{
		Name:      name,
		ID:        int64(cityId),
		CountryID: int64(country_id),
	})
	if errors.As(err, &sqliteErr) {
		http.Error(w, sqliteErr.ExtendedCode.Error(), http.StatusInternalServerError)
		return
		// switch sqliteErr.ExtendedCode {
		// 	case sql.ErrNoRows
		// }
	}
	http.Redirect(w, r, "/cities/"+strconv.Itoa(cityId), http.StatusSeeOther)
}

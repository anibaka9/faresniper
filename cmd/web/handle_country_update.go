package main

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/anibaka9/faresniper/internal/database"
	"github.com/go-chi/chi/v5"
	"github.com/mattn/go-sqlite3"
)

var sqliteErr sqlite3.Error

func (s Server) handleCountryUpdate(w http.ResponseWriter, r *http.Request) {
	countryId, err := strconv.Atoi(chi.URLParam(r, "country_id"))
	if err != nil {
		http.Error(w, "wrong counrty id", http.StatusBadRequest)
		return
	}
	name := r.FormValue("name")
	country_code := r.FormValue("country_code")

	if name == "" {
		http.Error(w, "name should not be empty", http.StatusBadRequest)
	}
	if country_code == "" {
		http.Error(w, "country_code should not be empty", http.StatusBadRequest)
	}
	_, err = s.db.Queries.UpdateCountry(r.Context(), database.UpdateCountryParams{
		CountryCode: country_code,
		Name:        name,
		ID:          int64(countryId),
	})
	if errors.As(err, &sqliteErr) {
		http.Error(w, sqliteErr.ExtendedCode.Error(), http.StatusInternalServerError)
		return
		// switch sqliteErr.ExtendedCode {
		// 	case sql.ErrNoRows
		// }
	}
	http.Redirect(w, r, "/countries/"+strconv.Itoa(countryId), http.StatusSeeOther)
}

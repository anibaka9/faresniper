package main

import (
	"errors"
	"net/http"

	"github.com/anibaka9/faresniper/internal/database"
	"github.com/go-chi/chi/v5"
)

func (s Server) handleCountryUpdate(w http.ResponseWriter, r *http.Request) {
	code := chi.URLParam(r, "country_id")
	if code == "" {
		http.Error(w, "wrong country id", http.StatusBadRequest)
		return
	}

	name := r.FormValue("name")
	if name == "" {
		http.Error(w, "name should not be empty", http.StatusBadRequest)
		return
	}

	_, err := s.db.Queries.UpdateCountry(r.Context(), database.UpdateCountryParams{
		Code: code,
		Name: name,
	})
	if errors.As(err, &sqliteErr) {
		http.Error(w, sqliteErr.ExtendedCode.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/countries/"+code, http.StatusSeeOther)
}

package main

import (
	"errors"
	"net/http"

	"github.com/anibaka9/faresniper/internal/database"
	"github.com/go-chi/chi/v5"
)

func (s Server) handleAirlineUpdate(w http.ResponseWriter, r *http.Request) {
	iata := chi.URLParam(r, "airline_id")
	if iata == "" {
		http.Error(w, "wrong airline id", http.StatusBadRequest)
		return
	}

	name := r.FormValue("name")
	isLowcost := r.FormValue("is_lowcost") == "on"

	if name == "" {
		http.Error(w, "name should not be empty", http.StatusBadRequest)
		return
	}

	_, err := s.db.Queries.UpdateAirline(r.Context(), database.UpdateAirlineParams{
		Iata:      iata,
		Name:      name,
		IsLowcost: isLowcost,
	})
	if errors.As(err, &sqliteErr) {
		http.Error(w, sqliteErr.ExtendedCode.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/airlines/"+iata, http.StatusSeeOther)
}

package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/anibaka9/faresniper/internal/database"
	databaseclient "github.com/anibaka9/faresniper/internal/database_client"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type data struct {
	NumCountries int
	Countries    []database.Country
}

func main() {
	c, err := databaseclient.NewClient()
	if err != nil {
		log.Fatal(err)
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles("./cmd/web/html/index.html", "./cmd/web/html/sidebar.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		stats, err := c.Queries.GetStats(r.Context())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = t.Execute(w, stats)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})
	r.Handle("/tailwind/*", http.StripPrefix("/tailwind/", http.FileServer(http.Dir("./cmd/web/tailwind"))))
	http.ListenAndServe(":3000", r)
}

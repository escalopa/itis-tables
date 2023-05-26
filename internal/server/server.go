package server

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/escalopa/itis-tables/internal/application"
)

type Server struct {
	ctx context.Context
	uc  *application.UseCase
}

func New(ctx context.Context, uc *application.UseCase) *Server {
	return &Server{
		ctx: ctx,
		uc:  uc,
	}
}

func (s *Server) Run(port string) error {
	http.HandleFunc("/tables", s.tableHandler)

	// Start the server
	log.Printf("Starting server on port %s...\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

func (s *Server) tableHandler(w http.ResponseWriter, r *http.Request) {
	// Extract the query parameters from the request
	group := r.URL.Query().Get("group")
	dateString := r.URL.Query().Get("date")

	// Parse the date string into a time.Time value
	date, err := time.Parse("02/01/2006", dateString)
	if err != nil {
		http.Error(w, "Invalid date format", http.StatusBadRequest)
		return
	}

	if group == "" {
		http.Error(w, "Missing group parameter", http.StatusBadRequest)
		return
	}

	// Call your function with the provided parameters
	tables, err := s.uc.GetSchedule(r.Context(), group, date)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tables)
}

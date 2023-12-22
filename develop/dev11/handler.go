package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Handler struct {
	mux      *http.ServeMux
	calendar *Calendar
}

func NewHandler(calendar *Calendar) *Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/calendar/add", func(w http.ResponseWriter, r *http.Request) {
		var event Event
		err := json.NewDecoder(r.Body).Decode(&event)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		calendar.NewEvent(event)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)

	})
	mux.HandleFunc("/calendar/show", func(w http.ResponseWriter, r *http.Request) {
		for i, item := range calendar.Events {
			fmt.Fprintln(w, i+1)
			fmt.Fprintf(w, "\tDate : %s\n", item.Date)
			fmt.Fprintf(w, "\tTitle : %s\n", item.Title)
			fmt.Fprintf(w, "\tDescription : %s\n", item.Description)
		}
	})
	return &Handler{
		mux:      mux,
		calendar: calendar,
	}
}

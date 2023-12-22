package main

import (
	"fmt"
	"net/http"
)

func main() {
	calendar := NewCalendar()
	handler := NewHandler(calendar)
	go func() {
		for {
			var input string
			fmt.Scanln(&input)
			if input == "show" {
				fmt.Println(calendar)
			}
		}
	}()
	http.ListenAndServe(":8080", handler.mux)
}

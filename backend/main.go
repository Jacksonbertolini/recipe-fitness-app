package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{"message": "Hello from FitMeals API", "status": "ok"}`)
	})

	fmt.Println("FitMeals API running on :8080")
	http.ListenAndServe(":8080", nil)
}

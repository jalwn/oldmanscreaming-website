package main

import (
	"fmt"
	"net/http"

	"screaming/templates"
)

func main() {

	router := http.NewServeMux()
	fileServer := http.FileServer(http.Dir("./static"))
	router.Handle("GET /static/", http.StripPrefix("/static/", fileServer))
	router.HandleFunc("GET /", defaultHandler)

	http.ListenAndServe(":8081", router)

	fmt.Println("Server listening on port 8081")
}

// GET /
func defaultHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	homeTemplate := templates.Home()
	err := templates.Layout(homeTemplate, "Spooktober party").Render(r.Context(), w)
	if err != nil {
		fmt.Printf("Error when rendering home: %v", err)
	}
}

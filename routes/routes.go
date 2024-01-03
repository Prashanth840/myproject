package routes

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func StartRoute() {

	r := mux.NewRouter()
	BuildHandler()
	r.HandleFunc("/upload", FileHandler.FileUpload).Methods("POST")

	err := http.ListenAndServe(":8000", r)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}
}

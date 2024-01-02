package routes

import (
	"fmt"
	"myproject/controller"
	"net/http"

	"github.com/gorilla/mux"
)

func StartRoute() {

	r := mux.NewRouter()

	r.HandleFunc("/upload", controller.HandleFileUpload).Methods("POST")

	err := http.ListenAndServe(":8000", r)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}
}

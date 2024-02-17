package routes

import (
	"github.com/gorilla/mux"
)

func Routes(r *mux.Router) {

	BuildHandler()
	r.HandleFunc("/upload", FileHandler.FileUpload).Methods("POST")

}

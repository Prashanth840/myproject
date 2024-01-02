package main

import (
	"myproject/controller"
	"myproject/data"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	data.InitRedisDB()
	r := mux.NewRouter()
	r.HandleFunc("/upload", controller.HandleFileUpload).Methods("POST")

	http.Handle("/", r)

	http.ListenAndServe(":8000", nil)
}

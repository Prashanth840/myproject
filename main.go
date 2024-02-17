package main

import (
	"myproject/data"
	"myproject/routes"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	data.InitRedisDB()
	router := mux.NewRouter()

	routes.Routes(router)
	http.ListenAndServe(":8000", router)

}

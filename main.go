package main

import (
	"myproject/data"
	"myproject/routes"
)

func main() {
	data.InitRedisDB()

	routes.StartRoute()
}

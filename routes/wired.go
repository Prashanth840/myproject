package routes

import (
	"myproject/controller"
	"myproject/data"
	"myproject/services"
)

var (
	FileHandler controller.FileHandler
)
var (
	FileService services.FileService
)

var (
	FileRepo data.FileRepo
)

func BuildDb() {
	FileRepo = data.NewFileRepo()
}
func BuildService() {
	FileService = services.NewFileService(FileRepo)
}
func BuildHandler() {
	BuildDb()
	BuildService()

	FileHandler = controller.NewFileHandler(FileService)
}

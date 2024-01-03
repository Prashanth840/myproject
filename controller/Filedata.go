package controller

import (
	"fmt"
	"myproject/services"
	"net/http"
)

type FileHandler interface {
	FileUpload(w http.ResponseWriter, r *http.Request)
}
type fileHandler struct {
	fileService services.FileService
}

func NewFileHandler(fileService services.FileService) *fileHandler {
	return &fileHandler{fileService: fileService}
}

func (h *fileHandler) FileUpload(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	file, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Error retrieving the file", http.StatusBadRequest)
		return
	}

	defer file.Close()

	res, sts := h.fileService.FileHandler(handler.Filename, file)
	if sts != "" {
		http.Error(w, sts, http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, res)

}

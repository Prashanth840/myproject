package controller

import (
	"fmt"
	"io"
	"myproject/services"
	"net/http"
	"os"
)

func HandleFileUpload(w http.ResponseWriter, r *http.Request) {
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

	dst, err := os.Create(handler.Filename)
	if err != nil {
		http.Error(w, "Error creating the file", http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	if _, err := io.Copy(dst, file); err != nil {
		http.Error(w, "Error copying the file", http.StatusInternalServerError)
		return
	}
	fileInfo, err := os.Stat(handler.Filename)
	if err != nil {
		http.Error(w, "Error getting file information", http.StatusInternalServerError)
		return
	}
	if fileInfo.Size() == 0 {
		http.Error(w, "File is empty", http.StatusBadRequest)
		return
	}
	filename, size := services.Filehandler(handler.Filename, fileInfo.Size())
	fmt.Println(filename, size)
	fmt.Fprintf(w, "Uploaded Successfully")

}

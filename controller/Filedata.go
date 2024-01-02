package controller

import (
	"fmt"
	"myproject/services"
	"net/http"
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

	res, sts := services.Filehandler(handler.Filename, file)
	if sts != "" {
		http.Error(w, sts, http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, res)

}

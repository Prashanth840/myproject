package controller

import (
	"encoding/csv"
	"fmt"
	"io"
	"myproject/services"
	"net/http"
	"os"
	"path/filepath"

	"github.com/tealeg/xlsx"
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
	fileExtension := filepath.Ext(handler.Filename)
	fmt.Fprintf(w, "File Extension: %s\n", fileExtension)

	if fileExtension == ".csv" {

		file, err := os.Open(handler.Filename)
		if err != nil {
			http.Error(w, "Error opening CSV file", http.StatusInternalServerError)

			return
		}
		defer file.Close()

		reader := csv.NewReader(file)

		records, err := reader.ReadAll()
		if err != nil {
			http.Error(w, "Error reading CSV records", http.StatusInternalServerError)

			return
		}
		if len(records) == 0 {

			fmt.Fprintf(w, "Uploaded file is empty")
			return
		}
	} else if fileExtension == ".xlsx" {
		xlFile, err := xlsx.OpenFile(handler.Filename)
		if err != nil {
			http.Error(w, "Error opening XLSX file", http.StatusInternalServerError)

			return
		}

		for _, sheet := range xlFile.Sheets {
			fmt.Printf("Sheet Name: %s\n", sheet.Name)

			if len(sheet.Rows) == 0 {
				fmt.Fprintf(w, "Uploaded file is empty")
				return
			}
		}
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

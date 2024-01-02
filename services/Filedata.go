package services

import (
	"encoding/csv"
	"fmt"
	"io"
	"myproject/data"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/tealeg/xlsx"
)

func Filehandler(Filename string, File io.Reader) (string, string) {

	dst, err := os.Create(Filename)
	if err != nil {

		return "", "Error creating the file"
	}
	defer dst.Close()

	if _, err := io.Copy(dst, File); err != nil {

		return "", "Error copying the file"
	}

	fileInfo, err := os.Stat(Filename)
	if err != nil {

		return "", "Error getting file information"
	}

	fileExtension := filepath.Ext(Filename)

	if fileExtension == ".csv" {

		file, err := os.Open(Filename)
		if err != nil {

			return "", "Error opening CSV file"
		}
		defer file.Close()

		reader := csv.NewReader(file)

		records, err := reader.ReadAll()
		if err != nil {

			return "", "Error reading CSV records"
		}
		if len(records) == 0 {

			return "Uploaded file is empty", ""
		}
	} else if fileExtension == ".xlsx" {
		xlFile, err := xlsx.OpenFile(Filename)
		if err != nil {

			return "", "Error opening XLSX file"
		}

		for _, sheet := range xlFile.Sheets {
			fmt.Printf("Sheet Name: %s\n", sheet.Name)

			if len(sheet.Rows) == 0 {

				return "Uploaded file is empty", ""
			}
		}
	} else {
		if fileInfo.Size() == 0 {

			return "File is empty", ""
		}
	}
	size := strconv.Itoa(int(fileInfo.Size()))
	Filename = strings.ReplaceAll(Filename, " ", "")
	data.RedisSetExp(Filename, size, time.Duration(time.Second*120))
	result := data.RedisGet(Filename)
	fmt.Println(result)
	return "Uploaded Successfully", ""
}

package services

import (
	"encoding/csv"
	"fmt"
	"io"
	"mime/multipart"
	"myproject/data"
	resources "myproject/models"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/tealeg/xlsx"
)

type FileService interface {
	FileHandler(Filename string, File multipart.File) (resources.Output, string, string)
}

type fileService struct {
	fileDb data.FileRepo
}

func NewFileService(fileDb data.FileRepo) *fileService {
	return &fileService{fileDb: fileDb}
}

func (s *fileService) FileHandler(Filename string, File multipart.File) (resources.Output, string, string) {
	var res resources.Output
	dst, err := os.Create(Filename)
	if err != nil {

		return res, "", "Error creating the file"
	}
	defer dst.Close()

	if _, err := io.Copy(dst, File); err != nil {

		return res, "", "Error copying the file"
	}

	fileInfo, err := os.Stat(Filename)
	if err != nil {

		return res, "", "Error getting file information"
	}

	fileExtension := filepath.Ext(Filename)

	if fileExtension == ".csv" {

		file, err := os.Open(Filename)
		if err != nil {

			return res, "", "Error opening CSV file"
		}
		defer file.Close()

		reader := csv.NewReader(file)

		records, err := reader.ReadAll()
		if err != nil {

			return res, "", "Error reading CSV records"
		}
		if len(records) == 0 {

			return res, "Uploaded file is empty", ""
		}
	} else if fileExtension == ".xlsx" {
		xlFile, err := xlsx.OpenFile(Filename)
		if err != nil {

			return res, "", "Error opening XLSX file"
		}

		for _, sheet := range xlFile.Sheets {
			fmt.Printf("Sheet Name: %s\n", sheet.Name)

			if len(sheet.Rows) == 0 {

				return res, "Uploaded file is empty", ""
			}
		}
	} else {
		if fileInfo.Size() == 0 {

			return res, "File is empty", ""
		}
	}
	size := strconv.Itoa(int(fileInfo.Size()))
	Filename = strings.ReplaceAll(Filename, " ", "")
	s.fileDb.RedisSetExp(Filename, size, time.Duration(time.Second*120))
	result := s.fileDb.RedisGet(Filename)
	res.Filename = Filename
	res.Filesize = result
	return res, "", ""
}

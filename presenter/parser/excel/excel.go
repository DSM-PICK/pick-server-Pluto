package excel

import (
	"pluto/usecase/dto"
	"github.com/tealeg/xlsx/v3"
	"mime/multipart"
)

type parser func(row *xlsx.Row) dto.Element

func parseDefaultForm(formData *multipart.FileHeader, f parser, startIndex int) []dto.Element {
	request := *new([]dto.Element)

	workbook := getWorkbook(formData)
	for _, sheet := range workbook.Sheets {
		for index := startIndex ;; index++ {
			row, _ := sheet.Row(index)
			if row.GetCell(0).Value == "" { break }
			request = append(request, f(row))
		}
	}

	return request
}

func getWorkbook(formData *multipart.FileHeader) *xlsx.File {
	file, _ := formData.Open()
	defer file.Close()

	workbook, _ := xlsx.OpenReaderAt(file, formData.Size)
	return workbook
}



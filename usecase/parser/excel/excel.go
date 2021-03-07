package excel

import (
	"github.com/tealeg/xlsx/v3"
	"pluto/usecase/dto"
)

type parser func(row *xlsx.Row) dto.Element

func parseDefaultForm(file *xlsx.File, f parser, startIndex int) []dto.Element {
	request := *new([]dto.Element)

	for _, sheet := range file.Sheets {
		for index := startIndex ;; index++ {
			row, _ := sheet.Row(index)
			if row.GetCell(0).Value == "" { break }
			request = append(request, f(row))
		}
	}

	return request
}

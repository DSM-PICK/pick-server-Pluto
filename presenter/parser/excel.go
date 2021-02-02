package parser

import (
	"excel-import/usecase/dto"
	"github.com/tealeg/xlsx/v3"
	"mime/multipart"
	"strings"
)

func Activity(file multipart.File, fileSize int64) dto.SetActivityRequest {
	activityRequest := new(dto.SetActivityRequest)

	workbook, _ := xlsx.OpenReaderAt(file, fileSize)
	for _, sheet := range workbook.Sheets {
		for index := 1 ;; index++ {
			row, _ := sheet.Row(index)
			date := row.GetCell(0).Value
			if date == "" { break }
			if !strings.Contains(date, "-") {
				parsed, _ := row.GetCell(0).GetTime(false)
				date = parsed.Format("2006-01-02")
			}


			activityRequest.Activities = append(activityRequest.Activities, dto.Activity{
				Date: date,
				Schedule: row.GetCell(1).Value,
				SecondFloorTeacherName: row.GetCell(2).Value,
				ThirdFloorTeacherName: row.GetCell(3).Value,
				FourthFloorTeacherName: row.GetCell(4).Value,
			})
		}
	}
	return *activityRequest
}

package excel

import (
	"github.com/tealeg/xlsx/v3"
	"pluto/usecase/dto"
)

func ParseStudents(file *xlsx.File) dto.SetStudentRequest {
	request := new(dto.SetStudentRequest)

	for _, element := range parseDefaultForm(file, student, 3) {
		request.Students = append(request.Students, element.(dto.Student))
	}

	return *request
}

func student(row *xlsx.Row) dto.Element {
	return dto.Student{
		Num:  row.GetCell(0).Value,
		Name: row.GetCell(1).Value,
	}
}

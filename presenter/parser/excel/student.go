package excel

import (
	"excel-import/usecase/dto"
	"github.com/tealeg/xlsx/v3"
	"mime/multipart"
)

func ParseStudents(formData *multipart.FileHeader) dto.SetStudentRequest {
	request := new(dto.SetStudentRequest)

	for _, element := range parseDefaultForm(formData, student, 3) {
		request.Students = append(request.Students, element.(dto.Student))
	}

	return *request
}

func student(row *xlsx.Row) dto.Element {
	return dto.Student{
		Num: row.GetCell(0).Value,
		Name: row.GetCell(1).Value,
	}
}

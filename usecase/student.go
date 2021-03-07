package usecase

import (
	"pluto/entity"
	"pluto/usecase/dto"
	"pluto/usecase/parser"
)

var StudentRepository entity.StudentRepository

func SetStudents(request dto.SetStudentRequest) error {
	if len(request.Students) == 0 { return parser.InvalidRequest }

	for _, student := range request.Students {
		setStudent(student)
	}

	return nil
}

func setStudent(student dto.Student) {
	if !entity.NumCheck(student.Num) { return }

	StudentRepository.CreateStudent(entity.Student{
		Num: student.Num,
		Name: student.Name,
		ClassName: string(student.Num[0]) + "학년 " + string(student.Num[1]) + "반",
		ClubName: "창조실",
	})
}

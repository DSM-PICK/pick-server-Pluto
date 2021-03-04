package usecase

import (
	"excel-import/entity"
	"excel-import/infrastructure/database"
	"excel-import/usecase/dto"
)

var studentRepository = database.StudentRepository

func SetStudent(request dto.SetStudentRequest) {
	for _, student := range request.Students {
		setStudent(student)
	}
}

func setStudent(student dto.Student) {
	if !entity.NumCheck(student.Num) { return }

	studentRepository.CreateStudent(entity.Student{
		Num: student.Num,
		Name: student.Name,
		ClassName: string(student.Num[0]) + "학년 " + string(student.Num[1]) + "반",
		ClubName: "창조실",
	})
}

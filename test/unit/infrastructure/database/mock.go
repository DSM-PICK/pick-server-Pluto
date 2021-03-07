package database

import (
	"github.com/stretchr/testify/mock"
	"pluto/entity"
	"pluto/usecase/dto"
	"time"
)

var MockConnection = new(TestConnection)

var TestActivitiesRequest = dto.SetActivityRequest{
	Activities: []dto.Activity{
		{
			Date: "2021-03-01",
			Weekday: "Monday",
			SecondFloorTeacherName: "정우영",
			ThirdFloorTeacherName: "정우영",
			FourthFloorTeacherName: "정우영",
		},
	},
}

var TestStudentsRequest = dto.SetStudentRequest{
	Students: []dto.Student{
		{
			Num: "1418",
			Name: "정우영",
		},
	},
}

var TestTeacherEntity = &entity.Teacher{
	Id: "test",
	Pw: "test",
	Name: "정우영",
	Office: "316호",
}

func (testConnection *TestConnection) TestActivityEntity() entity.Activity {
	activity := TestActivitiesRequest.Activities[0]

	parsed, _ := time.Parse("2006-01-02", activity.Date)
	weekday, _ := entity.WeekdayCheck(activity.Weekday)


	return entity.Activity{
		Date: parsed,
		Schedule: weekday,
		SecondFloorTeacher: *testConnection.FindTeacherByName(activity.SecondFloorTeacherName),
		ThirdFloorTeacher: *testConnection.FindTeacherByName(activity.ThirdFloorTeacherName),
		ForthFloorTeacher: *testConnection.FindTeacherByName(activity.FourthFloorTeacherName),
	}
}

func (testConnection *TestConnection) TestStudentEntity() entity.Student {
	student := TestStudentsRequest.Students[0]

	return entity.Student{
		Num:       student.Num,
		Name:      student.Name,
		ClassName: string(student.Num[0]) + "학년 " + string(student.Num[1]) + "반",
		ClubName: "창조실",
	}
}

type TestConnection struct {
	mock.Mock
}

func (testConnection *TestConnection) CreateActivity(activity entity.Activity) entity.Activity {
	testConnection.On("CreateActivity", testConnection.TestActivityEntity()).Return(activity)
	return testConnection.Called(activity).Get(0).(entity.Activity)
}

func (testConnection *TestConnection) CreateStudent(student entity.Student) entity.Student {
	testConnection.On("CreateStudent", testConnection.TestStudentEntity()).Return(student)
	return testConnection.Called(student).Get(0).(entity.Student)
}

func (testConnection *TestConnection) FindTeacherByName(name string) *entity.Teacher {
	testConnection.On("FindTeacherByName", "정우영").Return(TestTeacherEntity)
	return testConnection.Called(name).Get(0).(*entity.Teacher)
}


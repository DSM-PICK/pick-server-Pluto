package usecase

import (
	"pluto/entity"
	"pluto/usecase/dto"
	"pluto/usecase/parser"
	"time"
)

var ActivityRepository entity.ActivityRepository
var TeacherRepository entity.TeacherRepository

func SetActivities(request dto.SetActivityRequest) error {
	if request.Activities == nil {
		return parser.InvalidRequest
	}

	for _, activity := range request.Activities {
		setActivity(activity)
	}

	return nil
}

func setActivity(activity dto.Activity) {
	date, e := time.Parse("2006-01-02", activity.Date)
	if e != nil {
		return
	}
	schedule, e := entity.WeekdayCheck(activity.Weekday)
	if e != nil {
		return
	}
	secondFloorTeacher := TeacherRepository.FindTeacherByName(activity.SecondFloorTeacherName)
	thirdFloorTeacher := TeacherRepository.FindTeacherByName(activity.ThirdFloorTeacherName)
	fourthFloorTeacher := TeacherRepository.FindTeacherByName(activity.FourthFloorTeacherName)
	if secondFloorTeacher == nil || thirdFloorTeacher == nil || fourthFloorTeacher == nil {
		return
	}

	ActivityRepository.CreateActivity(entity.Activity{
		Date:               date,
		Schedule:           schedule,
		SecondFloorTeacher: *secondFloorTeacher,
		ThirdFloorTeacher:  *thirdFloorTeacher,
		ForthFloorTeacher:  *fourthFloorTeacher,
	})
}

package usecase

import (
	"excel-import/domain"
	"excel-import/infrastructure/database"
	"excel-import/usecase/dto"
	"time"
)

var activityRepository domain.ActivityRepository = database.Initialize()
var teacherRepository domain.TeacherRepository = database.Initialize()

func SetActivities(request dto.SetActivityRequest) {
	for _, activity := range request.Activities {
		setActivity(activity)
	}
}

func setActivity(activity dto.Activity)  {
	date, e := time.Parse("2006-01-02", activity.Date)
	if e != nil { return }
	schedule, e := domain.ScheduleCheck(activity.Schedule)
	if e != nil { return }
	secondFloorTeacher := teacherRepository.FindTeacherByName(activity.SecondFloorTeacherName)
	thirdFloorTeacher := teacherRepository.FindTeacherByName(activity.ThirdFloorTeacherName)
	fourthFloorTeacher := teacherRepository.FindTeacherByName(activity.FourthFloorTeacherName)
	if secondFloorTeacher == nil || thirdFloorTeacher == nil || fourthFloorTeacher == nil {
		return
	}

	activityRepository.CreateActivity(domain.Activity{
		Date: date,
		Schedule: schedule,
		SecondFloorTeacher: *secondFloorTeacher,
		ThirdFloorTeacher: *thirdFloorTeacher,
		ForthFloorTeacher: *fourthFloorTeacher,
	})
}

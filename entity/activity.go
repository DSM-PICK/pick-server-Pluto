package entity

import (
	"errors"
	"time"
)

type Activity struct {
	Date time.Time `gorm:"PrimaryKey"`
	Schedule string `gorm:"size:28"`

	SecondFloorTeacher Teacher `gorm:"ForeignKey:SecondFloorTeacherId"`
	SecondFloorTeacherId string `gorm:"size:16"`
	ThirdFloorTeacher Teacher `gorm:"ForeignKey:ThirdFloorTeacherId"`
	ThirdFloorTeacherId string `gorm:"size:16"`
	ForthFloorTeacher Teacher `gorm:"ForeignKey:ForthFloorTeacherId"`
	ForthFloorTeacherId string `gorm:"size:16"`
}

func ScheduleCheck(scheduleType string) (string, error) {
	switch scheduleType {
	case "동아리":
		return "club", nil
	case "자습":
		return "self-study", nil
	case "방과후":
		return "after-school", nil
	}
	return "", errors.New("invalid schedule")
}

func WeekdayCheck(weekday string) (string, error) {
	switch weekday {
	case "Monday":
		return "after-school", nil
	case "Tuesday":
		return "after-school", nil
	case "Wednesday":
		return "club", nil
	case "Thursday":
		return "self-study", nil
	case "Friday":
		return "club", nil
	}
	return "", errors.New("invalid schedule")
}

type ActivityRepository interface {
	CreateActivity(activity Activity) Activity
}

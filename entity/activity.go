package entity

import (
	"fmt"
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

func WeekdayCheck(weekday string) (string, error) {
	switch weekday {
	case "Monday":
		return "after-school", nil
	case "Tuesday":
		return "club", nil
	case "Wednesday":
		return "club", nil
	case "Thursday":
		return "self-study", nil
	case "Friday":
		return "club", nil
	}
	return "", fmt.Errorf("invalid schedule")
}

type ActivityRepository interface {
	CreateActivity(activity Activity) Activity
}

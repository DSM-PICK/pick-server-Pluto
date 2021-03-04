package entity

import (
	"strconv"
)

type Student struct {
	Num string `gorm:"size:4;PrimaryKey"`
	Name string `gorm:"size:12"`
	ClubName string `gorm:"size:20"`
	ClassName string `gorm:"size:20"`
	IsMondaySelfStudy bool `gorm:"default:false"`
	IsTuesdaySelfStudy bool `gorm:"default:false"`
}

func NumCheck(num string) bool {
	_, e := strconv.Atoi(num)
	return len(num) == 4 && e == nil
}

type StudentRepository interface {
	CreateStudent(student Student) Student
}

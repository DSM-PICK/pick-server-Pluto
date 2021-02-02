package domain

type Teacher struct {
	Id string `gorm:"size:16;PrimaryKey"`
	Pw string `gorm:"size:128"`
	Name string `gorm:"size:12"`
	Office string `gorm:"size:40"`
}

type TeacherRepository interface {
	FindTeacherByName(name string) *Teacher
}

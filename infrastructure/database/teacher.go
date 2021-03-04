package database

import (
	"excel-import/entity"
)

var TeacherRepository entity.TeacherRepository = Initialize()

func (db *Database) FindTeacherByName(name string) *entity.Teacher {
	teacher := new(entity.Teacher)
	db.connection.First(teacher, "name = ?", name)
	return teacher
}

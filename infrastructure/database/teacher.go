package database

import (
	"pluto/entity"
)

func (db *Database) FindTeacherByName(name string) *entity.Teacher {
	teacher := new(entity.Teacher)
	db.Connection.First(teacher, "name = ?", name)
	return teacher
}

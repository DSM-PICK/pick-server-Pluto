package database

import (
	"excel-import/domain"
)

func (db *Database) FindTeacherByName(name string) *domain.Teacher {
	teacher := new(domain.Teacher)
	db.connection.First(teacher, "name = ?", name)
	return teacher
}

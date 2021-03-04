package database

import "excel-import/entity"

var StudentRepository entity.StudentRepository = Initialize()

func (db *Database) CreateStudent(student entity.Student) entity.Student {
	db.connection.FirstOrCreate(&student)
	return student
}

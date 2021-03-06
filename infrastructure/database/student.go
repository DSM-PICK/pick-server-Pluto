package database

import (
	"gorm.io/gorm/clause"
	"pluto/entity"
)

func (db *Database) CreateStudent(student entity.Student) entity.Student {
	db.Connection.Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: "num"}},
		DoUpdates: clause.AssignmentColumns([]string{"name",
													 "club_name",
													 "class_name",
													 "is_monday_self_study",
													 "is_tuesday_self_study"}),
	}).Create(&student)
	return student
}

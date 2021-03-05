package database

import (
	"gorm.io/gorm/clause"
	"pluto/entity"
)

var ActivityRepository entity.ActivityRepository = Initialize()

func (db *Database) CreateActivity(activity entity.Activity) entity.Activity {
	db.connection.Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: "date"}},
		DoUpdates: clause.AssignmentColumns([]string{"schedule",
											 		 "second_floor_teacher_id",
											 		 "third_floor_teacher_id",
													 "forth_floor_teacher_id"}),
	}).Create(&activity)
	return activity
}

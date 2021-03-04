package database

import "excel-import/entity"

var ActivityRepository entity.ActivityRepository = Initialize()

func (db *Database) CreateActivity(activity entity.Activity) entity.Activity {
	db.connection.FirstOrCreate(&activity)
	return activity
}

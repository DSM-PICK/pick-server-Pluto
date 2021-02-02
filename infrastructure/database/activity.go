package database

import "excel-import/domain"

func (db *Database) CreateActivity(activity domain.Activity) domain.Activity {
	db.connection.FirstOrCreate(&activity)
	return activity
}

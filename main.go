package main

import (
	"pluto/infrastructure/database"
	"pluto/presenter/web"
	"pluto/usecase"
)

func main() {
	usecase.ActivityRepository = database.DefaultInitialize()
	usecase.StudentRepository = database.DefaultInitialize()
	usecase.TeacherRepository = database.DefaultInitialize()

	web.Main()
}

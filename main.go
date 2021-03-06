package main

import (
	"pluto/infrastructure/database"
	"pluto/presenter/web"
	"pluto/usecase"
)

func main() {
	usecase.ActivityRepository = database.InitializeMysql()
	usecase.StudentRepository = database.InitializeMysql()
	usecase.TeacherRepository = database.InitializeMysql()

	web.Main()
}

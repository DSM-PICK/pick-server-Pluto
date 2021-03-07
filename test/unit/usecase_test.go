package unit

import (
	"github.com/stretchr/testify/assert"
	"pluto/test/unit/infrastructure/database"
	"pluto/usecase"
	"testing"
)

func TestSetActivities(t *testing.T) {
	usecase.ActivityRepository = database.MockConnection
	usecase.TeacherRepository = database.MockConnection

	assert.Equal(t, usecase.SetActivities(database.TestActivitiesRequest), nil)
}

func TestSetStudents(t *testing.T) {
	usecase.StudentRepository = database.MockConnection

	assert.Equal(t, usecase.SetStudents(database.TestStudentsRequest), nil)
}

package unit

import (
	"github.com/stretchr/testify/assert"
	"pluto/usecase"
	"testing"
)

func TestSetActivities(t *testing.T) {
	usecase.ActivityRepository = MockConnection
	usecase.TeacherRepository = MockConnection

	assert.Equal(t, usecase.SetActivities(TestActivitiesRequest), nil)
}

func TestSetStudents(t *testing.T) {
	usecase.StudentRepository = MockConnection

	assert.Equal(t, usecase.SetStudents(TestStudentsRequest), nil)
}

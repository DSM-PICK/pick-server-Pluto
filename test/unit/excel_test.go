package unit

import (
	"github.com/stretchr/testify/assert"
	"github.com/tealeg/xlsx/v3"
	"pluto/usecase/parser/excel"
	"testing"
)

func TestParseActivities(t *testing.T) {
	file, _ := xlsx.OpenFile("./activity-test.xlsx")

	assert.Equal(t, excel.ParseActivities(file), TestActivitiesRequest)
}

func TestParseStudents(t *testing.T) {
	file, _ := xlsx.OpenFile("./student-test.xlsx")

	assert.Equal(t, excel.ParseStudents(file), TestStudentsRequest)
}

package unit

import (
	"github.com/stretchr/testify/assert"
	"pluto/entity"
	"testing"
)

func TestWeekdayCheck(t *testing.T) {
	monday, _ := entity.WeekdayCheck("Monday")
	tuesday, _ := entity.WeekdayCheck("Tuesday")
	wednesday, _ := entity.WeekdayCheck("Wednesday")
	thursday, _ := entity.WeekdayCheck("Thursday")
	friday, _ := entity.WeekdayCheck("Friday")

	assert.Equal(t, monday, "after-school")
	assert.Equal(t, tuesday, "club")
	assert.Equal(t, wednesday, "club")
	assert.Equal(t, thursday, "self-study")
	assert.Equal(t, friday, "club")
}

func TestNumCheck(t *testing.T) {
	assert.Equal(t, entity.NumCheck("1418"), true)
	assert.Equal(t, entity.NumCheck("invalid"), false)
	assert.Equal(t, entity.NumCheck("012"), false)
}

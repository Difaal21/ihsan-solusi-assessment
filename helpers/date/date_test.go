package date_test

import (
	"difaal21/ihsan-solusi-assessment/helpers/date"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCurrentUTCTime(t *testing.T) {
	now := date.CurrentUTCTime()
	assert.NotNil(t, now, "Expected non-nil time")

	utcNow := time.Now().UTC()
	assert.Equal(t, utcNow.Year(), now.Year(), "Expected same year")
	assert.Equal(t, utcNow.Month(), now.Month(), "Expected same month")
	assert.Equal(t, utcNow.Day(), now.Day(), "Expected same day")
	assert.Equal(t, utcNow.Hour(), now.Hour(), "Expected same hour")
	assert.Equal(t, utcNow.Minute(), now.Minute(), "Expected same minute")
	assert.Equal(t, utcNow.Second(), now.Second(), "Expected same second")
}

func TestCurrentUTCTimeNotNil(t *testing.T) {
	now := date.CurrentUTCTime()
	assert.NotNil(t, now, "Expected non-nil time")
}

func TestCurrentUTCTimeIsUTC(t *testing.T) {
	now := date.CurrentUTCTime()
	assert.Equal(t, time.UTC, now.Location(), "Expected time to be in UTC")
}

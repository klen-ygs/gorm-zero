package sql

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestGetDbName(t *testing.T) {
	tag := "autoUpdateTime:milli;column:id"
	field, err := getDbField(tag)
	assert.Nil(t, err, "err not nil")
	assert.Equal(t, "id", field)
}

type TestStruct struct {
	F1 string    `gorm:"column:f1"`
	F2 int       `gorm:"column:f2"`
	F3 bool      `gorm:"column:f3"`
	F4 time.Time `gorm:"column:f4"`
}

var d TestStruct

func TestGetDbNameFromStruct(t *testing.T) {
	InitField(&d)

	assert.Equal(t, "f1", Field(&d.F1))
}

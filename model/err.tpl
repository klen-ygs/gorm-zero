package {{.pkg}}

import (
    "gorm.io/gorm"
    "github.com/klen-ygs/gorm-zero/gormc"
)

var ErrNotFound = gorm.ErrRecordNotFound

func field[T any](fieldPointer *T) string {
	return gormc.Field(fieldPointer)
}

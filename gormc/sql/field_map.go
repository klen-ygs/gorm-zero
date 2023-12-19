package sql

import (
	"errors"
	"github.com/zeromicro/go-zero/core/logx"
	"reflect"
	"strings"
	"unsafe"
)

var fieldMap = map[uintptr]string{}

func getDbField(tag string) (string, error) {
	if tag == "" {
		return "", errors.New("gorm tag is empty")
	}

	const key = "column"
	index := strings.Index(tag, key)
	if index == -1 {
		return "", errors.New("no match key " + key)
	}

	dbNameStart := index + len(":") + len(key)

	dbname := strings.Split(tag[dbNameStart:], ";")[0]
	return dbname, nil
}

type Tabler interface {
	TableName() string
}

func InitField(model Tabler) {
	typePointerInfo := reflect.TypeOf(model)
	if typePointerInfo.Kind() != reflect.Pointer {
		logx.Must(errors.New("model must be pointer"))
	}

	valueInfo := reflect.ValueOf(model).Elem()
	typeInfo := typePointerInfo.Elem()
	fieldNum := typeInfo.NumField()

	tableName := model.TableName()

	for i := 0; i < fieldNum; i++ {
		tag := typeInfo.Field(i).Tag.Get("gorm")
		addr := valueInfo.Field(i).UnsafeAddr()
		fieldName, err := getDbField(tag)
		if err != nil {
			logx.Must(err)
		}

		fieldMap[addr] = tableName + "." + fieldName
	}
}

func Field[T any](fieldPointer *T) string {
	addr := unsafe.Pointer(fieldPointer)

	name, ok := fieldMap[uintptr(addr)]
	if !ok {
		panic("model need to InitField to register")
	}

	return name
}

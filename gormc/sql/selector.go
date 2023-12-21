package sql

import "fmt"

func Eq[FieldType any](field *FieldType, query FieldType) (string, FieldType) {
	return Field(field), query
}

func Ne[FieldType any](field *FieldType, query FieldType) (string, FieldType) {
	return Field(field), query
}

func Gt[FieldType any](field *FieldType, query FieldType) (string, FieldType) {
	return Field(field) + " > ?", query
}

func Ge[FieldType any](field *FieldType, query FieldType) (string, FieldType) {
	return Field(field) + " >= ?", query
}

func Lt[FieldType any](field *FieldType, query FieldType) (string, FieldType) {
	return Field(field) + " < ?", query
}

func Le[FieldType any](field *FieldType, query FieldType) (string, FieldType) {
	return Field(field) + " <= ?", query
}

func Between[FieldType any](field *FieldType, start, end FieldType) (string, FieldType, FieldType) {
	return Field(field) + " BETWEEN ? AND ?", start, end
}

func Like[FieldType any](field *FieldType, str string) (string, string) {
	return Field(field) + " LIKE ?", fmt.Sprintf("%%%s%%", str)
}

func LikeLeft[FieldType any](field *FieldType, str string) (string, string) {
	return Field(field) + " LIKE ?", fmt.Sprintf("%%%s", str)
}

func LikeRight[FieldType any](field *FieldType, str string) (string, string) {
	return Field(field) + " LIKE ?", fmt.Sprintf("%s%%", str)
}

func IsNil[FieldType any](field *FieldType) string {
	return Field(field) + " IS NULL"
}

func In[FieldType any](field *FieldType, query []FieldType) (string, []FieldType) {
	return Field(field), query
}

func On[FieldType any](tabler Tabler, left *FieldType, right *FieldType) string {
	return fmt.Sprintf("JOIN %s ON %s = %s", tabler.TableName(), Field(left), Field(right))
}

func LeftOn[FieldType any](tabler Tabler, left *FieldType, right *FieldType) string {
	return fmt.Sprintf("Left JOIN %s ON %s = %s", tabler.TableName(), Field(left), Field(right))
}

func RightOn[FieldType any](tabler Tabler, left *FieldType, right *FieldType) string {
	return fmt.Sprintf("Right JOIN %s ON %s = %s", tabler.TableName(), Field(left), Field(right))
}

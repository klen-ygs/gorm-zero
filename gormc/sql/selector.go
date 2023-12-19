package sql

import "fmt"

func Eq[FieldType any, Selector any](field *FieldType, query Selector) (string, Selector) {
	return Field(field), query
}

func Ne[FieldType any, Selector any](field *FieldType, query Selector) (string, Selector) {
	return Field(field), query
}

func Gt[FieldType any, Selector any](field *FieldType, query Selector) (string, Selector) {
	return Field(field) + " > ?", query
}

func Ge[FieldType any, Selector any](field *FieldType, query Selector) (string, Selector) {
	return Field(field) + " >= ?", query
}

func Lt[FieldType any, Selector any](field *FieldType, query Selector) (string, Selector) {
	return Field(field) + " < ?", query
}

func Le[FieldType any, Selector any](field *FieldType, query Selector) (string, Selector) {
	return Field(field) + " <= ?", query
}

func Between[FieldType any, Selector any](field *FieldType, start, end Selector) (string, Selector, Selector) {
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

func In[FieldType any, Selector any](field *FieldType, query []Selector) (string, []Selector) {
	return Field(field), query
}

func On[LeftField any, RightField any](tabler Tabler, left *LeftField, right *RightField) string {
	return fmt.Sprintf("JOIN %s ON %s = %s", tabler.TableName(), Field(left), Field(right))
}

func LeftOn[LeftField any, RightField any](tabler Tabler, left *LeftField, right *RightField) string {
	return fmt.Sprintf("Left JOIN %s ON %s = %s", tabler.TableName(), Field(left), Field(right))
}

func RightOn[LeftField any, RightField any](tabler Tabler, left *LeftField, right *RightField) string {
	return fmt.Sprintf("Right JOIN %s ON %s = %s", tabler.TableName(), Field(left), Field(right))
}

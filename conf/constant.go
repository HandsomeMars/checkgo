package conf

type commentType string
type fieldType string
type layoutType string

const (
	Form  commentType = "form"
	Table commentType = "table"
)

const (
	Horizontal layoutType = "horizontal"
	Vertical   layoutType = "vertical"
)

const (
	String fieldType = "string"
	Int    fieldType = "int"
	Float  fieldType = "float"
	Array  fieldType = "array"
)

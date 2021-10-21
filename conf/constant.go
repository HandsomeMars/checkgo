package conf

type commentType string

const (
	Form   commentType = "form"
	Table  commentType = "table"
	String commentType = "string"
	Int    commentType = "int"
	Float  commentType = "float"
	Array  commentType = "array"
)

package models

// type alias declaration

type (
	Model = interface {
		TableName() string //回傳表格名稱
	}

	ModelSlice = []Model
)

type (
	WheresMap    = map[string]interface{}
	UpdatesMap   = map[string]interface{}
	CreatesMap   = map[string]interface{}
	SelectsSlice = []string
	OrderBySlice = []string  // 依序排序, 例如 {"column1 desc", "column2"}
)

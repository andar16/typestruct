package typestruct

// Model Operator
type Operator struct {
	Id       int    `orm:"column(op_id);pk"`
	OpName   string `orm:"size(100)"`
	OpCreate string
	OpUpdate string
	OpType   string
}

type OperatorResponse struct {
	Status  int      `json:"status"`
	Message string   `json:"message"`
	Data    Operator `json:"data"`
}

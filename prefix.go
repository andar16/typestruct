package typestruct

// Struct OperatorPrefix
type OperatorPrefix struct {
	Id       int    `orm:"column(op_id)"`
	PrPrefix string `orm:"size(5)"`
	PrCreate string
	PrUpdate string
}

type OperatorPrefixResponse struct {
	Status  int            `json:"status"`
	Message string         `json:"message"`
	Data    OperatorPrefix `json:"data"`
}

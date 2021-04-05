package typestruct

type Product struct {
	Id        int64 `orm:"column(prd_id);pk"`
	PrdCode   string
	PrdName   string
	PrdValue  int64
	PrdNote   string
	PrdStatus bool
	PrdCreate string
	PrdUpdate string
	PgId      int64
}

type ProductResponse struct {
	Status  int     `json:"status"`
	Message string  `json:"message"`
	Data    Product `json:"data"`
}

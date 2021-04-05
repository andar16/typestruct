package typestruct

type ProductGroup struct {
	Id       int64 `orm:"column(pg_id);pk"`
	PgName   string
	PgStatus bool
	PgCreate string
	PgUpdate string
	OpId     int
}

type ProductGroupResponse struct {
	Status  int          `json:"status"`
	Message string       `json:"message"`
	Data    ProductGroup `json:"data"`
}

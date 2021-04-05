package typestruct

import "github.com/beego/beego/v2/client/orm"

type GeneralResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Count   int64  `json:"count"`
}

type ArrayResponse struct {
	Status  int          `json:"status"`
	Message string       `json:"message"`
	Count   int64        `json:"count"`
	Data    []orm.Params `json:"data"`
}

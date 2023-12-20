package model

type Filter struct {
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
}

type StrFilter struct {
	Filter string `json:"filter"`
}

type Pagination struct {
	Length int `json:"length" form:"length"`
	Start  int `json:"start" form:"start"`
}

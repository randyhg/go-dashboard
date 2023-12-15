package model

type Filter struct {
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
}

type StrFilter struct {
	Filter string `json:"filter"`
}

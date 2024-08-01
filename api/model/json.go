package model

type Filter struct {
	Filter   map[string]interface{} `json:"filter"`
	KeyWord  string                 `json:"keyword"`
	PriceMin float64                `json:"price_min"`
	PriceMax float64                `json:"price_max"`
	Itmid    string                 `json:"item_id"`
	Wid      string                 `json:"warehouse_id"`
}

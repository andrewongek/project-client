package structs

type ItemData struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type OrderData struct {
	Item     ItemData `json:"item"`
	Quantity int32    `json:"quantity"`
}

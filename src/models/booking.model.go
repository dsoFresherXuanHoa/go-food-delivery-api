package models

import "time"

type BookingResponse struct {
	Table       TableResponse         `json:"table"`
	OrderID     uint                  `json:"orderId"`
	CreatedTime time.Time             `json:"CreatedTime"`
	Items       []BookingItemResponse `json:"items"`
	Note        string                `json:"note"`
	Status      bool                  `json:"status"`
	Accepted    bool                  `json:"accepted"`
	Compensate  bool                  `json:"compensate"`
}

type BookingCreatable struct {
	Note       *string `json:"note"`
	TableId    *uint   `json:"tableId"`
	SecretCode *int    `json:"secretCode"`
	Quantities []int   `json:"quantities"`
	ProductsId []uint  `json:"productsId"`
}

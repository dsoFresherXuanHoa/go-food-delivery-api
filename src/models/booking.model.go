package models

import "time"

type BookingResponse struct {
	Table            TableResponse         `json:"table"`
	OrderID          uint                  `json:"orderId"`
	CreatedTime      time.Time             `json:"createdTime"`
	AcceptedTime     *time.Time            `json:"acceptedTime"`
	Items            []BookingItemResponse `json:"items"`
	Note             string                `json:"note"`
	Status           bool                  `json:"status"`
	Accepted         bool                  `json:"accepted"`
	Compensate       bool                  `json:"compensate"`
	IncludeTableCost bool                  `json:"includeTableCost"`
	Rejected         bool                  `json:"rejected"`
	Reason           string                `json:"reason"`
	TotalPrice       int                   `json:"totalPrice"`
}

type BookingCreatable struct {
	Note             *string `json:"note"`
	TableId          *uint   `json:"tableId"`
	SecretCode       *int    `json:"secretCode"`
	Quantities       []int   `json:"quantities"`
	ProductsId       []uint  `json:"productsId"`
	IncludeTableCost bool    `json:"includeTableCost"`
	Reason           *string `json:"reason"`
}

package models

type BookingResponse struct {
	Employee   EmployeeResponse      `json:"employee"`
	Table      TableResponse         `json:"table"`
	Items      []BookingItemResponse `json:"items"`
	Note       string                `json:"note"`
	Status     bool                  `json:"status"`
	Accepted   bool                  `json:"accepted"`
	Compensate bool                  `json:"compensate"`
}

type BookingCreatable struct {
	Note       *string `json:"note"`
	TableId    *uint   `json:"tableId"`
	Quantities []int   `json:"quantities"`
	ProductsId []uint  `json:"productsId"`
}
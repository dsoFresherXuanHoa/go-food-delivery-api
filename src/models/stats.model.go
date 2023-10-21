package models

type StatsTopEmployeeByOrder struct {
	EmployeeId  uint `json:"employeeId"`
	TotalOrders uint `json:"totalOrders"`
	StartTime   int  `json:"-"`
	EndTime     int  `json:"-"`
}

type StatsTopGoodsByReorderLevel struct {
	Id           uint `json:"productId"`
	ReorderLevel uint `json:"reorderLevel"`
	StartTime    int  `json:"-"`
	EndTime      int  `json:"-"`
}

func (StatsTopEmployeeByOrder) GetTableName() string     { return "orders" }
func (StatsTopGoodsByReorderLevel) GetTableName() string { return "products" }

package models

type StatsTopEmployeeByOrder struct {
	EmployeeId  uint `json:"employeeId"`
	TotalOrders uint `json:"totalOrders"`
	StartTime   int  `json:"startTime"`
	EndTime     int  `json:"endTime"`
}

func (StatsTopEmployeeByOrder) GetTableName() string { return "orders" }

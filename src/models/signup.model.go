package models

type SignUp struct {
	FullName   string `json:"fullName"`
	Tel        string `json:"tel"`
	Gender     bool   `json:"gender"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	Email      string `json:"email"`
	SecretCode int    `json:"-" `
	EmployeeId uint   `json:"employeeId"`
	RoleId     uint   `json:"roleId"`
}

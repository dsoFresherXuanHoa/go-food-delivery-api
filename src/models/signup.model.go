package models

type SignUp struct {
	FullName   string `json:"fullName"`
	Tel        string `json:"tel"`
	Gender     bool   `json:"gender"`
	Password   string `json:"password"`
	Email      string `json:"email"`
	SecretCode int    `json:"-" `
	RoleId     uint   `json:"roleId"`
}

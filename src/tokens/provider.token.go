package tokens

import "time"

type Provider interface {
	Generate(payload TokenPayload, exp int) (*Token, error)
	Validate(token string) (*TokenPayload, error)
}

type Token struct {
	Token     string    `json:"token"`
	CreatedAt time.Time `json:"createdAt"`
	Exp       int       `json:"exp"`
}

type TokenPayload struct {
	AccountId  int `json:"accountId"`
	EmployeeId int `json:"employeeId"`
	RoleId     int `json:"roleId"`
}

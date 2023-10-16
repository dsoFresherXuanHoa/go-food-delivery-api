package models

import (
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	Accounts Accounts `json:"-"`
	Name     string   `json:"name" gorm:"column:name;unique;not null"`
}

type RoleResponse struct {
	gorm.Model
	Accounts Accounts `json:"accounts"`
	Name     string   `json:"name"`
}

type RoleCreatable struct {
	gorm.Model
	Name *string `json:"name" gorm:"column:name;unique;not null"`
}

type RoleUpdatable struct {
	gorm.Model
	Name *string `json:"name" gorm:"column:name;unique;not null"`
}

type Roles []Role

func (Role) GetTableName() string          { return "roles" }
func (RoleCreatable) GetTableName() string { return Role{}.GetTableName() }
func (RoleUpdatable) GetTableName() string { return Role{}.GetTableName() }
func (Roles) GetTableName() string         { return Role{}.GetTableName() }

package services

import (
	"context"
	"fmt"
	"go-food-delivery-api/src/models"
)

type RoleStorage interface {
	CreateRole(ctx context.Context, role *models.RoleCreatable) (*uint, error)
}

type RoleBusiness struct {
	storage RoleStorage
}

func NewRoleBusiness(storage RoleStorage) *RoleBusiness {
	return &RoleBusiness{storage: storage}
}

func (business *RoleBusiness) CreateRole(ctx context.Context, role *models.RoleCreatable) (*uint, error) {
	if roleId, err := business.storage.CreateRole(ctx, role); err != nil {
		fmt.Println("Error while save Role in service: " + err.Error())
		return nil, err
	} else {
		return roleId, nil
	}
}

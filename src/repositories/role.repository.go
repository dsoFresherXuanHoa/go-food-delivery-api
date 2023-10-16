package repositories

import (
	"context"
	"fmt"
	"go-food-delivery-api/src/models"
)

func (s *sqlStorage) CreateRole(ctx context.Context, role *models.RoleCreatable) (*uint, error) {
	if err := s.db.Table(models.RoleCreatable{}.GetTableName()).Create(&role).Error; err != nil {
		fmt.Println("Error while create role in repository: " + err.Error())
		return nil, err
	} else {
		return &role.ID, nil
	}
}

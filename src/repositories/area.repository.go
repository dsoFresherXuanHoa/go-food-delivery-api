package repositories

import (
	"context"
	"fmt"
	"go-food-delivery-api/src/models"
)

func (s *sqlStorage) CreateArea(ctx context.Context, area *models.AreaCreatable) (*uint, error) {
	if err := s.db.Table(models.AreaCreatable{}.GetTableName()).Create(&area).Error; err != nil {
		fmt.Println("Error while create area in repository: " + err.Error())
		return nil, err
	}
	return &area.ID, nil
}

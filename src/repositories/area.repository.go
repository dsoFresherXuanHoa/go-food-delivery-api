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

func (s *sqlStorage) ReadAreaById(ctx context.Context, areaId uint) (*models.Area, error) {
	var area models.Area
	if err := s.db.Unscoped().Where("id = ?", areaId).Find(&area).Error; err != nil {
		fmt.Println("Error while find area by id in repository: " + err.Error())
		return nil, err
	}
	return &area, nil
}

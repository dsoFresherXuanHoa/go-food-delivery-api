package repositories

import (
	"context"
	"fmt"
	"go-food-delivery-api/src/models"
)

func (s *sqlStorage) CreateCategory(ctx context.Context, category *models.CategoryCreatable) (*uint, error) {
	if err := s.db.Table(models.CategoryCreatable{}.GetTableName()).Create(&category).Error; err != nil {
		fmt.Println("Error while create category in repository: " + err.Error())
		return nil, err
	}
	return &category.ID, nil
}

func (s *sqlStorage) ReadCategory(ctx context.Context) (models.Categories, error) {
	var categories models.Categories
	if err := s.db.Table(models.Categories{}.GetTableName()).Find(&categories).Error; err != nil {
		fmt.Println("Error while read category in repository: " + err.Error())
		return nil, err
	}
	return categories, nil
}

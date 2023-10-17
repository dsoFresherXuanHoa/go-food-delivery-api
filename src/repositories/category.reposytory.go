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

func (s *sqlStorage) ReadCategoryById(ctx context.Context, categoryId uint) (*models.Category, error) {
	var category models.Category
	if err := s.db.Where("id = ?", categoryId).Find(&category).Error; err != nil {
		fmt.Println("Error while find category by id in repository: " + err.Error())
		return nil, err
	}
	return &category, nil
}

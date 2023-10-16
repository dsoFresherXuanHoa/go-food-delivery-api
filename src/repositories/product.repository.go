package repositories

import (
	"context"
	"fmt"
	"go-food-delivery-api/src/models"
)

func (s *sqlStorage) CreateProduct(ctx context.Context, product *models.ProductCreatable) (*uint, error) {
	if err := s.db.Table(models.ProductCreatable{}.GetTableName()).Create(&product).Error; err != nil {
		fmt.Println("Error while create product in repository: " + err.Error())
		return nil, err
	}
	return &product.ID, nil
}

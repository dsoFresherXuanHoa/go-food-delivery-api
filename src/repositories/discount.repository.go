package repositories

import (
	"context"
	"fmt"
	"go-food-delivery-api/src/models"
)

func (s *sqlStorage) CreateDiscount(ctx context.Context, discount *models.DiscountCreatable) (*uint, error) {
	if err := s.db.Table(models.DiscountCreatable{}.GetTableName()).Create(&discount).Error; err != nil {
		fmt.Println("Error while create discount in repository: " + err.Error())
		return nil, err
	}
	return &discount.ID, nil
}

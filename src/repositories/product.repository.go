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

func (s *sqlStorage) ReadProduct(ctx context.Context) (models.Products, error) {
	var products models.Products
	if err := s.db.Table(models.Products{}.GetTableName()).Find(&products).Error; err != nil {
		fmt.Println("Error while read product in repository: " + err.Error())
		return nil, err
	}
	return products, nil
}

func (s *sqlStorage) ReadProductByCategoryId(ctx context.Context, categoryId uint) (models.Products, error) {
	var products models.Products
	if err := s.db.Where("category_id = ?", categoryId).Find(&products).Error; err != nil {
		fmt.Println("Error while read product by categoryId in repository: " + err.Error())
		return nil, err
	}
	return products, nil
}

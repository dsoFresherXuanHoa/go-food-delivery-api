package repositories

import (
	"context"
	"fmt"
	"go-food-delivery-api/src/models"
	"go-food-delivery-api/src/services"
)

func (s *sqlStorage) CreateProduct(ctx context.Context, product *models.ProductCreatable) (*uint, error) {
	if err := s.db.Table(models.ProductCreatable{}.GetTableName()).Create(&product).Error; err != nil {
		fmt.Println("Error while create product in repository: " + err.Error())
		return nil, err
	}
	return &product.ID, nil
}

func (s *sqlStorage) ReadProduct(ctx context.Context) ([]models.ProductResponse, error) {
	var products models.Products
	if err := s.db.Table(models.Products{}.GetTableName()).Find(&products).Error; err != nil {
		fmt.Println("Error while read product in repository: " + err.Error())
		return nil, err
	}

	var res = make([]models.ProductResponse, len(products))
	for i, product := range products {
		res[i] = s.GetDetailProduct(ctx, &product)
	}
	return res, nil
}

func (s *sqlStorage) ReadProductByCategoryId(ctx context.Context, categoryId uint) ([]models.ProductResponse, error) {
	var products models.Products
	if err := s.db.Where("category_id = ?", categoryId).Find(&products).Error; err != nil {
		fmt.Println("Error while read product by categoryId in repository: " + err.Error())
		return nil, err
	}

	var res = make([]models.ProductResponse, len(products))
	for i, product := range products {
		res[i] = s.GetDetailProduct(ctx, &product)
	}
	return res, nil
}

func (s *sqlStorage) GetDetailProduct(ctx context.Context, product *models.Product) models.ProductResponse {
	categoryService := services.NewCategoryBusiness(s)
	discountService := services.NewDiscountBusiness(s)
	embedCategory, _ := categoryService.ReadCategoryById(ctx, product.CategoryId)
	embedDiscount, _ := discountService.ReadDiscountById(ctx, product.DiscountId)

	discountPercent := embedDiscount.DiscountPercent
	salePercent := 1 - float64(discountPercent)/100
	product.DiscountPrice = salePercent * product.Price
	return models.ProductResponse{Model: product.Model, Bills: nil, Name: &product.Name, Description: &product.Description, Price: &product.Price, DiscountPrice: &product.DiscountPrice, ReorderLevel: &product.ReorderLevel, Thumb: &product.Thumb, StockAmount: &product.StockAmount, Category: *embedCategory, Discount: *embedDiscount}
}

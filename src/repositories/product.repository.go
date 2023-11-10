package repositories

import (
	"context"
	"fmt"
	"go-food-delivery-api/src/models"
	"go-food-delivery-api/src/services"
	"time"
)

func (s *sqlStorage) GetDetailProduct(ctx context.Context, product models.Product) models.ProductResponse {
	categoryService := services.NewCategoryBusiness(s)
	discountService := services.NewDiscountBusiness(s)
	embedCategory, _ := categoryService.ReadCategoryById(ctx, product.CategoryId)
	embedDiscount, _ := discountService.ReadDiscountById(ctx, product.DiscountId)

	discountPercent := embedDiscount.DiscountPercent
	salePercent := 1 - float64(discountPercent)/100
	product.DiscountPrice = salePercent * product.Price

	return models.ProductResponse{Model: product.Model, Bills: nil, Name: &product.Name, Description: &product.Description, Price: &product.Price, DiscountPrice: &product.DiscountPrice, ReorderLevel: &product.ReorderLevel, Thumb: &product.Thumb, StockAmount: &product.StockAmount, Category: *embedCategory, Discount: *embedDiscount, Uint: &product.Uint, ActualDiscountPrice: nil}
}

func (s *sqlStorage) GetCreatableProduct(ctx context.Context, product models.ProductResponse) models.ProductCreatable {
	return models.ProductCreatable{Model: product.Model, Name: product.Name, Description: product.Description, Price: *product.Price, Thumb: product.Thumb, StockAmount: product.StockAmount, ReorderLevel: product.ReorderLevel, CategoryId: &product.Category.ID, DiscountId: &product.Discount.ID, Uint: product.Uint}
}

func (s *sqlStorage) GetUpdatableProduct(ctx context.Context, product models.ProductCreatable) models.ProductUpdatable {
	return models.ProductUpdatable{Model: product.Model, Name: product.Name, Description: product.Description, ReorderLevel: product.ReorderLevel, StockAmount: product.StockAmount}
}

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
		res[i] = s.GetDetailProduct(ctx, product)
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
		res[i] = s.GetDetailProduct(ctx, product)
	}
	return res, nil
}

func (s *sqlStorage) ReadProductById(ctx context.Context, id uint) (*models.ProductResponse, error) {
	var product models.Product
	if err := s.db.Where("id = ?", id).First(&product).Error; err != nil {
		fmt.Println("Error while read product by id in repository: " + err.Error())
		return nil, err
	}
	res := s.GetDetailProduct(ctx, product)
	return &res, nil
}

func (s *sqlStorage) ReadRecommendProduct(ctx context.Context, limit int) ([]models.ProductResponse, error) {
	var products models.Products
	if err := s.db.Table(models.Products{}.GetTableName()).Order("reorder_level desc").Limit(limit).Find(&products).Error; err != nil {
		fmt.Println("Error while read recommend product in repository: " + err.Error())
		return nil, err
	} else {
		var res = make([]models.ProductResponse, len(products))
		for i, product := range products {
			res[i] = s.GetDetailProduct(ctx, product)
		}
		return res, nil
	}
}

func (s *sqlStorage) UpdateProductById(ctx context.Context, id int, product *models.ProductUpdatable) (*int64, error) {
	if result := s.db.Table(models.ProductUpdatable{}.GetTableName()).Where("id = ?", id).Updates(product); result.Error != nil {
		fmt.Println("Error while update product in repository: " + result.Error.Error())
		return nil, result.Error
	} else {
		return &result.RowsAffected, nil
	}
}

func (s *sqlStorage) ReadTopGoodsByReorderLevel(ctx context.Context, startTime int, endTime int) ([]models.StatsTopGoodsByReorderLevel, error) {
	var res []models.StatsTopGoodsByReorderLevel
	startTimeUnix := time.Unix(int64(startTime), 0)
	endTimeUnix := time.Unix(int64(endTime), 0)
	if err := s.db.Table(models.StatsTopGoodsByReorderLevel{}.GetTableName()).Order("reorder_level DESC").Select("id, reorder_level").Where("created_at >= ? AND created_at <= ?", startTimeUnix, endTimeUnix).Find(&res).Error; err != nil {
		fmt.Println("Error while get all product by reorderLevel: " + err.Error())
		return nil, err
	}
	fmt.Println(res[0].Id)
	return res, nil
}

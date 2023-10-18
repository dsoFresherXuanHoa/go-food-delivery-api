package repositories

import (
	"context"
	"fmt"
	"go-food-delivery-api/src/models"
)

func (s *sqlStorage) CreateBill(ctx context.Context, bill *models.BillCreatable) (*uint, error) {
	if err := s.db.Table(models.BillCreatable{}.GetTableName()).Create(&bill).Error; err != nil {
		fmt.Println("Error while create bill in repository: " + err.Error())
		return nil, err
	}
	return &bill.ID, nil
}

func (s *sqlStorage) ReadBillById(ctx context.Context, billId uint) (*models.Bill, error) {
	var bill models.Bill
	if err := s.db.Where("id = ?", billId).First(&bill).Error; err != nil {
		fmt.Println("Error while find bill by id in repository: " + err.Error())
		return nil, err
	}
	return &bill, nil
}
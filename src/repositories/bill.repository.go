package repositories

import (
	"context"
	"fmt"
	"go-food-delivery-api/src/models"
	"go-food-delivery-api/src/services"
)

func (s *sqlStorage) Bill2Creatable(ctx context.Context, bill *models.Bill) models.BillCreatable {
	return models.BillCreatable{Model: bill.Model, OrderId: &bill.OrderId, Quantity: &bill.Quantity, ProductId: &bill.ProductId}
}

func (s *sqlStorage) Bill2Updatable(ctx context.Context, bill *models.Bill) models.BillUpdatable {
	return models.BillUpdatable{Model: bill.Model, Status: &bill.Status}
}

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

func (s *sqlStorage) ReadBillsByOrderId(ctx context.Context, orderId uint) ([]models.Bill, error) {
	var bills []models.Bill
	if err := s.db.Where("order_id = ?", orderId).Find(&bills).Error; err != nil {
		fmt.Println("Error while find bill by order id in repository: ", err.Error())
		return nil, err
	}
	fmt.Println(bills[0], bills[1])
	return bills, nil
}

func (s *sqlStorage) UpdateBillById(ctx context.Context, billId int, bill *models.BillUpdatable) (*int64, error) {
	if result := s.db.Table(models.BillUpdatable{}.GetTableName()).Where("id = ?", billId).Updates(bill); result.Error != nil {
		fmt.Println("Error while update bill by id in repository: " + result.Error.Error())
		return nil, result.Error
	} else {
		return &result.RowsAffected, nil
	}
}

func (s *sqlStorage) FinishBillById(ctx context.Context, billId uint) (*uint, error) {
	billService := services.NewBillBusiness(s)
	if bill, err := billService.ReadBillById(ctx, billId); err != nil {
		fmt.Println("Error while read bill by id in repository: " + err.Error())
		return nil, err
	} else {
		finished := true
		billUpdatable := s.Bill2Updatable(ctx, bill)
		billUpdatable.Status = &finished
		if _, err := billService.UpdateBillById(ctx, int(billId), &billUpdatable); err != nil {
			fmt.Println("Error while finish a bill in repository: " + err.Error())
			return nil, err
		}
		return &bill.ID, nil
	}
}

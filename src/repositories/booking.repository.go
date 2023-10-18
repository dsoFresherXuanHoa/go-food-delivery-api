package repositories

import (
	"context"
	"fmt"
	"go-food-delivery-api/src/models"
	"go-food-delivery-api/src/services"
)

func (s *sqlStorage) CreateBooking(ctx context.Context, order *models.OrderCreatable, bills []models.BillCreatable) (*uint, error) {
	repository := NewSQLStore(s.db)
	orderService := services.NewOrderBusiness(repository)
	billService := services.NewBillBusiness(repository)
	tableService := services.NewTableBusiness(repository)
	productService := services.NewProductBusiness(repository)

	if orderId, err := orderService.CreateOrder(ctx, *order); err != nil {
		fmt.Println("Error while create order in repository: " + err.Error())
		return nil, err
	} else {
		tableId := *order.TableId
		for _, bill := range bills {
			bill.OrderId = orderId
			if _, err := billService.CreateBill(ctx, &bill); err != nil {
				fmt.Println("Error while create bill in repository: " + err.Error())
				return nil, err
			} else {
				product, _ := productService.ReadProductById(ctx, *bill.ProductId)
				productCreatable := s.GetCreatableProduct(ctx, *product)
				productUpdatable := s.GetUpdatableProduct(ctx, productCreatable)

				orderQuantity := bill.Quantity
				oldReorderLevel := productUpdatable.ReorderLevel
				oldStockAmount := productUpdatable.StockAmount
				newReorderLevel := *oldReorderLevel + *orderQuantity
				newStockAmount := *oldStockAmount - *orderQuantity
				productUpdatable.StockAmount = &newStockAmount
				productUpdatable.ReorderLevel = &newReorderLevel
				if _, err := productService.UpdateProductById(ctx, int(productCreatable.ID), &productUpdatable); err != nil {
					fmt.Println("Error while update product in repository: " + err.Error())
					return nil, err
				}
			}
		}
		if table, err := tableService.ReadTableById(ctx, uint(tableId)); err != nil {
			fmt.Println("Error while read table by id in booking repository: " + err.Error())
			return nil, err
		} else {
			table.Available = false
			tableUpdatable := repository.Table2Updatable(ctx, table)
			if _, err := tableService.UpdateTable(ctx, int(tableId), &tableUpdatable); err != nil {
				fmt.Println("Error while update table status in repository: " + err.Error())
				return nil, err
			}
		}
		return orderId, nil
	}
}

func (s *sqlStorage) AcceptOrder(ctx context.Context, orderId int) (*uint, error) {
	repository := NewSQLStore(s.db)
	orderService := services.NewOrderBusiness(repository)
	if order, err := orderService.ReadOrderById(ctx, uint(orderId)); err != nil {
		fmt.Println("Error while update order in repository: " + err.Error())
		return nil, err
	} else {
		acceptedOrder := true
		rejectedOrder := false
		orderUpdatable := s.GetUpdatableOrder(ctx, *order)
		orderUpdatable.Accepted = &acceptedOrder
		orderUpdatable.Rejected = &rejectedOrder
		if _, err := orderService.UpdateOrderById(ctx, orderId, &orderUpdatable); err != nil {
			fmt.Println("Error while update order by id in repository: " + err.Error())
			return nil, err
		}
		return &order.ID, nil
	}
}

func (s *sqlStorage) RejectOrder(ctx context.Context, orderId int) (*uint, error) {
	repository := NewSQLStore(s.db)
	orderService := services.NewOrderBusiness(repository)
	if order, err := orderService.ReadOrderById(ctx, uint(orderId)); err != nil {
		fmt.Println("Error while update order in repository: " + err.Error())
		return nil, err
	} else {
		acceptedOrder := false
		rejectedOrder := true
		orderUpdatable := s.GetUpdatableOrder(ctx, *order)
		orderUpdatable.Rejected = &rejectedOrder
		orderUpdatable.Accepted = &acceptedOrder
		if _, err := orderService.UpdateOrderById(ctx, orderId, &orderUpdatable); err != nil {
			fmt.Println("Error while update order by id in repository: " + err.Error())
			return nil, err
		}
		return &order.ID, nil
	}
}

func (s *sqlStorage) FinishOrder(ctx context.Context, orderId int) (*uint, error) {
	repository := NewSQLStore(s.db)
	orderService := services.NewOrderBusiness(repository)
	if order, err := orderService.ReadOrderById(ctx, uint(orderId)); err != nil {
		fmt.Println("Error while update order in repository: " + err.Error())
		return nil, err
	} else {
		finished := true
		orderUpdatable := s.GetUpdatableOrder(ctx, *order)
		orderUpdatable.Status = &finished
		if _, err := orderService.UpdateOrderById(ctx, orderId, &orderUpdatable); err != nil {
			fmt.Println("Error while update order by id in repository: " + err.Error())
			return nil, err
		}
		return &order.ID, nil
	}
}

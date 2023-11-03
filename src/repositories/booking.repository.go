package repositories

import (
	"context"
	"fmt"
	exceptions "go-food-delivery-api/src/errors"
	"go-food-delivery-api/src/models"
	"go-food-delivery-api/src/services"
)

func (s *sqlStorage) CreateBooking(ctx context.Context, order *models.OrderCreatable, bills []models.BillCreatable, secretCode int) (*uint, error) {
	repository := NewSQLStore(s.db)
	orderService := services.NewOrderBusiness(repository)
	billService := services.NewBillBusiness(repository)
	tableService := services.NewTableBusiness(repository)
	productService := services.NewProductBusiness(repository)
	accountService := services.NewAccountBusiness(repository)
	accountId := ctx.Value("accountId").(int)
	account, _ := accountService.ReadAccountById(ctx, uint(accountId))
	if account.SecretCode != secretCode {
		fmt.Println("Error while create booking in repository: invalid secretCode")
		return nil, exceptions.ErrorInvalidSecretCode
	}
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
			tableUpdatable := repository.GetUpdatableTable(ctx, table)
			if _, err := tableService.UpdateTable(ctx, int(tableId), &tableUpdatable); err != nil {
				fmt.Println("Error while update table status in repository: " + err.Error())
				return nil, err
			}
		}
		return orderId, nil
	}
}

func (s *sqlStorage) GetDetailBooking(ctx context.Context, orderId int) (*models.BookingResponse, error) {
	repository := NewSQLStore(s.db)
	orderService := services.NewOrderBusiness(repository)
	tableService := services.NewTableBusiness(repository)
	billService := services.NewBillBusiness(repository)
	productService := services.NewProductBusiness(repository)
	if order, err := orderService.ReadOrderById(ctx, uint(orderId)); err != nil {
		fmt.Println("Error while update order in repository: " + err.Error())
		return nil, err
	} else {
		bills, _ := billService.ReadBillsByOrderId(ctx, uint(orderId))
		var embedItems []models.BookingItemResponse
		for _, bill := range bills {
			product, _ := productService.ReadProductById(ctx, bill.ProductId)
			embedItems = append(embedItems, models.BookingItemResponse{Product: *product, Quantity: bill.Quantity})
		}
		embedTable, _ := tableService.ReadTableById(ctx, order.TableId)
		detailEmbedTable := s.GetDetailTable(ctx, *embedTable)
		return &models.BookingResponse{Table: detailEmbedTable, Items: embedItems, Note: order.Note, Status: order.Status, Accepted: order.Accepted, Compensate: order.Compensate}, nil
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
	} else if order.Rejected {
		fmt.Println("Can't finish rejected order: ")
		return nil, exceptions.ErrorFinishRejectedOrder
	} else {
		tableService := services.NewTableBusiness(repository)
		table, _ := tableService.ReadTableById(ctx, order.TableId)
		finishedOrder := true
		availableTable := true
		orderUpdatable := s.GetUpdatableOrder(ctx, *order)
		tableUpdatable := s.GetUpdatableTable(ctx, table)
		orderUpdatable.Status = &finishedOrder
		tableUpdatable.Available = &availableTable
		if _, err := orderService.UpdateOrderById(ctx, orderId, &orderUpdatable); err != nil {
			fmt.Println("Error while update order by id in repository: " + err.Error())
			return nil, err
		} else if _, err := tableService.UpdateTable(ctx, int(order.TableId), &tableUpdatable); err != nil {
			fmt.Println("Error while update table in repository: " + err.Error())
			return nil, err
		}
		return &order.ID, nil
	}
}

func (s *sqlStorage) CompensatedOrder(ctx context.Context, orderId int, employeeId int) (*uint, error) {
	repository := NewSQLStore(s.db)
	orderService := services.NewOrderBusiness(repository)
	bookingService := services.NewBookingBusiness(repository)
	if order, err := orderService.ReadOrderById(ctx, uint(orderId)); err != nil {
		fmt.Println("Error while update order in repository: " + err.Error())
		return nil, err
	} else if order.Rejected || order.Status {
		fmt.Println("Error while update order in repository: ")
		return nil, exceptions.ErrorCompensatedFinishedOrRejectedOrder
	} else {
		compensated := true
		orderUpdatable := s.GetUpdatableOrder(ctx, *order)
		orderUpdatable.Compensate = &compensated
		if _, err := orderService.UpdateOrderById(ctx, orderId, &orderUpdatable); err != nil {
			fmt.Println("Error while update order by id in repository: " + err.Error())
			return nil, err
		} else if booking, err := bookingService.GetDetailBooking(ctx, orderId); err != nil {
			fmt.Println("Error while get order detail in booking repository: " + err.Error())
			return nil, err
		} else {
			compensatedOrder := models.OrderCreatable{Note: &booking.Note, EmployeeId: &order.EmployeeId, TableId: &order.TableId}
			compensatedBills := make([]models.BillCreatable, len(booking.Items))
			for i, item := range booking.Items {
				bill := models.Bill{Quantity: item.Quantity, OrderId: order.ID, ProductId: item.Product.ID}
				billCreatable := s.Bill2Creatable(ctx, &bill)
				compensatedBills[i] = billCreatable
			}
			secretCode := ctx.Value("secretCode").(int)
			if _, err := bookingService.CreateBooking(ctx, &compensatedOrder, compensatedBills, secretCode); err != nil {
				fmt.Println("Error while create compensated booking: " + err.Error())
				return nil, err
			}
		}
		return &order.ID, nil
	}
}

func (s *sqlStorage) GetOrdersByEmployeeId(ctx context.Context, employeeId int) ([]models.BookingResponse, error) {
	repository := NewSQLStore(s.db)
	orderService := services.NewOrderBusiness(repository)
	bookingService := services.NewBookingBusiness(repository)
	if orders, err := orderService.ReadOrdersByEmployeeId(ctx, uint(employeeId)); err != nil {
		fmt.Println("Error while get all order by employeeId in repository: " + err.Error())
		return nil, err
	} else {
		var res = make([]models.BookingResponse, len(orders))
		for i, order := range orders {
			orderId := order.ID
			orderDetail, _ := bookingService.GetDetailBooking(ctx, int(orderId))
			res[i] = *orderDetail
		}
		return res, nil
	}
}

func (s *sqlStorage) GetServeBookingsByTableId(ctx context.Context, tableId int) ([]models.BookingResponse, error) {
	repository := NewSQLStore(s.db)
	orderService := services.NewOrderBusiness(repository)
	bookingService := services.NewBookingBusiness(repository)
	if orders, err := orderService.GetServeOrdersByTableId(ctx, tableId); err != nil {
		fmt.Println("Error while get serve order by tableId in repository: " + err.Error())
		return nil, err
	} else {
		var res = make([]models.BookingResponse, len(orders))
		for i, order := range orders {
			orderId := order.ID
			orderDetail, _ := bookingService.GetDetailBooking(ctx, int(orderId))
			res[i] = *orderDetail
		}
		return res, nil
	}
}

func (s *sqlStorage) GetPreparingBookingsByTableId(ctx context.Context, tableId int) ([]models.BookingResponse, error) {
	repository := NewSQLStore(s.db)
	orderService := services.NewOrderBusiness(repository)
	bookingService := services.NewBookingBusiness(repository)
	if orders, err := orderService.GetPreparingOrdersByTableId(ctx, tableId); err != nil {
		fmt.Println("Error while get preparing order by tableId in repository: " + err.Error())
		return nil, err
	} else {
		var res = make([]models.BookingResponse, len(orders))
		for i, order := range orders {
			orderId := order.ID
			orderDetail, _ := bookingService.GetDetailBooking(ctx, int(orderId))
			res[i] = *orderDetail
		}
		return res, nil
	}
}

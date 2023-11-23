package repositories

import (
	"context"
	"fmt"
	exceptions "go-food-delivery-api/src/errors"
	"go-food-delivery-api/src/models"
	"go-food-delivery-api/src/services"
	"time"
)

func (s *sqlStorage) CreateBooking(ctx context.Context, order *models.OrderCreatable, bills []models.BillCreatable, secretCode int) (*uint, error) {
	repository := NewSQLStore(s.db)
	orderService := services.NewOrderBusiness(repository)
	billService := services.NewBillBusiness(repository)
	tableService := services.NewTableBusiness(repository)
	productService := services.NewProductBusiness(repository)
	accountService := services.NewAccountBusiness(repository)
	discountService := services.NewDiscountBusiness(repository)
	accountId := ctx.Value("accountId").(int)
	account, _ := accountService.ReadAccountById(ctx, uint(accountId))
	// Create Order without Total Price
	if account.SecretCode != secretCode {
		fmt.Println("Error while create booking in repository: invalid secretCode")
		return nil, exceptions.ErrorInvalidSecretCode
	}
	if orderId, err := orderService.CreateOrder(ctx, *order); err != nil {
		fmt.Println("Error while create order in repository: " + err.Error())
		return nil, err
	} else {
		tableId := *order.TableId
		// Get Product and Discount
		// Get Discount
		var discounts = make([]models.Discount, len(bills))
		var products = make([]models.Product, len(bills))
		for i := 0; i < len(bills); i++ {
			productId := bills[i].ProductId
			targetDiscount, _ := discountService.ReadDiscountById(ctx, *productId)
			targetProduct, _ := productService.ReadProductById(ctx, *productId)
			discounts[i] = *targetDiscount
			products[i] = models.Product{Price: *targetProduct.Price}
		}

		// Create Bill
		totalPrice := 0
		for i, bill := range bills {
			// Calc Total Price
			discountPercent := discounts[i].DiscountPercent
			minQuantity := discounts[i].MinQuantity
			originPrice := products[i].Price
			orderQuantity := *bill.Quantity
			if orderQuantity >= minQuantity {
				totalPrice += int(float64(float64(100-discountPercent)/100)*originPrice) * orderQuantity
			} else {
				totalPrice += int(originPrice) * orderQuantity
			}

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
			// Add Table Cost
			includeTableCost := *order.IncludeTableCost
			if includeTableCost {
				totalPrice += int(table.BasePrice)
			}

			// Update Total Price
			order.TotalPrice = &totalPrice
			if _, err := orderService.UpdateOrderById(ctx, int(*orderId), &models.OrderUpdatable{TotalPrice: &totalPrice}); err != nil {
				fmt.Println("Error while update total price of order: " + err.Error())
				return nil, err
			} else {
				table.Available = false
				tableUpdatable := repository.GetUpdatableTable(ctx, table)
				if _, err := tableService.UpdateTable(ctx, int(tableId), &tableUpdatable); err != nil {
					fmt.Println("Error while update table status in repository: " + err.Error())
					return nil, err
				}
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
			if bill.Quantity >= product.Discount.MinQuantity {
				product.ActualDiscountPrice = product.DiscountPrice
			} else {
				product.ActualDiscountPrice = product.Price
			}
			embedItems = append(embedItems, models.BookingItemResponse{Product: *product, Quantity: bill.Quantity, BillId: bill.ID, Status: bill.Status})
		}
		embedTable, _ := tableService.ReadTableById(ctx, order.TableId)
		detailEmbedTable := s.GetDetailTable(ctx, *embedTable)
		var acceptedTime *time.Time
		if order.Accepted {
			acceptedTime = &order.UpdatedAt
		} else {
			acceptedTime = nil
		}
		return &models.BookingResponse{Table: detailEmbedTable, OrderID: uint(orderId), CreatedTime: order.CreatedAt, AcceptedTime: acceptedTime, Items: embedItems, Note: order.Note, Status: order.Status, Accepted: order.Accepted, TotalPrice: order.TotalPrice}, nil
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

func (s *sqlStorage) RejectOrder(ctx context.Context, orderId int, reason string) (*uint, error) {
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
		orderUpdatable.Reason = &reason
		if _, err := orderService.UpdateOrderById(ctx, orderId, &orderUpdatable); err != nil {
			fmt.Println("Error while update order by id in repository: " + err.Error())
			return nil, err
		}

		tableService := services.NewTableBusiness(repository)
		table, _ := tableService.ReadTableById(ctx, order.TableId)
		availableTable := true
		tableUpdatable := s.GetUpdatableTable(ctx, table)
		if serveOrder, err := orderService.GetServeOrdersByTableId(ctx, int(order.TableId)); err == nil && len(serveOrder) <= 1 {
			fmt.Println("Incoming: ", err, len(serveOrder))
			tableUpdatable.Available = &availableTable
			if preparingTable, err := orderService.GetPreparingOrdersByTableId(ctx, int(order.TableId)); err == nil && len(preparingTable) == 0 {
				tableUpdatable.Available = &availableTable
				if _, err := tableService.UpdateTable(ctx, int(order.TableId), &tableUpdatable); err != nil {
					fmt.Println("Error while update table in repository: " + err.Error())
					return nil, err
				}
			}
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

		if serveOrder, err := orderService.GetServeOrdersByTableId(ctx, int(order.TableId)); err == nil && len(serveOrder) <= 1 {
			fmt.Println("Incoming: ", err, len(serveOrder))
			tableUpdatable.Available = &availableTable
			if preparingTable, err := orderService.GetPreparingOrdersByTableId(ctx, int(order.TableId)); err == nil && len(preparingTable) == 0 {
				tableUpdatable.Available = &availableTable
			}
		} else {
			fmt.Println("Incoming: ", err, len(serveOrder))
		}
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

func (s *sqlStorage) GetTop10OrdersByEmployeeId(ctx context.Context, employeeId int) ([]models.BookingResponse, error) {
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

func (s *sqlStorage) GetRejectedBookingsByTableId(ctx context.Context, tableId int) ([]models.BookingResponse, error) {
	repository := NewSQLStore(s.db)
	orderService := services.NewOrderBusiness(repository)
	bookingService := services.NewBookingBusiness(repository)
	if orders, err := orderService.GetRejectedOrdersByTableId(ctx, tableId); err != nil {
		fmt.Println("Error while get rejected order by tableId in repository: " + err.Error())
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

func (s *sqlStorage) RefundOrderById(ctx context.Context, orderId int, order *models.OrderCreatable, bills []models.BillCreatable, secretCode int) (*uint, *uint, error) {
	var deletedOrder models.Order
	if result := s.db.Table(models.OrderCreatable{}.GetTableName()).Where("id = ?", orderId).Delete(&deletedOrder); result.Error != nil {
		fmt.Println("Error while delete a refund order in repository: " + result.Error.Error())
		return nil, nil, result.Error
	} else if newOrderId, err := s.CreateBooking(ctx, order, bills, secretCode); err != nil {
		fmt.Println("Error while create new order after refunding in repository: " + err.Error())
		return nil, nil, err
	} else {
		return &deletedOrder.ID, newOrderId, nil
	}
}

func (s *sqlStorage) GetAllBooking(ctx context.Context) ([]models.BookingResponse, error) {
	var bookings []models.BookingResponse
	var orderIds []uint
	if result := s.db.Table(models.OrderCreatable{}.GetTableName()).Select("id").Find(&orderIds); result.Error != nil {
		fmt.Println("Error while get all order id: " + result.Error.Error())
		return nil, result.Error
	} else {
		for _, orderId := range orderIds {
			detailBooking, _ := s.GetDetailBooking(ctx, int(orderId))
			bookings = append(bookings, *detailBooking)
		}
		return bookings, nil
	}
}

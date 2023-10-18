package repositories

// func (s *sqlStorage) CreateBooking(ctx context.Context, order *models.OrderCreatable, bills []models.BillCreatable) (*uint, error) {
// 	repository := services.NewSQLStore(s.db)
// 	orderService := services.NewOrderBusiness(repository)
// 	billService := services.NewBillBusiness(repository)
// 	tableService := services.NewTableBusiness(repository)

// 	if orderId, err := orderService.CreateOrder(ctx, *order); err != nil {
// 		fmt.Println("Error while create order in repository: " + err.Error())
// 		return nil, err
// 	} else {
// 		tableId := int(order.TableId)
// 		for _, bill := range bills {
// 			*bill.OrderId = *orderId
// 			if _, err := createBillBusiness.CreateBill(ctx, &bill, db); err != nil {
// 				fmt.Println(err)
// 				return nil, err
// 			}
// 		}
// 		if table, err := tableService.ReadTableById(ctx, uint(tableId)); err != nil {
// 			fmt.Println("Error while read table by id in booking repository: " + err.Error())
// 			return nil, err
// 		} else {
// 			*table.Available = false
// 			if _, err := updateTableBusiness.UpdateTable(ctx, tableId, table); err != nil {
// 				fmt.Println(err)
// 				return nil, err
// 			}
// 		}
// 		return orderId, nil
// 	}
// }

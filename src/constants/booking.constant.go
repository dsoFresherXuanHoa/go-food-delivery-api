package constants

var (
	InvalidOrderIDQueryString = "Invalid orderId in query string: orderId must not be empty or must be numeric!"
	InvalidBillIDQueryString  = "Invalid billId in query string: billId must not be empty or must be numeric!"
	InvalidTableIdQueryString = "Invalid tableId in query string: tableId must not be empty or must be numeric!"
	InvalidReasonQueryString  = "Invalid reason in query string: reason cannot be empty!"

	CannotCreateBooking              = "Cannot create booking right now: Check your information or permission again! If you think nothing went wrong, try again later!"
	CannotAcceptBooking              = "Cannot accept booking right now: Check your information or permission again! If you think nothing went wrong, try again later!"
	CannotRejectBooking              = "Cannot reject booking right now: Check your information or permission again! If you think nothing went wrong, try again later!"
	CannotCompensatedBill            = "Cannot compensated bill right now: Check your information or permission again! If you think nothing went wrong, try again later!"
	CannotGetDetailBooking           = "Cannot get booking detail information right now: Check your information or permission again! If you think nothing went wrong, try again later!"
	CannotGetAllOrderByEmployeeId    = "Cannot get all booking by employeeId right now: Check your information or permission again! If you think nothing went wrong, try again later!"
	CannotGetServeOrderByTableId     = "Cannot get serve booking by tableId right now: Check your information or permission again! If you think nothing went wrong, try again later!"
	CannotGetPreparingOrderByTableId = "Cannot get preparing booking by tableId right now: Check your information or permission again! If you think nothing went wrong, try again later!"
	CannotRefundOrderByOrderId       = "Cannot refund booking by orderId right now: Check your information or permission again! If you think nothing went wrong, try again later!"
	RefundTargetNotFound             = "Cannot find bills by orderId before refunding: Can not refund un exists order"

	CreateBookingSuccess              = "Create booking success: congrats!"
	AcceptBookingSuccess              = "Accept booking success: congrats!"
	RejectBookingSuccess              = "Reject booking success: congrats!"
	CompensatedBillSuccess            = "Compensated bill success: congrats!"
	GetDetailBookingSuccess           = "Get detail booking information success: congrats!"
	GetAllOrderByEmployeeIdSuccess    = "Get all booking by employeeId success: congrats!"
	GetServeOrderByTableIdSuccess     = "Get serve booking by tableId success: congrats!"
	GetPreparingOrderByTableIdSuccess = "Get preparing booking by tableId success: congrats!"
	RefundOrderByOrderIdSuccess       = "Refund booking by orderId success: congrats!"
)

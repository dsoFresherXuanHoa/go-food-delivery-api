package constants

var (
	InvalidOrderIDQueryString = "Invalid orderId in query string: orderId must not be empty or must be numeric!"
	InvalidTableIdQueryString = "Invalid tableId in query string: tableId must not be empty or must be numeric!"

	CannotCreateBooking              = "Cannot create booking right now: Check your information or permission again! If you think nothing went wrong, try again later!"
	CannotAcceptBooking              = "Cannot accept booking right now: Check your information or permission again! If you think nothing went wrong, try again later!"
	CannotRejectBooking              = "Cannot reject booking right now: Check your information or permission again! If you think nothing went wrong, try again later!"
	CannotGetDetailBooking           = "Cannot get booking detail information right now: Check your information or permission again! If you think nothing went wrong, try again later!"
	CannotCompensateBooking          = "Cannot compensated an booking right now: Check your information or permission again! If you think nothing went wrong, try again later!"
	CannotGetAllOrderByEmployeeId    = "Cannot get all booking by employeeId right now: Check your information or permission again! If you think nothing went wrong, try again later!"
	CannotGetServeOrderByTableId     = "Cannot get serve booking by tableId right now: Check your information or permission again! If you think nothing went wrong, try again later!"
	CannotGetPreparingOrderByTableId = "Cannot get preparing booking by tableId right now: Check your information or permission again! If you think nothing went wrong, try again later!"

	CreateBookingSuccess              = "Create booking success: congrats!"
	AcceptBookingSuccess              = "Accept booking success: congrats!"
	RejectBookingSuccess              = "Reject booking success: congrats!"
	GetDetailBookingSuccess           = "Get detail booking information success: congrats!"
	CompensatedBookingSuccess         = "Compensated booking success: congrats!"
	GetAllOrderByEmployeeIdSuccess    = "Get all booking by employeeId success: congrats!"
	GetServeOrderByTableIdSuccess     = "Get serve booking by tableId success: congrats!"
	GetPreparingOrderByTableIdSuccess = "Get preparing booking by tableId success: congrats!"
)

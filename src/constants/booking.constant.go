package constants

var (
	InvalidOrderIDQueryString = "Invalid orderId in query string: orderId must not be empty or must be numeric!"

	CannotCreateBooking    = "Cannot create booking right now: Check your information or permission again! If you think nothing went wrong, try again later!"
	CannotAcceptBooking    = "Cannot accept booking right now: Check your information or permission again! If you think nothing went wrong, try again later!"
	CannotRejectBooking    = "Cannot reject booking right now: Check your information or permission again! If you think nothing went wrong, try again later!"
	CannotGetDetailBooking = "Cannot get booking detail information right now: Check your information or permission again! If you think nothing went wrong, try again later!"

	CreateBookingSuccess    = "Create booking success: congrats!"
	AcceptBookingSuccess    = "Accept booking success: congrats!"
	RejectBookingSuccess    = "Reject booking success: congrats!"
	GetDetailBookingSuccess = "Get detail booking information success: congrats!"
)

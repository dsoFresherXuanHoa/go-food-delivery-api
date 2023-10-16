package constants

var (
	CannotCreateTable                    = "Cannot create Table right now: Check your information or permission again! If you think nothing went wrong, try again later!"
	CannotReadTableByEmployeeId          = "Cannot read Table by employeeId right now: Check your information or permission again! If you think nothing went wrong, try again later!"
	CannotReadTableByEmployeeIdAndStatus = "Cannot read Table by employeeId and status right now: Check your information or permission again! If you think nothing went wrong, try again later!"

	CreateTableSuccess                    = "Create Table success: congrats!"
	ReadTableByEmployeeIdSuccess          = "Read Table by employeeId success: congrats!"
	ReadTableByEmployeeIdAndStatusSuccess = "Read Table by employeeId and status success: congrats!"

	InvalidTableStatusQueryString = "Invalid table status in query string: table status must be true or false!"
)

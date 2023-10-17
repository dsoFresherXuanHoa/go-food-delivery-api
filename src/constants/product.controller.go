package constants

var (
	InvalidEmployeeIDQueryString = "Invalid employeeId in query string: employeeId must not be empty or must be numeric!"
	InvalidLimitQueryString      = "Invalid limit in query string: employeeId must not be empty or must be numeric!"

	CannotReadProduct          = "Cannot read Product right now: Check your information or permission again! If you think nothing went wrong, try again later!"
	CannotReadRecommendProduct = "Cannot read recommend product right now: Check your information or permission again! If you think nothing went wrong, try again later!"

	ReadProductSuccess          = "Read Product success: congrats!"
	ReadProductRecommendSuccess = "Read recommend Product success: congrats!"
)

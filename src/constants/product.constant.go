package constants

var (
	InvalidEmployeeIDQueryString = "Invalid employeeId in query string: employeeId must not be empty or must be numeric!"
	InvalidLimitQueryString      = "Invalid limit in query string: employeeId must not be empty or must be numeric!"
	InvalidProductFormDataStruct = "Invalid form data struct: can't access form data, check your request again or try again later!"
	InvalidProductFormFile       = "Invalid product thumbnail in form data: check image again or try another one!"

	CannotCreateProduct                       = "Cannot create Product right now: check your information or permission again! If you think nothing went wrong, try again later!"
	CannotReadProduct                         = "Cannot read Product right now: check your information or permission again! If you think nothing went wrong, try again later!"
	CannotReadRecommendProduct                = "Cannot read recommend product right now: check your information or permission again! If you think nothing went wrong, try again later!"
	CannotSaveProductThumbnailsToLocalStorage = "Cannot save product thumbnail to local storage: check your request file or try again later"

	CreateProductSuccess        = "Create Product success: congrats!"
	ReadProductSuccess          = "Read Product success: congrats!"
	ReadProductRecommendSuccess = "Read recommend Product success: congrats!"
)

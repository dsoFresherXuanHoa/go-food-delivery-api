package constants

var (
	EmptyBearerHeader  = "Empty Bearer Header: you must contain accessToken in Header to validate your permission"
	InvalidAccessToken = "Invalid Access Token: accessToken has been expired or damaged"
	PermissionDenied   = "Permission Denied: you don't has permission to access this resource, contact manager to more detail"
)

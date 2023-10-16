package constants

var (
	InvalidRequestBody   = "Invalid Request Body: cannot parse request body to struct!"
	CannotSignUpRightNow = "Cannot sign up right now: Check your information or try again later!"
	CannotCreateRole     = "Cannot create role right now: Check your information or permission again! If you think nothing went wrong, try again later!"
	CannotGetUserInfo    = "Cannot get user information right now: may be your accessToken has been expired or damage, try log out and log in again!"

	SignUpSuccess      = "Sign up success: welcome!"
	CreateRoleSuccess  = "Create role success: congrats!"
	GetUserInfoSuccess = "Get user info success: congrats!"
)

package constants

var (
	InvalidRequestBody = "Invalid Request Body: cannot parse request body to struct!"

	CannotSignUpRightNow = "Cannot sign up right now: Check your information or try again later!"
	CannotGetUserInfo    = "Cannot get user information right now: may be your accessToken has been expired or damage, try log out and log in again!"
	CannotResetPassword  = "Cannot reset password right now: may be your accessToken has been expired or damage, try log out and log in again!"

	SignUpSuccess        = "Sign up success: welcome!"
	GetUserInfoSuccess   = "Get user info success: congrats!"
	ResetPasswordSuccess = "Reset user password success: congrats!"
)

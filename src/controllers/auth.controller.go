package controllers

import (
	"fmt"
	"go-food-delivery-api/src/constants"
	"go-food-delivery-api/src/models"
	"go-food-delivery-api/src/repositories"
	"go-food-delivery-api/src/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// SignUp godoc
//
//	@Summary		Sign up
//	@Description	Sign up
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			signUp	body		models.SignUp	true	"Sign Up"
//	@Success		200	{object}		models.response
//	@Failure		400	{object}	models.response
//	@Failure		500	{object}	models.response
//	@Router			/auth/sign-up [post]
func SignUp(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var signUp models.SignUp
		if err := ctx.ShouldBind(&signUp); err != nil {
			fmt.Println("Error while try parse request body to struct: " + err.Error())
			ctx.JSON(http.StatusBadRequest, models.NewStandardResponse(nil, http.StatusBadRequest, err.Error(), constants.InvalidRequestBody))
		} else {
			employee := models.EmployeeCreatable{FullName: &signUp.FullName, Tel: &signUp.Tel, Gender: &signUp.Gender}
			account := models.AccountCreatable{Password: &signUp.Email, Email: &signUp.Email, RoleId: &signUp.RoleId, EmployeeId: &employee.ID}

			repositories := repositories.NewSQLStore(db)
			authService := services.NewAuthBusiness(repositories)

			if employeeId, accountId, err := authService.SignUp(ctx.Request.Context(), &account, &employee); err != nil {
				fmt.Println("Error while sign up in auth controller: " + err.Error())
				ctx.JSON(http.StatusInternalServerError, models.NewStandardResponse(nil, http.StatusInternalServerError, err.Error(), constants.CannotSignUpRightNow))
			} else {
				ctx.JSON(http.StatusOK, models.NewStandardResponse(gin.H{
					"employeeId": employeeId,
					"accountId":  accountId,
				}, http.StatusOK, "", constants.SignUpSuccess))
			}
		}
	}
}

// SignIn godoc
//
//	@Summary		SignIn
//	@Description	SignIn
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			signIn	body		models.SignIn	true	"Sign In"
//	@Success		200		{object}	models.response
//	@Failure		400		{object}	models.response
//	@Failure		500		{object}	models.response
//	@Router			/auth/sign-in [post]
func SignIn(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var signIn models.SignIn
		if err := ctx.ShouldBind(&signIn); err != nil {
			fmt.Println("Error while try parse request body to struct: " + err.Error())
			ctx.JSON(http.StatusBadRequest, models.NewStandardResponse(nil, http.StatusBadRequest, err.Error(), constants.InvalidRequestBody))
		} else {
			repositories := repositories.NewSQLStore(db)
			authService := services.NewAuthBusiness(repositories)

			if token, err := authService.SignIn(ctx, &signIn); err != nil {
				fmt.Println("Error while sign in in auth controller: " + err.Error())
				ctx.JSON(http.StatusInternalServerError, models.NewStandardResponse(nil, http.StatusInternalServerError, err.Error(), constants.CannotSignUpRightNow))
			} else {
				ctx.JSON(http.StatusOK, models.NewStandardResponse(token, http.StatusOK, "", constants.SignUpSuccess))
			}
		}
	}
}

// Me godoc
//
//	@Summary		Me
//	@Description	Me
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param  Authorization  header  string  required  "Bearer Token"
//	@Success		200	{object}	models.response
//	@Failure		403	{object}	models.response
//	@Router			/auth/me [get]
func Me(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var employee models.Employee
		if err := ctx.ShouldBind(&employee); err != nil {
			fmt.Println("Error while try parse request body to struct: " + err.Error())
			ctx.JSON(http.StatusBadRequest, models.NewStandardResponse(nil, http.StatusBadRequest, err.Error(), constants.InvalidRequestBody))
		} else {
			repositories := repositories.NewSQLStore(db)
			employeeService := services.NewEmployeeBusiness(repositories)

			employeeId := ctx.Value("employeeId").(int)
			if token, err := employeeService.ReadEmployeeById(ctx, uint(employeeId)); err != nil {
				fmt.Println("Error while get user information in auth controller: " + err.Error())
				ctx.JSON(http.StatusInternalServerError, models.NewStandardResponse(nil, http.StatusInternalServerError, err.Error(), constants.CannotGetUserInfo))
			} else {
				ctx.JSON(http.StatusOK, models.NewStandardResponse(token, http.StatusOK, "", constants.GetUserInfoSuccess))
			}
		}
	}
}

// ResetPassword godoc
//
//	@Summary		Reset Password
//	@Description	Reset Password
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param  Authorization  header  string  required  "Bearer Token"
//	@Param			resetPassword	body		models.ResetPasswordCreatable	true	"Reset Password"
//	@Success		200		{object}	models.response
//	@Failure		400		{object}	models.response
//	@Failure		500		{object}	models.response
//	@Router			/auth/reset-password [patch]
func ResetPassword(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var resetPassword models.ResetPasswordCreatable
		if err := ctx.ShouldBind(&resetPassword); err != nil {
			fmt.Println("Error while try parse request body to struct: " + err.Error())
			ctx.JSON(http.StatusBadRequest, models.NewStandardResponse(nil, http.StatusBadRequest, err.Error(), constants.InvalidRequestBody))
		} else {
			repository := repositories.NewSQLStore(db)
			authService := services.NewAuthBusiness(repository)
			if resetPasswordId, err := authService.ResetPassword(ctx, &resetPassword); err != nil {
				fmt.Println("Error while reset password in auth controller: " + err.Error())
				ctx.JSON(http.StatusInternalServerError, models.NewStandardResponse(nil, http.StatusInternalServerError, err.Error(), constants.CannotResetPassword))
			} else {
				ctx.JSON(http.StatusOK, models.NewStandardResponse(resetPasswordId, http.StatusOK, "", constants.ResetPasswordSuccess))
			}
		}
	}
}

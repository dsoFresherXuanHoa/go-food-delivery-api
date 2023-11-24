package middlewares

import (
	"context"
	"fmt"
	"go-food-delivery-api/src/constants"
	"go-food-delivery-api/src/models"
	"go-food-delivery-api/src/repositories"
	"go-food-delivery-api/src/services"
	"go-food-delivery-api/src/tokens/jwt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	exception "go-food-delivery-api/src/errors"
)

func GetTokenFromHeader(c *gin.Context, key string) (token *string, err error) {
	if authHeader := c.Request.Header.Get(key); authHeader == "" {
		fmt.Println("Error while get Bearer header: missing Token")
		return nil, exception.ErrorEmptyBearer
	} else {
		accessToken := strings.Split(authHeader, " ")[1]
		return &accessToken, nil
	}
}

func RequiredWaiterPermissionOrMore(secretKey string, db *gorm.DB) gin.HandlerFunc {
	jwtTokenProvider := jwt.NewJWTProvider(secretKey)
	return func(c *gin.Context) {
		if authToken, err := GetTokenFromHeader(c, "Authorization"); err != nil {
			fmt.Println("Error while get Bearer header: " + err.Error())
			c.AbortWithStatusJSON(http.StatusUnauthorized, models.NewStandardResponse(nil, http.StatusUnauthorized, err.Error(), constants.EmptyBearerHeader))
		} else {
			if jwtPayload, err := jwtTokenProvider.Validate(*authToken); err != nil {
				fmt.Println("Error while validate accessToken: " + err.Error())
				c.AbortWithStatusJSON(http.StatusUnauthorized, models.NewStandardResponse(nil, http.StatusUnauthorized, err.Error(), constants.InvalidAccessToken))
			} else {
				// TODO: Strict validate by query from database
				if roleId := jwtPayload.RoleId; roleId != 1 && roleId != 2 {
					c.AbortWithStatusJSON(http.StatusForbidden, models.NewStandardResponse(nil, http.StatusForbidden, constants.PermissionDenied, constants.PermissionDenied))
				} else {
					accountId := jwtPayload.AccountId
					employeeId := jwtPayload.EmployeeId
					repository := repositories.NewSQLStore(db)
					accountService := services.NewAccountBusiness(repository)
					account, _ := accountService.ReadAccountById(context.Background(), uint(accountId))
					secretCode := account.SecretCode
					c.Set("accountId", accountId)
					c.Set("employeeId", employeeId)
					c.Set("secretCode", secretCode)
					c.Next()
				}
			}
		}
	}
}

func RequiredChiefPermissionOrMore(secretKey string) gin.HandlerFunc {
	jwtTokenProvider := jwt.NewJWTProvider(secretKey)
	return func(c *gin.Context) {
		if authToken, err := GetTokenFromHeader(c, "Authorization"); err != nil {
			fmt.Println("Error while get Bearer header: " + err.Error())
			c.AbortWithStatusJSON(http.StatusUnauthorized, models.NewStandardResponse(nil, http.StatusUnauthorized, err.Error(), constants.EmptyBearerHeader))
		} else {
			if jwtPayload, err := jwtTokenProvider.Validate(*authToken); err != nil {
				fmt.Println("Error while validate accessToken: " + err.Error())
				c.AbortWithStatusJSON(http.StatusUnauthorized, models.NewStandardResponse(nil, http.StatusUnauthorized, err.Error(), constants.InvalidAccessToken))
			} else {
				// TODO: Strict validate by query from database
				if roleId := jwtPayload.RoleId; roleId != 1 && roleId != 3 {
					c.AbortWithStatusJSON(http.StatusForbidden, models.NewStandardResponse("", http.StatusForbidden, constants.PermissionDenied, constants.PermissionDenied))
				} else {
					accountId := jwtPayload.AccountId
					employeeId := jwtPayload.EmployeeId
					c.Set("accountId", accountId)
					c.Set("employeeId", employeeId)
					c.Next()
				}
			}
		}
	}
}

func RequiredManagerPermission(secretKey string) gin.HandlerFunc {
	jwtTokenProvider := jwt.NewJWTProvider(secretKey)
	return func(c *gin.Context) {
		if authToken, err := GetTokenFromHeader(c, "Authorization"); err != nil {
			fmt.Println("Error while get Bearer header: " + err.Error())
			c.AbortWithStatusJSON(http.StatusUnauthorized, models.NewStandardResponse(nil, http.StatusUnauthorized, err.Error(), constants.EmptyBearerHeader))
		} else {
			if jwtPayload, err := jwtTokenProvider.Validate(*authToken); err != nil {
				fmt.Println("Error while validate accessToken: " + err.Error())
				c.AbortWithStatusJSON(http.StatusUnauthorized, models.NewStandardResponse(nil, http.StatusUnauthorized, err.Error(), constants.InvalidAccessToken))
			} else {
				// TODO: Strict validate by query from database
				if roleId := jwtPayload.RoleId; roleId != 1 {
					c.AbortWithStatusJSON(http.StatusForbidden, models.NewStandardResponse(nil, http.StatusForbidden, constants.PermissionDenied, constants.PermissionDenied))
				} else {
					accountId := jwtPayload.AccountId
					employeeId := jwtPayload.EmployeeId
					c.Set("accountId", accountId)
					c.Set("employeeId", employeeId)
					c.Next()
				}
			}
		}
	}
}

func RequiredAuthorize(secretKey string) gin.HandlerFunc {
	jwtTokenProvider := jwt.NewJWTProvider(secretKey)
	return func(c *gin.Context) {
		if authToken, err := GetTokenFromHeader(c, "Authorization"); err != nil {
			fmt.Println("Error while get Bearer header: " + err.Error())
			c.AbortWithStatusJSON(http.StatusUnauthorized, models.NewStandardResponse(nil, http.StatusUnauthorized, err.Error(), constants.EmptyBearerHeader))
		} else {
			if jwtPayload, err := jwtTokenProvider.Validate(*authToken); err != nil {
				fmt.Println("Error while validate accessToken: " + err.Error())
				c.AbortWithStatusJSON(http.StatusUnauthorized, models.NewStandardResponse(nil, http.StatusUnauthorized, err.Error(), constants.InvalidAccessToken))
			} else {
				// TODO: Strict validate by query from database
				if roleId := jwtPayload.RoleId; roleId == 0 {
					c.AbortWithStatusJSON(http.StatusForbidden, models.NewStandardResponse(nil, http.StatusForbidden, constants.PermissionDenied, constants.PermissionDenied))
				} else {
					accountId := jwtPayload.AccountId
					employeeId := jwtPayload.EmployeeId
					c.Set("accountId", accountId)
					c.Set("employeeId", employeeId)
					c.Next()
				}
			}
		}
	}
}

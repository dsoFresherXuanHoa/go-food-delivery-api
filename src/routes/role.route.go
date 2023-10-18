package routes

import (
	"go-food-delivery-api/src/controllers"

	"github.com/gin-gonic/gin"
)

func RoleRouteConfig(router *gin.Engine) {
	roles := router.Group("/api/v1/roles")
	{
		roles.POST("/", controllers.CreateRole())
	}
}

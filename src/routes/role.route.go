package routes

import (
	"go-food-delivery-api/src/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RoleRouteConfig(router *gin.Engine, db *gorm.DB) {
	roles := router.Group("/api/v1/roles")
	{
		roles.POST("/", controllers.CreateRole(db))
	}
}

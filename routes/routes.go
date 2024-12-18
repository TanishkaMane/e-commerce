package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/tanishkamane/ecommerce/controllers"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("users/signup", controllers.SignUp())
	incomingRoutes.POST("users/login", controllers.Login())
	incomingRoutes.POST("admin/addProduct", controllers.ProductViewerAdmin())
	incomingRoutes.GET("admin/productView", controllers.SearchProduct())
	incomingRoutes.GET("admin/search", controllers.SearchProductByQuery())
}

package router

import (
	"github.com/gin-gonic/gin"
	"jwt/controllers"
	"jwt/middlewares"
)

func StarApp() *gin.Engine {
	r := gin.Default()

	userRouter := r.Group("/users")
	{
		userRouter.POST("/register", controllers.UserRegister)

		userRouter.POST("/login", controllers.UserLogin)
	}
	return r

	productRouter := r.Group("/products")
	{
		productRouter.Use(middlewares.Auhemtication())
		productRouter.POST("/", controllers.CreateProduct)

		productRouter.PUT("/:productId", middlewares.ProductAuthorization(), controllers.UpdateProduct)
	}
	return r

	return r
}

package main

import (
	"github.com/gin-gonic/gin"
	v1handler "minhbao.com/hoc-golang/internal/api/v1/handler"
	v2handler "minhbao.com/hoc-golang/internal/api/v2/handler"
)

func main() {

	r := gin.Default()

	v1 := r.Group("/api/v1")
	{
		user := v1.Group("/users")
		{
			userHandlerv1 := v1handler.NewUserHandler()
			user.GET("", userHandlerv1.GetUsersV1)
			user.GET("/:id", userHandlerv1.GetUsersByIdV1)
			user.GET("/admin/:uuid", userHandlerv1.GetUsersByUuIdV1)
			user.POST("", userHandlerv1.PostUsersV1)
			user.PUT("/:id", userHandlerv1.PutUsersV1)
			user.DELETE("/:id", userHandlerv1.DeleteUsersV1)
		}

		product := v1.Group("/products")
		{
			productHandlerv1 := v1handler.NewProductHandler()
			product.GET("", productHandlerv1.GetProductsV1)
			product.GET("/:slug", productHandlerv1.GetProductsBySlugV1)
			product.POST("", productHandlerv1.PostProductsV1)
			product.PUT("/:id", productHandlerv1.PutProductsV1)
			product.DELETE("/:id", productHandlerv1.DeleteProductsV1)
		}

		category := v1.Group("/categories")
		{
			categoryHandlerv1 := v1handler.NewCategoryHandler()
			category.GET("/:category", categoryHandlerv1.GetCategoryByCategoryV1)
		}

	}

	v2 := r.Group("/api/v2")
	{
		user := v2.Group("/users")
		{
			userHandlerv2 := v2handler.NewUserHandler()
			user.GET("", userHandlerv2.GetUsersV2)
			user.GET("/:id", userHandlerv2.GetUsersByIdV2)
			user.POST("", userHandlerv2.PostUsersV2)
			user.PUT("/:id", userHandlerv2.PutUsersV2)
			user.DELETE("/:id", userHandlerv2.DeleteUsersV2)
		}

	}

	r.Run(":8080")
}

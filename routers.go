package main

import (
	"awesomeProject/component/appctx"
	"awesomeProject/module/restaurant/transport/ginrestaurant"
	"awesomeProject/module/upload/transport/ginupload"
	"github.com/gin-gonic/gin"
)

func setupRoute(appContext appctx.AppContext, v1 *gin.RouterGroup) {
	restaurants := v1.Group("/restaurants")
	{
		restaurants.POST("", ginrestaurant.CreateRestaurant(appContext))
		restaurants.GET("", ginrestaurant.ListRestaurant(appContext))
		restaurants.DELETE("/:id", ginrestaurant.DeleteRestaurant(appContext))
	}

	upload := v1.Group("/upload")
	{
		upload.POST("", ginupload.UploadImage(appContext))
	}
}

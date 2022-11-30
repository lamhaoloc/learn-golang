package main

import (
	"awesomeProject/component/appctx"
	"awesomeProject/middleware"
	"awesomeProject/module/restaurant/transport/gin_restaurant"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

func main() {
	dsn := os.Getenv("MYSQL_CONFIG")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}

	appCtx := appctx.NewAppContext(db)

	r := gin.Default()
	r.Use(middleware.Recover(appCtx))

	v1 := r.Group("/v1")
	restaurants := v1.Group("/restaurants")
	restaurants.POST("/", gin_restaurant.CreateRestaurant(appCtx))

	r.Run()
}

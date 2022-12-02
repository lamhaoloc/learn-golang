package main

import (
	"awesomeProject/component/appctx"
	"awesomeProject/component/uploadprovider"
	"awesomeProject/middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
)

func main() {
	dsn := os.Getenv("MYSQL_CONFIG")

	s3BucketName := os.Getenv("S3_BUCKET_NAME")
	s3Region := os.Getenv("S3_REGION")
	s3APIKey := os.Getenv("S3_API_KEY")
	s3SecretKey := os.Getenv("S3_SECRET_KEY")
	s3Domain := os.Getenv("S3_DOMAIN")

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db = db.Debug()

	s3Provider := uploadprovider.NewS3Provider(s3BucketName, s3Region, s3APIKey, s3SecretKey, s3Domain)
	appCtx := appctx.NewAppContext(db, s3Provider)

	r := gin.Default()
	r.Use(middleware.Recover(appCtx))
	r.Static("/static", "./static")

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	v1 := r.Group("/v1")
	setupRoute(appCtx, v1)

	r.Run()
}

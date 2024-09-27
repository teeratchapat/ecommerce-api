package main

import (
	"log"
	"menu-api/handlers"
	"menu-api/models"
	"menu-api/seed"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func main() {
	dsn := "user:password@tcp(127.0.0.1:3307)/ecommerce_db?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	models.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database!", err)
	}

	models.DB.AutoMigrate(&models.Menu{})
	seed.SeedMenus(models.DB)

	r := gin.Default()
	r.GET("/", handlers.GetMenus)
	r.GET("/menus", handlers.GetMenus)
	r.Run(":8080")
}

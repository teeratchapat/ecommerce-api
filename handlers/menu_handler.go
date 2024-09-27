package handlers

// import (
// 	"menu-api/models"
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )

// func GetMenus(c *gin.Context) {
// 	var menus []models.Menu
// 	models.DB.Preload("Children").Where("parent_id IS NULL").Find(&menus) // ดึงเมนูพร้อมกับ children
// 	c.JSON(http.StatusOK, gin.H{"status": true, "data": menus})
// }

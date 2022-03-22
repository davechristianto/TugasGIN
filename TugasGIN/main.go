package main

import (
	controllers "TugasGIN/Controller"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	router := gin.Default()

	router.GET("/menus", controllers.GetMenus)
	router.POST("/menus", controllers.AddMenus)
	router.PUT("/menus", controllers.UpdateMenus)
	router.DELETE("/menus", controllers.DeleteMenus)

	router.Run()
}

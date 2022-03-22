package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetMenus(c *gin.Context) {

	db := Connect()
	defer db.Close()

	results, err := db.Query("SELECT * FROM menu")
	if err != nil {
		fmt.Println("Err", err.Error())
	}

	var menu Menu
	var menus []Menu

	for results.Next() {
		err = results.Scan(&menu.ID, &menu.Nama, &menu.Harga)
		if err != nil {
			panic(err.Error())
		}
		menus = append(menus, menu)
	}

	if len(menus) == 0 {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.IndentedJSON(http.StatusCreated, menus)
	}

}

func AddMenus(c *gin.Context) {
	db := Connect()
	defer db.Close()

	var menu Menu

	if err := c.Bind(&menu); err != nil {
		fmt.Println(err)
		return
	}

	db.Exec(`INSERT INTO menu(Nama, Harga) VALUES(?, ?)`, menu.Nama, menu.Harga)

	c.IndentedJSON(http.StatusOK, "Success")
}
func DeleteMenus(c *gin.Context) {
	db := Connect()
	defer db.Close()

	id := c.Query("id")

	result, errQuery := db.Exec(`DELETE FROM menu WHERE id=?`, id)

	num, _ := result.RowsAffected()

	if errQuery == nil {
		if num == 0 {
			c.AbortWithStatusJSON(400, "Failed")
			return
		} else {
			c.IndentedJSON(http.StatusOK, "Success")
		}
	}
}
func UpdateMenus(c *gin.Context) {
	db := Connect()
	defer db.Close()

	var menu Menu

	if err := c.Bind(&menu); err != nil {
		fmt.Println(err)
		return
	}
	result, errQuery := db.Exec(`UPDATE menu SET Nama=?, Harga=? WHERE ID=?`, menu.Nama, menu.Harga, menu.ID)

	num, _ := result.RowsAffected()

	if errQuery == nil {
		if num == 0 {
			c.AbortWithStatusJSON(400, "Failed")
			return
		} else {
			c.IndentedJSON(http.StatusOK, "Success")
		}
	}
}

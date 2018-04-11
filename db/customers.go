package db

import (
	"github.com/gin-gonic/gin"
	"net/http"

)

type customers struct {
		Id        string    `json:"id" gorm:"column:id;primary_key;AUTO_INCREMENT "`
		Type_name string `json:"type_name" gorm:"column:type_name;not null"`

	}


func CrudCustomers(app *gin.RouterGroup) {

	app.GET("/", getAllCustomers)
	app.GET("/:id", getOneCustomers)
	app.POST("/", addCustomers)
	app.DELETE("/:id", deleteCustomers)
	app.PUT("/:id", updateCustomers)

}


func getAllCustomers(c *gin.Context) {

	var items [] customers
	db, err := OpenDb()
	defer db.Close()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	} else {
		db.Find(&items)

		if len(items) <= 0 {
			c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No Customers found!"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": items})
	}
}

func getOneCustomers(c *gin.Context) {
	truckID := c.Param("id")
	db, err := OpenDb()
	defer db.Close()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	} else {
		var items [] customers

		db.First(&items, truckID)

		if len(items) <= 0 {

			c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No Customers found!"})

			return
		}

		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": items})
	}
}


func addCustomers(c *gin.Context) {

	var item customers

	db, err := OpenDb()
	defer db.Close()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	} else {
		c.BindJSON(&item);
		if err := db.Create(&item).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		} else {
			c.JSON(http.StatusCreated, item)
		}
	}
}


func deleteCustomers(c *gin.Context) {

	var item customers

	db, err := OpenDb()
	defer db.Close()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	} else {
		id := c.Params.ByName("id")

		if err := db.Where("id = ?", id).First(&item).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		} else {
			db.Delete(&item)
			c.JSON(http.StatusOK, "Object "+id+" Deleted")
		}
	}
}

func updateCustomers(c *gin.Context) {

	var item customers
	id := c.Params.ByName("id")

	db, err := OpenDb()
	defer db.Close()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	} else {
		if err := db.Where("id = ?", id).First(&item).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		} else {
			if err := c.BindJSON(&item); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			} else {
				db.Save(&item)
				c.JSON(http.StatusOK, item)
			}
		}
	}
}

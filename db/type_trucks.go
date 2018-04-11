package db

import (
	"github.com/gin-gonic/gin"
	"net/http"

)

type type_trucks struct {
		Id        string    `json:"id" gorm:"column:id;primary_key;AUTO_INCREMENT "`
		Type_name string `json:"type_name" gorm:"column:type_name;not null"`

	}


func CrudType_Trucks(app *gin.RouterGroup) {

	app.GET("/", getAllType_Trucks)
	app.GET("/:id", getOneType_Trucks)
	app.POST("/", addType_Trucks)
	app.DELETE("/:id", deleteType_Trucks)
	app.PUT("/:id", updateType_Trucks)

}


func getAllType_Trucks(c *gin.Context) {

	var items [] type_trucks
	db, err := OpenDb()
	defer db.Close()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	} else {
		db.Find(&items)

		if len(items) <= 0 {
			c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No Type_Trucks found!"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": items})
	}
}

func getOneType_Trucks(c *gin.Context) {
	truckID := c.Param("id")
	db, err := OpenDb()
	defer db.Close()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	} else {
		var items [] type_trucks

		db.First(&items, truckID)

		if len(items) <= 0 {

			c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No Type_Trucks found!"})

			return
		}

		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": items})
	}
}


func addType_Trucks(c *gin.Context) {

	var item type_trucks

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


func deleteType_Trucks(c *gin.Context) {

	var item type_trucks

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

func updateType_Trucks(c *gin.Context) {

	var item type_trucks
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

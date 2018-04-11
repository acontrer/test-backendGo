package db

import (
	"github.com/gin-gonic/gin"
	"net/http"

)

type adresses struct {
		Id        string    `json:"id" gorm:"column:id;primary_key;AUTO_INCREMENT "`
		Type_name string `json:"type_name" gorm:"column:type_name;not null"`

	}


func CrudAdresses(app *gin.RouterGroup) {

	app.GET("/", getAllAdresses)
	app.GET("/:id", getOneAdresses)
	app.POST("/", addAdresses)
	app.DELETE("/:id", deleteAdresses)
	app.PUT("/:id", updateAdresses)

}


func getAllAdresses(c *gin.Context) {

	var items [] adresses
	db, err := OpenDb()
	defer db.Close()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	} else {
		db.Find(&items)

		if len(items) <= 0 {
			c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No Adresses found!"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": items})
	}
}

func getOneAdresses(c *gin.Context) {
	truckID := c.Param("id")
	db, err := OpenDb()
	defer db.Close()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	} else {
		var items [] adresses

		db.First(&items, truckID)

		if len(items) <= 0 {

			c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No Adresses found!"})

			return
		}

		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": items})
	}
}


func addAdresses(c *gin.Context) {

	var item adresses

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


func deleteAdresses(c *gin.Context) {

	var item adresses

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

func updateAdresses(c *gin.Context) {

	var item adresses
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

package db

import (
	"github.com/gin-gonic/gin"
	"net/http"

)

type Trucks struct {
		Id        string    `json:"id" gorm:"column:id;primary_key;AUTO_INCREMENT "`
		Truck_number string `json:"truck_number" gorm:"column:truck_number;not null"`
		Type_truck_id string `json:"type_truck_id" gorm:"column:type_truck_id;not null"`
		Deliver_list []Deliver_lists`gorm:"foreignkey:Trucks_id;association_foreignkey:Id"`



	}


func CrudTrucks(app *gin.RouterGroup) {

	app.GET("/", getAllTrucks)
	app.GET("/:id", getOneTrucks)
	app.POST("/", addTrucks)
	app.DELETE("/:id", deleteTrucks)
	app.PUT("/:id", updateTrucks)
	// has many services
	app.GET("/:id/deliver_list", getTrucksDeliver)
}

func getTrucksDeliver(c *gin.Context) {

	var items [] Trucks
	truckID := c.Param("id")
	db, err := OpenDb()
	defer db.Close()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	} else {
		db.Preload("Deliver_list").First(&items,truckID)

		if len(items) <= 0 {
			c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No Trucks found!"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": items})
	}
}

func getAllTrucks(c *gin.Context) {

	var items [] Trucks
	db, err := OpenDb()
	defer db.Close()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	} else {
		db.Find(&items)

		if len(items) <= 0 {
			c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No Trucks found!"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": items})
	}
}

func getOneTrucks(c *gin.Context) {
	truckID := c.Param("id")
	db, err := OpenDb()
	defer db.Close()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	} else {
		var items [] Trucks

		db.First(&items, truckID)

		if len(items) <= 0 {

			c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No Trucks found!"})

			return
		}

		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": items})
	}
}


func addTrucks(c *gin.Context) {

	var item Trucks

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


func deleteTrucks(c *gin.Context) {

	var item Trucks

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

func updateTrucks(c *gin.Context) {

	var item Trucks
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

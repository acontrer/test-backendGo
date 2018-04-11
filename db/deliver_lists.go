package db

import (
	"github.com/gin-gonic/gin"
	"net/http"

)

type Deliver_lists struct {
		Id        string    `json:"id" gorm:"column:id;primary_key;AUTO_INCREMENT "`
		Trucks_id int `json:"trucks_id" gorm:"column:trucks_id;not null"`
		Worker_id int `json:"worker_id" gorm:"column:worker_id;not null"`
		Date_init string `json:"date_init" gorm:"column:date_init;not null"`
		Date_end string `json:"date_end" gorm:"column:date_end;not null"`
		Truck Trucks `gorm:"foreignkey:Trucks_id"`
		Driver Drivers `gorm:"foreignkey:Worker_id"`

//	Profile Profile `gorm:"foreignkey:UserID;association_foreignkey:Refer"`
		}

func CrudDeliver_Lists(app *gin.RouterGroup) {

	app.GET("/", getAllDeliver_Lists)
	app.GET("/:id", getOneDeliver_Lists)
	app.POST("/", addDeliver_Lists)
	app.DELETE("/:id", deleteDeliver_Lists)
	app.PUT("/:id", updateDeliver_Lists)

}


func getAllDeliver_Lists(c *gin.Context) {

	var items [] Deliver_lists
	db, err := OpenDb()
	defer db.Close()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	} else {

		db.Preload("Truck").Preload("Driver").Find(&items)

		if len(items) <= 0 {
			c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No Deliver_Lists found!"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": items})
	}
}

func getOneDeliver_Lists(c *gin.Context) {
	truckID := c.Param("id")
	db, err := OpenDb()
	defer db.Close()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	} else {
		var items [] Deliver_lists

		db.Preload("Truck").Preload("Driver").First(&items, truckID)

		if len(items) <= 0 {

			c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No Deliver_Lists found!"})

			return
		}

		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": items})
	}
}


func addDeliver_Lists(c *gin.Context) {

	var item Deliver_lists

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


func deleteDeliver_Lists(c *gin.Context) {

	var item Deliver_lists

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

func updateDeliver_Lists(c *gin.Context) {

	var item Deliver_lists
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

package main

import (

	gin "github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	_ "github.com/acontreras/Test/db"
	"github.com/acontreras/Test/db"
)


func main() {

	router := gin.Default()

	db.CrudDrivers(router.Group("/api/v1/drivers"))
	db.CrudPackages(router.Group("/api/v1/packages"))
	db.CrudAdresses(router.Group("/api/v1/adresses"))
	db.CrudCustomers(router.Group("/api/v1/customers"))
	db.CrudDeliver_Lists(router.Group("/api/v1/deliver_lists"))
	db.CrudPurchase_Orders(router.Group("/api/v1/purchase_orders"))
	db.CrudTrucks(router.Group("/api/v1/trucks"))
	db.CrudType_Packages(router.Group("/api/v1/type_packages"))
	db.CrudType_Trucks(router.Group("/api/v1/type_trucks"))


	router.Run(":3000")
}

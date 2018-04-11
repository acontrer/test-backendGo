package db

import (
	"github.com/jinzhu/gorm"
	"fmt"
)

//var db *gorm.DB

func init() {
	//open a db connection
	var err error
	//db, err = gorm.Open("postgres", "postgres://adriancontreras@localhost:5432/transport?sslmode=disable")

	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}

	//Migrate the schema
	//	db.AutoMigrate(&Type_truck{})
}
func OpenDb()(db *gorm.DB,err error){
	return  gorm.Open("postgres", "postgres://adriancontreras@localhost:5432/transport?sslmode=disable")
}
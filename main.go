package main

import (
	"log"

	"github.com/mrb-haqee/go-crud-product/api"
	"github.com/mrb-haqee/go-crud-product/database"
	"github.com/mrb-haqee/go-crud-product/model"
)

var dbname model.DatabaseName = model.DatabaseName{
	Host:         "localhost",
	Username:     "postgres",
	Password:     "mrb28",
	DatabaseName: "crud_product",
	Port:         5432,
	Schema:       "public",
}

func main() {
	db := database.NewDb()
	conn, err := db.Connect(dbname)
	if err != nil {
		panic(err)
	}
	log.Print("berhasil connect")
	conn.AutoMigrate(&model.Product{})

	dbconn:=database.NewDbProduct(conn)
	mainAPI := api.NewAPI(dbconn)
	mainAPI.Start()
}

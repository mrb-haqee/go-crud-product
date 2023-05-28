package database

import (
	"fmt"

	"github.com/mrb-haqee/go-crud-product/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type ConnectDB struct {
	db *gorm.DB
}

func NewDb() *ConnectDB {
	return &ConnectDB{}
}

var err error

func (p *ConnectDB) Connect(dbname model.DatabaseName) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Jakarta", dbname.Host, dbname.Username, dbname.Password, dbname.DatabaseName, dbname.Port)
	dbconn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	return dbconn, err
}

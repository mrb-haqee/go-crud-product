package database

import (
	"log"

	"github.com/mrb-haqee/go-crud-product/model"
	"gorm.io/gorm"
)

type DbProduct interface {
	FindAll() ([]model.Product, error)
	FindById(id int, product *model.Product) error
	Add(product model.Product) error
	Update(id int, product model.Product) error
	Delete(id int) error
}

type dbProduct struct {
	db *gorm.DB
}

func NewDbProduct(db *gorm.DB) *dbProduct {
	return &dbProduct{db}
}

func (p *dbProduct) FindAll() ([]model.Product, error) {
	var products []model.Product
	err := p.db.Find(&products).Error
	log.Println(products)
	return products, err
}

func (p *dbProduct) FindById(id int, product *model.Product) error {
	return p.db.First(product, "id = ?", id).Error
}

func (p *dbProduct) Add(product model.Product) error {
	return p.db.Create(&product).Error
}

func (p *dbProduct) Update(id int, product model.Product) error {
	return p.db.Model(model.Product{}).Where("id = ?", id).Updates(product).Error
}

func (p *dbProduct) Delete(id int) error {
	return p.db.Where("id = ?", id).Delete(model.Product{}).Error
}

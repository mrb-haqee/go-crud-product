package model

type DatabaseName struct {
	Host         string
	Username     string
	Password     string
	DatabaseName string
	Port         int
	Schema       string
}

type Product struct {
	ID    uint `gorm:"PrimaryKey" json:"id"`
	Name  string `gorm:"type:varchar" json:"name"`
	Qty   uint64 `gorm:"type:smallint" json:"qty"`
	Harga float32 `gorm:"type:numeric(14,2)" json:"harga"`
}


type SuccesMessage struct{
	Pesan string
}


type ErrorResponse struct{
	Error string
}
package model

type Product struct {
	Id    int    `gorm:"primary_key" json:"id"`
	Name  string `json:"name"`
	Stock int    `json:"stock"`
}

type TicketTrx struct {
	Id        int    `gorm:"primary_key" json:"id"`
	IdProduct int    `json:"idProduct"`
	Name      string `json:"name"`
}

func (Product) TableName() string {
	return "product"
}

func (TicketTrx) TableName() string {
	return "ticket"
}

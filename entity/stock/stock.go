package stock

import "time"

type Product struct {
	ID             uint      `json:"id"`
	Name           string    `json:"name"`
	Category       string    `json:"category"`
	Producer       string    `json:"producer"`
	Price          uint      `json:"price"`
	Quantity       uint      `json:"qty"`
	ProductionDate time.Time `json:"production_date"`
	ExpiredDate    time.Time `json:"expired_date"`
	Timestamp
}

type Timestamp struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

type Request struct {
	Data []*Product `json:"data"`
}

package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"

	"github.com/google/uuid"
)

type Order struct {
	ID         uuid.UUID `json:"id"`
	CustomerID uuid.UUID `json:"customer_id"`
	Products   []Product `json:"products"`
	Status     status    `json:"status"`
	CreatedAt  string    `json:"-"`
	UpdatedAt  string    `json:"-"`
	DeletedAt  string    `json:"-"`
}

type status int

const (
	InProgress status = iota + 1
	Placed
	Delivered
)

func (o Order) Value() (driver.Value, error) {
	return json.Marshal(o)
}

func (o *Order) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &o)
}

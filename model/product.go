package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"

	"github.com/google/uuid"
)

type Product struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Price int       `json:"price"`
}

func (p Product) Value() (driver.Value, error) {
	return json.Marshal(p)
}

func (p *Product) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &p)
}

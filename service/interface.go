package service

import (
	"github.com/google/uuid"
	"gofr.dev/pkg/gofr"

	"github.com/srijan-27/order-service/model"
)

type Order interface {
	Create(ctx *gofr.Context, order *model.Order) (*model.Order, error)
	GetAll(ctx *gofr.Context) ([]model.Order, error)
	GetByID(ctx *gofr.Context, id uuid.UUID) (*model.Order, error)
	Update(ctx *gofr.Context, order *model.Order) (*model.Order, error)
	Delete(ctx *gofr.Context, id uuid.UUID) error
}

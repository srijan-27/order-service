package service

import (
	"github.com/google/uuid"
	"gofr.dev/pkg/gofr"

	"github.com/srijan-27/order-service/model"
	"github.com/srijan-27/order-service/store"
)

type service struct {
	store store.Order
}

// New - is a factory function to inject store in service.
func New(s store.Order) Order {
	return service{store: s}
}

func (s service) Create(ctx *gofr.Context, order *model.Order) (*model.Order, error) {
	return s.store.Create(ctx, order)
}

func (s service) GetAll(ctx *gofr.Context) ([]model.Order, error) {
	return s.store.GetAll(ctx)
}

func (s service) GetByID(ctx *gofr.Context, id uuid.UUID) (*model.Order, error) {
	return s.store.GetByID(ctx, id)
}

func (s service) Update(ctx *gofr.Context, order *model.Order) (*model.Order, error) {
	return s.store.Update(ctx, order)
}

func (s service) Delete(ctx *gofr.Context, id uuid.UUID) error {
	return s.store.Delete(ctx, id)
}

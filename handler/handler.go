package handler

import (
	"errors"
	"strings"

	"github.com/google/uuid"
	"gofr.dev/pkg/gofr"

	"github.com/srijan-27/order-service/model"
	"github.com/srijan-27/order-service/service"
)

type handler struct {
	service service.Order
}

// New - is a factory function to inject service in handler.
//
//nolint:revive // handler has to be unexported
func New(s service.Order) handler {
	return handler{service: s}
}

func (h handler) Create(ctx *gofr.Context) (interface{}, error) {
	var orders model.Order

	err := ctx.Bind(&orders)
	if err != nil {
		return nil, errors.New("invalid param: body")
	}

	resp, err := h.service.Create(ctx, &orders)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

func (h handler) GetAll(ctx *gofr.Context) (interface{}, error) {
	resp, err := h.service.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (h handler) GetByID(ctx *gofr.Context) (interface{}, error) {
	id := ctx.PathParam("id")
	if id == "" {
		return nil, errors.New("missing param: ID")
	}

	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, errors.New("invalid param: ID")
	}

	resp, err := h.service.GetByID(ctx, uid)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (h handler) Update(ctx *gofr.Context) (interface{}, error) {
	var order model.Order

	id := ctx.PathParam("id")
	if strings.TrimSpace(id) == "" {
		return nil, errors.New("missing param: ID")
	}

	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, errors.New("invalid param: ID")
	}

	order.ID = uid

	err = ctx.Bind(&order)
	if err != nil {
		return nil, errors.New("invalid param: body")
	}

	resp, err := h.service.Update(ctx, &order)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (h handler) Delete(ctx *gofr.Context) (interface{}, error) {
	id := ctx.PathParam("id")
	if id == "" {
		return nil, errors.New("missing param: ID")
	}

	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, errors.New("invalid param: ID")
	}

	_, err = h.service.Delete(ctx, uid), nil
	if err != nil {
		return nil, err
	}

	return nil, nil
}

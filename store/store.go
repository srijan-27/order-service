package store

import (
	"database/sql"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"gofr.dev/pkg/gofr"

	"github.com/srijan-27/order-service/model"
)

type store struct{}

// New is a factory function for store layer.
func New() Order {
	return store{}
}

func (s store) Create(ctx *gofr.Context, order *model.Order) (*model.Order, error) {
	const uniqueViolation = "23505"

	id, err := uuid.NewUUID()
	if err != nil {
		return nil, errors.New("error in creating uuid")
	}

	createdAt := time.Now().UTC().Format(time.RFC3339)

	_, err = ctx.SQL.ExecContext(ctx, "INSERT INTO orders (id, cust_id, products, status, created_at, updated_at) VALUES ($1,$2,$3,$4,$5,$6)",
		id, order.CustomerID, pq.Array(order.Products), order.Status, createdAt, createdAt)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			if pqErr.Code == uniqueViolation { // unique_violation
				return nil, errors.New("entity already exists")
			}
		}

		return nil, errors.New("DB error")
	}

	order.ID = id

	return order, nil
}

func (s store) GetAll(ctx *gofr.Context) ([]model.Order, error) {
	rows, err := ctx.SQL.QueryContext(ctx, "SELECT id, cust_id, products, status FROM orders WHERE deleted_at IS NULL")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var orders []model.Order

	for rows.Next() {
		var o model.Order

		err = rows.Scan(&o.ID, &o.CustomerID, pq.Array(&o.Products), &o.Status)
		if err != nil {
			return nil, err
		}

		orders = append(orders, o)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return orders, nil
}

func (s store) GetByID(ctx *gofr.Context, id uuid.UUID) (*model.Order, error) {
	var order model.Order

	err := ctx.SQL.QueryRowContext(ctx, "SELECT id, cust_id, products, status FROM orders WHERE id=$1 and deleted_at IS NULL", id).
		Scan(&order.ID, &order.CustomerID, pq.Array(&order.Products), &order.Status)

	switch {
	case err == sql.ErrNoRows:
		return nil, err
	case err != nil:
		return nil, err
	}

	return &order, nil
}

func (s store) Update(ctx *gofr.Context, order *model.Order) (*model.Order, error) {
	updatedAt := time.Now().UTC().Format(time.RFC3339)

	res, err := ctx.SQL.ExecContext(ctx, "UPDATE orders SET cust_id=$1, products=$2, status=$3, updated_at=$4 WHERE id=$5 and deleted_at IS NULL",
		order.CustomerID, pq.Array(order.Products), order.Status, updatedAt, order.ID)
	if err != nil {
		return nil, errors.New("DB error")
	}

	rowsAffected, _ := res.RowsAffected()

	if rowsAffected == 0 {
		return nil, errors.New("entity not found")
	}

	return order, nil
}

func (s store) Delete(ctx *gofr.Context, id uuid.UUID) error {
	deletedAt := time.Now().UTC().Format(time.RFC3339)
	updatedAt := deletedAt

	res, err := ctx.SQL.ExecContext(ctx, "UPDATE orders SET deleted_at=$1, updated_at=$2 WHERE id=$3 and deleted_at IS NULL", deletedAt, updatedAt, id)
	if err != nil {
		return err
	}

	rowsAffected, _ := res.RowsAffected()

	if rowsAffected == 0 {
		return err
	}

	return nil
}

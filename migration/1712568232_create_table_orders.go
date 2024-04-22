package migration

import (
	"gofr.dev/pkg/gofr/migration"
)

const createTable = `CREATE TABLE IF NOT EXISTS orders
(
    id          UUID        not null primary key,
    cust_id     UUID        not null,
    products    varchar[]   not null,
    status      varchar(10) not null,
    created_at  TIMESTAMP   not null,
    updated_at  TIMESTAMP   not null,
    deleted_at  TIMESTAMP
);`

func createTableOrders() migration.Migrate {
	return migration.Migrate{
		UP: func(d migration.Datasource) error {
			_, err := d.SQL.Exec(createTable)
			if err != nil {
				return err
			}

			return nil
		},
	}
}

package models

import (
	"context"
	"database/sql"
	"fmt"
	db "github.com/HiBang15/golang-api-base-tulpo.git/database/sqlc"
	"sync"
)

type Connector struct {
	db *sql.DB
	*db.Queries
}

var connector *Connector
var once sync.Once

func NewConnector(cnt *sql.DB) *Connector {
	once.Do(func() {
		connector = &Connector{
			db:      cnt,
			Queries: db.New(cnt),
		}
	})
	return connector
}

// execTx executes a function with transaction DB
func (conn *Connector) execTx(ctx context.Context, fn func(queries *db.Queries) error) error {
	tx, err := conn.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := db.New(tx)
	err = fn(q)
	if err != nil {
		rbErr := tx.Rollback()
		if rbErr != nil {
			return fmt.Errorf("transaction error: %v, rollback error: %v", err, rbErr)
		}
		return err
	}
	return tx.Commit()
}

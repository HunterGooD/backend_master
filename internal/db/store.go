package db

import (
	"context"
	"database/sql"
	"fmt"
)

// Store хранилище для выполения запросов в бд а так же их комбинаций в рамках транзакций
type Store struct {
	*Queries
	db *sql.DB
}

// NewStore создает новое хранилище
func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}

// execTx выполнение транзакции создает новый Queries с транзакцией и передает его в функцию обработчик
func (s *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := s.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)
	if err != nil {
		// объединить ошибки отката и транзакции
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}

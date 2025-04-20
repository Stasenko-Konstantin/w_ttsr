package repository

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// PgxDb - экземпляр [repository.Db]
type PgxDb struct {
	pool *pgxpool.Pool
	rows pgx.Rows
}

// NewPgxDb - конструктор PgxDb.
// На вход принимает [*pgxpool.Pool]
func NewPgxDb(pool *pgxpool.Pool) *PgxDb {
	return &PgxDb{pool: pool}
}

// Query - метод PgxDb для отправки запроса postgres.
// На вход принимает строку запроса query и список аргументов args.
func (p *PgxDb) Query(query string, args ...any) error {
	rows, err := p.pool.Query(context.Background(), query, args...)
	if err != nil {
		return err
	}
	p.rows = rows
	return nil
}

// Scan - метод PgxDb для чтения ответа СУБД и записи данных в [domain.Task]
func (p *PgxDb) Scan(args ...any) error {
	return p.rows.Scan(args...)
}

// Next - метод PgxDb для использования Scan в цикле на списке [domain.Task]
func (p *PgxDb) Next() bool {
	ok := p.rows.Next()
	if !ok {
		p.rows.Close()
	}
	return ok
}

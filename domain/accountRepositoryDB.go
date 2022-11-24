package domain

import (
	"context"
	"github.com/akshanshgusain/Hexagonal-Architecture/errs"
	"github.com/akshanshgusain/Hexagonal-Architecture/logger"
	"github.com/jackc/pgx/v5/pgxpool"
)

type AccountRepositoryDB struct {
	pool *pgxpool.Pool
}

func (a AccountRepositoryDB) Save(acc Account) (*Account, *errs.AppError) {
	insertSql := "insert into accounts " +
		"(customer_id, opening_date, account_type, amount, status) " +
		"values ($1,$2,$3,$4,$5)"

	err := a.pool.QueryRow(context.Background(), insertSql, acc.CustomerId,
		acc.OpeningDate, acc.AccountType, acc.Amount, acc.Status).Scan(&acc.AccountId)

	if err != nil {
		logger.Error("error creating account " + err.Error())
		return nil, errs.NewUnexpectedError("unexpected error from the database" + err.Error())
	}

	return &acc, nil
}

// Helper functions

func NewAccountRepositoryDb(pool *pgxpool.Pool) AccountRepositoryDB {
	return AccountRepositoryDB{pool: pool}
}

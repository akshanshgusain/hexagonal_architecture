package domain

import (
	"context"
	"fmt"
	"github.com/akshanshgusain/Hexagonal-Architecture/errs"
	"github.com/akshanshgusain/Hexagonal-Architecture/logger"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

type AccountRepositoryDB struct {
	pool *pgxpool.Pool
}

func (d AccountRepositoryDB) Save(acc Account) (*Account, *errs.AppError) {
	insertSql := "insert into accounts " +
		"(customer_id, opening_date, account_type, amount, status) " +
		"values ($1,$2,$3,$4,$5) returning account_id"

	err := d.pool.QueryRow(context.Background(), insertSql, acc.CustomerId,
		acc.OpeningDate, acc.AccountType, acc.Amount, acc.Status).Scan(&acc.AccountId)

	if err != nil {
		logger.Error("error creating account " + err.Error())
		return nil, errs.NewUnexpectedError("unexpected error from the database" + err.Error())
	}

	return &acc, nil
}

func (d AccountRepositoryDB) SaveTransaction(t Transaction) (*Transaction, *errs.AppError) {

	// starting the database transaction block
	tx, err := d.pool.Begin(context.Background())
	if err != nil {
		logger.Error("Error while starting a new transaction for bank account transaction: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	// inserting bank account transaction
	var transactionId string
	err = tx.QueryRow(context.Background(), "INSERT INTO transactions (account_id, amount, transaction_type, "+
		"transaction_date) values ($1, $2, $3, $4) returning transaction_id", t.AccountId, t.Amount, t.TransactionType, t.TransactionDate).Scan(&transactionId)

	// in case of error Rollback, and changes from both the tables will be reverted
	if err != nil {
		err := tx.Rollback(context.Background())
		if err != nil {
			logger.Error("Error while rolling back transaction: " + err.Error())
			return nil, errs.NewUnexpectedError("Unexpected database error")
		}
		logger.Error("Error while saving transaction: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	// updating account balance
	if t.IsWithdrawal() {
		_, err = tx.Exec(context.Background(), "UPDATE accounts SET amount = amount - $1 where account_id = $2",
			t.Amount, t.AccountId)
	} else {
		_, err = tx.Exec(context.Background(), "UPDATE accounts SET amount = amount + $1 where account_id = $2",
			t.Amount, t.AccountId)
	}

	// in case of error Rollback, and changes from both the tables will be reverted
	if err != nil {
		err := tx.Rollback(context.Background())
		if err != nil {
			logger.Error("Error while rolling back transaction: " + err.Error())
			return nil, errs.NewUnexpectedError("Unexpected database error")
		}
		logger.Error("Error while saving transaction: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}
	// commit the transaction when all is good
	err = tx.Commit(context.Background())

	if err != nil {
		err := tx.Rollback(context.Background())
		if err != nil {
			logger.Error("Error while rolling back transaction: " + err.Error())
			return nil, errs.NewUnexpectedError("Unexpected database error")
		}
		logger.Error("Error while committing transaction for bank account: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	if err != nil {
		logger.Error("Error while getting the last transaction id: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	fmt.Println("Before FindBy")
	// Getting the latest account information from the accounts table
	account, appErr := d.FindBy(t.AccountId)
	if appErr != nil {
		return nil, appErr
	}
	//t.TransactionId = strconv.FormatInt(transactionId, 10)
	t.TransactionId = transactionId

	// updating the transaction struct with the latest balance
	t.Amount = account.Amount
	return &t, nil
}

func (d AccountRepositoryDB) FindBy(accountId string) (*Account, *errs.AppError) {
	sqlGetAccount := "SELECT account_id, customer_id, opening_date, account_type, amount from accounts where account_id = $1"
	var account Account
	var dt pgtype.Date
	err := d.pool.QueryRow(context.Background(), sqlGetAccount, accountId).Scan(&account.AccountId, &account.CustomerId,
		&dt, &account.AccountType, &account.Amount)
	account.OpeningDate = dateToString(dt)
	if err != nil {
		logger.Error("Error while fetching account information: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}
	return &account, nil
}

// Helper functions

func NewAccountRepositoryDb(pool *pgxpool.Pool) AccountRepositoryDB {
	return AccountRepositoryDB{pool: pool}
}

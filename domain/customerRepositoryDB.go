package domain

import (
	"context"
	"github.com/akshanshgusain/Hexagonal-Architecture/errs"
	"github.com/akshanshgusain/Hexagonal-Architecture/logger"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
)

// Server Side Adapter

type CustomerRepositoryDB struct {
	pool *pgxpool.Pool
}

func (d CustomerRepositoryDB) FindAll() ([]Customer, *errs.AppError) {

	findAllSql := "select customer_id, name, date_of_birth ,city, zipcode, status from customers"
	rows, err := d.pool.Query(context.Background(), findAllSql)
	if err != nil {
		log.Println("Error while querying customer table " + err.Error())
		return nil, errs.NewUnexpectedError("unexpected database error")
	}
	defer rows.Close()

	// loop over rows
	var customers = make([]Customer, 0)

	for rows.Next() {
		var c Customer
		var dt pgtype.Date

		err := rows.Scan(&c.Id, &c.Name, &dt, &c.City, &c.Zipcode, &c.Status)
		if err != nil {
			log.Println("Error while scanning data from row " + err.Error())
			return nil, errs.NewUnexpectedError("unexpected database error")
		}
		c.DateOfBirth = dateToString(dt) // covert date to string
		customers = append(customers, c)
	}

	return customers, nil
}

func (d CustomerRepositoryDB) ById(id string) (*Customer, *errs.AppError) {
	var c Customer
	var dt pgtype.Date
	byIdSql := "select customer_id, name, date_of_birth ,city, zipcode, status from customers where customer_id = $1"
	err := d.pool.QueryRow(context.Background(), byIdSql, id).Scan(&c.Id, &c.Name, &dt, &c.City, &c.Zipcode, &c.Status)
	c.DateOfBirth = dateToString(dt)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, errs.NewNotFoundError("customer not found")
		}
		logger.Error("Error while scanning data from row " + err.Error())
		return nil, errs.NewUnexpectedError("unexpected database error")
	}
	return &c, nil
}

// Helper Functions

func NewCustomerRepositoryDb(pool *pgxpool.Pool) CustomerRepositoryDB {
	return CustomerRepositoryDB{pool: pool}
}

func dateToString(dt pgtype.Date) string {
	return dt.Time.String()
}

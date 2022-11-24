package domain

import (
	"context"
	"fmt"
	"github.com/akshanshgusain/Hexagonal-Architecture/errs"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"os"
)

// Server Side Adapter

type CustomerRepositoryDB struct {
	customers []Customer
	pool      *pgxpool.Pool
}

func (d CustomerRepositoryDB) FindAll() ([]Customer, error) {

	findAllSql := "select customer_id, name, date_of_birth ,city, zipcode, status from customers"
	rows, err := d.pool.Query(context.Background(), findAllSql)
	if err != nil {
		log.Println("Error while querying customer table " + err.Error())
		return nil, err
	}
	defer rows.Close()

	// loop over rows

	for rows.Next() {
		var c Customer
		var dt pgtype.Date
		err := rows.Scan(&c.Id, &c.Name, &dt, &c.City, &c.Zipcode, &c.Status)
		if err != nil {
			log.Println("Error while scanning data from row " + err.Error())
			return nil, err
		}
		c.DateOfBirth = dateToString(dt) // covert date to string
		d.customers = append(d.customers, c)
	}

	return d.customers, nil
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
		log.Println("Error while scanning data from row " + err.Error())
		return nil, errs.NewUnexpectedError("unexpected database error")
	}
	return &c, nil
}

// Helper Functions

func NewCustomerRepositoryDb() CustomerRepositoryDB {
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbAddress := os.Getenv("DB_ADDRESS")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dbSource := fmt.Sprintf("postgres://%v:%v@%v:%v/%v", dbUser, dbPassword, dbAddress, dbPort, dbName)

	//DB_USER=hello_fastapi DB_PASSWORD=hello_fastapi DB_ADDRESS=localhost DB_PORT=5432 DB_NAME=banking
	
	pool, err := pgxpool.New(context.Background(), dbSource)
	if err != nil {
		log.Println("Error connecting to DB" + err.Error())
		panic(err)
	}
	return CustomerRepositoryDB{pool: pool}
}

func dateToString(dt pgtype.Date) string {
	return dt.Time.String()
}

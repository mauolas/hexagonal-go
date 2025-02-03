package domain

import (
	"database/sql"
	"log"
	"time"

	"example.com/hexagonal/errs"
	"example.com/hexagonal/logger"
	_ "github.com/go-sql-driver/mysql"
)

type CustomerRespositoryDb struct {
	client *sql.DB
}

func (d CustomerRespositoryDb) FindAll() ([]Customer, error) {

	findAllSql := "SELECT * FROM customers"

	rows, err := d.client.Query(findAllSql)

	if err != nil {
		log.Println("Error while querying customer table" + err.Error())
	}

	var customers []Customer
	for rows.Next() {
		var c Customer
		err := rows.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateofBirth, &c.Status)
		if err != nil {
			log.Println("Error while scanning customer table" + err.Error())
		}
		customers = append(customers, c)
	}

	return customers, nil
}

func (d CustomerRespositoryDb) ById(id string) (*Customer, *errs.AppError) {
	findByIdSql := "SELECT * FROM customers WHERE customer_id = ?"

	row := d.client.QueryRow(findByIdSql, id)

	var c Customer
	err := row.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateofBirth, &c.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Customer not found")
		} else {
			log.Println("Error while scanning customer table" + err.Error())
			return nil, errs.NewUnexpectedError("Unexpected database error")
		}
	}

	return &c, nil

}

func NewCustomerRepositoryDb() CustomerRespositoryDb {
	client, err := sql.Open("mysql", "test:test@tcp(localhost:3306)/banking")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)

	return CustomerRespositoryDb{client}
}

func (d CustomerRespositoryDb) FindByStatus(status string) ([]Customer, *errs.AppError) {

	var final_status int

	if status == "active" {
		final_status = 1
	} else {
		final_status = 0
	}

	findByStatusSql := "SELECT * FROM customers WHERE status = ?"

	rows, err := d.client.Query(findByStatusSql, final_status)

	if err != nil {
		logger.Error("Error while querying customer table" + err.Error())
	}

	var customers []Customer
	for rows.Next() {
		var c Customer
		err := rows.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateofBirth, &c.Status)
		if err != nil {
			logger.Error("Error while scanning customer table" + err.Error())
		}
		customers = append(customers, c)
	}

	return customers, nil
}

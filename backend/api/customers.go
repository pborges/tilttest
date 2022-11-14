package api

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Customer struct {
	Number           int    `db:"customerNumber" json:"number"`
	Name             string `db:"customerName" json:"name"`
	ContactFirstName string `db:"contactFirstName" json:"contactFirstName"`
	ContactLastName  string `db:"contactLastName" json:"contactLastName"`
	Phone            string `db:"phone" json:"phone"`
	AddressLine1     string `db:"addressLine1" json:"addressLine1"`
}

func GetCustomers(db *sqlx.DB) ([]Customer, error) {
	customers := []Customer{}

	if err := db.Select(&customers, "select customerNumber,customerName,contactFirstName,contactLastName,phone,addressLine1 from customers"); err != nil {
		return nil, err
	}

	return customers, nil
}

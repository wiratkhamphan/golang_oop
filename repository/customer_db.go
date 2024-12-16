package repository

import "github.com/jmoiron/sqlx"

// customerRepositoryDB implements the repository layer
type customerRepositoryDB struct {
	db *sqlx.DB
}

// NewCustomerRepository initializes a new repository instance
func NewCustomerRepository(db *sqlx.DB) customerRepositoryDB {
	return customerRepositoryDB{db: db}
}

// GetAll fetches all customers from the database
func (r customerRepositoryDB) GetAll() ([]Customer, error) {
	customers := []Customer{}
	query := `
		SELECT customer_id, name, date_of_birth, city, zipcode, status
		FROM customers
	`
	err := r.db.Select(&customers, query)
	if err != nil {
		return nil, err
	}
	return customers, nil
}

// GetById fetches a customer by ID
func (r customerRepositoryDB) GetById(id int) (*Customer, error) {
	customer := Customer{}
	query := `
		SELECT customer_id, name, date_of_birth, city, zipcode, status
		FROM customers
		WHERE customer_id = ?
	`
	err := r.db.Get(&customer, query, id)
	if err != nil {
		return nil, err
	}
	return &customer, nil
}

// Save inserts a new customer into the database
func (r customerRepositoryDB) Save(customer Customer) (*Customer, error) {
	query := `
		INSERT INTO customers (name, date_of_birth, city, zipcode, status)
		VALUES (?, ?, ?, ?, ?)
	`
	result, err := r.db.Exec(
		query,
		customer.Name,
		customer.DateOfBirth,
		customer.City,
		customer.ZipCode,
		customer.Status,
	)
	if err != nil {
		return nil, err
	}

	insertedID, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	customer.CustomerID = int(insertedID)

	return &customer, nil
}

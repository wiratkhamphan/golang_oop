package repository

type Customer struct {
	CustomerID  int    `db:"customer_id"`
	Name        string `db:"name"`
	DateOfBirth string `db:"date_of_birth"` // Renamed for consistency
	City        string `db:"city"`
	ZipCode     string `db:"zipcode"`
	Status      int    `db:"status"`
}

type CustomerRepository interface {
	GetAll() ([]Customer, error)
	GetById(int) (*Customer, error)
	Save(Customer) (*Customer, error) // Renamed to reflect the purpose of the method
}

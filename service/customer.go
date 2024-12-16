package service

type CustomerResponse struct {
	CustomerID int    `json:"customer_id"`
	Name       string `json:"name"`
	Date0Birth string `json:"date_of_birth"`
	City       string `json:"city"`
	ZipCode    string `json:"zipcode"`
	Status     int    `json:"status"`
}

type Customer struct {
	CustomerID int    `json:"customer_id"`
	Name       string `json:"name"`
	Date0Birth string `json:"date_of_birth"`
	City       string `json:"city"`
	ZipCode    string `json:"zipcode"`
	Status     int    `json:"status"`
}

//go:generate mockgen -destination=../mock/mock_service/mock_customer_service.go bank/service CustomerService
type CustomerService interface {
	NewCustomers(int CustomerResponse) (*Customer, error)
	GetCustomers() ([]CustomerResponse, error)
	GetCustomer(int) (*CustomerResponse, error)
}

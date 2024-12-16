package service

import (
	"database/sql"
	"go_programming/errs"
	"go_programming/logs"
	"go_programming/repository"
	"time"
)

type customerService struct {
	custRepo repository.CustomerRepository
}

func NewCustomerService(custRepo repository.CustomerRepository) CustomerService {
	return customerService{custRepo: custRepo}
}

func (s customerService) GetCustomers() ([]CustomerResponse, error) {

	customers, err := s.custRepo.GetAll()
	if err != nil {
		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}

	custResponses := []CustomerResponse{}
	for _, customer := range customers {
		custResponse := CustomerResponse{
			CustomerID: customer.CustomerID,
			Name:       customer.Name,
			Status:     customer.Status,
		}
		custResponses = append(custResponses, custResponse)
	}

	return custResponses, nil
}

func (s customerService) GetCustomer(id int) (*CustomerResponse, error) {
	customer, err := s.custRepo.GetById(id)
	if err != nil {

		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("customer not found")
		}

		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}

	custResponse := CustomerResponse{
		CustomerID: customer.CustomerID,
		Name:       customer.Name,
		Status:     customer.Status,
	}

	return &custResponse, nil
}
func (s customerService) NewCustomers(request CustomerResponse) (*Customer, error) {
	// Map the request to the repository customer structure
	customer := repository.Customer{
		CustomerID:  request.CustomerID,
		Name:        request.Name,
		DateOfBirth: time.Now().Format("2006-01-02 15:04:05"),
		City:        request.City,
		ZipCode:     request.ZipCode,
		Status:      1,
	}

	// Use the repository to save the customer
	newCustomer, err := s.custRepo.Save(customer) // Ensure Save is implemented in the repository
	if err != nil {
		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}

	// Prepare the response object
	response := Customer{
		CustomerID: newCustomer.CustomerID,
		Name:       newCustomer.Name,
		Date0Birth: newCustomer.DateOfBirth,
		City:       newCustomer.City,
		ZipCode:    newCustomer.ZipCode,
		Status:     newCustomer.Status,
	}

	return &response, nil
}

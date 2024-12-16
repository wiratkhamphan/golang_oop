package controller

import (
	"go_programming/repository"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// CustomerController ตัวจัดการคำขอ (HTTP Request) ที่เกี่ยวกับ Customer
type CustomerController struct {
	CustomerRepository repository.CustomerRepository
}

// NewCustomerController สร้างอินสแตนซ์ของ CustomerController
func NewCustomerController(customerRepository repository.CustomerRepository) *CustomerController {
	return &CustomerController{
		CustomerRepository: customerRepository,
	}
}

// GetAllCustomers ดึงข้อมูลลูกค้าทั้งหมด
func (cc *CustomerController) GetAllCustomers(c *fiber.Ctx) error {
	customers, err := cc.CustomerRepository.GetAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"user": customers,
	})
}

// Get_by_id ดึงข้อมูลลูกค้าด้วย id
func (cc *CustomerController) Get_by_id(c *fiber.Ctx) error {
	idStr := c.Params("ID")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID format",
		})
	}

	customer, err := cc.CustomerRepository.GetById(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(customer)
}

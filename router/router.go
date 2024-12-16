package router

import (
	"go_programming/repository"
	"go_programming/router/controller"
	"go_programming/service"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func SetupRoutes(app *fiber.App, db *sqlx.DB) {
	// Create repository and service
	customerRepository := repository.NewCustomerRepository(db)
	customerService := service.NewCustomerService(customerRepository)
	// สร้าง Controller
	customerController := controller.NewCustomerController(customerRepository)
	// Define routes
	app.Get("/GetAll", customerController.GetAllCustomers)

	app.Get("/GetById/:ID", customerController.Get_by_id)

	app.Post("/AppCustomer", func(c *fiber.Ctx) error {
		var request service.CustomerResponse
		if err := c.BodyParser(&request); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid JSON body",
			})
		}

		customer, err := customerService.NewCustomers(request)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.JSON(fiber.Map{
			"message":       "Customer added successfully",
			"customer_data": customer,
		})
	})
}

package http

import (
	"go-inventory/internal/delivery/http/response"
	"go-inventory/internal/domain"

	"github.com/gofiber/fiber/v2"
)

type ProductHandler struct {
	UseCase domain.ProductUseCase
}

func NewProductHandler(app *fiber.App, us domain.ProductUseCase) {
	handler := &ProductHandler{
		UseCase: us,
	}

	api := app.Group("/api")
	api.Post("/products", handler.Create)
	api.Get("/products", handler.GetAll)
	api.Get("/products/:id", handler.GetOne)
}

type createRequest struct {
	Name  string `json:"name"`
	Price int64  `json:"price"`
	Stock int32  `json:"stock"`
}

func (h *ProductHandler) Create(c *fiber.Ctx) error {
	var req createRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(response.Error(400, "Invalid Input Format"))
	}

	res, err := h.UseCase.Create(c.Context(), req.Name, req.Price, req.Stock)
	if err != nil {
		return c.Status(500).JSON(response.Error(500, err.Error()))
	}

	return c.Status(201).JSON(response.Success(201, "Product Created Successfully", res))
}

func (h *ProductHandler) GetAll(c *fiber.Ctx) error {
	products, err := h.UseCase.GetAll(c.Context())
	if err != nil {
		return c.Status(500).JSON(response.Error(500, err.Error()))
	}

	return c.Status(200).JSON(response.Success(200, "List of Products", products))
}

func (h *ProductHandler) GetOne(c *fiber.Ctx) error {
	id := c.Params("id")
	product, err := h.UseCase.GetOne(c.Context(), id)
	if err != nil {
		return c.Status(404).JSON(response.Error(404, "Product Not Found"))
	}

	return c.Status(200).JSON(response.Success(200, "Product Detail", product))
}

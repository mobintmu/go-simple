package controller

import (
	"net/http"
	"strconv"

	"go-simple/internal/config"
	"go-simple/internal/http/response"
	"go-simple/internal/middleware"
	"go-simple/internal/product/dto"
	"go-simple/internal/product/service"

	"github.com/gin-gonic/gin"
)

type AdminProduct struct {
	Service *service.Product
}

func NewAdmin(s *service.Product) *AdminProduct {
	return &AdminProduct{Service: s}
}

func (c *AdminProduct) RegisterRoutes(rg *gin.RouterGroup, cfg *config.Config) {
	auth := middleware.JWTAuth(cfg)

	rg.POST("/", auth, c.CreateProduct)
	rg.PUT("/:id", auth, c.UpdateProduct)
	rg.DELETE("/:id", auth, c.DeleteProduct)
	rg.GET("/:id", auth, c.GetProductByID)
	rg.GET("/", auth, c.ListProducts)
}

// CreateProduct godoc
// @Summary Create a new product
// @Description Create a new product with the provided details
// @Tags Admin Products
// @Accept json
// @Produce json
// @Param product body dto.AdminCreateProductRequest true "Product to create"
// @Success 201 {object} dto.ProductResponse
// @Failure 400 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /api/v1/admin/products [post]
func (c *AdminProduct) CreateProduct(ctx *gin.Context) {
	var req dto.AdminCreateProductRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.JSONError(ctx, http.StatusBadRequest, err)
		return
	}
	product, err := c.Service.Create(ctx, req)
	if err != nil {
		response.JSONError(ctx, http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusCreated, product)
}

// UpdateProduct godoc
// @Summary Update an existing product
// @Description Update product details by ID
// @Tags Admin Products
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Param product body dto.AdminUpdateProductRequest true "Updated product details"
// @Success 200 {object} dto.ProductResponse
// @Failure 400 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /api/v1/admin/products/{id} [put]
func (c *AdminProduct) UpdateProduct(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		response.JSONError(ctx, http.StatusBadRequest, response.ErrInvalidID)
		return
	}
	var req dto.AdminUpdateProductRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.JSONError(ctx, http.StatusBadRequest, err)
		return
	}
	req.ID = int32(id)
	product, err := c.Service.Update(ctx, req)
	if err != nil {
		response.JSONError(ctx, http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, product)
}

// DeleteProduct godoc
// @Summary Delete a product by ID
// @Description Delete a product by its ID
// @Tags Admin Products
// @Param id path int true "Product ID"
// @Success 204 "No Content"
// @Failure 400 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /api/v1/admin/products/{id} [delete]
func (c *AdminProduct) DeleteProduct(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		response.JSONError(ctx, http.StatusBadRequest, response.ErrInvalidID)
		return
	}
	if err := c.Service.Delete(ctx, int32(id)); err != nil {
		response.JSONError(ctx, http.StatusInternalServerError, err)
		return
	}
	ctx.Status(http.StatusNoContent)
}

// GetProductByID godoc
// @Summary Get a product by ID
// @Description Get a product by its ID
// @Tags Admin Products
// @Param id path int true "Product ID"
// @Success 200 {object} dto.ProductResponse
// @Failure 400 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /api/v1/admin/products/{id} [get]
func (c *AdminProduct) GetProductByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		response.JSONError(ctx, http.StatusBadRequest, response.ErrInvalidID)
		return
	}

	product, err := c.Service.GetProductByID(ctx, int32(id))
	if err != nil {
		response.JSONError(ctx, http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, product)
}

// ListProducts godoc
// @Summary List all products
// @Description Get a list of all products
// @Tags Admin Products
// @Produce json
// @Success 200 {array} dto.ClientListProductsResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /api/v1/admin/products [get]
func (c *AdminProduct) ListProducts(ctx *gin.Context) {
	products, err := c.Service.ListProducts(ctx)
	if err != nil {
		response.JSONError(ctx, http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, products)
}

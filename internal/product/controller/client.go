package controller

import (
	"go-simple/internal/http/response"
	"go-simple/internal/product/dto"
	"go-simple/internal/product/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ClientProduct struct {
	Service *service.Product
}

func NewClient(s *service.Product) *ClientProduct {
	return &ClientProduct{Service: s}
}

func (c *ClientProduct) RegisterRoutes(rg *gin.RouterGroup) {
	rg.GET("/:id", c.GetProductByID)
	rg.GET("/", c.ListProducts)
}

// GetProductByID godoc
// @Summary Get a product by ID
// @Description Get a product by its ID
// @Tags Products
// @Param id path int true "Product ID"
// @Success 200 {object} dto.ProductResponse
// @Failure 400 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /api/v1/products/{id} [get]
func (c *ClientProduct) GetProductByID(ctx *gin.Context) {
	var product dto.ProductResponse
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		response.JSONError(ctx, http.StatusBadRequest, response.ErrInvalidID)
		return
	}
	product, err = c.Service.GetProductByID(ctx, int32(id))
	if err != nil {
		response.JSONError(ctx, http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, product)
}

// ListProducts godoc
// @Summary List all products
// @Description Get a list of all products
// @Tags Products
// @Produce json
// @Success 200 {array} dto.ClientListProductsResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /api/v1/products [get]
func (c *ClientProduct) ListProducts(ctx *gin.Context) {
	products, err := c.Service.ListProducts(ctx)
	if err != nil {
		response.JSONError(ctx, http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, products)
}

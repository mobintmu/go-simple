package service

import (
	"context"
	"go-simple/internal/db/sqlc"
	"go-simple/internal/product/dto"

	"go.uber.org/zap"
)

type Product struct {
	Q   *sqlc.Queries
	log *zap.Logger
}

func New(q *sqlc.Queries, log *zap.Logger) *Product {
	return &Product{
		Q:   q,
		log: log,
	}
}

func (s *Product) Create(ctx context.Context, req dto.AdminCreateProductRequest) (dto.ProductResponse, error) {
	arg := sqlc.CreateProductParams{
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		IsActive:    true,
	}
	product, err := s.Q.CreateProduct(ctx, arg)
	if err != nil {
		return dto.ProductResponse{}, err
	}
	s.log.Info("Product created", zap.Int32("id", product.ID))
	return dto.ProductResponse{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
	}, nil
}

func (s *Product) Update(ctx context.Context, req dto.AdminUpdateProductRequest) (dto.ProductResponse, error) {
	arg := sqlc.UpdateProductParams{
		ID:          int32(req.ID),
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		IsActive:    req.IsActive,
	}
	product, err := s.Q.UpdateProduct(ctx, arg)
	if err != nil {
		return dto.ProductResponse{}, err
	}
	return dto.ProductResponse{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
	}, nil
}

func (s *Product) Delete(ctx context.Context, id int32) error {
	return s.Q.DeleteProduct(ctx, id)
}

func (s *Product) GetProductByID(ctx context.Context, id int32) (dto.ProductResponse, error) {
	product, err := s.Q.GetProduct(ctx, id)
	if err != nil {
		return dto.ProductResponse{}, err
	}
	return dto.ProductResponse{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
	}, nil
}

func (s *Product) ListProducts(ctx context.Context) (dto.ClientListProductsResponse, error) {
	products, err := s.Q.ListProducts(ctx)
	if err != nil {
		return nil, err
	}
	var resp []dto.ProductResponse = make([]dto.ProductResponse, 0, len(products))
	for _, product := range products {
		resp = append(resp, dto.ProductResponse{
			ID:          product.ID,
			Name:        product.Name,
			Description: product.Description,
			Price:       product.Price,
		})
	}
	return resp, nil
}

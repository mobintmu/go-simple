package service

import (
	"context"
	"go-simple/internal/config"
	"go-simple/internal/product/dto"
	"go-simple/internal/storage/cache"
	"go-simple/internal/storage/sql/sqlc"

	"go.uber.org/zap"
)

type Product struct {
	query  *sqlc.Queries
	log    *zap.Logger
	memory *cache.Store
	cfg    *config.Config
}

func New(q *sqlc.Queries,
	log *zap.Logger,
	memory *cache.Store,
	cfg *config.Config) *Product {
	return &Product{
		query:  q,
		log:    log,
		memory: memory,
		cfg:    cfg,
	}
}

func (s *Product) Create(ctx context.Context, req dto.AdminCreateProductRequest) (dto.ProductResponse, error) {
	arg := sqlc.CreateProductParams{
		ProductName:        req.Name,
		ProductDescription: req.Description,
		Price:              req.Price,
		IsActive:           true,
	}
	product, err := s.query.CreateProduct(ctx, arg)
	if err != nil {
		return dto.ProductResponse{}, err
	}
	s.log.Info("Product created", zap.Int32("id", product.ID))
	s.memory.Set(ctx, s.memory.KeyProduct(product.ID), product, s.cfg.Redis.DefaultTTL)
	s.memory.Delete(ctx, s.memory.KeyAllProducts())
	return dto.ProductResponse{
		ID:          product.ID,
		Name:        product.ProductName,
		Description: product.ProductDescription,
		Price:       product.Price,
	}, nil
}

func (s *Product) Update(ctx context.Context, req dto.AdminUpdateProductRequest) (dto.ProductResponse, error) {
	arg := sqlc.UpdateProductParams{
		ID:                 int32(req.ID),
		ProductName:        req.Name,
		ProductDescription: req.Description,
		Price:              req.Price,
		IsActive:           req.IsActive,
	}
	product, err := s.query.UpdateProduct(ctx, arg)
	if err != nil {
		return dto.ProductResponse{}, err
	}
	s.memory.Set(ctx, s.memory.KeyProduct(product.ID), product, s.cfg.Redis.DefaultTTL)
	s.memory.Delete(ctx, s.memory.KeyAllProducts())
	return dto.ProductResponse{
		ID:          product.ID,
		Name:        product.ProductName,
		Description: product.ProductDescription,
		Price:       product.Price,
	}, nil
}

func (s *Product) Delete(ctx context.Context, id int32) error {
	s.memory.Delete(ctx, s.memory.KeyProduct(id))
	s.memory.Delete(ctx, s.memory.KeyAllProducts())
	return s.query.DeleteProduct(ctx, id)
}

func (s *Product) GetProductByID(ctx context.Context, id int32) (dto.ProductResponse, error) {
	var product sqlc.Product
	err := s.memory.Get(ctx, s.memory.KeyProduct(id), &product)
	if err != nil {
		product, err = s.query.GetProduct(ctx, id)
		if err != nil {
			return dto.ProductResponse{}, err
		}
		s.memory.Set(ctx, s.memory.KeyProduct(product.ID), product, s.cfg.Redis.DefaultTTL)
	}
	result := dto.ProductResponse{
		ID:          product.ID,
		Name:        product.ProductName,
		Description: product.ProductDescription,
		Price:       product.Price,
	}
	return result, nil
}

func (s *Product) ListProducts(ctx context.Context) (dto.ClientListProductsResponse, error) {
	var resp []dto.ProductResponse
	if err := s.memory.Get(ctx, s.memory.KeyAllProducts(), &resp); err == nil {
		return resp, nil
	}
	products, err := s.query.ListProducts(ctx)
	if err != nil {
		return nil, err
	}
	resp = make([]dto.ProductResponse, 0, len(products))
	for _, product := range products {
		resp = append(resp, dto.ProductResponse{
			ID:          product.ID,
			Name:        product.ProductName,
			Description: product.ProductDescription,
			Price:       product.Price,
		})
	}
	s.memory.Set(ctx, s.memory.KeyAllProducts(), resp, s.cfg.Redis.DefaultTTL)
	return resp, nil
}

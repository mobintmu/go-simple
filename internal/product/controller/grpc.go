package controller

import (
	"context"

	pb "go-simple/api/proto/product/v1"
	"go-simple/internal/product/dto"
	"go-simple/internal/product/service"

	"google.golang.org/protobuf/types/known/emptypb"
)

type ProductGRPC struct {
	pb.UnimplementedProductServiceServer
	svc *service.Product
}

func NewGRPC(svc *service.Product) pb.ProductServiceServer {
	return &ProductGRPC{
		svc: svc,
	}
}

func (h *ProductGRPC) GetProductByID(ctx context.Context, req *pb.ProductRequest) (*pb.ProductResponse, error) {
	p, err := h.svc.GetProductByID(ctx, req.Id)
	return &pb.ProductResponse{
		Id:          p.ID,
		Name:        p.Name,
		Description: p.Description,
		Price:       p.Price,
	}, err
}

func (h *ProductGRPC) CreateProduct(ctx context.Context, req *pb.CreateProductRequest) (*pb.ProductResponse, error) {
	p, err := h.svc.Create(ctx, dto.AdminCreateProductRequest{
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
	})
	if err != nil {
		return nil, err
	}
	return &pb.ProductResponse{
		Id:          p.ID,
		Name:        p.Name,
		Description: p.Description,
		Price:       p.Price,
	}, nil
}

func (h *ProductGRPC) UpdateProduct(ctx context.Context, req *pb.UpdateProductRequest) (*pb.ProductResponse, error) {
	p, err := h.svc.Update(ctx, dto.AdminUpdateProductRequest{
		ID:          req.Id,
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		IsActive:    req.IsActive,
	})
	if err != nil {
		return nil, err
	}
	return &pb.ProductResponse{
		Id:          p.ID,
		Name:        p.Name,
		Description: p.Description,
		Price:       p.Price,
	}, nil
}

func (h *ProductGRPC) DeleteProduct(ctx context.Context, req *pb.DeleteProductRequest) (*pb.ProductResponse, error) {
	err := h.svc.Delete(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.ProductResponse{Id: req.Id}, nil
}

func (h *ProductGRPC) ListProducts(ctx context.Context, _ *emptypb.Empty) (*pb.ListProductsResponse, error) {
	products, err := h.svc.ListProducts(ctx)
	if err != nil {
		return nil, err
	}
	var resp pb.ListProductsResponse
	for _, p := range products {
		resp.Products = append(resp.Products, &pb.ProductResponse{
			Id:          p.ID,
			Name:        p.Name,
			Description: p.Description,
			Price:       p.Price,
		})
	}
	return &resp, nil
}

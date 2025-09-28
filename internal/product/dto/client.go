package dto

type ProductResponse struct {
	ID          int32  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int64  `json:"price"`
}

type ClientListProductsResponse []ProductResponse

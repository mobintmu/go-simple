package dto

type AdminCreateProductRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int64  `json:"price"`
	IsActive    bool   `json:"is_active"`
}

type AdminUpdateProductRequest struct {
	ID          int32  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int64  `json:"price"`
	IsActive    bool   `json:"is_active"`
}

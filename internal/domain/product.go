package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
)

// Entity
type Product struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Price     int64     `json:"price"`
	Stock     int32     `json:"stock"`
	CreatedAt time.Time `json:"created_at"`
}

// Repository Interface
// UseCase tidak tahu datanya dari postgres, mongo atau file
type ProductRepository interface {
	Store(ctx context.Context, p *Product) (*Product, error)
	Fetch(ctx context.Context) ([]Product, error)
	GetByID(ctx context.Context, id uuid.UUID) (*Product, error)
}

// UseCase Interface
type ProductUseCase interface {
	Create(ctx context.Context, name string, price int64, stock int32) (*Product, error)
	GetAll(ctx context.Context) ([]Product, error)
	GetOne(ctx context.Context, idStr string) (*Product, error)
}

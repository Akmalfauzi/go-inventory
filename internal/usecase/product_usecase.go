package usecase

import (
	"context"
	"errors"
	"go-inventory/internal/domain"
	"time"

	"github.com/google/uuid"
)

type productUseCase struct {
	repo           domain.ProductRepository
	contextTimeout time.Duration
}

func NewProductUseCase(r domain.ProductRepository, timeout time.Duration) domain.ProductUseCase {
	return &productUseCase{
		repo:           r,
		contextTimeout: timeout,
	}
}

func (u *productUseCase) Create(c context.Context, name string, price int64, stock int32) (*domain.Product, error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	// Validasi harga
	if price < 0 {
		return nil, errors.New("price cannot be negative")
	}

	p := &domain.Product{
		Name:  name,
		Price: price,
		Stock: stock,
	}

	return u.repo.Store(ctx, p)
}

func (u *productUseCase) GetAll(c context.Context) ([]domain.Product, error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	return u.repo.Fetch(ctx)
}

func (u *productUseCase) GetOne(c context.Context, idStr string) (*domain.Product, error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	id, err := uuid.Parse(idStr)
	if err != nil {
		return nil, errors.New("invalid UUID format")
	}

	return u.repo.GetByID(ctx, id)
}

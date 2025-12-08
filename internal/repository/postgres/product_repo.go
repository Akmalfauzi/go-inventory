package postgres

import (
	"context"

	"go-inventory/internal/db"
	"go-inventory/internal/domain"

	"github.com/google/uuid"
)

type productRepo struct {
	queries *db.Queries
}

func NewProductRepository(q *db.Queries) domain.ProductRepository {
	return &productRepo{queries: q}
}

func (r *productRepo) Store(ctx context.Context, p *domain.Product) (*domain.Product, error) {
	// createProductParams sekarang mengharapkan uuid.UUID standar (bukan pgtype)
	// Jadi kita tidak perlu konversi aneh-aneh.

	// Catatan: created_at di DB digenerate otomatis, tapi sqlc return value-nya ada.
	res, err := r.queries.CreateProduct(ctx, db.CreateProductParams{
		Name:  p.Name,
		Price: p.Price,
		Stock: p.Stock,
	})
	if err != nil {
		return nil, err
	}

	return &domain.Product{
		ID:        res.ID, // Sekarang ini sudah compatible
		Name:      res.Name,
		Price:     res.Price,
		Stock:     res.Stock,
		CreatedAt: res.CreatedAt, // Ini juga sudah compatible (time.Time)
	}, nil
}

func (r *productRepo) Fetch(ctx context.Context) ([]domain.Product, error) {
	rows, err := r.queries.ListProducts(ctx)
	if err != nil {
		return nil, err
	}

	products := make([]domain.Product, len(rows))
	for i, row := range rows {
		products[i] = domain.Product{
			ID:        row.ID,
			Name:      row.Name,
			Price:     row.Price,
			Stock:     row.Stock,
			CreatedAt: row.CreatedAt,
		}
	}
	return products, nil
}

func (r *productRepo) GetByID(ctx context.Context, id uuid.UUID) (*domain.Product, error) {
	// sqlc sekarang menerima parameter uuid.UUID langsung
	res, err := r.queries.GetProduct(ctx, id)
	if err != nil {
		return nil, err
	}

	return &domain.Product{
		ID:        res.ID,
		Name:      res.Name,
		Price:     res.Price,
		Stock:     res.Stock,
		CreatedAt: res.CreatedAt,
	}, nil
}

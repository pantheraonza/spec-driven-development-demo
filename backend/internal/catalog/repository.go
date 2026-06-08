package catalog

import (
	"database/sql"
	"fmt"
)

// Repository handles product data access
type Repository interface {
	GetAll() ([]Product, error)
	GetByID(id string) (*Product, error)
	GetActiveProducts() ([]Product, error)
}

// SQLiteRepository implements Repository for SQLite
type SQLiteRepository struct {
	db *sql.DB
}

// NewSQLiteRepository creates a new repository
func NewSQLiteRepository(db *sql.DB) *SQLiteRepository {
	return &SQLiteRepository{db: db}
}

// GetAll retrieves all products
func (r *SQLiteRepository) GetAll() ([]Product, error) {
	rows, err := r.db.Query(`
		SELECT id, name, description, price_amount, currency, stock, active
		FROM products
		ORDER BY name ASC
	`)
	if err != nil {
		return nil, fmt.Errorf("failed to query products: %w", err)
	}
	defer rows.Close()

	var products []Product
	for rows.Next() {
		var p Product
		err := rows.Scan(&p.ID, &p.Name, &p.Description, &p.Price, &p.Currency, &p.Stock, &p.Active)
		if err != nil {
			return nil, fmt.Errorf("failed to scan product: %w", err)
		}
		products = append(products, p)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating products: %w", err)
	}

	return products, nil
}

// GetByID retrieves a product by ID
func (r *SQLiteRepository) GetByID(id string) (*Product, error) {
	var p Product
	err := r.db.QueryRow(`
		SELECT id, name, description, price_amount, currency, stock, active
		FROM products
		WHERE id = ?
	`, id).Scan(&p.ID, &p.Name, &p.Description, &p.Price, &p.Currency, &p.Stock, &p.Active)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("failed to query product: %w", err)
	}

	return &p, nil
}

// GetActiveProducts retrieves only active products
func (r *SQLiteRepository) GetActiveProducts() ([]Product, error) {
	rows, err := r.db.Query(`
		SELECT id, name, description, price_amount, currency, stock, active
		FROM products
		WHERE active = true
		ORDER BY name ASC
	`)
	if err != nil {
		return nil, fmt.Errorf("failed to query active products: %w", err)
	}
	defer rows.Close()

	var products []Product
	for rows.Next() {
		var p Product
		err := rows.Scan(&p.ID, &p.Name, &p.Description, &p.Price, &p.Currency, &p.Stock, &p.Active)
		if err != nil {
			return nil, fmt.Errorf("failed to scan product: %w", err)
		}
		products = append(products, p)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating products: %w", err)
	}

	return products, nil
}

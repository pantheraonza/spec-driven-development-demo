package catalog

import "fmt"

// Product represents a product in the catalog
type Product struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Currency    string  `json:"currency"`
	Stock       int     `json:"stock"`
	Active      bool    `json:"active"`
}

// IsAvailable checks if product is available for purchase
func (p *Product) IsAvailable() bool {
	return p.Active && p.Stock > 0
}

// CanAddToCart checks if requested quantity can be added
func (p *Product) CanAddToCart(quantity int) error {
	if !p.Active {
		return fmt.Errorf("product is inactive")
	}
	if p.Stock == 0 {
		return fmt.Errorf("product out of stock")
	}
	if quantity > p.Stock {
		return fmt.Errorf("insufficient stock")
	}
	return nil
}

// GetCatalogResponse wraps products for API response
type GetCatalogResponse struct {
	Products []Product `json:"products"`
}

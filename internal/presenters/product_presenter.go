package presenters

import (
	"context"
	"errors"
	"simple-goroutine-product/internal/models"
	"simple-goroutine-product/internal/repositories"
	"time"
)

// ProductPresenter interface for business logic
type ProductPresenter interface {
	CreateProduct(ctx context.Context, req models.ProductRequest) (*models.ProductResponse, error)
	GetProduct(ctx context.Context, id uint) (*models.ProductResponse, error)
	GetProducts(ctx context.Context, page, limit int) ([]models.ProductResponse, int64, error)
	UpdateProduct(ctx context.Context, id uint, req models.ProductRequest) (*models.ProductResponse, error)
	DeleteProduct(ctx context.Context, id uint) error
}

// productPresenter implements ProductPresenter
type productPresenter struct {
	productRepo repositories.ProductRepository
}

// NewProductPresenter creates a new product presenter
func NewProductPresenter(productRepo repositories.ProductRepository) ProductPresenter {
	return &productPresenter{
		productRepo: productRepo,
	}
}

// CreateProduct creates a new product using goroutine
func (p *productPresenter) CreateProduct(ctx context.Context, req models.ProductRequest) (*models.ProductResponse, error) {
	// Channel to receive result from goroutine
	resultChan := make(chan struct {
		product *models.Product
		err     error
	}, 1)

	// Execute create operation in goroutine
	go func() {
		product := &models.Product{
			Name:        req.Name,
			Description: req.Description,
			Price:       req.Price,
			Stock:       req.Stock,
		}

		err := p.productRepo.Create(product)
		resultChan <- struct {
			product *models.Product
			err     error
		}{product: product, err: err}
	}()

	// Wait for result with timeout
	select {
	case result := <-resultChan:
		if result.err != nil {
			return nil, result.err
		}
		response := result.product.ToResponse()
		return &response, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	case <-time.After(30 * time.Second):
		return nil, errors.New("operation timeout")
	}
}

// GetProduct gets a product by ID
func (p *productPresenter) GetProduct(ctx context.Context, id uint) (*models.ProductResponse, error) {
	product, err := p.productRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	if product == nil {
		return nil, nil
	}

	response := product.ToResponse()
	return &response, nil
}

// GetProducts gets all products with pagination
func (p *productPresenter) GetProducts(ctx context.Context, page, limit int) ([]models.ProductResponse, int64, error) {
	products, total, err := p.productRepo.GetAll(page, limit)
	if err != nil {
		return nil, 0, err
	}

	var responses []models.ProductResponse
	for _, product := range products {
		responses = append(responses, product.ToResponse())
	}

	return responses, total, nil
}

// UpdateProduct updates a product using goroutine
func (p *productPresenter) UpdateProduct(ctx context.Context, id uint, req models.ProductRequest) (*models.ProductResponse, error) {
	// Channel to receive result from goroutine
	resultChan := make(chan struct {
		product *models.Product
		err     error
	}, 1)

	// Execute update operation in goroutine
	go func() {
		// First get the existing product
		product, err := p.productRepo.GetByID(id)
		if err != nil {
			resultChan <- struct {
				product *models.Product
				err     error
			}{product: nil, err: err}
			return
		}

		// Update fields
		product.Name = req.Name
		product.Description = req.Description
		product.Price = req.Price
		product.Stock = req.Stock

		// Save updated product
		err = p.productRepo.Update(product)
		resultChan <- struct {
			product *models.Product
			err     error
		}{product: product, err: err}
	}()

	// Wait for result with timeout
	select {
	case result := <-resultChan:
		if result.err != nil {
			return nil, result.err
		}
		response := result.product.ToResponse()
		return &response, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	case <-time.After(30 * time.Second):
		return nil, errors.New("operation timeout")
	}
}

// DeleteProduct deletes a product
func (p *productPresenter) DeleteProduct(ctx context.Context, id uint) error {
	return p.productRepo.Delete(id)
}

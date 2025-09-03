package presenters

import (
	"context"
	"simple-goroutine-product/internal/models"
	"testing"
	"time"
)

// SimpleProductRepository is a simple mock implementation
type SimpleProductRepository struct {
	products []models.Product
	nextID   uint
}

func NewSimpleProductRepository() *SimpleProductRepository {
	return &SimpleProductRepository{
		products: make([]models.Product, 0),
		nextID:   1,
	}
}

func (r *SimpleProductRepository) Create(product *models.Product) error {
	product.ID = r.nextID
	product.CreatedAt = time.Now()
	product.UpdatedAt = time.Now()
	r.products = append(r.products, *product)
	r.nextID++
	return nil
}

func (r *SimpleProductRepository) GetByID(id uint) (*models.Product, error) {
	for _, product := range r.products {
		if product.ID == id {
			return &product, nil
		}
	}
	return nil, nil
}

func (r *SimpleProductRepository) GetAll(page, limit int) ([]models.Product, int64, error) {
	return r.products, int64(len(r.products)), nil
}

func (r *SimpleProductRepository) Update(product *models.Product) error {
	for i, p := range r.products {
		if p.ID == product.ID {
			product.UpdatedAt = time.Now()
			r.products[i] = *product
			return nil
		}
	}
	return nil
}

func (r *SimpleProductRepository) Delete(id uint) error {
	for i, product := range r.products {
		if product.ID == id {
			r.products = append(r.products[:i], r.products[i+1:]...)
			return nil
		}
	}
	return nil
}

func TestSimpleProductPresenter_CreateProduct(t *testing.T) {
	repo := NewSimpleProductRepository()
	presenter := NewProductPresenter(repo)

	req := models.ProductRequest{
		Name:        "Test Product",
		Description: "Test Description",
		Price:       99.99,
		Stock:       10,
	}

	ctx := context.Background()
	result, err := presenter.CreateProduct(ctx, req)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if result == nil {
		t.Error("Expected result, got nil")
	}

	if result.Name != req.Name {
		t.Errorf("Expected name %s, got %s", req.Name, result.Name)
	}

	if result.Price != req.Price {
		t.Errorf("Expected price %f, got %f", req.Price, result.Price)
	}
}

func TestSimpleProductPresenter_GetProduct(t *testing.T) {
	repo := NewSimpleProductRepository()
	presenter := NewProductPresenter(repo)

	// First create a product
	req := models.ProductRequest{
		Name:        "Test Product",
		Description: "Test Description",
		Price:       99.99,
		Stock:       10,
	}

	ctx := context.Background()
	created, err := presenter.CreateProduct(ctx, req)
	if err != nil {
		t.Fatalf("Failed to create product: %v", err)
	}

	// Now get the product
	result, err := presenter.GetProduct(ctx, created.ID)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if result == nil {
		t.Error("Expected result, got nil")
	}

	if result.ID != created.ID {
		t.Errorf("Expected ID %d, got %d", created.ID, result.ID)
	}
}

func TestSimpleProductPresenter_UpdateProduct(t *testing.T) {
	repo := NewSimpleProductRepository()
	presenter := NewProductPresenter(repo)

	// First create a product
	createReq := models.ProductRequest{
		Name:        "Original Product",
		Description: "Original Description",
		Price:       50.00,
		Stock:       5,
	}

	ctx := context.Background()
	created, err := presenter.CreateProduct(ctx, createReq)
	if err != nil {
		t.Fatalf("Failed to create product: %v", err)
	}

	// Update the product
	updateReq := models.ProductRequest{
		Name:        "Updated Product",
		Description: "Updated Description",
		Price:       99.99,
		Stock:       10,
	}

	result, err := presenter.UpdateProduct(ctx, created.ID, updateReq)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if result == nil {
		t.Error("Expected result, got nil")
	}

	if result.Name != updateReq.Name {
		t.Errorf("Expected name %s, got %s", updateReq.Name, result.Name)
	}

	if result.Price != updateReq.Price {
		t.Errorf("Expected price %f, got %f", updateReq.Price, result.Price)
	}
}

func TestSimpleProductPresenter_DeleteProduct(t *testing.T) {
	repo := NewSimpleProductRepository()
	presenter := NewProductPresenter(repo)

	// First create a product
	req := models.ProductRequest{
		Name:        "Test Product",
		Description: "Test Description",
		Price:       99.99,
		Stock:       10,
	}

	ctx := context.Background()
	created, err := presenter.CreateProduct(ctx, req)
	if err != nil {
		t.Fatalf("Failed to create product: %v", err)
	}

	// Delete the product
	err = presenter.DeleteProduct(ctx, created.ID)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// Verify product is deleted
	result, err := presenter.GetProduct(ctx, created.ID)
	if err != nil {
		t.Errorf("Expected no error when getting deleted product, got %v", err)
	}

	// For a deleted product, we expect nil result or an empty product
	// Since our simple repository returns nil for non-existent products
	if result != nil {
		t.Error("Expected nil result for deleted product")
	}
}

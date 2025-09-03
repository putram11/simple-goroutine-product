package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"simple-goroutine-product/internal/models"
	"simple-goroutine-product/internal/validators"
	"testing"

	"github.com/labstack/echo/v4"
)

// SimpleProductPresenter is a simple mock implementation
type SimpleProductPresenter struct {
	products []models.ProductResponse
	nextID   uint
}

func NewSimpleProductPresenter() *SimpleProductPresenter {
	return &SimpleProductPresenter{
		products: make([]models.ProductResponse, 0),
		nextID:   1,
	}
}

func (p *SimpleProductPresenter) CreateProduct(ctx context.Context, req models.ProductRequest) (*models.ProductResponse, error) {
	product := &models.ProductResponse{
		ID:          p.nextID,
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Stock:       req.Stock,
	}
	p.products = append(p.products, *product)
	p.nextID++
	return product, nil
}

func (p *SimpleProductPresenter) GetProduct(ctx context.Context, id uint) (*models.ProductResponse, error) {
	for _, product := range p.products {
		if product.ID == id {
			return &product, nil
		}
	}
	return nil, nil
}

func (p *SimpleProductPresenter) GetProducts(ctx context.Context, page, limit int) ([]models.ProductResponse, int64, error) {
	return p.products, int64(len(p.products)), nil
}

func (p *SimpleProductPresenter) UpdateProduct(ctx context.Context, id uint, req models.ProductRequest) (*models.ProductResponse, error) {
	for i, product := range p.products {
		if product.ID == id {
			p.products[i].Name = req.Name
			p.products[i].Description = req.Description
			p.products[i].Price = req.Price
			p.products[i].Stock = req.Stock
			return &p.products[i], nil
		}
	}
	return nil, nil
}

func (p *SimpleProductPresenter) DeleteProduct(ctx context.Context, id uint) error {
	for i, product := range p.products {
		if product.ID == id {
			p.products = append(p.products[:i], p.products[i+1:]...)
			return nil
		}
	}
	return nil
}

func TestSimpleProductHandler_CreateProduct(t *testing.T) {
	presenter := NewSimpleProductPresenter()
	handler := NewProductHandler(presenter)

	req := models.ProductRequest{
		Name:        "Test Product",
		Description: "Test Description",
		Price:       99.99,
		Stock:       10,
	}

	// Setup Echo
	e := echo.New()
	e.Validator = validators.NewValidator()
	reqBody, _ := json.Marshal(req)
	httpReq := httptest.NewRequest(http.MethodPost, "/products", bytes.NewBuffer(reqBody))
	httpReq.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(httpReq, rec)

	// Test
	err := handler.CreateProduct(c)

	// Assertions
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if rec.Code != http.StatusCreated {
		t.Errorf("Expected status code %d, got %d", http.StatusCreated, rec.Code)
	}

	var response models.ProductResponse
	json.Unmarshal(rec.Body.Bytes(), &response)

	if response.Name != req.Name {
		t.Errorf("Expected name %s, got %s", req.Name, response.Name)
	}

	if response.Price != req.Price {
		t.Errorf("Expected price %f, got %f", req.Price, response.Price)
	}
}

func TestSimpleProductHandler_GetProduct(t *testing.T) {
	presenter := NewSimpleProductPresenter()
	handler := NewProductHandler(presenter)

	// First create a product
	createReq := models.ProductRequest{
		Name:        "Test Product",
		Description: "Test Description",
		Price:       99.99,
		Stock:       10,
	}

	ctx := context.Background()
	created, err := presenter.CreateProduct(ctx, createReq)
	if err != nil {
		t.Fatalf("Failed to create product: %v", err)
	}

	// Setup Echo
	e := echo.New()
	httpReq := httptest.NewRequest(http.MethodGet, "/products/1", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(httpReq, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	// Test
	err = handler.GetProduct(c)

	// Assertions
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if rec.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, rec.Code)
	}

	var response models.ProductResponse
	json.Unmarshal(rec.Body.Bytes(), &response)

	if response.ID != created.ID {
		t.Errorf("Expected ID %d, got %d", created.ID, response.ID)
	}

	if response.Name != created.Name {
		t.Errorf("Expected name %s, got %s", created.Name, response.Name)
	}
}

func TestSimpleProductHandler_UpdateProduct(t *testing.T) {
	presenter := NewSimpleProductPresenter()
	handler := NewProductHandler(presenter)

	// First create a product
	createReq := models.ProductRequest{
		Name:        "Original Product",
		Description: "Original Description",
		Price:       50.00,
		Stock:       5,
	}

	ctx := context.Background()
	_, err := presenter.CreateProduct(ctx, createReq)
	if err != nil {
		t.Fatalf("Failed to create product: %v", err)
	}

	// Update request
	updateReq := models.ProductRequest{
		Name:        "Updated Product",
		Description: "Updated Description",
		Price:       149.99,
		Stock:       15,
	}

	// Setup Echo
	e := echo.New()
	e.Validator = validators.NewValidator()
	reqBody, _ := json.Marshal(updateReq)
	httpReq := httptest.NewRequest(http.MethodPut, "/products/1", bytes.NewBuffer(reqBody))
	httpReq.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(httpReq, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	// Test
	err = handler.UpdateProduct(c)

	// Assertions
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if rec.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, rec.Code)
	}

	var response models.ProductResponse
	json.Unmarshal(rec.Body.Bytes(), &response)

	if response.Name != updateReq.Name {
		t.Errorf("Expected name %s, got %s", updateReq.Name, response.Name)
	}

	if response.Price != updateReq.Price {
		t.Errorf("Expected price %f, got %f", updateReq.Price, response.Price)
	}
}

func TestSimpleProductHandler_DeleteProduct(t *testing.T) {
	presenter := NewSimpleProductPresenter()
	handler := NewProductHandler(presenter)

	// First create a product
	createReq := models.ProductRequest{
		Name:        "Test Product",
		Description: "Test Description",
		Price:       99.99,
		Stock:       10,
	}

	ctx := context.Background()
	_, err := presenter.CreateProduct(ctx, createReq)
	if err != nil {
		t.Fatalf("Failed to create product: %v", err)
	}

	// Setup Echo
	e := echo.New()
	httpReq := httptest.NewRequest(http.MethodDelete, "/products/1", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(httpReq, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	// Test
	err = handler.DeleteProduct(c)

	// Assertions
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if rec.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, rec.Code)
	}

	var response map[string]string
	json.Unmarshal(rec.Body.Bytes(), &response)

	expectedMessage := "Product deleted successfully"
	if response["message"] != expectedMessage {
		t.Errorf("Expected message %s, got %s", expectedMessage, response["message"])
	}

	// Verify product is deleted
	result, err := presenter.GetProduct(ctx, uint(1))
	if err != nil {
		t.Errorf("Expected no error when getting deleted product, got %v", err)
	}

	if result != nil {
		t.Error("Expected nil result for deleted product")
	}
}

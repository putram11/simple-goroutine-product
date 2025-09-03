package presenters

import (
	"context"
	"simple-goroutine-product/internal/models"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockProductRepository is a mock implementation of ProductRepository
type MockProductRepository struct {
	mock.Mock
}

func (m *MockProductRepository) Create(product *models.Product) error {
	args := m.Called(product)
	return args.Error(0)
}

func (m *MockProductRepository) GetByID(id uint) (*models.Product, error) {
	args := m.Called(id)
	return args.Get(0).(*models.Product), args.Error(1)
}

func (m *MockProductRepository) GetAll(page, limit int) ([]models.Product, int64, error) {
	args := m.Called(page, limit)
	return args.Get(0).([]models.Product), args.Get(1).(int64), args.Error(2)
}

func (m *MockProductRepository) Update(product *models.Product) error {
	args := m.Called(product)
	return args.Error(0)
}

func (m *MockProductRepository) Delete(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func TestProductPresenter_CreateProduct(t *testing.T) {
	mockRepo := new(MockProductRepository)
	presenter := NewProductPresenter(mockRepo)

	mockRepo.On("Create", mock.AnythingOfType("*models.Product")).Return(nil).Run(func(args mock.Arguments) {
		arg := args.Get(0).(*models.Product)
		arg.ID = 1
		arg.CreatedAt = time.Now()
		arg.UpdatedAt = time.Now()
	})

	req := models.ProductRequest{
		Name:        "Test Product",
		Description: "Test Description",
		Price:       99.99,
		Stock:       10,
	}

	ctx := context.Background()
	result, err := presenter.CreateProduct(ctx, req)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, req.Name, result.Name)
	assert.Equal(t, req.Price, result.Price)
	mockRepo.AssertExpectations(t)
}

func TestProductPresenter_GetProduct(t *testing.T) {
	mockRepo := new(MockProductRepository)
	presenter := NewProductPresenter(mockRepo)

	product := &models.Product{
		ID:          1,
		Name:        "Test Product",
		Description: "Test Description",
		Price:       99.99,
		Stock:       10,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	mockRepo.On("GetByID", uint(1)).Return(product, nil)

	ctx := context.Background()
	result, err := presenter.GetProduct(ctx, 1)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, product.ID, result.ID)
	assert.Equal(t, product.Name, result.Name)
	mockRepo.AssertExpectations(t)
}

func TestProductPresenter_UpdateProduct(t *testing.T) {
	mockRepo := new(MockProductRepository)
	presenter := NewProductPresenter(mockRepo)

	existingProduct := &models.Product{
		ID:          1,
		Name:        "Old Product",
		Description: "Old Description",
		Price:       50.00,
		Stock:       5,
		CreatedAt:   time.Now().Add(-time.Hour),
		UpdatedAt:   time.Now().Add(-time.Hour),
	}

	mockRepo.On("GetByID", uint(1)).Return(existingProduct, nil)
	mockRepo.On("Update", mock.AnythingOfType("*models.Product")).Return(nil)

	req := models.ProductRequest{
		Name:        "Updated Product",
		Description: "Updated Description",
		Price:       99.99,
		Stock:       10,
	}

	ctx := context.Background()
	result, err := presenter.UpdateProduct(ctx, 1, req)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, req.Name, result.Name)
	assert.Equal(t, req.Price, result.Price)
	mockRepo.AssertExpectations(t)
}

func TestProductPresenter_DeleteProduct(t *testing.T) {
	mockRepo := new(MockProductRepository)
	presenter := NewProductPresenter(mockRepo)

	mockRepo.On("Delete", uint(1)).Return(nil)

	ctx := context.Background()
	err := presenter.DeleteProduct(ctx, 1)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestProductPresenter_GetProducts(t *testing.T) {
	mockRepo := new(MockProductRepository)
	presenter := NewProductPresenter(mockRepo)

	products := []models.Product{
		{
			ID:          1,
			Name:        "Product 1",
			Description: "Description 1",
			Price:       99.99,
			Stock:       10,
		},
		{
			ID:          2,
			Name:        "Product 2",
			Description: "Description 2",
			Price:       149.99,
			Stock:       5,
		},
	}

	mockRepo.On("GetAll", 1, 10).Return(products, int64(2), nil)

	ctx := context.Background()
	result, total, err := presenter.GetProducts(ctx, 1, 10)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, int64(2), total)
	assert.Len(t, result, 2)
	mockRepo.AssertExpectations(t)
}

package repositories

import (
	"simple-goroutine-product/internal/models"

	"gorm.io/gorm"
)

// ProductRepository interface for product data operations
type ProductRepository interface {
	Create(product *models.Product) error
	GetByID(id uint) (*models.Product, error)
	GetAll(page, limit int) ([]models.Product, int64, error)
	Update(product *models.Product) error
	Delete(id uint) error
}

// productRepository implements ProductRepository
type productRepository struct {
	db *gorm.DB
}

// NewProductRepository creates a new product repository
func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db: db}
}

// Create creates a new product
func (r *productRepository) Create(product *models.Product) error {
	return r.db.Create(product).Error
}

// GetByID gets a product by ID
func (r *productRepository) GetByID(id uint) (*models.Product, error) {
	var product models.Product
	err := r.db.First(&product, id).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}

// GetAll gets all products with pagination
func (r *productRepository) GetAll(page, limit int) ([]models.Product, int64, error) {
	var products []models.Product
	var total int64

	// Count total records
	r.db.Model(&models.Product{}).Count(&total)

	// Get paginated records
	offset := (page - 1) * limit
	err := r.db.Offset(offset).Limit(limit).Find(&products).Error

	return products, total, err
}

// Update updates a product
func (r *productRepository) Update(product *models.Product) error {
	return r.db.Save(product).Error
}

// Delete soft deletes a product
func (r *productRepository) Delete(id uint) error {
	return r.db.Delete(&models.Product{}, id).Error
}

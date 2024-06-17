package database

import (
	"fmt"
	"stock_manager_back-end/models"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

// ############################### PRODUCTS #######################################

// CreateProduct creates a new product in the database
func CreateProduct(name, function string) error {
	product := models.Product{Name: name, Function: function}

	// Validate the product before creating
	if err := validate.Struct(product); err != nil {
		return fmt.Errorf("validation error: %w", err)
	}

	if err := DB.Create(&product).Error; err != nil {
		return fmt.Errorf("failed to create product: %w", err)
	}
	return nil
}

// GetProducts retrieves all products from the database
func GetProducts() ([]models.Product, error) {
	var products []models.Product
	if err := DB.Find(&products).Error; err != nil {
		return nil, fmt.Errorf("failed to retrieve products: %w", err)
	}
	return products, nil
}

// GetProductByID retrieves a product by its ID
func GetProductByID(id uint) (*models.Product, error) {
	var product models.Product
	if err := DB.First(&product, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to retrieve product: %w", err)
	}
	return &product, nil
}

// UpdateProduct updates an existing product in the database
func UpdateProduct(id uint, updatedProduct models.Product) error {
	// Validate the updated product before updating
	if err := validate.Struct(updatedProduct); err != nil {
		return fmt.Errorf("validation error: %w", err)
	}

	if err := DB.Model(&models.Product{}).Where("id = ?", id).Updates(updatedProduct).Error; err != nil {
		return fmt.Errorf("failed to update product: %w", err)
	}
	return nil
}

// DeleteProduct deletes a product by its ID from the database
func DeleteProduct(id uint) error {
	if err := DB.Delete(&models.Product{}, id).Error; err != nil {
		return fmt.Errorf("failed to delete product: %w", err)
	}
	return nil
}

// ############################### SITES #######################################

// CreateSite creates a new site in the database
func CreateSite(name string) error {
	site := models.Site{Name: name}
	return DB.Create(&site).Error
}

// GetSites retrieves all sites from the database
func GetSites() ([]models.Site, error) {
	var sites []models.Site
	err := DB.Find(&sites).Error
	return sites, err
}

// GetSiteByID retrieves a site by its ID
func GetSiteByID(id uint) (*models.Site, error) {
	var site models.Site
	err := DB.First(&site, id).Error
	if err != nil {
		return nil, err
	}
	return &site, nil
}

// ############################## STOCKS ######################################

// CreateStock creates a new stock entry in the database
func CreateStock(productID, siteID uint, quantity float64, unitOfMeasure string) error {
	stock := models.Stock{ProductID: productID, SiteID: siteID, Quantity: quantity, UnitOfMeasure: unitOfMeasure}
	return DB.Create(&stock).Error
}

// GetStocks retrieves all stock entries from the database
func GetStocks() ([]models.Stock, error) {
	var stocks []models.Stock
	err := DB.Preload("Product").Preload("Site").Find(&stocks).Error
	return stocks, err
}

// GetStockByID retrieves a stock entry by its ID
func GetStockByID(id uint) (*models.Stock, error) {
	var stock models.Stock
	err := DB.Preload("Product").Preload("Site").First(&stock, id).Error
	if err != nil {
		return nil, err
	}
	return &stock, nil
}

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

	// Validate the site before creating
	if err := validate.Struct(site); err != nil {
		return fmt.Errorf("validation error: %w", err)
	}

	if err := DB.Create(&site).Error; err != nil {
		return fmt.Errorf("failed to create site: %w", err)
	}
	return nil
}

// GetSites retrieves all sites from the database
func GetSites() ([]models.Site, error) {
	var sites []models.Site
	if err := DB.Find(&sites).Error; err != nil {
		return nil, fmt.Errorf("failed to retrieve sites: %w", err)
	}
	return sites, nil
}

// GetSiteByID retrieves a site by its ID
func GetSiteByID(id uint) (*models.Site, error) {
	var site models.Site
	if err := DB.First(&site, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to retrieve site: %w", err)
	}
	return &site, nil
}

// UpdateSite updates an existing site in the database
func UpdateSite(id uint, updatedSite models.Site) error {
	// Validate the updated site before updating
	if err := validate.Struct(updatedSite); err != nil {
		return fmt.Errorf("validation error: %w", err)
	}

	if err := DB.Model(&models.Site{}).Where("id = ?", id).Updates(updatedSite).Error; err != nil {
		return fmt.Errorf("failed to update site: %w", err)
	}
	return nil
}

// DeleteSite deletes a site by its ID from the database
func DeleteSite(id uint) error {
	if err := DB.Delete(&models.Site{}, id).Error; err != nil {
		return fmt.Errorf("failed to delete site: %w", err)
	}
	return nil
}

// ############################### STOCKS #######################################

// CreateStock creates a new stock entry in the database
func CreateStock(productID, siteID uint, quantity float64, unitOfMeasure string) error {
	stock := models.Stock{ProductID: productID, SiteID: siteID, Quantity: quantity, UnitOfMeasure: unitOfMeasure}

	// Validate the stock before creating
	if err := validate.Struct(stock); err != nil {
		return fmt.Errorf("validation error: %w", err)
	}

	if err := DB.Create(&stock).Error; err != nil {
		return fmt.Errorf("failed to create stock: %w", err)
	}
	return nil
}

// GetStocks retrieves all stock entries from the database
func GetStocks() ([]models.Stock, error) {
	var stocks []models.Stock
	if err := DB.Preload("Product").Preload("Site").Find(&stocks).Error; err != nil {
		return nil, fmt.Errorf("failed to retrieve stocks: %w", err)
	}
	return stocks, nil
}

// GetStockByID retrieves a stock entry by its ID
func GetStockByID(id uint) (*models.Stock, error) {
	var stock models.Stock
	if err := DB.Preload("Product").Preload("Site").First(&stock, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to retrieve stock: %w", err)
	}
	return &stock, nil
}

// UpdateStock updates an existing stock entry in the database
func UpdateStock(id uint, updatedStock models.Stock) error {
	// Validate the updated stock before updating
	if err := validate.Struct(updatedStock); err != nil {
		return fmt.Errorf("validation error: %w", err)
	}

	if err := DB.Model(&models.Stock{}).Where("id = ?", id).Updates(updatedStock).Error; err != nil {
		return fmt.Errorf("failed to update stock: %w", err)
	}
	return nil
}

// DeleteStock deletes a stock entry by its ID from the database
func DeleteStock(id uint) error {
	if err := DB.Delete(&models.Stock{}, id).Error; err != nil {
		return fmt.Errorf("failed to delete stock: %w", err)
	}
	return nil
}

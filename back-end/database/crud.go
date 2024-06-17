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

// CreateProduct creates a new product in the database.
//
// It takes in the name and function of the product as parameters.
// The name and function parameters are both strings.
//
// The function returns an error if there was a validation error or if there was a failure to create the product.
// The error is of type error.
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

// GetProducts retrieves all products from the database.
//
// It returns a slice of models.Product and an error. If there was an error
// retrieving the products, the error will contain a detailed message.
func GetProducts() ([]models.Product, error) {
	var products []models.Product
	if err := DB.Find(&products).Error; err != nil {
		return nil, fmt.Errorf("failed to retrieve products: %w", err)
	}
	return products, nil
}

// GetProductByID retrieves a product by its ID.
//
// It takes an ID of type uint as a parameter.
// It returns a pointer to a models.Product and an error.
// If the product is not found, it returns nil and nil.
// If there is an error retrieving the product, it returns nil and an error with a detailed message.
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

// UpdateProduct updates an existing product in the database.
//
// It takes the ID of the product to update of type uint and the updated product details of type models.Product as parameters.
// It returns an error if there was a validation error or if there was a failure to update the product.
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

// DeleteProduct deletes a product by its ID from the database.
//
// It takes the ID of the product to delete of type uint as a parameter.
// It returns an error if there was a failure to delete the product.
func DeleteProduct(id uint) error {
	if err := DB.Delete(&models.Product{}, id).Error; err != nil {
		return fmt.Errorf("failed to delete product: %w", err)
	}
	return nil
}

// ############################### SITES #######################################

// CreateSite creates a new site in the database.
//
// It takes in the name of the site as a parameter.
// The name parameter is a string.
//
// The function returns an error if there was a validation error or if there was a failure to create the site.
// The error is of type error.
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

// GetSites retrieves all sites from the database.
//
// It returns a slice of models.Site and an error. If there was an error retrieving the sites, the error will contain a detailed message.
//
// Returns:
// - sites: a slice of models.Site containing all the sites in the database
// - error: an error if there was a failure to retrieve the sites, otherwise nil
func GetSites() ([]models.Site, error) {
	var sites []models.Site
	if err := DB.Find(&sites).Error; err != nil {
		return nil, fmt.Errorf("failed to retrieve sites: %w", err)
	}
	return sites, nil
}

// GetSiteByID retrieves a site by its ID.
//
// It takes an ID of type uint as a parameter.
// It returns a pointer to a models.Site and an error.
// If the site is not found, it returns nil and nil.
// If there is an error retrieving the site, it returns nil and an error with a detailed message.
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

// UpdateSite updates an existing site in the database.
//
// It takes an ID of type uint and an updated site of type models.Site as parameters.
// The ID parameter is the ID of the site to be updated.
// The updatedSite parameter contains the updated values for the site.
//
// The function validates the updated site using the validate.Struct function.
// If the validation fails, it returns an error with a detailed message.
//
// The function then updates the site in the database using the DB.Model and DB.Updates functions.
// If the update fails, it returns an error with a detailed message.
//
// If the update is successful, it returns nil.
//
// Example usage:
//
// siteID := uint(1)
// updatedSite := models.Site{Name: "New Site Name", Description: "New Description"}
// err := UpdateSite(siteID, updatedSite)
//
//	if err != nil {
//	    // Handle error
//	}
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

// DeleteSite deletes a site by its ID from the database.
//
// It takes an ID of type uint as a parameter.
// It returns an error if there was a failure to delete the site.
func DeleteSite(id uint) error {
	if err := DB.Delete(&models.Site{}, id).Error; err != nil {
		return fmt.Errorf("failed to delete site: %w", err)
	}
	return nil
}

// ############################### STOCKS #######################################

// CreateStock creates a new stock entry in the database.
//
// It takes the product ID, site ID, quantity, and unit of measure as parameters.
// The product ID and site ID are of type uint, the quantity is of type float64, and the unit of measure is of type string.
//
// The function returns an error if there was a validation error or if there was a failure to create the stock.
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

// GetStocks retrieves all stock entries from the database.
//
// It does not take any parameters.
// It returns a slice of models.Stock and an error.
func GetStocks() ([]models.Stock, error) {
	var stocks []models.Stock
	if err := DB.Preload("Product").Preload("Site").Find(&stocks).Error; err != nil {
		return nil, fmt.Errorf("failed to retrieve stocks: %w", err)
	}
	return stocks, nil
}

// GetStockByID retrieves a stock entry by its ID.
//
// It takes an ID of type uint as a parameter.
// It returns a pointer to a models.Stock and an error.
// If the stock is not found, it returns nil and nil.
// If there is an error retrieving the stock, it returns nil and an error with a detailed message.
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

// UpdateStock updates an existing stock entry in the database.
//
// It takes an ID of type uint and an updated stock of type models.Stock as parameters.
// The ID parameter is the ID of the stock to be updated.
// The updatedStock parameter contains the updated values for the stock.
//
// The function validates the updated stock using the validate.Struct function.
// If the validation fails, it returns an error with a detailed message.
//
// The function then updates the stock in the database using the DB.Model and DB.Updates functions.
// If the update fails, it returns an error with a detailed message.
//
// If the update is successful, it returns nil.
//
// Example usage:
//
// id := uint(1)
// updatedStock := models.Stock{Name: "New Stock Name", Quantity: 10}
// err := UpdateStock(id, updatedStock)
//
//	if err != nil {
//	    // Handle error
//	}
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

// DeleteStock deletes a stock entry by its ID from the database.
//
// It takes an ID of type uint as a parameter.
// It returns an error if there was a failure to delete the stock.
func DeleteStock(id uint) error {
	if err := DB.Delete(&models.Stock{}, id).Error; err != nil {
		return fmt.Errorf("failed to delete stock: %w", err)
	}
	return nil
}

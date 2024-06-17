package database

import (
	"stock_manager_back-end/models"
)

// CreateProduct creates a new product in the database
func CreateProduct(name, function string) error {
	product := models.Product{Name: name, Function: function}
	return DB.Create(&product).Error
}

// GetProducts retrieves all products from the database
func GetProducts() ([]models.Product, error) {
	var products []models.Product
	err := DB.Find(&products).Error
	return products, err
}

// GetProductByID retrieves a product by its ID
func GetProductByID(id uint) (*models.Product, error) {
	var product models.Product
	err := DB.First(&product, id).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}

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

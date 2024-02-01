package database

import (
	"database/sql"

	"github.com/michelbahia/ecomerce/internal/entity"
)

type ProductDB struct {
	db *sql.DB
}

func NewProductDB(db *sql.DB) *ProductDB {
	return &ProductDB{db: db}
}

func (cd *ProductDB) GetProducts() ([]*entity.Product, error) {
	rows, err := cd.db.Query("SELECT id, name, price, category_id FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*entity.Product
	for rows.Next() {
		var product entity.Product
		if err := rows.Scan(&product.ID, &product.Name, &product.Price, &product.CategoryID); err != nil {
			return nil, err
		}
		products = append(products, &product)
	}
	return products, nil
}

func (cd *ProductDB) GetProduct(id string) (*entity.Product, error) {
	var product entity.Product
	err := cd.db.QueryRow("SELECT id, name, price, category_id, imageURL FROM products WHERE id = ?", id).Scan(&product.ID, &product.Name, &product.Price, &product.CategoryID)
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (pd *ProductDB) GetProductByCategoryID(categoryID string) ([]*entity.Product, error) {
	rows, err := pd.db.Query("SELECT id, name, price, category_id, imageURL FROM products WHERE category_id = ?", categoryID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*entity.Product
	for rows.Next() {
		var product entity.Product
		if err := rows.Scan(&product.ID, &product.Name,&product.Description, &product.Price, &product.CategoryID, &product.ImageURL); err != nil {
			return nil, err
		}
		products = append(products, &product)
	}
	return products, nil
}

func (pd *ProductDB) CreateProduct(product *entity.Product) (*entity.Product, error) {
	_, err := pd.db.Exec("INSERT INTO  products (id, name, price, description, category_id, image_url) VALUES (?, ?, ?, ?, ?)", 
	product.ID, product.Name, product.Description, product.CategoryID, product.ImageURL)
	if err != nil {
		return nil, err
	}
	return product, nil
}
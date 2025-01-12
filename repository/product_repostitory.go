package repository

import (
	"database/sql"
	"fmt"
	"products/model"
)

type ProductRepository struct {
	connection *sql.DB
}

func NewProductRepository(c *sql.DB) ProductRepository {
	return ProductRepository{
		connection: c,
	}
}

func (pr *ProductRepository) GetProducts() ([]model.Product, error) {
	query := "SELECT id, product_name, price FROM product"

	rows, err := pr.connection.Query(query)

	if err != nil {
		fmt.Println(err)
		return []model.Product{}, err
	}

	var productList []model.Product
	var productObj model.Product

	for rows.Next() {
		err = rows.Scan(
			&productObj.ID,
			&productObj.Name,
			&productObj.Price,
		)

		if err != nil {
			fmt.Println(err)
			return []model.Product{}, err
		}

		productList = append(productList, productObj)

	}
	rows.Close()

	return productList, nil

}

func (pr *ProductRepository) CreateProduct(p model.Product) (int, error) {

	var id int
	query, err := pr.connection.Prepare("INSERT INTO product" +
		"(product_name, price) VALUES($1, $2) RETURNING id")

	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	err = query.QueryRow(p.Name, p.Price).Scan(&id)

	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	query.Close()

	return id, nil

}

func (pr *ProductRepository) GetProductById(id int) (*model.Product, error) {
	query, err := pr.connection.Prepare("SELECT * FROM product WHERE id = $1")

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var p model.Product

	err = query.QueryRow(id).Scan(
		&p.ID,
		&p.Name,
		&p.Price,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil

		}
		return nil, err

	}

	query.Close()

	return &p, nil
}

package data

import "github.com/pedroreiscoder/product-rest-api/models"

func GetProducts() ([]models.Product, error) {
	rows, err := db.Query("select id, name, price from products")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var products []models.Product

	for rows.Next() {
		var product models.Product

		err := rows.Scan(&product.ID, &product.Name, &product.Price)

		if err != nil {
			return nil, err
		}

		products = append(products, product)
	}

	return products, nil
}

func GetProduct(id uint64) (models.Product, error) {
	row := db.QueryRow("select name, price from products where id=$1", id)

	product := models.Product{ID: id}

	err := row.Scan(&product.Name, &product.Price)

	if err != nil {
		return models.Product{}, err
	}

	return product, nil
}

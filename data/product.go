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

func CreateProduct(product *models.Product) error {
	err := db.QueryRow("insert into products (name, price) values ($1, $2) returning id",
		product.Name, product.Price).Scan(&product.ID)

	return err
}

func UpdateProduct(id uint64, product models.Product) (int64, error) {
	result, err := db.Exec("update products set name=$1, price=$2 where id=$3",
		product.Name, product.Price, id)

	if err != nil {
		return 0, err
	}

	return result.RowsAffected()
}

func DeleteProduct(id uint64) (int64, error) {
	result, err := db.Exec("delete from products where id=$1", id)

	if err != nil {
		return 0, err
	}

	return result.RowsAffected()
}

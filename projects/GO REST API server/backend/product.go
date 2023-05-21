package backend

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type product struct {
	ID          int     `json:"id"`
	ProductCode string  `'json:productCode`
	Name        string  `'json:name`
	Inventory   int     `json:inventory`
	Price       float32 `json:price`
	Status      string  `json:status`
}

func getProducts(db *sql.DB) ([]product, error) {
	rows, err := db.Query("select ID, Inventory, Name, Price, ProductCode, Status from Products")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	products := []product{}

	for rows.Next() {
		var p product
		if err := rows.Scan(&p.ID, &p.Inventory, &p.Name, &p.Price, &p.ProductCode, &p.Status); err != nil {
			return nil, err
		}

		products = append(products, p)
	}

	return products, nil
}

func (p *product) getProduct(db *sql.DB) error {
	return db.QueryRow("Select ID, Inventory, Name, Price, ProductCode, Status from Products Where ID=?", p.ID).Scan(&p.ID, &p.Inventory, &p.Name, &p.Price, &p.ProductCode, &p.Status)
}

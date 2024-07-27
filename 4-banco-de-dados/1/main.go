// Banco de dados
package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" // o _ é para importar o pacote sem usar
	"github.com/google/uuid"
)

type Product struct {
	ID    string
	Name  string
	Price float64
}

func newProduct(name string, price float64) *Product {
	return &Product{
		ID:    uuid.New().String(),
		Name:  name,
		Price: price,
	}
}

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3307)/goexpert")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	p := newProduct("camiseta", 29.99)
	err = insertProduct(db, p)
	if err != nil {
		panic(err)
	}

	// p, err = selecProduct(db, p.ID)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("Produto %v possui o preço de %v\n", p.Name, p.Price)
	products, err := selectAllProducts(db)
	if err != nil {
		panic(err)
	}
	for _, p := range products {
		fmt.Printf("Produto %v possui o preço de %v\n", p.Name, p.Price)
	}
}

func insertProduct(db *sql.DB, p *Product) error {
	stmt, err := db.Prepare("insert into products (id, name, price) values (?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(p.ID, p.Name, p.Price)
	if err != nil {
		return err
	}

	p.Price = 39.99
	err = updateProduct(db, p)
	if err != nil {
		panic(err)
	}

	return nil
}

func updateProduct(db *sql.DB, p *Product) error {
	stmt, err := db.Prepare("update products set name = ?, price = ? where id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(p.Name, p.Price, p.ID)
	if err != nil {
		return err
	}
	return nil
}

func selecProduct(db *sql.DB, id string) (*Product, error) {
	stmt, err := db.Prepare("select id, name, price from products where id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	var p Product
	err = stmt.QueryRow(id).Scan(&p.ID, &p.Name, &p.Price) // ordem dos campos tem que ser igual a ordem do select
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func selectAllProducts(db *sql.DB) ([]Product, error) {
	rows, err := db.Query("select id, name, price from products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var products []Product
	for rows.Next() {
		var p Product
		err = rows.Scan(&p.ID, &p.Name, &p.Price)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return products, nil
}

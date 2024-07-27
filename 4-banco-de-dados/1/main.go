// Banco de dados
package main

import (
	"database/sql"

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
	return nil
}

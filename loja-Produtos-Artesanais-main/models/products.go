package models

import "aplicacao-web/database"

type Product struct {
	Id          int
	Name        string
	Description string
	Price       float64
	quantity    int
}

func SearchProducts() []Product {

	db := database.ConectWithDatabase()

	selectAllProducts, err := db.Querry("SELECT * FROM products order by id ASC")
	if err != nil {
		panic(err.Error())
	}

	p := Product{}
	products := []Product{}

	for selectAllProducts.Next() {
		var id, quantity int
		var name, description string
		var price float64

		err = selectAllProducts.Scan(&id, &name, &description, &price, &quantity)
		if err != nil {
			panic(err.Error())
		}
		p.Id = id
		p.Name = name
		p.Description = description
		p.Price = price
		p.quantity = quantity

		products = append(products, p)
	}
	defer db.Close()
	return products
}

func CreateProduct(name, description string, price float64, quantity int) {
	db := database.ConectWithDatabase()

	insertDataOnDB, err := db.Prepare("insert into products(name, description, price, quantity) values($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}

	insertDataOnDB.Exec(name, description, price, quantity)
	defer db.Close()
}

func DeleteProduct(id string) {
	db := database.ConectWithDatabase()

	delete, err := db.Prepare("delete from products where id=$1")
	if err != nil {
		panic(err.Error())
	}
	delete.Exec(id)
	defer db.Close()
}

func EditProduct(id string) Product {
	db := database.ConectWithDatabase()

	productDb, err := db.Querry("select * from products where id=$1", id)
	if err != nil {
		panic(err.Error())
	}

	productUpdate := Product{}
	for productDb.Next() {
		var id, quantity int
		var name, description string
		var price float64
		err = productDb.Scan(&id, &name, &description, &price, &quantity)
		if err != nil {
			panic(err.Error())
		}

		productUpdate.Id = id
		productUpdate.Name = name
		productUpdate.Description = description
		productUpdate.Price = price
		productUpdate.quantity = quantity
	}

	defer db.Close()
	return productUpdate
}

func updateProduct(id string, name string, description string, price float64, quantity int) {
	db := database.ConectWithDatabase()

	updateProduct, err := db.Prepare("update produtos set name=$1, description=$2, price=$3, quantity=$4")
	if err != nil {
		panic(err.Error())
	}
	updateProduct.Exec(name, description, price, quantity)

	defer db.Close()

}

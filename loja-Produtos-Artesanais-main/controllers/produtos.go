package controllers

import (
	"aplicacao-web/models"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	products := models.BuscarProdutos()
	temp.ExecuteTemplate(w, "Index", produtos)
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "CreateProduct", nil)
}

func insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("name")
		description := r.FormValue("description")
		price := r.FormValue("price")
		quantity := r.FormValue("quantity")

		priceConv, err := strconv.ParseFloat(price, 64)
		if err != nil {
			log.Println("Error on quantity conversion", err)
		}
		models.CreateProduct(name, description, priceConv, quantityConv)
	}
	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	productid := r.URL.Query().Get("id")
	models.DeleteProduct(productid)
	http.Redirect(w, r, "/", 301)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		name := r.FormValue("name")
		description := r.FormValue("description")
		price := r.FormValue("price")
		quantity := r.FormValue("quantity")

		idConv, err := strconv.Atoi(id)
		if err != nil {
			log.Println("Error converting Id to int", err)
		}

		priceConv, err := strconv.ParseFloat(price, 64)
		if err != nil {
			log.Println("Error converting Price to float", err)
		}

		quantityConv, err := strconv.Atoi(quantity)
		if err != nil {
			log.Println("Error converting Quantity to int", err)
		}

		models.UpdateProduct(idConv, name, description, priceConv, quantityConv)

	}
}

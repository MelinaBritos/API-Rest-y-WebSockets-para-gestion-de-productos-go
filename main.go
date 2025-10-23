package main

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	"github.com/MelinaBritos/API-REST-y-WebSockets-para-gestion-de-productos/Database"
	"github.com/MelinaBritos/API-REST-y-WebSockets-para-gestion-de-productos/Handler"
	"github.com/MelinaBritos/API-REST-y-WebSockets-para-gestion-de-productos/Repository"
	"github.com/MelinaBritos/API-REST-y-WebSockets-para-gestion-de-productos/Service"
	_ "github.com/lib/pq"
)

func main() {

	db := Database.Conexiondb()
	Database.CrearTablas(db)

	// Inyección de dependencias crear funcion
	productRepository := Repository.NewProductRepository(db)
	productService := Service.NewProductService(productRepository)
	productHandler := Handler.NewProductHandler(productService)

	categoryRepository := Repository.NewCategoryRepository(db)
	categoryService := Service.NewCategoryService(categoryRepository)
	categoryHandler := Handler.NewCategoryHandler(categoryService)

	// Crear router
	r := mux.NewRouter()

	port, err := CargarPuerto()
	if err != nil {
		println(err.Error())
	}

	// Rutas productos
	r.HandleFunc("/api/products", productHandler.GetProducts).Methods("GET")
	r.HandleFunc("/api/products/{id}", productHandler.GetProductByID).Methods("GET")
	r.HandleFunc("/api/products", productHandler.CreateProduct).Methods("POST")
	r.HandleFunc("/api/products/{id}", productHandler.UpdateProduct).Methods("PUT")
	r.HandleFunc("/api/products/{id}", productHandler.DeleteProduct).Methods("DELETE")

	// Rutas categorías
	r.HandleFunc("/api/categories", categoryHandler.GetCategories).Methods("GET")
	r.HandleFunc("/api/categories", categoryHandler.CreateCategory).Methods("POST")
	r.HandleFunc("/api/categories/{id}", categoryHandler.UpdateCategory).Methods("PUT")
	r.HandleFunc("/api/categories/{id}", categoryHandler.DeleteCategory).Methods("DELETE")

	// Ruta de historial de productos
	r.HandleFunc("/api/products/{id}/history", productHandler.GetProductHistory).Methods("GET")

	http.ListenAndServe(":"+port, r)
}

func CargarPuerto() (string, error) {

	err := godotenv.Load(".env.example")
	if err != nil {
		return os.Getenv("PORT"), err
	}
	return os.Getenv("PORT"), nil
}

package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"

	handler "github.com/MelinaBritos/API-REST-y-WebSockets-para-gestion-de-productos/Handler"
	repository "github.com/MelinaBritos/API-REST-y-WebSockets-para-gestion-de-productos/Repository"
	service "github.com/MelinaBritos/API-REST-y-WebSockets-para-gestion-de-productos/Service"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	//conexion a PostgreSQL
	DSN, err := ObtenerDSN()
	if err != nil {
		log.Fatal(err.Error())
	}

	db, err := sql.Open("postgres", DSN)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()

	//Crear tablas
	query := `CREATE TABLE IF NOT EXISTS Products(id SERIAL PRIMARY KEY, name varchar(100) NOT NULL, description varchar(500),
	 	price NUMERIC(10, 2) NOT NULL, stock int NOT NULL, created_at TIMESTAMPTZ, updated_at TIMESTAMPTZ);`

	_, err = db.Exec(query)
	if err != nil {
		log.Fatal(err.Error())
	}

	// Inyección de dependencias
	productRepository := repository.NewProductRepository(db)
	productService := service.NewProductService(productRepository)
	productHandler := handler.NewProductHandler(productService)

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

	http.ListenAndServe(":"+port, r)
}

func ObtenerDSN() (string, error) {

	err := godotenv.Load(".env.example")
	if err != nil {
		return os.Getenv("DSN"), err
	}
	return os.Getenv("DSN"), nil

}

func CargarPuerto() (string, error) {

	err := godotenv.Load(".env.example")
	if err != nil {
		return os.Getenv("PORT"), err
	}
	return os.Getenv("PORT"), nil
}

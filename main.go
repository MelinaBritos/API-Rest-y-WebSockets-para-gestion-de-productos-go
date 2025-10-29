package main

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	"github.com/MelinaBritos/API-REST-y-WebSockets-para-gestion-de-productos/Database"
	"github.com/MelinaBritos/API-REST-y-WebSockets-para-gestion-de-productos/Handler"
	"github.com/MelinaBritos/API-REST-y-WebSockets-para-gestion-de-productos/Middleware"
	"github.com/MelinaBritos/API-REST-y-WebSockets-para-gestion-de-productos/Repository"
	"github.com/MelinaBritos/API-REST-y-WebSockets-para-gestion-de-productos/Service"
	"github.com/MelinaBritos/API-REST-y-WebSockets-para-gestion-de-productos/WebSocket"
	_ "github.com/lib/pq"
)

func main() {

	// Conexión a la base de datos
	db := Database.Conexiondb()
	Database.CrearTablas(db)

	// Inyección de dependencias
	productRepository := Repository.NewProductRepository(db)
	productService := Service.NewProductService(productRepository)
	productHandler := Handler.NewProductHandler(productService)

	categoryRepository := Repository.NewCategoryRepository(db)
	categoryService := Service.NewCategoryService(categoryRepository)
	categoryHandler := Handler.NewCategoryHandler(categoryService)

	userRepository := Repository.NewUserRepository(db)
	userService := Service.NewUserService(userRepository)
	userHandler := Handler.NewUserHandler(userService)

	searchRepository := Repository.NewSearchRepository(db)
	searchService := Service.NewSearchService(searchRepository)
	searchHandler := Handler.NewSearchHandler(searchService)

	// Crear router
	r := mux.NewRouter()

	port, err := CargarPuerto()
	if err != nil {
		println(err.Error())
	}

	// Rutas productos
	r.HandleFunc("/api/products", Middleware.SetMiddlewareAuthentication(productHandler.GetProducts)).Methods("GET")
	r.HandleFunc("/api/products/{id}", Middleware.SetMiddlewareAuthentication(productHandler.GetProductByID)).Methods("GET")
	r.HandleFunc("/api/products", Middleware.SetMiddlewareAuthentication(Middleware.RequireAdmin(productHandler.CreateProduct))).Methods("POST")
	r.HandleFunc("/api/products/lote", Middleware.SetMiddlewareAuthentication(Middleware.RequireAdmin(productHandler.CreateProducts))).Methods("POST")
	r.HandleFunc("/api/products/{id}", Middleware.SetMiddlewareAuthentication(Middleware.RequireAdmin(productHandler.UpdateProduct))).Methods("PUT")
	r.HandleFunc("/api/products/{id}", Middleware.SetMiddlewareAuthentication(Middleware.RequireAdmin(productHandler.DeleteProduct))).Methods("DELETE")

	// Rutas categorías
	r.HandleFunc("/api/categories", Middleware.SetMiddlewareAuthentication(categoryHandler.GetCategories)).Methods("GET")
	r.HandleFunc("/api/categories", Middleware.SetMiddlewareAuthentication(Middleware.RequireAdmin(categoryHandler.CreateCategory))).Methods("POST")
	r.HandleFunc("/api/categories/lote", Middleware.SetMiddlewareAuthentication(Middleware.RequireAdmin(categoryHandler.CreateCategories))).Methods("POST")
	r.HandleFunc("/api/categories/{id}", Middleware.SetMiddlewareAuthentication(Middleware.RequireAdmin(categoryHandler.UpdateCategory))).Methods("PUT")
	r.HandleFunc("/api/categories/{id}", Middleware.SetMiddlewareAuthentication(Middleware.RequireAdmin(categoryHandler.DeleteCategory))).Methods("DELETE")

	// Ruta de historial de productos
	r.HandleFunc("/api/products/{id}/history", Middleware.SetMiddlewareAuthentication(Middleware.RequireAdmin(productHandler.GetProductHistory))).Methods("GET")

	// Ruta de filtros, ordenamiento, paginación y busqueda
	r.HandleFunc("/api/search", searchHandler.Search).Methods("GET")

	// Ruta de autenticación
	r.HandleFunc("/api/login", userHandler.Login).Methods("POST")

	// Ruta web sockets
	r.HandleFunc("/api/ws", WebSocket.HandleConnection)
	WebSocket.Init()

	http.ListenAndServe(":"+port, r)
}

func CargarPuerto() (string, error) {

	err := godotenv.Load(".env.example")
	if err != nil {
		return os.Getenv("PORT"), err
	}
	return os.Getenv("PORT"), nil
}

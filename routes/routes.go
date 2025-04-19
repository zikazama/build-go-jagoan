package routes

import (
	"build-go-jagoan/handlers"
	"net/http"
)

func SetupRoutes() {
	http.HandleFunc("/products", handlers.GetProducts)         // Mendapatkan daftar produk
	http.HandleFunc("/product", handlers.AddProduct)           // Menambahkan produk
	http.HandleFunc("/product/delete", handlers.DeleteProduct) // Menghapus produk
}

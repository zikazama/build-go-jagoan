package main

import (
	"build-go-jagoan/models"
	"build-go-jagoan/routes"
	"log"
	"net/http"
)

// Database in-memory (menggunakan map untuk menyimpan produk)
var products = make(map[string]models.Product)

func main() {
	// Menambahkan beberapa produk contoh
	products["1"] = models.Product{ID: "1", Name: "Laptop", Price: 1500.00}
	products["2"] = models.Product{ID: "2", Name: "Smartphone", Price: 700.00}

	// Setup Routes
	routes.SetupRoutes()

	// Menjalankan server di port 8080
	log.Fatal(http.ListenAndServe(":8080", nil))
}

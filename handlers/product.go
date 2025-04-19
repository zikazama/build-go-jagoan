package handlers

import (
	"build-go-jagoan/models"
	"encoding/json"
	"net/http"
	"sync"
)

// Database in-memory (menggunakan map untuk menyimpan produk)
var products = make(map[string]models.Product)
var mu sync.Mutex // untuk sinkronisasi akses ke map

// Handler untuk menampilkan semua produk
func GetProducts(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	// Menampilkan semua produk dalam bentuk JSON
	var productList []models.Product
	for _, product := range products {
		productList = append(productList, product)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(productList)
}

// Handler untuk menambahkan produk
func AddProduct(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	var newProduct models.Product
	// Parsing body request untuk mendapatkan data produk
	if err := json.NewDecoder(r.Body).Decode(&newProduct); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Menambahkan produk ke database in-memory
	products[newProduct.ID] = newProduct

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newProduct)
}

// Handler untuk menghapus produk berdasarkan ID
func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	id := r.URL.Query().Get("id")
	if _, exists := products[id]; exists {
		delete(products, id)
		w.WriteHeader(http.StatusNoContent)
	} else {
		http.Error(w, "Product not found", http.StatusNotFound)
	}
}

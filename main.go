package main

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
)

// Struktur Produk
type Product struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

// Database in-memory (menggunakan map untuk menyimpan produk)
var products = make(map[string]Product)
var mu sync.Mutex // untuk sinkronisasi akses ke map

// Handler untuk menampilkan semua produk
func getProducts(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	// Menampilkan semua produk dalam bentuk JSON
	var productList []Product
	for _, product := range products {
		productList = append(productList, product)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(productList)
}

// Handler untuk menambahkan produk
func addProduct(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	var newProduct Product
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
func deleteProduct(w http.ResponseWriter, r *http.Request) {
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

func main() {
	// Menambahkan beberapa produk contoh
	products["1"] = Product{ID: "1", Name: "Laptop", Price: 1500.00}
	products["2"] = Product{ID: "2", Name: "Smartphone", Price: 700.00}

	// Menyiapkan handler untuk setiap endpoint
	http.HandleFunc("/products", getProducts)         // Mendapatkan daftar produk
	http.HandleFunc("/product", addProduct)           // Menambahkan produk
	http.HandleFunc("/product/delete", deleteProduct) // Menghapus produk

	// Menjalankan server di port 8080
	log.Fatal(http.ListenAndServe(":8080", nil))
}

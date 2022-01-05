package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/vicmanbrile/test/models"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/almacen", Almacen)
	r.HandleFunc("/almacen/{id}", AlmacenFilter)

	fmt.Println("http://localhost:8080")
	http.ListenAndServe(":8080", r)
}

func Almacen(w http.ResponseWriter, r *http.Request) {
	var result []byte

	switch r.Method {
	case http.MethodGet:
		{
			// Header
			w.Header().Set("Content-Type", "application/json")
			// Body
			result = models.SelectAllStock()
		}
	default:
		{
			w.WriteHeader(http.StatusNotFound)
		}
	}

	fmt.Fprintf(w, "%s", result)

}

func AlmacenFilter(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)

	w.Header().Set("Content-Type", "application/json")
	result := models.FilterStock(id["id"])

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s", result)
}

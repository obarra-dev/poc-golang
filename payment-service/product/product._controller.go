package product

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"poc-golang/payment-service/webservices/middleware"
)

const productsPath = "products"

// SetupRoutes :
func SetupRoutes(apiBasePath string) {
	productsHandler := http.HandlerFunc(handleProducts)
	http.Handle(fmt.Sprintf("%s/%s", apiBasePath, productsPath), middleware.Middleware(productsHandler))

	productHandler := http.HandlerFunc(handleProduct)
	http.Handle(fmt.Sprintf("%s/%s/", apiBasePath, productsPath), middleware.Middleware(productHandler))
}

func handleProducts(w http.ResponseWriter, r *http.Request) {
	log.Println("Product ALL")
	switch r.Method {
	case http.MethodGet:
		getAll(w, r)
	case http.MethodPost:
		save(w, r)
	case http.MethodOptions:
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func handleProduct(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getOne(w, r)
	case http.MethodPut:
		update(w, r)
	case http.MethodDelete:
		delete(w, r)
	case http.MethodOptions:
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func getAll(w http.ResponseWriter, r *http.Request) {
	productList, err := getProductList()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	j, err := json.Marshal(productList)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	//w.WriteHeader(http.StatusOK)
	_, err = w.Write(j)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func getRequestParm(path string) (int, error) {
	urlPathSegments := strings.Split(path, fmt.Sprintf("%s/", productsPath))
	if len(urlPathSegments[1:]) > 1 {
		return 0, fmt.Errorf("%s", "ORROR")
	}
	productID, err := strconv.Atoi(urlPathSegments[len(urlPathSegments)-1])
	if err != nil {
		log.Print(err)
		return 0, fmt.Errorf("%s", "ORROR")
	}

	return productID, nil
}

func getOne(w http.ResponseWriter, r *http.Request) {
	productID, err := getRequestParm(r.URL.Path)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	product, err := getOneProduct(productID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if product == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	j, err := json.Marshal(product)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	_, err = w.Write(j)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func save(w http.ResponseWriter, r *http.Request) {
	var product Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	productID, err := insertProduct(product)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf(`{"productId":%d}`, productID)))
}

func update(w http.ResponseWriter, r *http.Request) {
	productID, err := getRequestParm(r.URL.Path)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var product Product
	err = json.NewDecoder(r.Body).Decode(&product)
	if err != nil || product.ProductID == nil {
		log.Print(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if *product.ProductID != productID {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = updateProduct(product)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func delete(w http.ResponseWriter, r *http.Request) {
	productID, err := getRequestParm(r.URL.Path)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = removeProduct(productID)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

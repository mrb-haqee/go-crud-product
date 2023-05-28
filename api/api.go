package api

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mrb-haqee/go-crud-product/database"
)

type API struct {
	db  database.DbProduct
	mux *mux.Router
}

func NewAPI(db database.DbProduct) API {
	r := mux.NewRouter()
	api := API{mux: r, db: db}

	r.HandleFunc("/product", api.Index).Methods("GET")
	r.HandleFunc("/product/{id}", api.FindId).Methods("GET")
	r.HandleFunc("/product/add", api.AddProduct).Methods("POST")
	r.HandleFunc("/product/update/{id}", api.UpdateProduct).Methods("PUT")
	r.HandleFunc("/product/delete/{id}", api.DeleteProduct).Methods("DELETE")
	return api
}

func (api *API) Handler() http.Handler {
	return api.mux
}

func (api *API) Start() error {
	log.Println("Running server at http://localhost:8080")
	return http.ListenAndServe(":8080", api.Handler())
}

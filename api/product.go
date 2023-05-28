package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/mrb-haqee/go-crud-product/model"
)

func (api *API) Index(w http.ResponseWriter, r *http.Request) {

	result, err := api.db.FindAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)

}
func (api *API) FindId(w http.ResponseWriter, r *http.Request) {

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	var product model.Product
	err := api.db.FindById(id, &product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)

}
func (api *API) AddProduct(w http.ResponseWriter, r *http.Request) {

	var product model.Product
	json.NewDecoder(r.Body).Decode(&product)

	err := api.db.Add(product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	mess := model.SuccesMessage{Pesan: "Berhasil Input ke Database"}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(mess)
}
func (api *API) UpdateProduct(w http.ResponseWriter, r *http.Request) {

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	var product model.Product
	json.NewDecoder(r.Body).Decode(&product)
	err := api.db.Update(id, product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	mess := model.SuccesMessage{Pesan: "Berhasil Update ke Database"}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(mess)

}
func (api *API) DeleteProduct(w http.ResponseWriter, r *http.Request) {

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	err := api.db.Delete(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	mess := model.SuccesMessage{Pesan: "Berhasil Delete Product"}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(mess)

}

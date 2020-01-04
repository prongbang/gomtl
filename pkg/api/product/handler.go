package product

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Handler interface {
	GetProduct(w http.ResponseWriter, r *http.Request)
}

type handler struct {
	Uc UseCase
}

func (h *handler) GetProduct(w http.ResponseWriter, r *http.Request) {
	accept := r.Header.Get("Accept-Language")
	product := h.Uc.GetProductList(accept)
	log.Println("product", product)
	js, _ := json.Marshal(product)
	_, _ = fmt.Fprintf(w, "%s", js)
}

func NewHandler(uc UseCase) Handler {
	return &handler{
		Uc: uc,
	}
}

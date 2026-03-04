package http

import (
	"encoding/json"
	"net/http"
)

type Controller struct {
}

/**/

func NewHttpController() *Controller {
	return &Controller{}
}

/**/

func (c *Controller) SetRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/order/create", c.CreateOrder)
}

/**/
func (c *Controller) CreateOrder(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode("TODO:")

}

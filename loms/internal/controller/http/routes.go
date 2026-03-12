package httpcontroller

import "net/http"

func (this *LomsHttpController) SetRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/order/create", this.CreateOrder)
}

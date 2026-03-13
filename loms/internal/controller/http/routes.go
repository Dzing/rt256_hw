package httpcontroller

import "net/http"

func (this *LomsHttpController) SetRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/order/create", this.CreateOrder)
	mux.HandleFunc("/order/info", this.OrderInfo)
	mux.HandleFunc("/order/info", this.OrderPay)
	mux.HandleFunc("/order/cancel", this.OrderCancel)
	mux.HandleFunc("/stock/imfo", this.StockInfo)
}

package httpcontroller

import "net/http"

func (c *LomsHttpController) SetupRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/order/create", c.CreateOrder)
	mux.HandleFunc("/order/info", c.OrderInfo)
	mux.HandleFunc("/order/pay", c.OrderPay)
	mux.HandleFunc("/order/cancel", c.OrderCancel)
	mux.HandleFunc("/stock/imfo", c.StockInfo)
}

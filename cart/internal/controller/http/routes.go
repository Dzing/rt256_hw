package httpcontroller

import "net/http"

func (c *CartHttpController) SetupRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /cart/item/add", c.CartItemAdd)
	mux.HandleFunc("POST /cart/item/delete", c.CartItemDelete)
	mux.HandleFunc("POST /cart/list", c.CartList)
	mux.HandleFunc("POST /cart/clear", c.CartClear)
	mux.HandleFunc("POST /cart/checkout", c.CartCheckout)
}

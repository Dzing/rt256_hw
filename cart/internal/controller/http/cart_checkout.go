package httpcontroller

import (
	"encoding/json"
	"log"
	"net/http"
)

type (
	cartCheckoutRequestBody struct {
		User uint64 `json:"user"`
	}

	cartCheckoutResponseBody struct {
		OrderId uint64 `json:"orderId"`
	}
)

func (c *CartHttpController) CartCheckout(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := r.Body.Close(); err != nil {
			log.Printf("error closing response body: %v\n", err)
		}
	}()

	var reqBody cartCheckoutRequestBody

	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	order, err := c.cartService.CartCheckout(reqBody.User)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	respBody := cartCheckoutResponseBody{
		OrderId: order.OrderId,
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(respBody); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
}

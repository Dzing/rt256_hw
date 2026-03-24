package httpcontroller

import (
	"encoding/json"
	"log"
	"net/http"
)

type (
	cartClearRequestBody struct {
		User uint64 `json:"user"`
	}
)

func (c *CartHttpController) CartClear(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := r.Body.Close(); err != nil {
			log.Printf("error closing response body: %v\n", err)
		}
	}()

	var reqBody cartClearRequestBody

	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := c.cartService.CartClear(reqBody.User); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

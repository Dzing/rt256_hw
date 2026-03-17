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

}

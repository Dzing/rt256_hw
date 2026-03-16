package httpcontroller

import (
	"encoding/json"
	"net/http"
)

type (
	cartCheckoutRequestBody struct {
		User uint64 `json:"user"`
	}
)

func (c *CartHttpController) CartCheckout(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var reqBody cartCheckoutRequestBody

	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

}

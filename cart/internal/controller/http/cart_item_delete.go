package httpcontroller

import (
	"encoding/json"
	"net/http"
)

type (
	сartItemDeleteRequestBody struct {
		User uint64 `json:"user"`
		Sku  uint32 `json:"sku"`
	}
)

func (c *CartHttpController) CartItemDelete(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var reqBody сartItemDeleteRequestBody

	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := c.cartService.DeleteCartItem(reqBody.User, reqBody.Sku); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

}

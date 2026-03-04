package httpcontroller

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/vaa/hw/cart/internal/usecase"
)

type (
	cartAddItemRequestBody struct {
		User  uint64 `json:"user"`
		Sku   uint32 `json:"sku"`
		Count uint16 `json:"count"`
	}
)

func (c *CartHttpController) CartItemAdd(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var err error
	var reqBody cartAddItemRequestBody

	err = json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = c.cartService.AddCartItem(reqBody.User, reqBody.Sku, reqBody.Count)

	if err != nil {
		if errors.Is(err, usecase.ErrInsufficientStock) {
			w.WriteHeader(http.StatusPreconditionFailed)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

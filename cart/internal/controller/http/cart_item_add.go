package httpcontroller

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"log/slog"
	"net/http"

	"route/cart/internal/usecase"
)

type (
	cartAddItemRequestBody struct {
		User  uint64 `json:"user"`
		Sku   uint32 `json:"sku"`
		Count uint16 `json:"count"`
	}
)

func (c *CartHttpController) CartItemAdd(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := r.Body.Close(); err != nil {
			log.Printf("error closing response body: %v\n", err)
		}
	}()

	var reqBody cartAddItemRequestBody

	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := c.cartService.AddCartItem(reqBody.User, reqBody.Sku, reqBody.Count); err != nil {
		slog.Error(fmt.Sprintf("failed to add cart item: %+v\n", err))
		if errors.Is(err, usecase.ErrInsufficientStock) {
			w.WriteHeader(http.StatusPreconditionFailed)
			_ = json.NewEncoder(w).Encode(fmt.Sprint(err))
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(fmt.Sprint(err))
		return
	}

	w.WriteHeader(http.StatusOK)
}

package httpcontroller

import (
	"encoding/json"
	"net/http"
)

type (
	cartListRequestBody struct {
		User uint64 `json:"user"`
	}

	cartListResponseItem struct {
		Sku   uint32 `json:"sku"`
		Count uint16 `json:"count"`
		Name  string `json:"name"`
		Price uint64 `json:"price"`
	}

	cartListResponseBody struct {
		Items      []*cartListResponseItem `json:"items"`
		TotalPrice uint64                  `json:"totalPrice"`
	}
)

func (c *CartHttpController) CartList(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var reqBody cartListRequestBody
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	cart, err := c.cartService.FindCart(reqBody.User)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	responseItems := make([]*cartListResponseItem, 1)

	for _, cartItem := range cart.Items {

		responseItems = append(responseItems, &cartListResponseItem{
			Sku:   cartItem.Product.Sku,
			Name:  cartItem.Product.Name,
			Price: cartItem.Product.Price,
			Count: cartItem.Count,
		})
	}

	respBody := cartListResponseBody{
		Items:      responseItems,
		TotalPrice: cart.TotalPrice(),
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(respBody); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)

}

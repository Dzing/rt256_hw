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

	productItemList, err := c.cartService.ProductItemList(reqBody.User)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var totalPrice uint64 = 0
	responseItems := make([]*cartListResponseItem, 1)

	for _, prodItem := range productItemList {
		totalItemPrice := prodItem.TradeItem.Price * uint64(prodItem.Count)
		totalPrice += totalItemPrice
		responseItems = append(responseItems, &cartListResponseItem{
			Sku:   prodItem.TradeItem.Sku,
			Name:  prodItem.TradeItem.Name,
			Price: prodItem.TradeItem.Price,
			Count: prodItem.Count,
		})
	}

	respBody := cartListResponseBody{
		Items:      responseItems,
		TotalPrice: totalPrice,
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(respBody); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)

}

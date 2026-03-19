package httpcontroller

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"route/loms/internal/usecase"
)

type (
	createOrderRequestBodyItemRecord struct {
		Sku   uint32 `json:"sku"`
		Count uint16 `json:"count"`
	}
	createOrderRequestBody struct {
		UserId uint64                             `json:"user"`
		Items  []createOrderRequestBodyItemRecord `json:"items"`
	}
	createOrderResponseBody struct {
		OrderId uint64 `json:"orderId"`
	}
)

func (c *LomsHttpController) CreateOrder(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := r.Body.Close(); err != nil {
			log.Printf("error closing response body: %v\n", err)
		}
	}()

	var reqBody createOrderRequestBody

	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	itemList := usecase.ItemCountListDTO{
		Items: make([]*usecase.SkuCountRecord, 1),
	}

	for _, data := range reqBody.Items {
		itemList.Items = append(
			itemList.Items,
			&usecase.SkuCountRecord{
				Sku:   usecase.TSku(data.Sku),
				Count: usecase.TCount(data.Count),
			},
		)
	}

	order, err := c.lomsService.CreateOrder(
		usecase.TUserId(reqBody.UserId),
		&itemList,
	)

	if err != nil {
		if errors.As(err, &usecase.ErrInsufficientStock) {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
	}

	respBody := createOrderResponseBody{
		OrderId: uint64(order.Id),
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(respBody); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

}

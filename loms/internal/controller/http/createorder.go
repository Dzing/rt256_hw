package httpcontroller

import (
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
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
			slog.Error(fmt.Sprintf("failed to close response body: %v\n", err))
		}
	}()

	slog.Info("handling request CreateOrder")
	var reqBody createOrderRequestBody

	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		slog.Error(fmt.Sprintf("unable to decode body : %+v\n", err))
		return
	}

	slog.Info("decoded request", "body", fmt.Sprintf("%++v", reqBody))

	items := make([]*usecase.SkuCountRecord, 0)

	for _, data := range reqBody.Items {
		items = append(
			items,
			&usecase.SkuCountRecord{
				Sku:   usecase.TSku(data.Sku),
				Count: usecase.TCount(data.Count),
			},
		)
	}

	itemList := &usecase.ItemCountListDTO{
		Items: items,
	}

	order, err := c.lomsService.CreateOrder(
		usecase.TUserId(reqBody.UserId),
		itemList,
	)

	if err != nil {
		if errors.As(err, &usecase.ErrInsufficientStock) {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(err.Error())
			slog.Error(fmt.Sprintf("failed to create order : %+v\n", err))
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
		slog.Error(fmt.Sprintf("failed to encode response : %+v\n", respBody))
		return
	}

	w.WriteHeader(http.StatusOK)

}

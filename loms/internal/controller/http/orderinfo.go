package httpcontroller

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"route/loms/internal/usecase"
)

type (
	orderInfoRequestBody struct {
		OrderId uint64 `json:"orderId"`
	}

	orderInfoItemRecord struct {
		Sku   uint32 `json:"sku"`
		Count uint16 `json:"count"`
	}
	orderInfoResponseBody struct {
		Status string                 `json:"status"`
		UserId uint64                 `json:"user"`
		Items  []*orderInfoItemRecord `json:"items"`
	}
)

func (c *LomsHttpController) OrderInfo(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := r.Body.Close(); err != nil {
			slog.Error(fmt.Sprintf("failed to close response body: %v\n", err))
		}
	}()

	var reqBody orderInfoRequestBody

	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(err.Error())
		slog.Error(fmt.Sprintf("unable to decode body : %+v\n", err))
		return
	}

	order, err := c.lomsService.FindOrder(usecase.TOrderId(reqBody.OrderId))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(fmt.Sprint(err))
		slog.Error(fmt.Sprintf("failed to find order : %+v\n", err))
		return
	}

	items := make([]*orderInfoItemRecord, 0)
	for _, itemData := range order.Items {
		items = append(items, &orderInfoItemRecord{
			Sku:   uint32(itemData.Sku),
			Count: uint16(itemData.Count),
		})
	}

	respBody := orderInfoResponseBody{
		Status: usecase.OrderStateToString(usecase.EOrderState(order.State)),
		UserId: uint64(order.UserId),
		Items:  items,
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(respBody); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		slog.Error(fmt.Sprintf("failed to encode response : %+v\n", respBody))
		return
	}
	w.WriteHeader(http.StatusOK)
}

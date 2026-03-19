package httpcontroller

import (
	"encoding/json"
	"log"
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
			log.Printf("error closing response body: %v\n", err)
		}
	}()

	var reqBody orderInfoRequestBody

	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	order, err := c.lomsService.FindOrder(usecase.TOrderId(reqBody.OrderId))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	items := make([]*orderInfoItemRecord, 1)
	for _, itemData := range order.Items {
		items = append(items, &orderInfoItemRecord{
			Sku:   uint32(itemData.Sku),
			Count: uint16(itemData.Count),
		})
	}

	respBody := orderInfoResponseBody{
		Status: OrderStateToString(order.State),
		UserId: uint64(order.UserId),
		Items:  items,
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(respBody); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

package httpcontroller

import (
	"encoding/json"
	"net/http"

	"atlas.chr/vaa/route-hw/loms/internal/usecase"
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

func (this *LomsHttpController) OrderInfo(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var err error
	var reqBody orderInfoRequestBody

	err = json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	order, err := this.lomsService.FindOrder(usecase.TOrderId(reqBody.OrderId))

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

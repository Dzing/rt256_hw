package httpcontroller

import (
	"encoding/json"
	"net/http"

	"atlas.chr/vaa/route-hw/loms/internal/usecase"
)

type (
	orderPayRequestBody struct {
		OrderId uint64 `json:"orderId"`
	}
)

func (c *LomsHttpController) OrderPay(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var err error
	var reqBody orderInfoRequestBody

	err = json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = c.lomsService.PayOrder(usecase.TOrderId(reqBody.OrderId))

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

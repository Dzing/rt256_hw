package httpcontroller

import (
	"encoding/json"
	"log"
	"net/http"

	"atlas.chr/vaa/route-hw/loms/internal/usecase"
)

type (
	orderPayRequestBody struct {
		OrderId uint64 `json:"orderId"`
	}
)

func (c *LomsHttpController) OrderPay(w http.ResponseWriter, r *http.Request) {
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

	if err := c.lomsService.PayOrder(usecase.TOrderId(reqBody.OrderId)); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// TODO: отправить ответ

	w.WriteHeader(http.StatusOK)
}

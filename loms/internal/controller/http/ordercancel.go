package httpcontroller

import (
	"encoding/json"
	"log"
	"net/http"

	"route/loms/internal/usecase"
)

type (
	orderCancelRequestBody struct {
		OrderId uint64 `json:"orderId"`
	}
)

func (c *LomsHttpController) OrderCancel(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := r.Body.Close(); err != nil {
			log.Printf("error closing response body: %v\n", err)
		}
	}()

	var reqBody orderCancelRequestBody

	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := c.lomsService.CancelOrder(usecase.TOrderId(reqBody.OrderId)); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

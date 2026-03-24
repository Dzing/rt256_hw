package httpcontroller

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"log/slog"
	"net/http"

	"route/loms/internal/usecase"
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

	var reqBody orderPayRequestBody

	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := c.lomsService.PayOrder(usecase.TOrderId(reqBody.OrderId)); err != nil {
		if errors.As(err, &usecase.ErrOrderStateMismatch) {
			w.WriteHeader(http.StatusPreconditionFailed)
			_ = json.NewEncoder(w).Encode(fmt.Sprint(err))
			slog.Error(fmt.Sprintf("failed to pay order : %+v\n", err))
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(fmt.Sprint(err))
		return
	}

	w.WriteHeader(http.StatusOK)
}

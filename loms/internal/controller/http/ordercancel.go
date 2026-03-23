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
		if errors.As(err, &usecase.ErrOrderStateMismatch) {
			w.WriteHeader(http.StatusPreconditionFailed)
			_ = json.NewEncoder(w).Encode(fmt.Sprint(err))
			slog.Error(fmt.Sprintf("failed to cancel order : %+v\n", err))
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(fmt.Sprint(err))
		return
	}

	w.WriteHeader(http.StatusOK)
}

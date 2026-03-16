package httpcontroller

import (
	"encoding/json"
	"net/http"

	"atlas.chr/vaa/route-hw/loms/internal/usecase"
)

type (
	stockInfoRequestBody struct {
		Sku uint32 `json:"sku"`
	}

	stockInfoResponseBody struct {
		Count uint64 `json:"count"`
	}
)

func (c *LomsHttpController) StockInfo(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var reqBody stockInfoRequestBody

	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	info, err := c.lomsService.StockInfo(usecase.TSku(reqBody.Sku))

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	respBody := stockInfoResponseBody{
		Count: uint64(info.Available),
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(respBody); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

package clienthttp

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	uc "route/cart/internal/usecase"
)

type (
	stockInfoRequestPayload struct {
		Sku uint32
	}

	stockInfoResponsePayload struct {
		Count uint16 `json:"count"`
	}
)

func (s *LomsHttpClient) StockInfo(sku uint32) (*uc.StockInfoDTO, error) {
	path := "/stock/info"

	body := stockInfoRequestPayload{
		Sku: sku,
	}

	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("error marshalling JSON: %v", err)
	}

	resp, err := http.Post(s.addr+path, "application/json", bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, fmt.Errorf("error POST request execution: %v", err)
	}

	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Printf("error closing response body: %v\n", err)
		}
	}()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("POST request failed with status: %s", resp.Status)
	}

	var respData stockInfoResponsePayload
	if err = json.NewDecoder(resp.Body).Decode(&respData); err != nil {
		return nil, fmt.Errorf("error decoding JSON response: %v", err)
	}

	return &uc.StockInfoDTO{Sku: sku, Count: respData.Count}, nil
}

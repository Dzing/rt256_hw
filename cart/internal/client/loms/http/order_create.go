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
	itemRecord struct {
		Sku   uint32 `json:"sku"`
		Count uint16 `json:"count"`
	}
	orderCreateRequestBody struct {
		User  uint64       `json:"user"`
		Items []itemRecord `json:"items"`
	}
	orderCreateResponseBody struct {
		OrderID uint64 `json:"orderId"`
	}
)

func (s *LomsHttpClient) OrderCreate(user uint64, cartContent *uc.OrderContentDTO) (*uc.OrderDto, error) {
	path := "/order/create"

	items := make([]itemRecord, 1)

	for _, cartItemData := range cartContent.Items {
		items = append(items, itemRecord{
			Sku:   cartItemData.Sku,
			Count: cartItemData.Count,
		})
	}
	requestBody := orderCreateRequestBody{
		User:  user,
		Items: items,
	}

	jsonBody, err := json.Marshal(requestBody)
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

	var respData orderCreateResponseBody
	if err = json.NewDecoder(resp.Body).Decode(&respData); err != nil {
		return nil, fmt.Errorf("error decoding JSON response: %v", err)
	}

	return &uc.OrderDto{OrderId: respData.OrderID}, nil
}

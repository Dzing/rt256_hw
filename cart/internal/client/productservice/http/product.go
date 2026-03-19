package clienthttp

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"atlas.chr/vaa/hw/cart/internal/usecase"
)

type (
	ProductInfoDTO struct {
		Name  string
		Price uint32
	}
	requestPayload struct {
		sku   uint32
		token string
	}

	responsePayload struct {
		Name  string `json:"name"`
		Price uint64 `json:"price"`
	}
)

func (s *ProductServiceHttpClient) Product(sku uint32) (*usecase.ProductDTO, error) {
	path := "/get_product"

	body := requestPayload{
		sku:   sku,
		token: s.token,
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

	var respData responsePayload
	if err = json.NewDecoder(resp.Body).Decode(&respData); err != nil {
		return nil, fmt.Errorf("error decoding JSON response: %v", err)
	}

	return &usecase.ProductDTO{Sku: sku, Name: respData.Name, Price: uint64(respData.Price)}, nil
}

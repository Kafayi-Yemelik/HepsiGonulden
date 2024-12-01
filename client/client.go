package client

import (
	"HepsiGonulden/internal/types"
	"HepsiGonulden/pkg/authentication"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type HttpOrderClient struct {
	BaseURL string
	client  *http.Client
}

func NewHttpOrderClient(baseURL string) *HttpOrderClient {
	return &HttpOrderClient{
		BaseURL: baseURL,
		client:  &http.Client{},
	}
}

func (c *HttpOrderClient) UpdateOrder(ctx context.Context, orderId string, updateModel types.OrderUpdateModel) error {
	url := fmt.Sprintf("%s/orders/%s", c.BaseURL, orderId)
	payload, err := json.Marshal(updateModel)
	if err != nil {
		return fmt.Errorf("failed to marshal update model: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPut, url, bytes.NewBuffer(payload))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	/*
		TODO(Gonul): Token 72 saatlik, yani ornegin 70 saatte bir yenilenmesi problem yaratmaz.
			ayri bir goroutine ile token her request icin degil, 70 saatte bir guncellenecek sekilde kod revize edilebilir
	*/
	token, err := authentication.JwtGenerator("f31bef0d-ee18-425e-b8c1-ab66bde5e07a", "kafka", "order_create_consumer")
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

	resp, err := c.client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to make request: %w", err)
	} else if !(resp.StatusCode >= 200 && resp.StatusCode < 300) {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		resp.Body.Close()
		return fmt.Errorf("operation failed: %s", string(body))
	}
	defer resp.Body.Close()

	return nil
}

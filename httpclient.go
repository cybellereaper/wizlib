package wizlib

import (
	"context"
	"errors"
	"io"
	"net/http"
	"sync"
	"time"

	cloudflarebp "github.com/DaRealFreak/cloudflare-bp-go"
)

// APIClient provides methods for making HTTP requests.
type APIClient struct {
	Client *http.Client
}

// NewAPIClient creates a new instance of APIClient.
func NewAPIClient() *APIClient {
	return &APIClient{
		Client: &http.Client{
			Timeout:   10 * time.Second,
			Transport: cloudflarebp.AddCloudFlareByPass(&http.Transport{}),
		},
	}
}

// Get makes a GET request to the specified URL.
func (c *APIClient) Get(url string) ([]byte, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	bufferPool := sync.Pool{
		New: func() interface{} {
			return make([]byte, 32*1024)
		},
	}

	buffer := bufferPool.Get().([]byte)
	defer bufferPool.Put(&buffer)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	n, err := io.ReadFull(resp.Body, buffer)
	if err != nil && !errors.Is(err, io.ErrUnexpectedEOF) {
		return nil, err
	}

	return buffer[:n], nil
}

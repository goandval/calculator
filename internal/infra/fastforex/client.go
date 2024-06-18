package fastforex

import (
	"os"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/goandval/calculator/internal/config"
)

type Client struct {
	driver         *resty.Client
	updateInterval time.Duration
}

func New(cfg config.ClientConfig) *Client {
	client := resty.
		New().
		SetBaseURL(cfg.BaseURL).
		SetTimeout(cfg.Timeout)
	return &Client{
		driver:         client,
		updateInterval: cfg.Timeout,
	}
}

func (c Client) Run(stop <-chan os.Signal) {
	ticker := time.NewTicker(c.updateInterval)

	for {
		select {
		case <-stop:
			break
		}
	}
	<-ticker.C
}

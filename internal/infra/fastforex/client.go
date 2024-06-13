package fastforex

import "github.com/go-resty/resty/v2"

type Client struct {
	*resty.Client
}

func New() *Client { // убрать * в будущем
	return nil
}

func (c Client) Run() {
	
}

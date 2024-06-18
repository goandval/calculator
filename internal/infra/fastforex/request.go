package fastforex

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/goandval/calculator/internal/domain"
)

const requestURI = "/fetch-multi"

type FFResponse struct {
	Base    string          `json:"base"`
	Results json.RawMessage `json:"results"`
	Updated string          `json:"updated"`
}

type ErrorResp struct {
	Msg string `json:"error"`
}

func (e *ErrorResp) Error() string {
	return e.Msg
}

func (c Client) makeRequest(ctx context.Context, from string, to []string) (domain.RateMap, error) {
	const op = "fastforex.makeRequest"
	from = strings.ToUpper(from)
	for i := range to {
		to[i] = strings.ToUpper(to[i])
	}

	params := map[string]string{
		"from": from,
		"to":   strings.Join(to, ","),
	}

	response, err := c.driver.R().
		SetContext(ctx).
		SetResult(&FFResponse{}).
		SetError(&ErrorResp{}).
		SetQueryParams(params).
		Get(requestURI)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	if response.StatusCode() == http.StatusOK {
		resp, ok := response.Result().(*FFResponse)
		if !ok {
			return nil, fmt.Errorf("%s: got 200 result but schema is invalid: %w", op, err)
		}
		return makeMapFromFFResponse(resp), nil
	}

	errResp, ok := response.Error().(*ErrorResp)
	if !ok {
		return nil, fmt.Errorf("%s: got not 200 result but schema is invalid: %w", op, err)
	}
	return nil, errResp
}

func makeMapFromFFResponse(r *FFResponse) domain.RateMap {
	return nil
}

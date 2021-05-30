package client

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type request struct {
	method      string
	apiCall     string
	queryParams url.Values
	data        url.Values
}

func (c *Client) doRequest(ctx context.Context, reqConfig request, data interface{}) error {
	var params string
	if len(reqConfig.queryParams) != 0 {
		params = fmt.Sprintf("?%s", reqConfig.queryParams.Encode())
	}
	reqUrl := fmt.Sprintf("%s/%s%s", c.url, reqConfig.apiCall, params)

	var body io.Reader
	if len(reqConfig.data) != 0 {
		body = strings.NewReader(reqConfig.data.Encode())
	}

	req, err := http.NewRequestWithContext(ctx, reqConfig.method, reqUrl, body)
	if err != nil {
		return err
	}
	if reqConfig.method == "POST" {
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		req.Header.Add("Content-Length", strconv.Itoa(len(reqConfig.data.Encode())))
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("http request failed: %s", resp.Status)
	}

	if err := json.NewDecoder(resp.Body).Decode(data); err != nil {
		return err
	}

	return nil
}

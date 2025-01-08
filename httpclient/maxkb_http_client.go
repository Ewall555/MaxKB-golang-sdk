package httpclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/Ewall555/MaxKB-golang-sdk/api/constant"
)

type MaxKBHTTPClient struct {
	BaseURL    string `json:"baseURL"` // baseURL
	ApiKey     string `json:"apiKey"`  // apiKey
	HTTPClient *http.Client
}

func NewDefaultMaxKBHttpClient(baseURL, apiKey string) IMaxKBHttpClient {
	return &MaxKBHTTPClient{
		BaseURL:    strings.TrimRight(baseURL, "/"),
		ApiKey:     apiKey,
		HTTPClient: http.DefaultClient,
	}
}

// SetHTTPClient set HTTPClient
func (c *MaxKBHTTPClient) SetHTTPClient(client *http.Client) {
	c.HTTPClient = client
}

func (c *MaxKBHTTPClient) DoRequest(method, endpoint string, body interface{}, result interface{}) error {
	resp, err := c.DoRequestStream(method, endpoint, body)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if result != nil {
		respBody, err := io.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("failed to read response body: %w", err)
		}

		if err := json.Unmarshal(respBody, result); err != nil {
			return fmt.Errorf("failed to unmarshal response: %w", err)
		}
	}

	return nil
}

func (c *MaxKBHTTPClient) DoRequestStream(method, endpoint string, body interface{}) (*http.Response, error) {
	parsedURL, err := url.Parse(c.BaseURL)
	if err != nil {
		return nil, fmt.Errorf("error parsing URL: %w", err)
	}
	baseURL := fmt.Sprintf("%s://%s", parsedURL.Scheme, parsedURL.Host)
	url := baseURL + constant.ApiBase + endpoint
	var reqBody io.Reader
	if body != nil {
		jsonData, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal JSON body: %w", err)
		}
		reqBody = bytes.NewReader(jsonData)
	}

	req, err := http.NewRequest(method, url, reqBody)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Authorization", c.ApiKey)
	req.Header.Set("Content-Type", "application/json;charset=utf-8")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}

	if resp.StatusCode >= 400 {
		defer resp.Body.Close()
		respBody, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("API error: StatusCode=%d, Body=%s", resp.StatusCode, string(respBody))
	}

	return resp, nil
}

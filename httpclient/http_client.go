package httpclient

import (
	"net/http"
)

// IMaxKBHttpClient
type IMaxKBHttpClient interface {
	SetHTTPClient(client *http.Client)
	DoRequest(method, endpoint string, body interface{}, result interface{}) error
	DoRequestStream(method, endpoint string, body interface{}) (*http.Response, error)
}

// IMaxKBHttpClientContext
type IMaxKBHttpClientContext interface {
}

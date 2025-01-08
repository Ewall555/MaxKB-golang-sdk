package maxkb

import (
	"net/http"

	"github.com/Ewall555/MaxKB-golang-sdk/api/application"
	"github.com/Ewall555/MaxKB-golang-sdk/config"
	"github.com/Ewall555/MaxKB-golang-sdk/context"
	"github.com/Ewall555/MaxKB-golang-sdk/httpclient"
)

type MaxKB struct {
	ctx             *context.Context
	applicationChat *application.ApplicationChat
}

func NewMaxKB(cfg *config.Config) *MaxKB {
	defaultMaxKBHttpClient := httpclient.NewDefaultMaxKBHttpClient(cfg.BaseURL, cfg.ApiKey)
	ctx := &context.Context{
		Config:           cfg,
		IMaxKBHttpClient: defaultMaxKBHttpClient,
	}
	return &MaxKB{
		ctx: ctx,
	}
}

// SetHTTPClient set HTTPClient
func (maxkb *MaxKB) SetHTTPClient(client *http.Client) {
	maxkb.ctx.SetHTTPClient(client)
}

// GetContext get Context
func (maxkb *MaxKB) GetContext() *context.Context {
	return maxkb.ctx
}

// GetApplicationChat get applicationChat
func (maxkb *MaxKB) GetApplicationChat() *application.ApplicationChat {
	if maxkb.applicationChat == nil {
		maxkb.applicationChat = application.NewApplicationChat(maxkb.ctx)
	}
	return maxkb.applicationChat
}

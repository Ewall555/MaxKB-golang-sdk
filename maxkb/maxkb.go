package maxkb

import (
	"github.com/Ewall555/MaxKB-golang-sdk/api/application"
	"github.com/Ewall555/MaxKB-golang-sdk/client"
)

type MaxKB struct {
	ApplicationChat *application.ApplicationChat
}

func New(baseURL, apiKey string) *MaxKB {
	cli := client.NewClient(baseURL, apiKey)
	return &MaxKB{
		ApplicationChat: application.NewApplicationChat(cli),
	}
}

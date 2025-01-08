package context

import (
	"github.com/Ewall555/MaxKB-golang-sdk/config"
	"github.com/Ewall555/MaxKB-golang-sdk/httpclient"
)

// Context struct
type Context struct {
	*config.Config
	httpclient.IMaxKBHttpClient
}

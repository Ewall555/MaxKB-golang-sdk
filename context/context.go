package context

import (
	"github.com/Ewall555/MaxKB-golang-sdk/v2/config"
	"github.com/Ewall555/MaxKB-golang-sdk/v2/httpclient"
)

// Context struct
type Context struct {
	*config.Config
	httpclient.IMaxKBHttpClient
}

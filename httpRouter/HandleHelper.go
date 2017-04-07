package httpRouter

import (
	"httpTools/httpHelper"
	"net/url"
)

type HandleHelper struct {
	Responder httpHelper.Writer
	Request   httpHelper.Request
	Params    url.Values
}

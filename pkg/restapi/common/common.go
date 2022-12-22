/*
Go-lang-Vcs-NEXTClan
*/

package common

import "net/http"

// HTTPRequestHandler is an HTTP handler.
type HTTPRequestHandler func(http.ResponseWriter, *http.Request)

// HTTPHandler is an HTTP handler descriptor containing the context path, method, and request handler.
type HTTPHandler interface {
	Path() string
	Method() string
	Handler() HTTPRequestHandler
}

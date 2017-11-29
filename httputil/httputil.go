package httputil

import (
	"net/http"
	"strings"
)

// IsHTTPS check a request for HTTPS protocol
func IsHTTPS(r *http.Request) bool {
	switch {
	case r.URL.Scheme == "https":
		return true
	case r.TLS != nil:
		return true
	case strings.HasPrefix(r.Proto, "HTTPS"):
		return true
	case r.Header.Get("X-Forwarded-Proto") == "https":
		return true
	default:
		return false
	}
}

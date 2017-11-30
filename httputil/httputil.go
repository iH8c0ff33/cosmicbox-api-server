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

// SetCookie sets a cookie to the client
func SetCookie(w http.ResponseWriter, r *http.Request, name, value string) {
	cookie := http.Cookie{
		Name: name,
		Value: value,
		Path: "/",
		Domain: r.URL.Host,
		HttpOnly: true,
		Secure: IsHTTPS(r),
		MaxAge: 2147483647,
	}

	http.SetCookie(w, &cookie)
}

// DelCookie deletes a cookie
func DelCookie(w http.ResponseWriter, r *http.Request, name string) {
	cookie := http.Cookie{
		Name: name,
		Value: "DELETED",
		Path: "/",
		Domain: r.URL.Host,
		MaxAge: -1,
	}

	http.SetCookie(w, &cookie)
}

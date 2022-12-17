package helpers

import "net/http"

// HTTPKey is the key used to extract the Http struct.
type HTTPKey string

// HTTP is the struct used to inject the response writer and request http structs.
type HTTP struct {
	W *http.ResponseWriter
	R *http.Request
}

package svc

import (
	"net/http"

	"github.com/owncloud/ocis-pkg/v2/middleware"
)

// NewTracing returns a service that instruments traces.
func NewTracing(next Service) Service {
	return tracing{
		next: next,
	}
}

type tracing struct {
	next Service
}

// ServeHTTP implements the Service interface.
func (t tracing) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	middleware.Trace(t.next).ServeHTTP(w, r)
}

// Dummy implements the Service interface.
func (t tracing) Dummy(w http.ResponseWriter, r *http.Request) {
	t.next.Dummy(w, r)
}

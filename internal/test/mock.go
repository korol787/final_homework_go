package test

import (
	"net/http"
	"net/http/httptest"

	routing "github.com/go-ozzo/ozzo-routing/v2"
	"github.com/go-ozzo/ozzo-routing/v2/content"
	"github.com/go-ozzo/ozzo-routing/v2/cors"
	"users-balance-microservice/internal/errors"
	"users-balance-microservice/pkg/accesslog"
	"users-balance-microservice/pkg/log"
)

// MockRoutingContext creates a routing.Context for testing handlers.
func MockRoutingContext(req *http.Request) (*routing.Context, *httptest.ResponseRecorder) {
	res := httptest.NewRecorder()
	if req.Header.Get("Content-Type") == "" {
		req.Header.Set("Content-Type", "application/json")
	}
	ctx := routing.NewContext(res, req)
	ctx.SetDataWriter(&content.JSONDataWriter{})
	return ctx, res
}

// MockRouter creates a routing.Router for testing APIs.
func MockRouter(logger log.Logger) *routing.Router {
	router := routing.New()
	router.Use(
		accesslog.Handler(logger),
		errors.Handler(logger),
		content.TypeNegotiator(content.JSON),
		cors.Handler(cors.AllowAll),
	)
	return router
}
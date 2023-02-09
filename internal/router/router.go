package router

import (
	"github.com/gin-gonic/gin"
	ohlc "github.com/teezzan/ohlc/internal/controller/ohlc"
)

type Router struct {
	router *gin.Engine

	healthHandler   gin.HandlerFunc
	ohlcHttpHandler *ohlc.HTTPHandler
}

// New initializes a new router
func New(
	healthHandler gin.HandlerFunc,
	ohlcHttpHandler *ohlc.HTTPHandler,
) *Router {
	return &Router{
		healthHandler:   healthHandler,
		ohlcHttpHandler: ohlcHttpHandler,
	}
}

// SetupRouter implements the gin.Server interface.
func (r *Router) SetupRouter(router *gin.Engine) error {
	r.router = router

	r.setupRoutes()

	return nil
}

// Run implements the gin.Server interface.
func (r *Router) Run(addr ...string) error {
	return r.router.Run(addr...)
}

// GetRouter returns the gin router.
func (r *Router) GetRouter() *gin.Engine {
	return r.router
}

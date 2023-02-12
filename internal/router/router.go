package router

import (
	"github.com/gin-gonic/gin"
	"github.com/teezzan/candles/docs"
	ohlc "github.com/teezzan/candles/internal/controller/ohlc"
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
	r.router.Static("/docs", "./docs")
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Schemes = []string{"https"}

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

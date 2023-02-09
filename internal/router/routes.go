package router

func (r *Router) setupRoutes() {
	r.router.GET("/health", r.healthHandler)

	r.ohlcHttpHandler.SetupRouter(r.router.Group("/ohlc"))
}

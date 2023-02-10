package ohlc

import (
	"encoding/csv"

	"github.com/gin-gonic/gin"
	"github.com/teezzan/ohlc/internal/controller/ohlc/data"
	"github.com/teezzan/ohlc/internal/httputil"
	"go.uber.org/zap"
)

// HTTPHandler is the HTTP handler for the ohlc service.
type HTTPHandler struct {
	logger      *zap.Logger
	ohlcService Service
}

// NewHTTPHandler initializes a new HTTP Handler.
func NewHTTPHandler(
	logger *zap.Logger,
	ohlcService Service,
) *HTTPHandler {
	return &HTTPHandler{
		logger:      logger,
		ohlcService: ohlcService,
	}
}

// SetupRouter sets up the router for the ohlc service.
func (h *HTTPHandler) SetupRouter(r *gin.RouterGroup) error {
	handler := httputil.NewHandlerWrapper(h.logger)

	r.POST("/data", handler(h.processCSVHandler))
	r.GET("/data", handler(h.getOHLCPointsHandler))

	return nil
}

// processCSVHandler processes the CSV file.
func (h *HTTPHandler) processCSVHandler(c *gin.Context) error {
	file, err := c.FormFile("file")
	if err != nil {
		return httputil.BadRequest(c, err)
	}

	src, err := file.Open()
	if err != nil {
		return httputil.BadRequest(c, err)
	}

	r := csv.NewReader(src)
	csvData, err := r.ReadAll()
	if err != nil {
		return httputil.BadRequest(c, err)
	}

	err = h.ohlcService.CreateDataPoints(c, csvData)
	if err != nil {
		return err
	}

	return httputil.OK(c, nil)
}

// getOHLCPointsHandler gets the OHLC points for the given time range.
func (h *HTTPHandler) getOHLCPointsHandler(c *gin.Context) error {
	var query data.GetOHLCRequest
	if err := c.ShouldBindQuery(&query); err != nil {
		return httputil.BadRequest(c, err)
	}
	dp, err := h.ohlcService.GetDataPoints(c, query)
	if err != nil {
		return err
	}

	var p = []data.OHLC{}
	for _, point := range dp {
		p = append(p, point.ToOHLC())
	}

	return httputil.OK(c, p)
}

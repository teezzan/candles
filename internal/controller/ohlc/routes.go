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
	r.GET("/data", handler(h.getOHLCDataHandler))
	r.GET("/generate_url", handler(h.generatePreSignedURLHandler))

	return nil
}

// processCSVHandler processes the CSV file.
// ProcessCSVFile godoc
// @Summary      Takes a CSV file upload and processes it
// @Description  The endpoint takes a small CSV file upload and processes it. Max file size is 30MB.
// @Tags         Upload, CSV
// @Accept       mpfd
// @Produce      json
// @Success      200
// @Failure      400  {object}  httputil.ErrorResponse
// @Failure      500  {object}  httputil.ErrorResponse
// @Router       /data [post]
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
// GetOHLCData godoc
// @Summary      returns the OHLC points for the given time range
// @Description  The endpoint returns the OHLC points for a particular Symbol for  the given time range
// @Tags         fetch, OHLC
// @Produce      json
// @Param   	 symbol     query     string     true  "HAKO"     default(A)
// @Param   	 from     query     string     true  "123363781282"     example(string)
// @Param   	 to     query     string     false  "123363212332"     example(string)
// @Param   	 page     query     int     false  1     example(string)
// @Param   	 page_size     query     int     false  2     example(string)
// @Success      200  {object}  data.GetOHLCResponse
// @Failure      400  {object}  httputil.ErrorResponse
// @Failure      500  {object}  httputil.ErrorResponse
// @Router       /data [get]
func (h *HTTPHandler) getOHLCDataHandler(c *gin.Context) error {
	var query data.GetOHLCRequest
	if err := c.ShouldBindQuery(&query); err != nil {
		return httputil.BadRequest(c, err)
	}
	dp, page, err := h.ohlcService.GetDataPoints(c, query)
	if err != nil {
		return err
	}

	var p = []data.OHLC{}
	for _, point := range dp {
		p = append(p, point.ToOHLC())
	}

	resp := data.GetOHLCResponse{
		DataPoints: p,
		Page:       *page,
	}

	return httputil.OK(c, resp)
}

// generatePreSignedURLHandler generates a pre-signed URL for the given file name.
func (h *HTTPHandler) generatePreSignedURLHandler(c *gin.Context) error {
	result, err := h.ohlcService.GeneratePreSignedURL(c)
	if err != nil {
		return err
	}

	return httputil.OK(c, result)
}

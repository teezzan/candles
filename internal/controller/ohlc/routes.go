package ohlc

import (
	"encoding/csv"

	"github.com/gin-gonic/gin"
	"github.com/teezzan/candles/internal/controller/ohlc/data"
	"github.com/teezzan/candles/internal/httputil"
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
//
//	@Summary		Takes a CSV file upload and processes it
//	@Description	The endpoint takes a small CSV file upload and processes it. Max file size is 30MB.
//	@Accept			multipart/form-data
//	@Produce		json
//	@Param			file	formData	file	true	"account image"
//	@Success		200
//	@Failure		400	{object}	httputil.ErrorResponse
//	@Failure		500	{object}	httputil.ErrorResponse
//	@Router			/data [post]
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
//
//	@Summary		returns the OHLC points for the given time range
//	@Description	The endpoint returns the OHLC points for a particular Symbol for  the given time range
//	@Produce		json
//	@Param			symbol		query		string	true	"This is the symbol of the OHLC token"			example(BTC)
//	@Param			from		query		string	true	"UNIX time representation of the start time"	example(10344553332)
//	@Param			to			query		string	false	"UNIX time representation of the end time"		example(101019283847)
//	@Param			page		query		int		false	"page of response"								example(1)
//	@Param			page_size	query		int		false	"Number of OHLC datapoints per page"			example(5)
//	@Success		200			{object}	data.GetOHLCResponse
//	@Failure		400			{object}	httputil.ErrorResponse
//	@Failure		500			{object}	httputil.ErrorResponse
//	@Router			/data [get]
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
//
//	@Summary		Generates a pre-signed URL for the given file name for uploading on S3
//	@Description	The endpoint generates a pre-signed URL for the given file name for uploading on S3, It supports huge files
//	@Success		200	{object}	data.GeneratePresignedURLResponse
//	@Failure		400	{object}	httputil.ErrorResponse
//	@Failure		500	{object}	httputil.ErrorResponse
//	@Router			/generate_url [get]
func (h *HTTPHandler) generatePreSignedURLHandler(c *gin.Context) error {
	result, err := h.ohlcService.GeneratePreSignedURL(c)
	if err != nil {
		return err
	}

	return httputil.OK(c, result)
}

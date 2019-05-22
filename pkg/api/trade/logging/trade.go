package trade

import (
	"time"

	"github.com/evzpav/gorsk/pkg/api/trade"
	gorsk "github.com/evzpav/gorsk/pkg/utl/model"
	"github.com/labstack/echo"
)

// New creates new trade logging service
func New(svc trade.Service, logger gorsk.Logger) *LogService {
	return &LogService{
		Service: svc,
		logger:  logger,
	}
}

// LogService represents user logging service
type LogService struct {
	trade.Service
	logger gorsk.Logger
}

const name = "trade"

// Create logging
func (ls *LogService) Create(c echo.Context, req gorsk.Trade) (resp *gorsk.Trade, err error) {
	defer func(begin time.Time) {

		ls.logger.Log(
			c,
			name, "Create Trade request", err,
			map[string]interface{}{
				"req":  req,
				"resp": resp,
				"took": time.Since(begin),
			},
		)
	}(time.Now())
	return ls.Service.Create(c, req)
}

// List logging
func (ls *LogService) List(c echo.Context, req *gorsk.Pagination) (resp []gorsk.Trade, err error) {
	defer func(begin time.Time) {
		ls.logger.Log(
			c,
			name, "List Trade request", err,
			map[string]interface{}{
				"req":    req,
				"resp":   resp,
				"nameok": time.Since(begin),
			},
		)
	}(time.Now())
	return ls.Service.List(c, req)
}

// View logging
func (ls *LogService) View(c echo.Context, req int) (resp *gorsk.Trade, err error) {
	defer func(begin time.Time) {
		ls.logger.Log(
			c,
			name, "View Trade request", err,
			map[string]interface{}{
				"req":  req,
				"resp": resp,
				"took": time.Since(begin),
			},
		)
	}(time.Now())
	return ls.Service.View(c, req)
}

// Delete logging
func (ls *LogService) Delete(c echo.Context, req int) (err error) {
	defer func(begin time.Time) {
		ls.logger.Log(
			c,
			name, "Delete Trade request", err,
			map[string]interface{}{
				"req":  req,
				"took": time.Since(begin),
			},
		)
	}(time.Now())
	return ls.Service.Delete(c, req)
}

// Update logging
func (ls *LogService) Update(c echo.Context, req *trade.Update) (resp *gorsk.Trade, err error) {
	defer func(begin time.Time) {
		ls.logger.Log(
			c,
			name, "Update Trade request", err,
			map[string]interface{}{
				"req":  req,
				"resp": resp,
				"took": time.Since(begin),
			},
		)
	}(time.Now())
	return ls.Service.Update(c, req)
}

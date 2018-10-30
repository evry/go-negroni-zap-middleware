package negroni

import (
	"net/http"
	"time"

	zapStackdriver "github.com/tommy351/zap-stackdriver"
	"github.com/urfave/negroni"
	"go.uber.org/zap"
)

// ZapLogger - Util struct for the request logger
type ZapLogger struct{}

// NewZapSDLogger - Just a thing to return a ZapLogger
func NewZapSDLogger() *ZapLogger {
	return &ZapLogger{}
}

func (h *ZapLogger) ServeHTTP(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	start := time.Now()

	next(rw, r)
	res := rw.(negroni.ResponseWriter)

	sdReq := &zapStackdriver.HTTPRequest{
		Method:             r.Method,
		Referrer:           r.Referer(),
		RemoteIP:           r.RemoteAddr,
		ResponseStatusCode: res.Status(),
		URL:                r.URL.Path,
		UserAgent:          r.UserAgent(),
	}

	zap.S().Infow("Request ", zapStackdriver.LogHTTPRequest(sdReq), zap.Duration("Duration", time.Since(start)))
}

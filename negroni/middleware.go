package negroni

import (
	"net/http"

	zapdriver "github.com/blendle/zapdriver"
	"go.uber.org/zap"
)

// ZapLogger - Util struct for the request logger
type ZapLogger struct{}

// NewZapSDLogger - Just a thing to return a ZapLogger
func NewZapSDLogger() *ZapLogger {
	return &ZapLogger{}
}

func (h *ZapLogger) ServeHTTP(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	next(rw, r)

	sdReq := zapdriver.NewHTTP(r, r.Response)

	zap.S().Infow("Request ", sdReq)
}

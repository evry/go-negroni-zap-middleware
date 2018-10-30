package negroni

import (
	"fmt"
	"net/http"
	"time"

	zapdriver "github.com/blendle/zapdriver"
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

	negroniRw := rw.(negroni.ResponseWriter)

	sdReq := zapdriver.NewHTTP(r, nil)
	sdReq.Status = negroniRw.Status()
	sdReq.Latency = fmt.Sprintf("%ds", time.Since(start))
	zap.S().Infow("Request ", zapdriver.HTTP(sdReq))
}

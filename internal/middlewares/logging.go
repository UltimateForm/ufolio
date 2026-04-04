package middlewares

import (
	"log"
	"net/http"
	"time"
)

type responseWriter struct {
	http.ResponseWriter
	statusCode int
	written    bool
}

func (rw *responseWriter) WriteHeader(code int) {
	if !rw.written {
		rw.statusCode = code
		rw.written = true
		rw.ResponseWriter.WriteHeader(code)
	}
}

func (rw *responseWriter) Write(b []byte) (int, error) {
	if !rw.written {
		rw.statusCode = http.StatusOK // Default if no explicit WriteHeader
		rw.written = true
	}
	return rw.ResponseWriter.Write(b)
}

var middlewareLogger *log.Logger

func init() {
	defaultLogger := log.Default()
	middlewareLogger = log.New(defaultLogger.Writer(), "[LoggingMidleware] ", defaultLogger.Flags())
}

func Logging(next http.HandlerFunc) http.HandlerFunc {
	// consider making next a http.Handler so we call like next.ServeHTTP(w, r)
	return func(w http.ResponseWriter, r *http.Request) {
		middlewareLogger.Printf("REQ %s %s %s\n", r.Method, r.URL.Path, r.RemoteAddr)

		// stategy from https://upgear.io/blog/golang-tip-wrapping-http-response-writer-for-middleware/
		wrapped := &responseWriter{ResponseWriter: w}
		start := time.Now()
		next(wrapped, r)
		duration := time.Since(start)
		middlewareLogger.Printf("RES %s %s %s %v %dms\n", r.Method, r.URL.Path, r.RemoteAddr, wrapped.statusCode, duration.Milliseconds())
	}
}

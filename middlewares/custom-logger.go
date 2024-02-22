package middlewares

import (
	"github.com/justinas/alice"
	"go.uber.org/zap"
	logger "main/errorLogger"
	"net/http"
	"net/http/httptest"
	"strings"
)

type ResponseLogger struct {
	http.ResponseWriter
	body strings.Builder
}

type MiddlewareLogger struct {
	zap    *zap.Logger
	logger *logger.Logger
}

func NewMiddlewareLogger(zap *zap.Logger, logger *logger.Logger) *MiddlewareLogger {
	return &MiddlewareLogger{
		zap:    zap,
		logger: logger,
	}
}

func (m *MiddlewareLogger) CentralisedLogger() alice.Constructor {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			rr := httptest.NewRecorder()
			next.ServeHTTP(rr, r.WithContext(r.Context()))

			respBody := rr.Body.String()
			for k, v := range rr.Header() {
				for _, vv := range v {
					w.Header().Add(k, vv)
				}
			}

			cusError := &logger.Error{}
			_ = logger.Unmarshal(respBody, cusError)
			if cusError.ErrorCode == 0 {
				w.WriteHeader(rr.Code)
				w.Write([]byte(respBody))
				return
			}
			rr.Code = cusError.ErrorCode

			userError := &logger.UserError{
				ErrorCode: cusError.ErrorCode,
				Reason:    cusError.Reason,
			}
			w.Write([]byte(logger.MarshalUserError(userError)))
			m.logger.LogError(cusError)
		})
	}
}

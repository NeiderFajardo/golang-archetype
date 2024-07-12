package middlewares

import (
	"fmt"
	"net/http"

	"github.com/NeiderFajardo/pkg/logger"
	"github.com/NeiderFajardo/pkg/utils"
)

func LogResponse(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ew := utils.ExtendResponseWriter(w)
		next.ServeHTTP(ew, r)
		info := fmt.Sprintf("%s %s responded %v", r.Method, r.URL.Path, ew.StatusCode)
		ew.Done()
		logger.Info(info)
	})
}

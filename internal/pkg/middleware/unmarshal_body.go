package middleware

import (
	"net/http"

	"booking/internal/pkg/logger"
	"booking/internal/pkg/webserver"
	"booking/internal/pkg/webserver/common_handlers"
)

func GetUnmarshalBodyMiddleware(l *logger.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Body == http.NoBody {
				next.ServeHTTP(w, r)
				return
			}

			body, rawReq, err := common_handlers.ParseRequest(r)
			if err != nil {
				common_handlers.SendError(r.Context(), w, err, l)
				return
			}

			ctx := r.Context()
			ctx = webserver.NewContextDataMeta(ctx, rawReq.Meta, rawReq.Data)
			ctx = webserver.NewContextRawBody(ctx, body)

			next.ServeHTTP(w, r.Clone(ctx))
		})
	}
}

package webserver

import (
	"context"
	"net/http"
	"time"

	"booking/internal/pkg/errors"
	"booking/internal/pkg/logger"
	"booking/internal/pkg/webserver/common_handlers"

	jsoniter "github.com/json-iterator/go"
)

func Request[Req any, Resp any](w http.ResponseWriter, r *http.Request, lg *logger.Logger, timeout time.Duration, do func(ctx context.Context, req Req) (Resp, *errors.HTTPError)) {
	ctx := r.Context()

	var req Req
	if err := extractFromContext(ctx, &req); err != nil {
		common_handlers.SendError(ctx, w, errors.NewRequestValidationError(err.Error()), lg)
		return
	}

	var ctxWithTimeout = ctx
	if timeout > 0 {
		var cancel context.CancelFunc
		ctxWithTimeout, cancel = context.WithTimeout(ctx, timeout)
		defer cancel()
	}

	res, dErr := do(ctxWithTimeout, req)
	if dErr != nil {
		common_handlers.SendError(ctx, w, dErr, lg)
		return
	}

	result, err := jsoniter.Marshal(res)
	if err != nil {
		jErr := errors.NewInvalidJSONError(err.Error())
		common_handlers.SendError(ctx, w, jErr, lg)
		return
	}

	common_handlers.SendSuccess(ctx, w, result, lg)
}

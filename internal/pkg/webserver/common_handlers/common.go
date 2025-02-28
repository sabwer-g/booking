package common_handlers

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"booking/internal/pkg/errors"
	"booking/internal/pkg/logger"

	dtoRequest "booking/internal/pkg/dto/request"
	dtoResponse "booking/internal/pkg/dto/response"
)

const (
	contentTypeHeader = "Content-Type"
	contentTypeValue  = "application/json"
)

type ResponseWriter struct {
	http.ResponseWriter
	success bool
	errCode string
}

func NewResponseWriter(w http.ResponseWriter) *ResponseWriter {
	return &ResponseWriter{
		ResponseWriter: w,
		success:        false,
		errCode:        "na",
	}
}

func (r *ResponseWriter) WriteSuccess(success bool) {
	r.success = success
}

func (r *ResponseWriter) SetErrCode(code string) {
	r.errCode = code
}

func SendError(_ context.Context, w http.ResponseWriter, err *errors.HTTPError, lg *logger.Logger) {
	w.Header().Set(contentTypeHeader, contentTypeValue)

	msg := dtoResponse.Response{
		Success: 0,
		Meta:    []byte("{}"),
		Data:    []byte("{}"),
		Error:   err,
	}

	data, _ := json.Marshal(msg)

	if _, err := w.Write(data); err != nil {
		lg.LogErrorf("error occurred while writing error response: %s", err.Error())
		return
	}

	if rw, ok := w.(*ResponseWriter); ok {
		rw.SetErrCode(err.Code)
	}

	lg.LogErrorf("sent not successful response")
}

func SendSuccess(_ context.Context, w http.ResponseWriter, data []byte, lg *logger.Logger) {
	w.Header().Set(contentTypeHeader, contentTypeValue)

	resp := dtoResponse.Response{
		Success: 1,
		Meta:    []byte(`{}`),
		Data:    data,
	}

	respData, err := json.Marshal(resp)
	if err != nil {
		lg.LogErrorf("error occurred while json marshal success response")
	}

	if _, err := w.Write(respData); err != nil {
		lg.LogErrorf("error occurred while writing success response: %s \n%v", err.Error(), data)
	}

	if rw, ok := w.(*ResponseWriter); ok {
		rw.WriteSuccess(true)
	}

	lg.LogInfo("sent successful response")
}

func ParseRequest(r *http.Request) ([]byte, *dtoRequest.Request, *errors.HTTPError) {
	var body []byte
	var err error
	if r.Body != nil {
		if body, err = ioutil.ReadAll(r.Body); err != nil {
			dErr := errors.NewInternalError(err.Error())
			return nil, nil, dErr
		}
	}

	var rawReq dtoRequest.Request
	if err := json.Unmarshal(body, &rawReq); err != nil {
		dErr := errors.NewInvalidJSONError(err.Error())
		return body, nil, dErr
	}

	return body, &rawReq, nil
}

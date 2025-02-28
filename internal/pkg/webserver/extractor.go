package webserver

import (
	"context"
	"encoding/json"

	"github.com/go-playground/validator"
)

const (
	FieldCtxData = iota + 1
	FieldCtxMeta
	FieldCtxRawBody
)

type Validatable interface {
	Validate() error
}

func extractFromContext(ctx context.Context, d interface{}) error {
	rawData := fromContextData(ctx)

	if err := json.Unmarshal(rawData, d); err != nil {
		return err
	}

	if err := validator.New().Struct(d); err != nil {
		return err
	}

	if v, ok := d.(Validatable); ok {
		if err := v.Validate(); err != nil {
			return err
		}
	}

	return nil
}

func fromContextData(ctx context.Context) json.RawMessage {
	if value, ok := ctx.Value(FieldCtxData).(json.RawMessage); ok {
		return value
	}
	return json.RawMessage{}
}

func NewContextRawBody(ctx context.Context, rawBody []byte) context.Context {
	return context.WithValue(ctx, FieldCtxRawBody, rawBody)
}

func NewContextDataMeta(ctx context.Context, meta, data json.RawMessage) context.Context {
	return context.WithValue(
		context.WithValue(ctx, FieldCtxData, data),
		FieldCtxMeta,
		meta,
	)
}

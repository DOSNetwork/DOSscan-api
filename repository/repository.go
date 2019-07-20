package repository

import (
	"context"
)

type Onchain interface {
	GetBalance(ctx context.Context) string
}

type DB interface {
	SaveModel(ctx context.Context, modelType int, eventc chan []interface{}) (err error, errc chan error)
	GetEventsByModelType(ctx context.Context, modelType int, limit, offset int) (result []interface{})
	GetGroupByID(ctx context.Context, id string) interface{}
	GetRequestByID(ctx context.Context, id string) interface{}
	GetNodeByID(ctx context.Context, id string) interface{}
}

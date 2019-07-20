package repository

import (
	"context"
)

type Onchain interface {
	GetBalance(ctx context.Context) string
}

type DB interface {
	SaveModel(ctx context.Context, modelType int, eventc chan []interface{}) chan error
	GetEventsByModel(ctx context.Context, model interface{}) []interface{}
	GetGroupByID(ctx context.Context, id string) interface{}
	GetRequestByID(ctx context.Context, id string) interface{}
	GetNodeByID(ctx context.Context, id string) interface{}
}

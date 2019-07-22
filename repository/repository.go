package repository

import (
	"context"

	"github.com/DOSNetwork/DOSscan-api/models"
)

type Onchain interface {
	GetBalance(ctx context.Context) string
}

type DB interface {
	SaveModel(ctx context.Context, modelType int, eventc chan []interface{}) (err error, errc chan error)
	EventsByModelType(ctx context.Context, modelType int, limit, offset int) (results []interface{}, err error)
	NodeByID(ctx context.Context, id string) (node models.Node, err error)
	GroupByID(ctx context.Context, id string) (group models.Group, err error)
	UrlRequestByID(ctx context.Context, id string) (urlRequest models.UrlRequest, err error)
	RandomRequestByID(ctx context.Context, id string) (randRequest models.UserRandomRequest, err error)
}

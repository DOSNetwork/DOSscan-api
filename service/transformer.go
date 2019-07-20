package service

import (
	"context"

	"github.com/DOSNetwork/DOSscan-api/repository"
)

type Transformer struct {
	repoOnchain repository.onchain
	repoDB      repository.db
}

func NewTransformer(onchainRepo repository.Onchain, dbRepo repository.DB) *transformer {
	return &Transformer{
		onchainRepo: onchainRepo,
		dbRepo:      dbRepo,
	}
}

func (t *Transformer) Transformer(ctx context.Context) {
	_ = ctx
	fmt.Println("Transformer")
}

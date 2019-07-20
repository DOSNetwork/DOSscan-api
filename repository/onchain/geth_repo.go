package onchain

import (
	"context"
	"fmt"

	"github.com/DOSNetwork/DOSscan-api/repository"
	"github.com/ethereum/go-ethereum/ethclient"
)

type gethRepo struct {
	client *ethclient.Client
}

func NewGethRepo(client *ethclient.Client) repository.Onchain {
	return &gethRepo{
		client: client,
	}
}

func (g *gethRepo) GetBalance(ctx context.Context) string {
	_ = ctx
	fmt.Println("GetBalance")
	return ""
}

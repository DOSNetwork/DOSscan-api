package service

import (
	"context"
	"fmt"

	"github.com/DOSNetwork/DOSscan-api/repository"
	"github.com/DOSNetwork/DOSscan-api/utils"
)

const (
	logsLimit = 1000
)

type Transformer struct {
	onchainRepo   repository.Onchain
	dbRepo        repository.DB
	updatedBlknum uint64
}

func NewTransformer(onchainRepo repository.Onchain, dbRepo repository.DB) *Transformer {
	return &Transformer{
		onchainRepo:   onchainRepo,
		dbRepo:        dbRepo,
		updatedBlknum: 4468400,
	}
}
func (t *Transformer) BuildRelations(ctx context.Context) {
	t.dbRepo.BuildRelation(ctx)
}

func (t *Transformer) FetchHistoricalLogs(ctx context.Context, modelsTypes ...int) (error, <-chan error) {
	var errcList []<-chan error
	toBlock, err := t.onchainRepo.CurrentBlockNum(ctx)
	if err != nil {
		fmt.Println("Transformer err ", err)
		return err, nil
	}
	if toBlock-t.updatedBlknum > logsLimit {
		toBlock = t.updatedBlknum + logsLimit
	}

	for _, mType := range modelsTypes {

		err, logsc, fetchErrc := t.onchainRepo.FetchLogs(ctx, mType, t.updatedBlknum, toBlock, logsLimit)
		if err != nil {
			fmt.Println("FetchLogs Err ", err)
			continue
		}
		errcList = append(errcList, fetchErrc)
		err, saveErrc := t.dbRepo.SaveModel(ctx, mType, logsc)
		if err != nil {
			fmt.Println("SaveModel Err ", err)
			continue
		}
		errcList = append(errcList, saveErrc)
	}
	t.updatedBlknum = toBlock
	return nil, utils.MergeErrors(ctx, errcList...)
}

func (t *Transformer) WatchLogs(ctx context.Context, modelsTypes ...int) <-chan error {
	var errcList []<-chan error
	for _, mType := range modelsTypes {
		fromBlock, err := t.dbRepo.LastBlockNum(ctx, mType)
		if err != nil {
			fmt.Println("Transformer err ", err)
		}
		fmt.Println("From blkNum", fromBlock)

		toBlock, err := t.onchainRepo.CurrentBlockNum(ctx)
		if err != nil {
			fmt.Println("Transformer err ", err)
		}
		fmt.Println("To blkNum", toBlock)
		err, logsc, fetchErrc := t.onchainRepo.SubscribeLogs(ctx, mType)
		if err != nil {
			fmt.Println("FetchLogs Err ", err)
			continue
		}
		errcList = append(errcList, fetchErrc)
		err, saveErrc := t.dbRepo.SaveModel(ctx, mType, logsc)
		if err != nil {
			fmt.Println("SaveModel Err ", err)
			continue
		}
		errcList = append(errcList, saveErrc)
	}
	return utils.MergeErrors(ctx, errcList...)
}

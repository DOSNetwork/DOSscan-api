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
	modelsTypes   []int
}

func lastBlcokNum(ctx context.Context, dbRepo repository.DB, modelsTypes []int) (lastBlk uint64) {
	for _, mType := range modelsTypes {
		blkNum, err := dbRepo.LastBlockNum(ctx, mType)
		if err != nil {
			fmt.Println("Transformer err ", err)
			continue
		}
		if lastBlk == 0 {
			lastBlk = blkNum
		} else if blkNum != 0 && lastBlk > blkNum {
			lastBlk = blkNum
		}
	}
	fmt.Println("lastBlcokNum ", lastBlk)
	return
}

func NewTransformer(onchainRepo repository.Onchain, dbRepo repository.DB, updatedBlknum uint64, modelsTypes []int) *Transformer {
	t := &Transformer{
		onchainRepo:   onchainRepo,
		dbRepo:        dbRepo,
		updatedBlknum: updatedBlknum,
		modelsTypes:   modelsTypes,
	}

	if lastBlk := lastBlcokNum(context.Background(), dbRepo, modelsTypes); lastBlk != 0 {
		if t.updatedBlknum > lastBlk {
			t.updatedBlknum = lastBlk
		}
	}
	return t
}

func (t *Transformer) BuildRelations(ctx context.Context) {
	t.dbRepo.BuildRelation(ctx)
}

func (t *Transformer) FetchHistoricalLogs(ctx context.Context) (error, <-chan error) {
	var errcList []<-chan error
	toBlock, err := t.onchainRepo.CurrentBlockNum(ctx)
	if err != nil {
		fmt.Println("Transformer err ", err)
		return err, nil
	}
	if toBlock-t.updatedBlknum > logsLimit {
		toBlock = t.updatedBlknum + logsLimit
	}
	fmt.Println("FetchHistoricalLogs from ", t.updatedBlknum, " to ", toBlock)
	for _, mType := range t.modelsTypes {
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

/*
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
*/

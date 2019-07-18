package dosproxy

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	//	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

var FetchTable = []func(ctx context.Context, fromBlkc chan uint64, toBlock uint64, blockLimit uint64, filter *DosproxyFilterer) (chan interface{}, chan error){
	LogGrouping: func(ctx context.Context, fromBlkc chan uint64, toBlock uint64, blockLimit uint64, filter *DosproxyFilterer) (chan interface{}, chan error) {
		out := make(chan interface{})
		errc := make(chan error)
		go func() {
			defer close(out)
			defer close(errc)
			defer fmt.Println("Done LogGrouping")
			select {
			case <-ctx.Done():
				return
			case fromBlock, ok := <-fromBlkc:
				if !ok {
					return
				}
				if toBlock <= fromBlock {
					return
				}
				for fromBlock <= toBlock {
					nextBlk := toBlock
					if nextBlk-fromBlock > blockLimit {
						nextBlk = fromBlock + blockLimit
					}
					fmt.Println("LogGrouping ", fromBlock, " - ", nextBlk)
					//2) get the historic data from proxy that start from lastBlkNum to latest
					logs, err := filter.FilterLogGrouping(&bind.FilterOpts{Start: fromBlock, End: &nextBlk, Context: ctx})
					if err != nil {
						fmt.Println("FilterLogGrouping err ", err)
					}
					for logs.Next() {
						select {
						case <-ctx.Done():
						case out <- logs.Event:
						}
					}
					fromBlock = nextBlk + 1
				}
				fmt.Println("LogGrouping Done fetch")
			}
		}()
		return out, errc
	},
	LogPublicKeySuggested: func(ctx context.Context, fromBlkc chan uint64, toBlock uint64, blockLimit uint64, filter *DosproxyFilterer) (chan interface{}, chan error) {
		out := make(chan interface{})
		errc := make(chan error)
		go func() {
			defer close(out)
			defer close(errc)
			defer fmt.Println("Done LogPublicKeySuggested")
			select {
			case <-ctx.Done():
				return
			case fromBlock, ok := <-fromBlkc:
				if !ok {
					return
				}
				if toBlock <= fromBlock {
					return
				}
				for fromBlock <= toBlock {
					nextBlk := toBlock
					if nextBlk-fromBlock > blockLimit {
						nextBlk = fromBlock + blockLimit
					}
					fmt.Println("LogPublicKeySuggested ", fromBlock, " - ", nextBlk)
					//get the historic data from proxy that start from lastBlkNum to latest
					logs, err := filter.FilterLogPublicKeySuggested(&bind.FilterOpts{Start: fromBlock, End: &nextBlk, Context: ctx})
					if err != nil {
						fmt.Println("FilterLogPublicKeySuggested err ", err)
					}
					for logs.Next() {
						select {
						case <-ctx.Done():
							return
						case out <- logs.Event:
						}
					}
					fromBlock = nextBlk + 1
				}
			}
		}()
		return out, errc
	},
	LogPublicKeyAccepted: func(ctx context.Context, fromBlkc chan uint64, toBlock uint64, blockLimit uint64, filter *DosproxyFilterer) (chan interface{}, chan error) {
		out := make(chan interface{})
		errc := make(chan error)
		go func() {
			defer close(out)
			defer close(errc)
			defer fmt.Println("Done LogPublicKeyAccepted")
			select {
			case <-ctx.Done():
				return
			case fromBlock, ok := <-fromBlkc:
				if !ok {
					return
				}
				if toBlock <= fromBlock {
					return
				}
				for fromBlock <= toBlock {
					nextBlk := toBlock
					if nextBlk-fromBlock > blockLimit {
						nextBlk = fromBlock + blockLimit
					}
					fmt.Println("LogPublicKeyAccepted ", fromBlock, " - ", nextBlk)
					//get the historic data from proxy that start from lastBlkNum to latest
					logs, err := filter.FilterLogPublicKeyAccepted(&bind.FilterOpts{Start: fromBlock, End: &nextBlk, Context: ctx})
					if err != nil {
						fmt.Println("FilterLogPublicKeyAccepted err ", err)
					}
					for logs.Next() {
						select {
						case <-ctx.Done():
						case out <- logs.Event:
						}
					}
					fromBlock = nextBlk + 1
				}
			}
		}()
		return out, errc
	},
	LogGroupDissolve: func(ctx context.Context, fromBlkc chan uint64, toBlock uint64, blockLimit uint64, filter *DosproxyFilterer) (chan interface{}, chan error) {
		out := make(chan interface{})
		errc := make(chan error)
		go func() {
			defer close(out)
			defer close(errc)
			defer fmt.Println("Done LogGroupDissolve")
			select {
			case <-ctx.Done():
				return
			case fromBlock, ok := <-fromBlkc:
				if !ok {
					return
				}
				if toBlock <= fromBlock {
					return
				}
				for fromBlock <= toBlock {
					nextBlk := toBlock
					if nextBlk-fromBlock > blockLimit {
						nextBlk = fromBlock + blockLimit
					}
					fmt.Println("LogGroupDissolve ", fromBlock, " - ", nextBlk)
					//get the historic data from proxy that start from lastBlkNum to latest
					logs, err := filter.FilterLogGroupDissolve(&bind.FilterOpts{Start: fromBlock, End: &nextBlk, Context: ctx})
					if err != nil {
						fmt.Println("FilterLogGroupDissolve err ", err)
					}
					for logs.Next() {
						select {
						case <-ctx.Done():
						case out <- logs.Event:
						}
					}
					fromBlock = nextBlk + 1

				}
			}
		}()
		return out, errc
	},
	LogRegisteredNewPendingNode: func(ctx context.Context, fromBlkc chan uint64, toBlock uint64, blockLimit uint64, filter *DosproxyFilterer) (chan interface{}, chan error) {
		out := make(chan interface{})
		errc := make(chan error)
		go func() {
			defer close(out)
			defer close(errc)
			defer fmt.Println("Done LogRegisteredNewPendingNode")
			select {
			case <-ctx.Done():
				return
			case fromBlock, ok := <-fromBlkc:
				if !ok {
					return
				}
				if toBlock <= fromBlock {
					return
				}
				for fromBlock <= toBlock {
					nextBlk := toBlock
					if nextBlk-fromBlock > blockLimit {
						nextBlk = fromBlock + blockLimit
					}
					fmt.Println("LogRegisteredNewPendingNode ", fromBlock, " - ", nextBlk)
					//get the historic data from proxy that start from lastBlkNum to latest
					logs, err := filter.FilterLogRegisteredNewPendingNode(&bind.FilterOpts{Start: fromBlock, End: &nextBlk, Context: ctx})
					if err != nil {
						fmt.Println("FilterLogRegisteredNewPendingNode err ", err)
					}
					for logs.Next() {
						select {
						case <-ctx.Done():
						case out <- logs.Event:
						}
					}
					fromBlock = nextBlk + 1

				}
			}
		}()
		return out, errc
	},
	LogUpdateRandom: func(ctx context.Context, fromBlkc chan uint64, toBlock uint64, blockLimit uint64, filter *DosproxyFilterer) (chan interface{}, chan error) {
		out := make(chan interface{})
		errc := make(chan error)
		go func() {
			defer close(out)
			defer close(errc)
			defer fmt.Println("Done LogUpdateRandom")
			select {
			case <-ctx.Done():
				return
			case fromBlock, ok := <-fromBlkc:
				if !ok {
					return
				}
				if toBlock <= fromBlock {
					return
				}
				for fromBlock <= toBlock {
					nextBlk := toBlock
					if nextBlk-fromBlock > blockLimit {
						nextBlk = fromBlock + blockLimit
					}
					fmt.Println("LogUpdateRandom ", fromBlock, " - ", nextBlk)
					logs, err := filter.FilterLogUpdateRandom(&bind.FilterOpts{Start: fromBlock, End: &nextBlk, Context: ctx})
					if err != nil {
						fmt.Println("FilterLogUpdateRandom err ", err)
					}
					for logs.Next() {
						select {
						case <-ctx.Done():
						case out <- logs.Event:
						}
					}
					fromBlock = nextBlk + 1

				}
			}
		}()
		return out, errc
	},
	LogUrl: func(ctx context.Context, fromBlkc chan uint64, toBlock uint64, blockLimit uint64, filter *DosproxyFilterer) (chan interface{}, chan error) {
		out := make(chan interface{})
		errc := make(chan error)
		go func() {
			defer close(out)
			defer close(errc)
			defer fmt.Println("Done LogUrl")
			select {
			case <-ctx.Done():
				return
			case fromBlock, ok := <-fromBlkc:
				if !ok {
					return
				}
				if toBlock <= fromBlock {
					return
				}
				for fromBlock <= toBlock {
					nextBlk := toBlock
					if nextBlk-fromBlock > blockLimit {
						nextBlk = fromBlock + blockLimit
					}
					fmt.Println("LogUrl ", fromBlock, " - ", nextBlk)
					//1) get the historic data from proxy that start from lastBlkNum to latest

					logs, err := filter.FilterLogUrl(&bind.FilterOpts{Start: fromBlock, End: &nextBlk, Context: ctx})
					if err != nil {
						fmt.Println("FilterLogUrl err ", err)
					}
					for logs.Next() {
						select {
						case <-ctx.Done():
						case out <- logs.Event:
						}
					}
					fromBlock = nextBlk + 1

				}
				fmt.Println("LogUrl Done")
			}
		}()
		return out, errc
	},
	LogRequestUserRandom: func(ctx context.Context, fromBlkc chan uint64, toBlock uint64, blockLimit uint64, filter *DosproxyFilterer) (chan interface{}, chan error) {
		out := make(chan interface{})
		errc := make(chan error)
		go func() {
			defer close(out)
			defer close(errc)
			defer fmt.Println("Done LogRequestUserRandom")
			select {
			case <-ctx.Done():
				return
			case fromBlock, ok := <-fromBlkc:
				if !ok {
					return
				}
				if toBlock <= fromBlock {
					return
				}
				for fromBlock <= toBlock {
					nextBlk := toBlock
					if nextBlk-fromBlock > blockLimit {
						nextBlk = fromBlock + blockLimit
					}
					fmt.Println("LogRequestUserRandom ", fromBlock, " - ", nextBlk)
					//get the historic data from proxy that start from lastBlkNum to latest
					logs, err := filter.FilterLogRequestUserRandom(&bind.FilterOpts{Start: fromBlock, End: &nextBlk, Context: ctx})
					if err != nil {
						fmt.Println("FilterLogRequestUserRandom err ", err)
					}
					for logs.Next() {
						select {
						case <-ctx.Done():
						case out <- logs.Event:
						}
					}
					fromBlock = nextBlk + 1

				}
			}
		}()
		return out, errc
	},
	LogValidationResult: func(ctx context.Context, fromBlkc chan uint64, toBlock uint64, blockLimit uint64, filter *DosproxyFilterer) (chan interface{}, chan error) {
		out := make(chan interface{})
		errc := make(chan error)
		go func() {
			defer close(out)
			defer close(errc)
			defer fmt.Println("Done LogValidationResult")
			select {
			case <-ctx.Done():
				return
			case fromBlock, ok := <-fromBlkc:
				if !ok {
					return
				}
				if toBlock <= fromBlock {
					return
				}
				for fromBlock <= toBlock {
					nextBlk := toBlock
					if nextBlk-fromBlock > blockLimit {
						nextBlk = fromBlock + blockLimit
					}
					fmt.Println("LogValidationResult ", fromBlock, " - ", nextBlk)
					//2) get the historic data from proxy that start from lastBlkNum to latest
					logs, err := filter.FilterLogValidationResult(&bind.FilterOpts{Start: fromBlock, End: &nextBlk, Context: ctx})
					if err != nil {
						fmt.Println("FilterLogValidationResult err ", err)
					}
					for logs.Next() {
						select {
						case <-ctx.Done():
						case out <- logs.Event:
						}
					}
					fromBlock = nextBlk + 1

				}
			}
		}()
		return out, errc
	},
	LogCallbackTriggeredFor: func(ctx context.Context, fromBlkc chan uint64, toBlock uint64, blockLimit uint64, filter *DosproxyFilterer) (chan interface{}, chan error) {
		out := make(chan interface{})
		errc := make(chan error)
		go func() {
			defer close(out)
			defer close(errc)
			defer fmt.Println("Done LogCallbackTriggeredFor")
			select {
			case <-ctx.Done():
				return
			case fromBlock, ok := <-fromBlkc:
				if !ok {
					return
				}
				if toBlock <= fromBlock {
					return
				}
				for fromBlock <= toBlock {
					nextBlk := toBlock
					if nextBlk-fromBlock > blockLimit {
						nextBlk = fromBlock + blockLimit
					}
					fmt.Println("LogCallbackTriggeredFor ", fromBlock, " - ", nextBlk)
					logs, err := filter.FilterLogCallbackTriggeredFor(&bind.FilterOpts{Start: fromBlock, End: &nextBlk, Context: ctx})
					if err != nil {
						fmt.Println("FilterLogCallbackTriggeredFor err ", err)
					}
					for logs.Next() {
						select {
						case <-ctx.Done():
						case out <- logs.Event:
						}
					}
					fromBlock = nextBlk + 1

				}
			}
		}()
		return out, errc
	},
	LogError: func(ctx context.Context, fromBlkc chan uint64, toBlock uint64, blockLimit uint64, filter *DosproxyFilterer) (chan interface{}, chan error) {
		out := make(chan interface{})
		errc := make(chan error)
		go func() {
			defer close(out)
			defer close(errc)
			defer fmt.Println("Done LogError")
			select {
			case <-ctx.Done():
				return
			case fromBlock, ok := <-fromBlkc:
				if !ok {
					return
				}
				if toBlock <= fromBlock {
					return
				}
				for fromBlock <= toBlock {
					nextBlk := toBlock
					if nextBlk-fromBlock > blockLimit {
						nextBlk = fromBlock + blockLimit
					}
					fmt.Println("LogError ", fromBlock, " - ", nextBlk)
					logs, err := filter.FilterLogError(&bind.FilterOpts{Start: fromBlock, End: &nextBlk, Context: ctx})
					if err != nil {
						fmt.Println("FilterLogError err ", err)
					}
					for logs.Next() {
						select {
						case <-ctx.Done():
						case out <- logs.Event:
						}
					}
					fromBlock = nextBlk + 1
				}
			}
		}()
		return out, errc
	},
	GuardianReward: func(ctx context.Context, fromBlkc chan uint64, toBlock uint64, blockLimit uint64, filter *DosproxyFilterer) (chan interface{}, chan error) {
		out := make(chan interface{})
		errc := make(chan error)
		go func() {
			defer close(out)
			defer close(errc)
			defer fmt.Println("Done Guardianreward")

			select {
			case <-ctx.Done():
				return
			case fromBlock, ok := <-fromBlkc:
				if !ok {
					return
				}
				if toBlock <= fromBlock {
					return
				}
				for fromBlock <= toBlock {
					nextBlk := toBlock
					if nextBlk-fromBlock > blockLimit {
						nextBlk = fromBlock + blockLimit
					}
					fmt.Println("GuardianReward ", fromBlock, " - ", nextBlk)
					logs, err := filter.FilterGuardianReward(&bind.FilterOpts{Start: fromBlock, End: &nextBlk, Context: ctx}, nil)
					if err != nil {
						fmt.Println("FilterGuardianReward err ", err)
					}
					for logs.Next() {
						select {
						case <-ctx.Done():
						case out <- logs.Event:
						}
					}
					fromBlock = nextBlk + 1

				}
			}
		}()
		return out, errc
	},
}

/*
var SubscriptionTable = []func(ctx context.Context, fromBlkc chan uint64, toBlock uint64, filter *DosproxyFilterer) (chan interface{}, chan error){
	Logurl: func(ctx context.Context, fromBlkc chan uint64, toBlock uint64, filter *DosproxyFilterer) (chan interface{}, chan error) {
		out := make(chan interface{})
		errc := make(chan error)
		go func() {
			var fromBlock uint64
			select {
			case <-ctx.Done():
			case fromBlock = <-fromBlkc:
			}
			ch := make(chan *DosproxyLogUrl)
			opt := &bind.WatchOpts{Context: ctx}
			sub, err := filter.WatchLogUrl(opt, ch)
			if err != nil {
				fmt.Println("WatchLogUrl err ", err)
				return
			}

			//2) get the historic data from proxy that start from lastBlkNum to latest
			pageSize := uint64(1000)
			for fromBlock+pageSize < toBlock {
				toBlock := fromBlock + pageSize
				logs, err := filter.FilterLogUrl(&bind.FilterOpts{Start: fromBlock, End: &toBlock, Context: ctx})
				if err != nil {
					fmt.Println("FilterLogUrl err ", err)
				}
				for logs.Next() {
					select {
					case <-ctx.Done():
					case out <- logs.Event:
					}
				}
				fromBlock = toBlock
			}

			if fromBlock+pageSize >= toBlock {
				logs, err := filter.FilterLogUrl(&bind.FilterOpts{Start: fromBlock, End: &toBlock, Context: ctx})
				if err != nil {
					fmt.Println("FilterLogUrl err ", err)
				}
				for logs.Next() {
					fmt.Println("FilterLogUrl form ")
					select {
					case <-ctx.Done():
					case out <- logs.Event:
					}
				}
			}

			for {
				select {
				case <-ctx.Done():
					return
				case err := <-sub.Err():
					select {
					case <-ctx.Done():
					case errc <- err:
					}
					return
				case event := <-ch:
					select {
					case <-ctx.Done():
						return
					case out <- event:
					}
				}
			}
		}()
		return out, errc
	},
	Logrequestuserrandom: func(ctx context.Context, fromBlkc chan uint64, toBlock uint64, filter *DosproxyFilterer) (chan interface{}, chan error) {
		out := make(chan interface{})
		errc := make(chan error)
		go func() {
			var fromBlock uint64
			select {
			case <-ctx.Done():
			case fromBlock = <-fromBlkc:
			}
			ch := make(chan *DosproxyLogRequestUserRandom)
			opt := &bind.WatchOpts{Context: ctx}
			sub, err := filter.WatchLogRequestUserRandom(opt, ch)
			if err != nil {
				fmt.Println("WatchLogRequestUserRandom err ", err)
				return
			}

			//2) get the historic data from proxy that start from lastBlkNum to latest
			pageSize := uint64(1000)
			for fromBlock+pageSize < toBlock {
				toBlock := fromBlock + pageSize
				logs, err := filter.FilterLogRequestUserRandom(&bind.FilterOpts{Start: fromBlock, End: &toBlock, Context: ctx})
				if err != nil {
					fmt.Println("FilterLogRequestUserRandom err ", err)
				}
				for logs.Next() {
					select {
					case <-ctx.Done():
					case out <- logs.Event:
					}
				}
				fromBlock = toBlock
			}

			if fromBlock+pageSize >= toBlock {
				logs, err := filter.FilterLogRequestUserRandom(&bind.FilterOpts{Start: fromBlock, End: &toBlock, Context: ctx})
				if err != nil {
					fmt.Println("FilterLogRequestUserRandom err ", err)
				}
				for logs.Next() {
					select {
					case <-ctx.Done():
					case out <- logs.Event:
					}
				}
			}

			for {
				select {
				case <-ctx.Done():
					return
				case err := <-sub.Err():
					select {
					case <-ctx.Done():
					case errc <- err:
					}
					return
				case event := <-ch:
					select {
					case <-ctx.Done():
					case out <- event:
					}
				}
			}
		}()
		return out, errc
	},
	Lognonsupportedtype: func(ctx context.Context, fromBlkc chan uint64, toBlock uint64, filter *DosproxyFilterer) (chan interface{}, chan error) {
		out := make(chan interface{})
		errc := make(chan error)
		go func() {
			var fromBlock uint64
			select {
			case <-ctx.Done():
			case fromBlock = <-fromBlkc:
			}
			ch := make(chan *DosproxyLogNonSupportedType)
			opt := &bind.WatchOpts{Context: ctx}
			sub, err := filter.WatchLogNonSupportedType(opt, ch)
			if err != nil {
				fmt.Println("WatchLogNonSupportedType err ", err)
				return
			}

			//2) get the historic data from proxy that start from lastBlkNum to latest
			pageSize := uint64(1000)
			for fromBlock+pageSize < toBlock {
				toBlock := fromBlock + pageSize
				logs, err := filter.FilterLogNonSupportedType(&bind.FilterOpts{Start: fromBlock, End: &toBlock, Context: ctx})
				if err != nil {
					fmt.Println("FilterLogNonSupportedType err ", err)
				}
				for logs.Next() {
					select {
					case <-ctx.Done():
					case out <- logs.Event:
					}
				}
				fromBlock = toBlock
			}

			if fromBlock+pageSize >= toBlock {
				logs, err := filter.FilterLogNonSupportedType(&bind.FilterOpts{Start: fromBlock, End: &toBlock, Context: ctx})
				if err != nil {
					fmt.Println("FilterLogNonSupportedType err ", err)
				}
				for logs.Next() {
					select {
					case <-ctx.Done():
					case out <- logs.Event:
					}
				}
			}

			for {
				select {
				case <-ctx.Done():
					return
				case err := <-sub.Err():
					select {
					case <-ctx.Done():
					case errc <- err:
					}
					return
				case event := <-ch:
					select {
					case <-ctx.Done():
					case out <- event:
					}
				}
			}
		}()
		return out, errc
	},
	Lognoncontractcall: func(ctx context.Context, fromBlkc chan uint64, toBlock uint64, filter *DosproxyFilterer) (chan interface{}, chan error) {
		out := make(chan interface{})
		errc := make(chan error)
		go func() {
			var fromBlock uint64
			select {
			case <-ctx.Done():
			case fromBlock = <-fromBlkc:
			}
			ch := make(chan *DosproxyLogNonContractCall)
			opt := &bind.WatchOpts{Context: ctx}
			sub, err := filter.WatchLogNonContractCall(opt, ch)
			if err != nil {
				fmt.Println("WatchLogNonContractCallerr ", err)
				return
			}

			//2) get the historic data from proxy that start from lastBlkNum to latest
			pageSize := uint64(1000)
			for fromBlock+pageSize < toBlock {
				toBlock := fromBlock + pageSize
				logs, err := filter.FilterLogNonContractCall(&bind.FilterOpts{Start: fromBlock, End: &toBlock, Context: ctx})
				if err != nil {
					fmt.Println("FilterLogNonContractCall err ", err)
				}
				for logs.Next() {
					select {
					case <-ctx.Done():
					case out <- logs.Event:
					}
				}
				fromBlock = toBlock
			}

			if fromBlock+pageSize >= toBlock {
				logs, err := filter.FilterLogNonContractCall(&bind.FilterOpts{Start: fromBlock, End: &toBlock, Context: ctx})
				if err != nil {
					fmt.Println("FilterLogNonContractCall err ", err)
				}
				for logs.Next() {
					select {
					case <-ctx.Done():
					case out <- logs.Event:
					}
				}
			}

			for {
				select {
				case <-ctx.Done():
					return
				case err := <-sub.Err():
					select {
					case <-ctx.Done():
					case errc <- err:
					}
					return
				case event := <-ch:
					select {
					case <-ctx.Done():
					case out <- event:
					}
				}
			}
		}()
		return out, errc
	},
	Logcallbacktriggeredfor: func(ctx context.Context, fromBlkc chan uint64, toBlock uint64, filter *DosproxyFilterer) (chan interface{}, chan error) {
		out := make(chan interface{})
		errc := make(chan error)
		go func() {
			var fromBlock uint64
			select {
			case <-ctx.Done():
			case fromBlock = <-fromBlkc:
			}
			ch := make(chan *DosproxyLogCallbackTriggeredFor)
			opt := &bind.WatchOpts{Context: ctx}
			sub, err := filter.WatchLogCallbackTriggeredFor(opt, ch)
			if err != nil {
				fmt.Println("WatchLogCallbackTriggeredFor err ", err)
				return
			}

			//2) get the historic data from proxy that start from lastBlkNum to latest
			pageSize := uint64(1000)
			for fromBlock+pageSize < toBlock {
				toBlock := fromBlock + pageSize
				logs, err := filter.FilterLogCallbackTriggeredFor(&bind.FilterOpts{Start: fromBlock, End: &toBlock, Context: ctx})
				if err != nil {
					fmt.Println("FilterLogCallbackTriggeredFor err ", err)
				}
				for logs.Next() {
					select {
					case <-ctx.Done():
					case out <- logs.Event:
					}
				}
				fromBlock = toBlock
			}

			if fromBlock+pageSize >= toBlock {
				logs, err := filter.FilterLogCallbackTriggeredFor(&bind.FilterOpts{Start: fromBlock, End: &toBlock, Context: ctx})
				if err != nil {
					fmt.Println("FilterLogCallbackTriggeredFor err ", err)
				}
				for logs.Next() {
					select {
					case <-ctx.Done():
					case out <- logs.Event:
					}
				}
			}

			for {
				select {
				case <-ctx.Done():
					return
				case err := <-sub.Err():
					select {
					case <-ctx.Done():
					case errc <- err:
					}
					return
				case event := <-ch:
					select {
					case <-ctx.Done():
					case out <- event:
					}
				}
			}
		}()
		return out, errc
	},
	Logrequestfromnonexistentuc: func(ctx context.Context, fromBlkc chan uint64, toBlock uint64, filter *DosproxyFilterer) (chan interface{}, chan error) {
		out := make(chan interface{})
		errc := make(chan error)
		go func() {
			var fromBlock uint64
			select {
			case <-ctx.Done():
			case fromBlock = <-fromBlkc:
			}
			ch := make(chan *DosproxyLogRequestFromNonExistentUC)
			opt := &bind.WatchOpts{Context: ctx}
			sub, err := filter.WatchLogRequestFromNonExistentUC(opt, ch)
			if err != nil {
				fmt.Println("WatchLogRequestFromNonExistentUC err ", err)
				return
			}

			//2) get the historic data from proxy that start from lastBlkNum to latest
			pageSize := uint64(1000)
			for fromBlock+pageSize < toBlock {
				toBlock := fromBlock + pageSize
				logs, err := filter.FilterLogRequestFromNonExistentUC(&bind.FilterOpts{Start: fromBlock, End: &toBlock, Context: ctx})
				if err != nil {
					fmt.Println("FilterLogRequestFromNonExistentUC err ", err)
				}
				for logs.Next() {
					select {
					case <-ctx.Done():
					case out <- logs.Event:
					}
				}
				fromBlock = toBlock
			}

			if fromBlock+pageSize >= toBlock {
				logs, err := filter.FilterLogRequestFromNonExistentUC(&bind.FilterOpts{Start: fromBlock, End: &toBlock, Context: ctx})
				if err != nil {
					fmt.Println("FilterLogRequestFromNonExistentUC err ", err)
				}
				for logs.Next() {
					select {
					case <-ctx.Done():
					case out <- logs.Event:
					}
				}
			}

			for {
				select {
				case <-ctx.Done():
					return
				case err := <-sub.Err():
					select {
					case <-ctx.Done():
					case errc <- err:
					}
					return
				case event := <-ch:
					select {
					case <-ctx.Done():
					case out <- event:
					}
				}
			}
		}()
		return out, errc
	},
	Logupdaterandom: func(ctx context.Context, fromBlkc chan uint64, toBlock uint64, filter *DosproxyFilterer) (chan interface{}, chan error) {
		out := make(chan interface{})
		errc := make(chan error)
		go func() {
			var fromBlock uint64
			select {
			case <-ctx.Done():
			case fromBlock = <-fromBlkc:
			}
			ch := make(chan *DosproxyLogUpdateRandom)
			opt := &bind.WatchOpts{Context: ctx}
			sub, err := filter.WatchLogUpdateRandom(opt, ch)
			if err != nil {
				fmt.Println("WatchLogUpdateRandom err ", err)
				return
			}

			//2) get the historic data from proxy that start from lastBlkNum to latest
			pageSize := uint64(1000)
			for fromBlock+pageSize < toBlock {
				toBlock := fromBlock + pageSize
				logs, err := filter.FilterLogUpdateRandom(&bind.FilterOpts{Start: fromBlock, End: &toBlock, Context: ctx})
				if err != nil {
					fmt.Println("FilterLogUpdateRandom err ", err)
				}
				for logs.Next() {
					select {
					case <-ctx.Done():
					case out <- logs.Event:
					}
				}
				fromBlock = toBlock
			}

			if fromBlock+pageSize >= toBlock {
				logs, err := filter.FilterLogUpdateRandom(&bind.FilterOpts{Start: fromBlock, End: &toBlock, Context: ctx})
				if err != nil {
					fmt.Println("FilterLogUpdateRandom err ", err)
				}
				for logs.Next() {
					select {
					case <-ctx.Done():
					case out <- logs.Event:
					}
				}
			}

			for {
				select {
				case <-ctx.Done():
					return
				case err := <-sub.Err():
					select {
					case <-ctx.Done():
					case errc <- err:
					}
					return
				case event := <-ch:
					select {
					case <-ctx.Done():
					case out <- event:
					}
				}
			}
		}()
		return out, errc
	},
	Logvalidationresult: func(ctx context.Context, fromBlkc chan uint64, toBlock uint64, filter *DosproxyFilterer) (chan interface{}, chan error) {
		out := make(chan interface{})
		errc := make(chan error)
		go func() {
			var fromBlock uint64
			select {
			case <-ctx.Done():
			case fromBlock = <-fromBlkc:
			}

			ch := make(chan *DosproxyLogValidationResult)
			opt := &bind.WatchOpts{Context: ctx}
			sub, err := filter.WatchLogValidationResult(opt, ch)
			if err != nil {
				fmt.Println("WatchLogValidationResult err ", err)
				return
			}

			//2) get the historic data from proxy that start from lastBlkNum to latest
			pageSize := uint64(1000)
			for fromBlock+pageSize < toBlock {
				toBlock := fromBlock + pageSize
				logs, err := filter.FilterLogValidationResult(&bind.FilterOpts{Start: fromBlock, End: &toBlock, Context: ctx})
				if err != nil {
					fmt.Println("FilterLogValidationResult err ", err)
				}
				for logs.Next() {
					select {
					case <-ctx.Done():
					case out <- logs.Event:
					}
				}
				fromBlock = toBlock
			}

			if fromBlock+pageSize >= toBlock {
				logs, err := filter.FilterLogValidationResult(&bind.FilterOpts{Start: fromBlock, End: &toBlock, Context: ctx})
				if err != nil {
					fmt.Println("FilterLogValidationResult err ", err)
				}
				for logs.Next() {
					select {
					case <-ctx.Done():
					case out <- logs.Event:
					}
				}
			}

			for {
				select {
				case <-ctx.Done():
					return
				case err := <-sub.Err():
					select {
					case <-ctx.Done():
					case errc <- err:
					}
					return
				case event := <-ch:
					select {
					case <-ctx.Done():
					case out <- event:
					}
				}
			}
		}()
		return out, errc
	},
	Loginsufficientpendingnode: func(ctx context.Context, fromBlkc chan uint64, toBlock uint64, filter *DosproxyFilterer) (chan interface{}, chan error) {
		out := make(chan interface{})
		errc := make(chan error)
		go func() {
			var fromBlock uint64
			select {
			case <-ctx.Done():
			case fromBlock = <-fromBlkc:
			}
			ch := make(chan *DosproxyLogInsufficientPendingNode)
			opt := &bind.WatchOpts{Context: ctx}
			sub, err := filter.WatchLogInsufficientPendingNode(opt, ch)
			if err != nil {
				fmt.Println("WatchLogInsufficientPendingNode err ", err)
				return
			}

			//2) get the historic data from proxy that start from lastBlkNum to latest
			pageSize := uint64(1000)
			for fromBlock+pageSize < toBlock {
				toBlock := fromBlock + pageSize
				logs, err := filter.FilterLogInsufficientPendingNode(&bind.FilterOpts{Start: fromBlock, End: &toBlock, Context: ctx})
				if err != nil {
					fmt.Println("FilterLogInsufficientPendingNode err ", err)
				}
				for logs.Next() {
					select {
					case <-ctx.Done():
					case out <- logs.Event:
					}
				}
				fromBlock = toBlock
			}

			if fromBlock+pageSize >= toBlock {
				logs, err := filter.FilterLogInsufficientPendingNode(&bind.FilterOpts{Start: fromBlock, End: &toBlock, Context: ctx})
				if err != nil {
					fmt.Println("FilterLogInsufficientPendingNode err ", err)
				}
				for logs.Next() {
					select {
					case <-ctx.Done():
					case out <- logs.Event:
					}
				}
			}

			for {
				select {
				case <-ctx.Done():
					return
				case err := <-sub.Err():
					select {
					case <-ctx.Done():
					case errc <- err:
					}
					return
				case event := <-ch:
					select {
					case <-ctx.Done():
					case out <- event:
					}
				}
			}
		}()
		return out, errc
	},
	Loginsufficientworkinggroup: func(ctx context.Context, fromBlkc chan uint64, toBlock uint64, filter *DosproxyFilterer) (chan interface{}, chan error) {
		out := make(chan interface{})
		errc := make(chan error)
		go func() {
			var fromBlock uint64
			select {
			case <-ctx.Done():
			case fromBlock = <-fromBlkc:
			}
			ch := make(chan *DosproxyLogInsufficientWorkingGroup)
			opt := &bind.WatchOpts{Context: ctx}
			sub, err := filter.WatchLogInsufficientWorkingGroup(opt, ch)
			if err != nil {
				fmt.Println("WatchLogInsufficientWorkingGroup err ", err)
				return
			}

			//2) get the historic data from proxy that start from lastBlkNum to latest
			pageSize := uint64(1000)
			for fromBlock+pageSize < toBlock {
				toBlock := fromBlock + pageSize
				logs, err := filter.FilterLogInsufficientWorkingGroup(&bind.FilterOpts{Start: fromBlock, End: &toBlock, Context: ctx})
				if err != nil {
					fmt.Println("FilterLogInsufficientWorkingGroup err ", err)
				}
				for logs.Next() {
					select {
					case <-ctx.Done():
					case out <- logs.Event:
					}
				}
				fromBlock = toBlock
			}

			if fromBlock+pageSize >= toBlock {
				logs, err := filter.FilterLogInsufficientWorkingGroup(&bind.FilterOpts{Start: fromBlock, End: &toBlock, Context: ctx})
				if err != nil {
					fmt.Println("FilterLogInsufficientWorkingGroup err ", err)
				}
				for logs.Next() {
					select {
					case <-ctx.Done():
					case out <- logs.Event:
					}
				}
			}

			for {
				select {
				case <-ctx.Done():
					return
				case err := <-sub.Err():
					select {
					case <-ctx.Done():
					case errc <- err:
					}
					return
				case event := <-ch:
					select {
					case <-ctx.Done():
					case out <- event:
					}
				}
			}
		}()
		return out, errc
	},
	Loggrouping: func(ctx context.Context, fromBlkc chan uint64, toBlock uint64, filter *DosproxyFilterer) (chan interface{}, chan error) {
		out := make(chan interface{})
		errc := make(chan error)
		go func() {
			var fromBlock uint64
			select {
			case <-ctx.Done():
			case fromBlock = <-fromBlkc:
			}
			ch := make(chan *DosproxyLogGrouping)
			opt := &bind.WatchOpts{Context: ctx}
			sub, err := filter.WatchLogGrouping(opt, ch)
			if err != nil {
				fmt.Println("WatchLogGrouping err ", err)
				return
			}

			//2) get the historic data from proxy that start from lastBlkNum to latest
			pageSize := uint64(1000)
			for fromBlock+pageSize < toBlock {
				toBlock := fromBlock + pageSize
				logs, err := filter.FilterLogGrouping(&bind.FilterOpts{Start: fromBlock, End: &toBlock, Context: ctx})
				if err != nil {
					fmt.Println("FilterLogGrouping err ", err)
				}
				for logs.Next() {
					select {
					case <-ctx.Done():
					case out <- logs.Event:
					}
				}
				fromBlock = toBlock
			}

			if fromBlock+pageSize >= toBlock {
				logs, err := filter.FilterLogGrouping(&bind.FilterOpts{Start: fromBlock, End: &toBlock, Context: ctx})
				if err != nil {
					fmt.Println("FilterLogGrouping err ", err)
				}
				for logs.Next() {
					select {
					case <-ctx.Done():
					case out <- logs.Event:
					}
				}
			}

			for {
				select {
				case <-ctx.Done():
					return
				case err := <-sub.Err():
					fmt.Println("LogGrouping err", err)
					select {
					case <-ctx.Done():
					case errc <- err:
					}
					return
				case event := <-ch:
					select {
					case <-ctx.Done():
					case out <- event:
					}
				}
			}
		}()
		return out, errc
	},
	Logpublickeyaccepted: func(ctx context.Context, fromBlkc chan uint64, toBlock uint64, filter *DosproxyFilterer) (chan interface{}, chan error) {
		out := make(chan interface{})
		errc := make(chan error)
		go func() {
			var fromBlock uint64
			select {
			case <-ctx.Done():
			case fromBlock = <-fromBlkc:
			}
			ch := make(chan *DosproxyLogPublicKeyAccepted)
			opt := &bind.WatchOpts{Context: ctx}
			sub, err := filter.WatchLogPublicKeyAccepted(opt, ch)
			if err != nil {
				fmt.Println("WatchLogPublicKeyAccepted err ", err)
				return
			}

			//2) get the historic data from proxy that start from lastBlkNum to latest
			pageSize := uint64(1000)
			for fromBlock+pageSize < toBlock {
				toBlock := fromBlock + pageSize
				logs, err := filter.FilterLogPublicKeyAccepted(&bind.FilterOpts{Start: fromBlock, End: &toBlock, Context: ctx})
				if err != nil {
					fmt.Println("FilterLogPublicKeyAccepted err ", err)
				}
				for logs.Next() {
					select {
					case <-ctx.Done():
					case out <- logs.Event:
					}
				}
				fromBlock = toBlock
			}

			if fromBlock+pageSize >= toBlock {
				logs, err := filter.FilterLogPublicKeyAccepted(&bind.FilterOpts{Start: fromBlock, End: &toBlock, Context: ctx})
				if err != nil {
					fmt.Println("FilterLogPublicKeyAccepted err ", err)
				}
				for logs.Next() {
					select {
					case <-ctx.Done():
					case out <- logs.Event:
					}
				}
			}

			for {
				select {
				case <-ctx.Done():
					return
				case err := <-sub.Err():
					select {
					case <-ctx.Done():
					case errc <- err:
					}
					return
				case event := <-ch:
					select {
					case <-ctx.Done():
					case out <- event:
					}
				}
			}
		}()
		return out, errc
	},
	Logpublickeysuggested: func(ctx context.Context, fromBlkc chan uint64, toBlock uint64, filter *DosproxyFilterer) (chan interface{}, chan error) {
		out := make(chan interface{})
		errc := make(chan error)
		go func() {
			var fromBlock uint64
			select {
			case <-ctx.Done():
			case fromBlock = <-fromBlkc:
			}
			ch := make(chan *DosproxyLogPublicKeySuggested)
			opt := &bind.WatchOpts{Context: ctx}
			sub, err := filter.WatchLogPublicKeySuggested(opt, ch)
			if err != nil {
				fmt.Println("WatchLogPublicKeySuggested err ", err)
				return
			}

			//2) get the historic data from proxy that start from lastBlkNum to latest
			pageSize := uint64(1000)
			for fromBlock+pageSize < toBlock {
				toBlock := fromBlock + pageSize
				logs, err := filter.FilterLogPublicKeySuggested(&bind.FilterOpts{Start: fromBlock, End: &toBlock, Context: ctx})
				if err != nil {
					fmt.Println("FilterLogPublicKeySuggested err ", err)
				}
				for logs.Next() {
					select {
					case <-ctx.Done():
					case out <- logs.Event:
					}
				}
				fromBlock = toBlock
			}

			if fromBlock+pageSize >= toBlock {
				logs, err := filter.FilterLogPublicKeySuggested(&bind.FilterOpts{Start: fromBlock, End: &toBlock, Context: ctx})
				if err != nil {
					fmt.Println("FilterLogPublicKeySuggested err ", err)
				}
				for logs.Next() {
					select {
					case <-ctx.Done():
					case out <- logs.Event:
					}
				}
			}

			for {
				select {
				case <-ctx.Done():
					return
				case err := <-sub.Err():
					select {
					case <-ctx.Done():
					case errc <- err:
					}
					return
				case event := <-ch:
					select {
					case <-ctx.Done():
					case out <- event:
					}
				}
			}
		}()
		return out, errc
	},
	Loggroupdissolve: func(ctx context.Context, fromBlkc chan uint64, toBlock uint64, filter *DosproxyFilterer) (chan interface{}, chan error) {
		out := make(chan interface{})
		errc := make(chan error)
		go func() {
			var fromBlock uint64
			select {
			case <-ctx.Done():
			case fromBlock = <-fromBlkc:
			}
			ch := make(chan *DosproxyLogGroupDissolve)
			opt := &bind.WatchOpts{Context: ctx}
			sub, err := filter.WatchLogGroupDissolve(opt, ch)
			if err != nil {
				fmt.Println("WatchLogGroupDissolve err ", err)
				return
			}

			//2) get the historic data from proxy that start from lastBlkNum to latest
			pageSize := uint64(1000)
			for fromBlock+pageSize < toBlock {
				toBlock := fromBlock + pageSize
				logs, err := filter.FilterLogGroupDissolve(&bind.FilterOpts{Start: fromBlock, End: &toBlock, Context: ctx})
				if err != nil {
					fmt.Println("FilterLogGroupDissolve err ", err)
				}
				for logs.Next() {
					select {
					case <-ctx.Done():
					case out <- logs.Event:
					}
				}
				fromBlock = toBlock
			}

			if fromBlock+pageSize >= toBlock {
				logs, err := filter.FilterLogGroupDissolve(&bind.FilterOpts{Start: fromBlock, End: &toBlock, Context: ctx})
				if err != nil {
					fmt.Println("FilterLogGroupDissolve err ", err)
				}
				go func(logs *DosproxyLogGroupDissolveIterator) {
					for logs.Next() {
						select {
						case <-ctx.Done():
						case out <- logs.Event:
						}
					}
				}(logs)
			}

			for {
				select {
				case <-ctx.Done():
					return
				case err := <-sub.Err():
					select {
					case <-ctx.Done():
					case errc <- err:
					}
					return
				case event := <-ch:
					select {
					case <-ctx.Done():
					case out <- event:
					}
				}
			}
		}()
		return out, errc
	},
	Logregisterednewpendingnode: func(ctx context.Context, fromBlkc chan uint64, toBlock uint64, filter *DosproxyFilterer) (chan interface{}, chan error) {
		out := make(chan interface{})
		errc := make(chan error)
		go func() {
			var fromBlock uint64
			select {
			case <-ctx.Done():
			case fromBlock = <-fromBlkc:
			}
			ch := make(chan *DosproxyLogRegisteredNewPendingNode)
			opt := &bind.WatchOpts{Context: ctx}
			sub, err := filter.WatchLogRegisteredNewPendingNode(opt, ch)
			if err != nil {
				fmt.Println("WatchLogRegisteredNewPendingNode err ", err)
				return
			}

			//2) get the historic data from proxy that start from lastBlkNum to latest
			pageSize := uint64(1000)
			for fromBlock+pageSize < toBlock {
				toBlock := fromBlock + pageSize
				logs, err := filter.FilterLogRegisteredNewPendingNode(&bind.FilterOpts{Start: fromBlock, End: &toBlock, Context: ctx})
				if err != nil {
					fmt.Println("FilterLogRegisteredNewPendingNode err ", err)
				}
				for logs.Next() {
					select {
					case <-ctx.Done():
					case out <- logs.Event:
					}
				}
				fromBlock = toBlock
			}

			if fromBlock+pageSize >= toBlock {
				logs, err := filter.FilterLogRegisteredNewPendingNode(&bind.FilterOpts{Start: fromBlock, End: &toBlock, Context: ctx})
				if err != nil {
					fmt.Println("FilterLogRegisteredNewPendingNode err ", err)
				}
				for logs.Next() {
					select {
					case <-ctx.Done():
					case out <- logs.Event:
					}
				}
			}

			for {
				select {
				case <-ctx.Done():
					return
				case err := <-sub.Err():
					select {
					case <-ctx.Done():
					case errc <- err:
					}
					return
				case event := <-ch:
					select {
					case <-ctx.Done():
					case out <- event:
					}
				}
			}
		}()
		return out, errc
	},
	Loggroupinginitiated: func(ctx context.Context, fromBlkc chan uint64, toBlock uint64, filter *DosproxyFilterer) (chan interface{}, chan error) {
		out := make(chan interface{})
		errc := make(chan error)
		go func() {
			var fromBlock uint64
			select {
			case <-ctx.Done():
			case fromBlock = <-fromBlkc:
			}
			ch := make(chan *DosproxyLogGroupingInitiated)
			opt := &bind.WatchOpts{Context: ctx}
			sub, err := filter.WatchLogGroupingInitiated(opt, ch)
			if err != nil {
				fmt.Println("WatchLogGrouping err ", err)
				return
			}

			//2) get the historic data from proxy that start from lastBlkNum to latest
			pageSize := uint64(1000)
			for fromBlock+pageSize < toBlock {
				toBlock := fromBlock + pageSize
				logs, err := filter.FilterLogGroupingInitiated(&bind.FilterOpts{Start: fromBlock, End: &toBlock, Context: ctx})
				if err != nil {
					fmt.Println("FilterLogGroupingInitiated err ", err)
				}
				for logs.Next() {
					select {
					case <-ctx.Done():
					case out <- logs.Event:
					}
				}
				fromBlock = toBlock
			}

			if fromBlock+pageSize >= toBlock {
				logs, err := filter.FilterLogGroupingInitiated(&bind.FilterOpts{Start: fromBlock, End: &toBlock, Context: ctx})
				if err != nil {
					fmt.Println("FilterLogGroupingInitiated err ", err)
				}
				for logs.Next() {
					select {
					case <-ctx.Done():
					case out <- logs.Event:
					}
				}
			}

			for {
				select {
				case <-ctx.Done():
					return
				case err := <-sub.Err():
					select {
					case <-ctx.Done():
					case errc <- err:
					}
					return
				case event := <-ch:
					select {
					case <-ctx.Done():
					case out <- event:
					}
				}
			}
		}()
		return out, errc
	},
	Lognopendinggroup: func(ctx context.Context, fromBlkc chan uint64, toBlock uint64, filter *DosproxyFilterer) (chan interface{}, chan error) {
		out := make(chan interface{})
		errc := make(chan error)
		go func() {
			var fromBlock uint64
			select {
			case <-ctx.Done():
			case fromBlock = <-fromBlkc:
			}
			ch := make(chan *DosproxyLogNoPendingGroup)
			opt := &bind.WatchOpts{Context: ctx}
			sub, err := filter.WatchLogNoPendingGroup(opt, ch)
			if err != nil {
				fmt.Println("WatchLogNoPendingGroup err ", err)
				return
			}

			//2) get the historic data from proxy that start from lastBlkNum to latest
			pageSize := uint64(1000)
			for fromBlock+pageSize < toBlock {
				toBlock := fromBlock + pageSize
				logs, err := filter.FilterLogNoPendingGroup(&bind.FilterOpts{Start: fromBlock, End: &toBlock, Context: ctx})
				if err != nil {
					fmt.Println("FilterLogNoPendingGroup err ", err)
				}
				for logs.Next() {
					select {
					case <-ctx.Done():
					case out <- logs.Event:
					}
				}
				fromBlock = toBlock
			}

			if fromBlock+pageSize >= toBlock {
				logs, err := filter.FilterLogNoPendingGroup(&bind.FilterOpts{Start: fromBlock, End: &toBlock, Context: ctx})
				if err != nil {
					fmt.Println("FilterLogNoPendingGroup err ", err)
				}
				for logs.Next() {
					select {
					case <-ctx.Done():
					case out <- logs.Event:
					}
				}
			}

			for {
				select {
				case <-ctx.Done():
					return
				case err := <-sub.Err():
					select {
					case <-ctx.Done():
					case errc <- err:
					}
					return
				case event := <-ch:
					select {
					case <-ctx.Done():
					case out <- event:
					}
				}
			}
		}()
		return out, errc
	},
	Logpendinggroupremoved: func(ctx context.Context, fromBlkc chan uint64, toBlock uint64, filter *DosproxyFilterer) (chan interface{}, chan error) {
		out := make(chan interface{})
		errc := make(chan error)
		go func() {
			var fromBlock uint64
			select {
			case <-ctx.Done():
			case fromBlock = <-fromBlkc:
			}
			ch := make(chan *DosproxyLogPendingGroupRemoved)
			opt := &bind.WatchOpts{Context: ctx}
			sub, err := filter.WatchLogPendingGroupRemoved(opt, ch)
			if err != nil {
				fmt.Println("WatchLogPendingGroupRemoved err ", err)
				return
			}

			//2) get the historic data from proxy that start from lastBlkNum to latest
			pageSize := uint64(1000)
			for fromBlock+pageSize < toBlock {
				toBlock := fromBlock + pageSize
				logs, err := filter.FilterLogPendingGroupRemoved(&bind.FilterOpts{Start: fromBlock, End: &toBlock, Context: ctx})
				if err != nil {
					fmt.Println("FilterLogPendingGroupRemoved err ", err)
				}
				for logs.Next() {
					select {
					case <-ctx.Done():
					case out <- logs.Event:
					}
				}
				fromBlock = toBlock
			}

			if fromBlock+pageSize >= toBlock {
				logs, err := filter.FilterLogPendingGroupRemoved(&bind.FilterOpts{Start: fromBlock, End: &toBlock, Context: ctx})
				if err != nil {
					fmt.Println("FilterLogPendingGroupRemoved err ", err)
				}
				for logs.Next() {
					select {
					case <-ctx.Done():
					case out <- logs.Event:
					}
				}
			}

			for {
				select {
				case <-ctx.Done():
					return
				case err := <-sub.Err():
					select {
					case <-ctx.Done():
					case errc <- err:
					}
					return
				case event := <-ch:
					select {
					case <-ctx.Done():
					case out <- event:
					}
				}
			}
		}()
		return out, errc
	},
	Logerror: func(ctx context.Context, fromBlkc chan uint64, toBlock uint64, filter *DosproxyFilterer) (chan interface{}, chan error) {
		out := make(chan interface{})
		errc := make(chan error)
		go func() {
			var fromBlock uint64
			select {
			case <-ctx.Done():
			case fromBlock = <-fromBlkc:
			}
			ch := make(chan *DosproxyLogError)
			opt := &bind.WatchOpts{Context: ctx}
			sub, err := filter.WatchLogError(opt, ch)
			if err != nil {
				fmt.Println("WatchLogError err ", err)
				return
			}

			//2) get the historic data from proxy that start from lastBlkNum to latest
			pageSize := uint64(1000)
			for fromBlock+pageSize < toBlock {
				toBlock := fromBlock + pageSize
				logs, err := filter.FilterLogError(&bind.FilterOpts{Start: fromBlock, End: &toBlock, Context: ctx})
				if err != nil {
					fmt.Println("FilterLogError err ", err)
				}
				for logs.Next() {
					select {
					case <-ctx.Done():
					case out <- logs.Event:
					}
				}
				fromBlock = toBlock
			}

			if fromBlock+pageSize >= toBlock {
				logs, err := filter.FilterLogError(&bind.FilterOpts{Start: fromBlock, End: &toBlock, Context: ctx})
				if err != nil {
					fmt.Println("FilterLogError err ", err)
				}
				for logs.Next() {
					select {
					case <-ctx.Done():
					case out <- logs.Event:
					}
				}
			}

			for {
				select {
				case <-ctx.Done():
					return
				case err := <-sub.Err():
					select {
					case <-ctx.Done():
					case errc <- err:
					}
					return
				case event := <-ch:
					select {
					case <-ctx.Done():
					case out <- event:
					}
				}
			}
		}()
		return out, errc
	},
	Updategrouptopick: func(ctx context.Context, fromBlkc chan uint64, toBlock uint64, filter *DosproxyFilterer) (chan interface{}, chan error) {
		out := make(chan interface{})
		errc := make(chan error)
		go func() {
			var fromBlock uint64
			select {
			case <-ctx.Done():
			case fromBlock = <-fromBlkc:
			}
			ch := make(chan *DosproxyUpdateGroupToPick)
			opt := &bind.WatchOpts{Context: ctx}
			sub, err := filter.WatchUpdateGroupToPick(opt, ch)
			if err != nil {
				fmt.Println("WatchUpdateGroupToPick err ", err)
				return
			}

			//2) get the historic data from proxy that start from lastBlkNum to latest
			pageSize := uint64(1000)
			for fromBlock+pageSize < toBlock {
				toBlock := fromBlock + pageSize
				logs, err := filter.FilterUpdateGroupToPick(&bind.FilterOpts{Start: fromBlock, End: &toBlock, Context: ctx})
				if err != nil {
					fmt.Println("FilterUpdateGroupToPick err ", err)
				}
				for logs.Next() {
					select {
					case <-ctx.Done():
					case out <- logs.Event:
					}
				}
				fromBlock = toBlock
			}

			if fromBlock+pageSize >= toBlock {
				logs, err := filter.FilterUpdateGroupToPick(&bind.FilterOpts{Start: fromBlock, End: &toBlock, Context: ctx})
				if err != nil {
					fmt.Println("FilterUpdateGroupToPick err ", err)
				}
				for logs.Next() {
					select {
					case <-ctx.Done():
					case out <- logs.Event:
					}
				}
			}

			for {
				select {
				case <-ctx.Done():
					return
				case err := <-sub.Err():
					select {
					case <-ctx.Done():
					case errc <- err:
					}
					return
				case event := <-ch:
					select {
					case <-ctx.Done():
					case out <- event:
					}
				}
			}
		}()
		return out, errc
	},
	Updategroupsize: func(ctx context.Context, fromBlkc chan uint64, toBlock uint64, filter *DosproxyFilterer) (chan interface{}, chan error) {
		out := make(chan interface{})
		errc := make(chan error)
		go func() {
			var fromBlock uint64
			select {
			case <-ctx.Done():
			case fromBlock = <-fromBlkc:
			}
			ch := make(chan *DosproxyUpdateGroupSize)
			opt := &bind.WatchOpts{Context: ctx}
			sub, err := filter.WatchUpdateGroupSize(opt, ch)
			if err != nil {
				fmt.Println("WatchUpdateGroupSize err ", err)
				return
			}

			//2) get the historic data from proxy that start from lastBlkNum to latest
			pageSize := uint64(1000)
			for fromBlock+pageSize < toBlock {
				toBlock := fromBlock + pageSize
				logs, err := filter.FilterUpdateGroupSize(&bind.FilterOpts{Start: fromBlock, End: &toBlock, Context: ctx})
				if err != nil {
					fmt.Println("FilterUpdateGroupSize err ", err)
				}
				for logs.Next() {
					select {
					case <-ctx.Done():
					case out <- logs.Event:
					}
				}
				fromBlock = toBlock
			}

			if fromBlock+pageSize >= toBlock {
				logs, err := filter.FilterUpdateGroupSize(&bind.FilterOpts{Start: fromBlock, End: &toBlock, Context: ctx})
				if err != nil {
					fmt.Println("FilterUpdateGroupSize err ", err)
				}
				for logs.Next() {
					select {
					case <-ctx.Done():
					case out <- logs.Event:
					}
				}
			}

			for {
				select {
				case <-ctx.Done():
					return
				case err := <-sub.Err():
					select {
					case <-ctx.Done():
					case errc <- err:
					}
					return
				case event := <-ch:
					select {
					case <-ctx.Done():
					case out <- event:
					}
				}
			}
		}()
		return out, errc
	},
	Updategroupingthreshold: func(ctx context.Context, fromBlkc chan uint64, toBlock uint64, filter *DosproxyFilterer) (chan interface{}, chan error) {
		out := make(chan interface{})
		errc := make(chan error)
		go func() {
			var fromBlock uint64
			select {
			case <-ctx.Done():
			case fromBlock = <-fromBlkc:
			}
			ch := make(chan *DosproxyUpdateGroupingThreshold)
			opt := &bind.WatchOpts{Context: ctx}
			sub, err := filter.WatchUpdateGroupingThreshold(opt, ch)
			if err != nil {
				fmt.Println("WatchUpdateGroupingThreshold err ", err)
				return
			}

			//2) get the historic data from proxy that start from lastBlkNum to latest
			pageSize := uint64(1000)
			for fromBlock+pageSize < toBlock {
				toBlock := fromBlock + pageSize
				logs, err := filter.FilterUpdateGroupingThreshold(&bind.FilterOpts{Start: fromBlock, End: &toBlock, Context: ctx})
				if err != nil {
					fmt.Println("FilterUpdateGroupingThreshold err ", err)
				}
				for logs.Next() {
					select {
					case <-ctx.Done():
					case out <- logs.Event:
					}
				}
				fromBlock = toBlock
			}

			if fromBlock+pageSize >= toBlock {
				logs, err := filter.FilterUpdateGroupingThreshold(&bind.FilterOpts{Start: fromBlock, End: &toBlock, Context: ctx})
				if err != nil {
					fmt.Println("FilterUpdateGroupingThreshold err ", err)
				}
				for logs.Next() {
					select {
					case <-ctx.Done():
					case out <- logs.Event:
					}
				}
			}

			for {
				select {
				case <-ctx.Done():
					return
				case err := <-sub.Err():
					select {
					case <-ctx.Done():
					case errc <- err:
					}
					return
				case event := <-ch:
					select {
					case <-ctx.Done():
					case out <- event:
					}
				}
			}
		}()
		return out, errc
	},
	Updategroupmaturityperiod: func(ctx context.Context, fromBlkc chan uint64, toBlock uint64, filter *DosproxyFilterer) (chan interface{}, chan error) {
		out := make(chan interface{})
		errc := make(chan error)
		go func() {
			var fromBlock uint64
			select {
			case <-ctx.Done():
			case fromBlock = <-fromBlkc:
			}
			ch := make(chan *DosproxyUpdateGroupMaturityPeriod)
			opt := &bind.WatchOpts{Context: ctx}
			sub, err := filter.WatchUpdateGroupMaturityPeriod(opt, ch)
			if err != nil {
				fmt.Println("WatchUpdateGroupMaturityPeriod err ", err)
				return
			}

			//2) get the historic data from proxy that start from lastBlkNum to latest
			pageSize := uint64(1000)
			for fromBlock+pageSize < toBlock {
				toBlock := fromBlock + pageSize
				logs, err := filter.FilterUpdateGroupMaturityPeriod(&bind.FilterOpts{Start: fromBlock, End: &toBlock, Context: ctx})
				if err != nil {
					fmt.Println("FilterUpdateGroupMaturityPeriod err ", err)
				}
				for logs.Next() {
					select {
					case <-ctx.Done():
					case out <- logs.Event:
					}
				}
				fromBlock = toBlock
			}

			if fromBlock+pageSize >= toBlock {
				logs, err := filter.FilterUpdateGroupMaturityPeriod(&bind.FilterOpts{Start: fromBlock, End: &toBlock, Context: ctx})
				if err != nil {
					fmt.Println("FilterUpdateGroupMaturityPeriod err ", err)
				}
				for logs.Next() {
					select {
					case <-ctx.Done():
					case out <- logs.Event:
					}
				}
			}

			for {
				select {
				case <-ctx.Done():
					return
				case err := <-sub.Err():
					select {
					case <-ctx.Done():
					case errc <- err:
					}
					return
				case event := <-ch:
					select {
					case <-ctx.Done():
					case out <- event:
					}
				}
			}
		}()
		return out, errc
	},
	Updatebootstrapcommitduration: func(ctx context.Context, fromBlkc chan uint64, toBlock uint64, filter *DosproxyFilterer) (chan interface{}, chan error) {
		out := make(chan interface{})
		errc := make(chan error)
		go func() {
			var fromBlock uint64
			select {
			case <-ctx.Done():
			case fromBlock = <-fromBlkc:
			}
			ch := make(chan *DosproxyUpdateBootstrapCommitDuration)
			opt := &bind.WatchOpts{Context: ctx}
			sub, err := filter.WatchUpdateBootstrapCommitDuration(opt, ch)
			if err != nil {
				fmt.Println("WatchUpdateBootstrapCommitDuration err ", err)
				return
			}

			//2) get the historic data from proxy that start from lastBlkNum to latest
			pageSize := uint64(1000)
			for fromBlock+pageSize < toBlock {
				toBlock := fromBlock + pageSize
				logs, err := filter.FilterUpdateBootstrapCommitDuration(&bind.FilterOpts{Start: fromBlock, End: &toBlock, Context: ctx})
				if err != nil {
					fmt.Println("FilterUpdateBootstrapCommitDuration err ", err)
				}
				for logs.Next() {
					select {
					case <-ctx.Done():
					case out <- logs.Event:
					}
				}
				fromBlock = toBlock
			}

			if fromBlock+pageSize >= toBlock {
				logs, err := filter.FilterUpdateBootstrapCommitDuration(&bind.FilterOpts{Start: fromBlock, End: &toBlock, Context: ctx})
				if err != nil {
					fmt.Println("FilterUpdateBootstrapCommitDuration err ", err)
				}
				go func(logs *DosproxyUpdateBootstrapCommitDurationIterator) {
					for logs.Next() {
						select {
						case <-ctx.Done():
						case out <- logs.Event:
						}
					}
				}(logs)
			}

			for {
				select {
				case <-ctx.Done():
					return
				case err := <-sub.Err():
					select {
					case <-ctx.Done():
					case errc <- err:
					}
					return
				case event := <-ch:
					select {
					case <-ctx.Done():
					case out <- event:
					}
				}
			}
		}()
		return out, errc
	},
	Updatebootstraprevealduration: func(ctx context.Context, fromBlkc chan uint64, toBlock uint64, filter *DosproxyFilterer) (chan interface{}, chan error) {
		out := make(chan interface{})
		errc := make(chan error)
		go func() {
			var fromBlock uint64
			select {
			case <-ctx.Done():
			case fromBlock = <-fromBlkc:
			}
			ch := make(chan *DosproxyUpdateBootstrapRevealDuration)
			opt := &bind.WatchOpts{Context: ctx}
			sub, err := filter.WatchUpdateBootstrapRevealDuration(opt, ch)
			if err != nil {
				fmt.Println("WatchUpdateBootstrapRevealDuration err ", err)
				return
			}

			//2) get the historic data from proxy that start from lastBlkNum to latest
			pageSize := uint64(1000)
			for fromBlock+pageSize < toBlock {
				toBlock := fromBlock + pageSize
				logs, err := filter.FilterUpdateBootstrapRevealDuration(&bind.FilterOpts{Start: fromBlock, End: &toBlock, Context: ctx})
				if err != nil {
					fmt.Println("FilterUpdateBootstrapRevealDuration err ", err)
				}
				for logs.Next() {
					select {
					case <-ctx.Done():
					case out <- logs.Event:
					}
				}
				fromBlock = toBlock
			}

			if fromBlock+pageSize >= toBlock {
				logs, err := filter.FilterUpdateBootstrapRevealDuration(&bind.FilterOpts{Start: fromBlock, End: &toBlock, Context: ctx})
				if err != nil {
					fmt.Println("FilterUpdateBootstrapRevealDuration err ", err)
				}
				for logs.Next() {
					select {
					case <-ctx.Done():
					case out <- logs.Event:
					}
				}
			}

			for {
				select {
				case <-ctx.Done():
					return
				case err := <-sub.Err():
					select {
					case <-ctx.Done():
					case errc <- err:
					}
					return
				case event := <-ch:
					select {
					case <-ctx.Done():
					case out <- event:
					}
				}
			}
		}()
		return out, errc
	},
	Updatebootstrapstartthreshold: func(ctx context.Context, fromBlkc chan uint64, toBlock uint64, filter *DosproxyFilterer) (chan interface{}, chan error) {
		out := make(chan interface{})
		errc := make(chan error)
		go func() {
			var fromBlock uint64
			select {
			case <-ctx.Done():
			case fromBlock = <-fromBlkc:
			}
			ch := make(chan *DosproxyUpdatebootstrapStartThreshold)
			opt := &bind.WatchOpts{Context: ctx}
			sub, err := filter.WatchUpdatebootstrapStartThreshold(opt, ch)
			if err != nil {
				fmt.Println("WatchUpdatebootstrapStartThreshold err ", err)
				return
			}

			//2) get the historic data from proxy that start from lastBlkNum to latest
			pageSize := uint64(1000)
			for fromBlock+pageSize < toBlock {
				toBlock := fromBlock + pageSize
				logs, err := filter.FilterUpdatebootstrapStartThreshold(&bind.FilterOpts{Start: fromBlock, End: &toBlock, Context: ctx})
				if err != nil {
					fmt.Println("FilterUpdatebootstrapStartThreshold err ", err)
				}
				for logs.Next() {
					select {
					case <-ctx.Done():
					case out <- logs.Event:
					}
				}
				fromBlock = toBlock
			}

			if fromBlock+pageSize >= toBlock {
				logs, err := filter.FilterUpdatebootstrapStartThreshold(&bind.FilterOpts{Start: fromBlock, End: &toBlock, Context: ctx})
				if err != nil {
					fmt.Println("FilterUpdatebootstrapStartThreshold err ", err)
				}
				for logs.Next() {
					select {
					case <-ctx.Done():
					case out <- logs.Event:
					}
				}
			}

			for {
				select {
				case <-ctx.Done():
					return
				case err := <-sub.Err():
					select {
					case <-ctx.Done():
					case errc <- err:
					}
					return
				case event := <-ch:
					select {
					case <-ctx.Done():
					case out <- event:
					}
				}
			}
		}()
		return out, errc
	},
	UpdatependinggroupmaxLife: func(ctx context.Context, fromBlkc chan uint64, toBlock uint64, filter *DosproxyFilterer) (chan interface{}, chan error) {
		out := make(chan interface{})
		errc := make(chan error)
		go func() {
			var fromBlock uint64
			select {
			case <-ctx.Done():
			case fromBlock = <-fromBlkc:
			}
			ch := make(chan *DosproxyUpdatePendingGroupMaxLife)
			opt := &bind.WatchOpts{Context: ctx}
			sub, err := filter.WatchUpdatePendingGroupMaxLife(opt, ch)
			if err != nil {
				fmt.Println("WatchLogGrouping err ", err)
				return
			}

			//2) get the historic data from proxy that start from lastBlkNum to latest
			pageSize := uint64(1000)
			for fromBlock+pageSize < toBlock {
				toBlock := fromBlock + pageSize
				logs, err := filter.FilterUpdatePendingGroupMaxLife(&bind.FilterOpts{Start: fromBlock, End: &toBlock, Context: ctx})
				if err != nil {
					fmt.Println("FilterUpdatePendingGroupMaxLife err ", err)
				}
				for logs.Next() {
					select {
					case <-ctx.Done():
					case out <- logs.Event:
					}
				}
				fromBlock = toBlock
			}

			if fromBlock+pageSize >= toBlock {
				logs, err := filter.FilterUpdatePendingGroupMaxLife(&bind.FilterOpts{Start: fromBlock, End: &toBlock, Context: ctx})
				if err != nil {
					fmt.Println("FilterUpdatePendingGroupMaxLife err ", err)
				}
				for logs.Next() {
					select {
					case <-ctx.Done():
					case out <- logs.Event:
					}
				}
			}

			for {
				select {
				case <-ctx.Done():
					return
				case err := <-sub.Err():
					select {
					case <-ctx.Done():
					case errc <- err:
					}
					return
				case event := <-ch:
					select {
					case <-ctx.Done():
					case out <- event:
					}
				}
			}
		}()
		return out, errc
	},
	Guardianreward: func(ctx context.Context, fromBlkc chan uint64, toBlock uint64, filter *DosproxyFilterer) (chan interface{}, chan error) {
		out := make(chan interface{})
		errc := make(chan error)
		go func() {
			var fromBlock uint64
			select {
			case <-ctx.Done():
			case fromBlock = <-fromBlkc:
			}
			ch := make(chan *DosproxyGuardianReward)
			opt := &bind.WatchOpts{Context: ctx}
			sub, err := filter.WatchGuardianReward(opt, ch, []common.Address{})
			if err != nil {
				fmt.Println("WatchLogGrouping err ", err)
				return
			}

			//2) get the historic data from proxy that start from lastBlkNum to latest
			pageSize := uint64(1000)
			for fromBlock+pageSize < toBlock {
				toBlock := fromBlock + pageSize
				logs, err := filter.FilterGuardianReward(&bind.FilterOpts{Start: fromBlock, End: &toBlock, Context: ctx}, nil)
				if err != nil {
					fmt.Println("FilterGuardianReward err ", err)
				}
				for logs.Next() {
					select {
					case <-ctx.Done():
					case out <- logs.Event:
					}
				}
				fromBlock = toBlock
			}

			if fromBlock+pageSize >= toBlock {
				logs, err := filter.FilterGuardianReward(&bind.FilterOpts{Start: fromBlock, End: &toBlock, Context: ctx}, nil)
				if err != nil {
					fmt.Println("FilterGuardianReward err ", err)
				}
				for logs.Next() {
					select {
					case <-ctx.Done():
					case out <- logs.Event:
					}
				}
			}

			for {
				select {
				case <-ctx.Done():
					return
				case err := <-sub.Err():
					select {
					case <-ctx.Done():
					case errc <- err:
					}
					return
				case event := <-ch:
					select {
					case <-ctx.Done():
					case out <- event:
					}
				}
			}
		}()
		return out, errc
	},
}*/
func NewProxy(proxyAddr common.Address, client *ethclient.Client) (*DosproxySession, error) {
	p, err := NewDosproxy(proxyAddr, client)
	if err != nil {
		return nil, err
	}
	return &DosproxySession{Contract: p, CallOpts: bind.CallOpts{Context: context.Background()}}, nil
}

package dosproxy

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	//	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

var FetchTable = []func(ctx context.Context, lastBlkc chan uint64, filter *DosproxyFilterer) (chan interface{}, chan error){
	0: func(ctx context.Context, lastBlkc chan uint64, filter *DosproxyFilterer) (chan interface{}, chan error) {
		out := make(chan interface{})
		errc := make(chan error)
		go func() {
			blkNum := lastBlock(ctx, lastBlkc)

			ch := make(chan *DosproxyLogUrl)
			opt := &bind.WatchOpts{Context: ctx}
			sub, err := filter.WatchLogUrl(opt, ch)
			if err != nil {
				fmt.Println("WatchLogUrl err ", err)
				return
			}

			//2) get the historic data from proxy that start from lastBlkNum to latest
			logs, err := filter.FilterLogUrl(&bind.FilterOpts{Start: blkNum + 1, End: nil, Context: ctx})
			go func() {
				for logs.Next() {
					select {
					case <-ctx.Done():
					case out <- logs.Event:
					}
				}
			}()

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
	1: func(ctx context.Context, lastBlkc chan uint64, filter *DosproxyFilterer) (chan interface{}, chan error) {
		out := make(chan interface{})
		errc := make(chan error)
		go func() {
			blkNum := lastBlock(ctx, lastBlkc)

			ch := make(chan *DosproxyLogRequestUserRandom)
			opt := &bind.WatchOpts{Context: ctx}
			sub, err := filter.WatchLogRequestUserRandom(opt, ch)
			if err != nil {
				fmt.Println("WatchLogRequestUserRandom err ", err)
				return
			}

			//2) get the historic data from proxy that start from lastBlkNum to latest
			logs, err := filter.FilterLogRequestUserRandom(&bind.FilterOpts{Start: blkNum + 1, End: nil, Context: ctx})
			go func() {
				for logs.Next() {
					select {
					case <-ctx.Done():
					case out <- logs.Event:
					}
				}
			}()

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
	2: func(ctx context.Context, lastBlkc chan uint64, filter *DosproxyFilterer) (chan interface{}, chan error) {
		out := make(chan interface{})
		errc := make(chan error)
		go func() {
			blkNum := lastBlock(ctx, lastBlkc)

			ch := make(chan *DosproxyLogNonSupportedType)
			opt := &bind.WatchOpts{Context: ctx}
			sub, err := filter.WatchLogNonSupportedType(opt, ch)
			if err != nil {
				fmt.Println("WatchLogNonSupportedType err ", err)
				return
			}

			//2) get the historic data from proxy that start from lastBlkNum to latest
			logs, err := filter.FilterLogNonSupportedType(&bind.FilterOpts{Start: blkNum + 1, End: nil, Context: ctx})
			go func() {
				for logs.Next() {
					select {
					case <-ctx.Done():
					case out <- logs.Event:
					}
				}
			}()

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
	3: func(ctx context.Context, lastBlkc chan uint64, filter *DosproxyFilterer) (chan interface{}, chan error) {
		out := make(chan interface{})
		errc := make(chan error)
		go func() {
			blkNum := lastBlock(ctx, lastBlkc)

			ch := make(chan *DosproxyLogNonContractCall)
			opt := &bind.WatchOpts{Context: ctx}
			sub, err := filter.WatchLogNonContractCall(opt, ch)
			if err != nil {
				fmt.Println("WatchLogNonContractCallerr ", err)
				return
			}

			//2) get the historic data from proxy that start from lastBlkNum to latest
			logs, err := filter.FilterLogNonContractCall(&bind.FilterOpts{Start: blkNum + 1, End: nil, Context: ctx})
			go func() {
				for logs.Next() {
					select {
					case <-ctx.Done():
					case out <- logs.Event:
					}
				}
			}()

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
	4: func(ctx context.Context, lastBlkc chan uint64, filter *DosproxyFilterer) (chan interface{}, chan error) {
		out := make(chan interface{})
		errc := make(chan error)
		go func() {
			blkNum := lastBlock(ctx, lastBlkc)

			ch := make(chan *DosproxyLogCallbackTriggeredFor)
			opt := &bind.WatchOpts{Context: ctx}
			sub, err := filter.WatchLogCallbackTriggeredFor(opt, ch)
			if err != nil {
				fmt.Println("WatchLogCallbackTriggeredFor err ", err)
				return
			}

			//2) get the historic data from proxy that start from lastBlkNum to latest
			logs, err := filter.FilterLogCallbackTriggeredFor(&bind.FilterOpts{Start: blkNum + 1, End: nil, Context: ctx})
			go func() {
				for logs.Next() {
					select {
					case <-ctx.Done():
					case out <- logs.Event:
					}
				}
			}()

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
	5: func(ctx context.Context, lastBlkc chan uint64, filter *DosproxyFilterer) (chan interface{}, chan error) {
		out := make(chan interface{})
		errc := make(chan error)
		go func() {
			blkNum := lastBlock(ctx, lastBlkc)

			ch := make(chan *DosproxyLogRequestFromNonExistentUC)
			opt := &bind.WatchOpts{Context: ctx}
			sub, err := filter.WatchLogRequestFromNonExistentUC(opt, ch)
			if err != nil {
				fmt.Println("WatchLogRequestFromNonExistentUC err ", err)
				return
			}

			//2) get the historic data from proxy that start from lastBlkNum to latest
			logs, err := filter.FilterLogRequestFromNonExistentUC(&bind.FilterOpts{Start: blkNum + 1, End: nil, Context: ctx})
			go func() {
				for logs.Next() {
					select {
					case <-ctx.Done():
					case out <- logs.Event:
					}
				}
			}()

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
	6: func(ctx context.Context, lastBlkc chan uint64, filter *DosproxyFilterer) (chan interface{}, chan error) {
		out := make(chan interface{})
		errc := make(chan error)
		go func() {
			blkNum := lastBlock(ctx, lastBlkc)

			ch := make(chan *DosproxyLogUpdateRandom)
			opt := &bind.WatchOpts{Context: ctx}
			sub, err := filter.WatchLogUpdateRandom(opt, ch)
			if err != nil {
				fmt.Println("WatchLogUpdateRandom err ", err)
				return
			}

			//2) get the historic data from proxy that start from lastBlkNum to latest
			logs, err := filter.FilterLogUpdateRandom(&bind.FilterOpts{Start: blkNum + 1, End: nil, Context: ctx})
			go func() {
				for logs.Next() {
					select {
					case <-ctx.Done():
					case out <- logs.Event:
					}
				}
			}()

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
	7: func(ctx context.Context, lastBlkc chan uint64, filter *DosproxyFilterer) (chan interface{}, chan error) {
		out := make(chan interface{})
		errc := make(chan error)
		go func() {
			blkNum := lastBlock(ctx, lastBlkc)

			ch := make(chan *DosproxyLogValidationResult)
			opt := &bind.WatchOpts{Context: ctx}
			sub, err := filter.WatchLogValidationResult(opt, ch)
			if err != nil {
				fmt.Println("WatchLogValidationResult err ", err)
				return
			}

			//2) get the historic data from proxy that start from lastBlkNum to latest
			logs, err := filter.FilterLogValidationResult(&bind.FilterOpts{Start: blkNum + 1, End: nil, Context: ctx})
			go func() {
				for logs.Next() {
					select {
					case <-ctx.Done():
					case out <- logs.Event:
					}
				}
			}()

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
	8: func(ctx context.Context, lastBlkc chan uint64, filter *DosproxyFilterer) (chan interface{}, chan error) {
		out := make(chan interface{})
		errc := make(chan error)
		go func() {
			blkNum := lastBlock(ctx, lastBlkc)

			ch := make(chan *DosproxyLogInsufficientPendingNode)
			opt := &bind.WatchOpts{Context: ctx}
			sub, err := filter.WatchLogInsufficientPendingNode(opt, ch)
			if err != nil {
				fmt.Println("WatchLogInsufficientPendingNode err ", err)
				return
			}

			//2) get the historic data from proxy that start from lastBlkNum to latest
			logs, err := filter.FilterLogInsufficientPendingNode(&bind.FilterOpts{Start: blkNum + 1, End: nil, Context: ctx})
			go func() {
				for logs.Next() {
					select {
					case <-ctx.Done():
					case out <- logs.Event:
					}
				}
			}()

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
	9: func(ctx context.Context, lastBlkc chan uint64, filter *DosproxyFilterer) (chan interface{}, chan error) {
		out := make(chan interface{})
		errc := make(chan error)
		go func() {
			blkNum := lastBlock(ctx, lastBlkc)

			ch := make(chan *DosproxyLogInsufficientWorkingGroup)
			opt := &bind.WatchOpts{Context: ctx}
			sub, err := filter.WatchLogInsufficientWorkingGroup(opt, ch)
			if err != nil {
				fmt.Println("WatchLogInsufficientWorkingGroup err ", err)
				return
			}

			//2) get the historic data from proxy that start from lastBlkNum to latest
			logs, err := filter.FilterLogInsufficientWorkingGroup(&bind.FilterOpts{Start: blkNum + 1, End: nil, Context: ctx})
			go func() {
				for logs.Next() {
					select {
					case <-ctx.Done():
					case out <- logs.Event:
					}
				}
			}()

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
	10: func(ctx context.Context, lastBlkc chan uint64, filter *DosproxyFilterer) (chan interface{}, chan error) {
		out := make(chan interface{})
		errc := make(chan error)
		go func() {
			blkNum := lastBlock(ctx, lastBlkc)

			ch := make(chan *DosproxyLogGrouping)
			opt := &bind.WatchOpts{Context: ctx}
			sub, err := filter.WatchLogGrouping(opt, ch)
			if err != nil {
				fmt.Println("WatchLogGrouping err ", err)
				return
			}

			//2) get the historic data from proxy that start from lastBlkNum to latest
			logs, err := filter.FilterLogGrouping(&bind.FilterOpts{Start: blkNum + 1, End: nil, Context: ctx})
			go func() {
				for logs.Next() {
					select {
					case <-ctx.Done():
					case out <- logs.Event:
					}
				}
			}()

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
					fmt.Println("fetchTable get event")
					select {
					case <-ctx.Done():
					case out <- event:
					}
				}
			}
		}()
		return out, errc
	},
	11: func(ctx context.Context, lastBlkc chan uint64, filter *DosproxyFilterer) (chan interface{}, chan error) {
		out := make(chan interface{})
		errc := make(chan error)
		go func() {
			blkNum := lastBlock(ctx, lastBlkc)

			ch := make(chan *DosproxyLogPublicKeyAccepted)
			opt := &bind.WatchOpts{Context: ctx}
			sub, err := filter.WatchLogPublicKeyAccepted(opt, ch)
			if err != nil {
				fmt.Println("WatchLogPublicKeyAccepted err ", err)
				return
			}

			//2) get the historic data from proxy that start from lastBlkNum to latest
			logs, err := filter.FilterLogPublicKeyAccepted(&bind.FilterOpts{Start: blkNum + 1, End: nil, Context: ctx})
			go func() {
				for logs.Next() {
					select {
					case <-ctx.Done():
					case out <- logs.Event:
					}
				}
			}()

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
	12: func(ctx context.Context, lastBlkc chan uint64, filter *DosproxyFilterer) (chan interface{}, chan error) {
		out := make(chan interface{})
		errc := make(chan error)
		go func() {
			blkNum := lastBlock(ctx, lastBlkc)

			ch := make(chan *DosproxyLogPublicKeySuggested)
			opt := &bind.WatchOpts{Context: ctx}
			sub, err := filter.WatchLogPublicKeySuggested(opt, ch)
			if err != nil {
				fmt.Println("WatchLogPublicKeySuggested err ", err)
				return
			}

			//2) get the historic data from proxy that start from lastBlkNum to latest
			logs, err := filter.FilterLogPublicKeySuggested(&bind.FilterOpts{Start: blkNum + 1, End: nil, Context: ctx})
			go func() {
				for logs.Next() {
					select {
					case <-ctx.Done():
					case out <- logs.Event:
					}
				}
			}()

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
	13: func(ctx context.Context, lastBlkc chan uint64, filter *DosproxyFilterer) (chan interface{}, chan error) {
		out := make(chan interface{})
		errc := make(chan error)
		go func() {
			blkNum := lastBlock(ctx, lastBlkc)

			ch := make(chan *DosproxyLogGroupDissolve)
			opt := &bind.WatchOpts{Context: ctx}
			sub, err := filter.WatchLogGroupDissolve(opt, ch)
			if err != nil {
				fmt.Println("WatchLogGroupDissolve err ", err)
				return
			}

			//2) get the historic data from proxy that start from lastBlkNum to latest
			logs, err := filter.FilterLogGroupDissolve(&bind.FilterOpts{Start: blkNum + 1, End: nil, Context: ctx})
			go func() {
				for logs.Next() {
					select {
					case <-ctx.Done():
					case out <- logs.Event:
					}
				}
			}()

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
	14: func(ctx context.Context, lastBlkc chan uint64, filter *DosproxyFilterer) (chan interface{}, chan error) {
		out := make(chan interface{})
		errc := make(chan error)
		go func() {
			blkNum := lastBlock(ctx, lastBlkc)

			ch := make(chan *DosproxyLogRegisteredNewPendingNode)
			opt := &bind.WatchOpts{Context: ctx}
			sub, err := filter.WatchLogRegisteredNewPendingNode(opt, ch)
			if err != nil {
				fmt.Println("WatchLogRegisteredNewPendingNode err ", err)
				return
			}

			//2) get the historic data from proxy that start from lastBlkNum to latest
			logs, err := filter.FilterLogRegisteredNewPendingNode(&bind.FilterOpts{Start: blkNum + 1, End: nil, Context: ctx})
			go func() {
				for logs.Next() {
					select {
					case <-ctx.Done():
					case out <- logs.Event:
					}
				}
			}()

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
	15: func(ctx context.Context, lastBlkc chan uint64, filter *DosproxyFilterer) (chan interface{}, chan error) {
		out := make(chan interface{})
		errc := make(chan error)
		go func() {
			blkNum := lastBlock(ctx, lastBlkc)

			ch := make(chan *DosproxyLogGroupingInitiated)
			opt := &bind.WatchOpts{Context: ctx}
			sub, err := filter.WatchLogGroupingInitiated(opt, ch)
			if err != nil {
				fmt.Println("WatchLogGrouping err ", err)
				return
			}

			//2) get the historic data from proxy that start from lastBlkNum to latest
			logs, err := filter.FilterLogGroupingInitiated(&bind.FilterOpts{Start: blkNum + 1, End: nil, Context: ctx})
			go func() {
				for logs.Next() {
					select {
					case <-ctx.Done():
					case out <- logs.Event:
					}
				}
			}()

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
	16: func(ctx context.Context, lastBlkc chan uint64, filter *DosproxyFilterer) (chan interface{}, chan error) {
		out := make(chan interface{})
		errc := make(chan error)
		go func() {
			blkNum := lastBlock(ctx, lastBlkc)

			ch := make(chan *DosproxyLogNoPendingGroup)
			opt := &bind.WatchOpts{Context: ctx}
			sub, err := filter.WatchLogNoPendingGroup(opt, ch)
			if err != nil {
				fmt.Println("WatchLogNoPendingGroup err ", err)
				return
			}

			//2) get the historic data from proxy that start from lastBlkNum to latest
			logs, err := filter.FilterLogNoPendingGroup(&bind.FilterOpts{Start: blkNum + 1, End: nil, Context: ctx})
			go func() {
				for logs.Next() {
					select {
					case <-ctx.Done():
					case out <- logs.Event:
					}
				}
			}()

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
	17: func(ctx context.Context, lastBlkc chan uint64, filter *DosproxyFilterer) (chan interface{}, chan error) {
		out := make(chan interface{})
		errc := make(chan error)
		go func() {
			blkNum := lastBlock(ctx, lastBlkc)

			ch := make(chan *DosproxyLogPendingGroupRemoved)
			opt := &bind.WatchOpts{Context: ctx}
			sub, err := filter.WatchLogPendingGroupRemoved(opt, ch)
			if err != nil {
				fmt.Println("WatchLogPendingGroupRemoved err ", err)
				return
			}

			//2) get the historic data from proxy that start from lastBlkNum to latest
			logs, err := filter.FilterLogPendingGroupRemoved(&bind.FilterOpts{Start: blkNum + 1, End: nil, Context: ctx})
			go func() {
				for logs.Next() {
					select {
					case <-ctx.Done():
					case out <- logs.Event:
					}
				}
			}()

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
	18: func(ctx context.Context, lastBlkc chan uint64, filter *DosproxyFilterer) (chan interface{}, chan error) {
		out := make(chan interface{})
		errc := make(chan error)
		go func() {
			blkNum := lastBlock(ctx, lastBlkc)

			ch := make(chan *DosproxyLogError)
			opt := &bind.WatchOpts{Context: ctx}
			sub, err := filter.WatchLogError(opt, ch)
			if err != nil {
				fmt.Println("WatchLogError err ", err)
				return
			}

			//2) get the historic data from proxy that start from lastBlkNum to latest
			logs, err := filter.FilterLogError(&bind.FilterOpts{Start: blkNum + 1, End: nil, Context: ctx})
			go func() {
				for logs.Next() {
					select {
					case <-ctx.Done():
					case out <- logs.Event:
					}
				}
			}()

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
	19: func(ctx context.Context, lastBlkc chan uint64, filter *DosproxyFilterer) (chan interface{}, chan error) {
		out := make(chan interface{})
		errc := make(chan error)
		go func() {
			blkNum := lastBlock(ctx, lastBlkc)

			ch := make(chan *DosproxyUpdateGroupToPick)
			opt := &bind.WatchOpts{Context: ctx}
			sub, err := filter.WatchUpdateGroupToPick(opt, ch)
			if err != nil {
				fmt.Println("WatchUpdateGroupToPick err ", err)
				return
			}

			//2) get the historic data from proxy that start from lastBlkNum to latest
			logs, err := filter.FilterUpdateGroupToPick(&bind.FilterOpts{Start: blkNum + 1, End: nil, Context: ctx})
			go func() {
				for logs.Next() {
					select {
					case <-ctx.Done():
					case out <- logs.Event:
					}
				}
			}()

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
	20: func(ctx context.Context, lastBlkc chan uint64, filter *DosproxyFilterer) (chan interface{}, chan error) {
		out := make(chan interface{})
		errc := make(chan error)
		go func() {
			blkNum := lastBlock(ctx, lastBlkc)

			ch := make(chan *DosproxyUpdateGroupSize)
			opt := &bind.WatchOpts{Context: ctx}
			sub, err := filter.WatchUpdateGroupSize(opt, ch)
			if err != nil {
				fmt.Println("WatchUpdateGroupSize err ", err)
				return
			}

			//2) get the historic data from proxy that start from lastBlkNum to latest
			logs, err := filter.FilterUpdateGroupSize(&bind.FilterOpts{Start: blkNum + 1, End: nil, Context: ctx})
			go func() {
				for logs.Next() {
					select {
					case <-ctx.Done():
					case out <- logs.Event:
					}
				}
			}()

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
	21: func(ctx context.Context, lastBlkc chan uint64, filter *DosproxyFilterer) (chan interface{}, chan error) {
		out := make(chan interface{})
		errc := make(chan error)
		go func() {
			blkNum := lastBlock(ctx, lastBlkc)

			ch := make(chan *DosproxyUpdateGroupingThreshold)
			opt := &bind.WatchOpts{Context: ctx}
			sub, err := filter.WatchUpdateGroupingThreshold(opt, ch)
			if err != nil {
				fmt.Println("WatchUpdateGroupingThreshold err ", err)
				return
			}

			//2) get the historic data from proxy that start from lastBlkNum to latest
			logs, err := filter.FilterUpdateGroupingThreshold(&bind.FilterOpts{Start: blkNum + 1, End: nil, Context: ctx})
			go func() {
				for logs.Next() {
					select {
					case <-ctx.Done():
					case out <- logs.Event:
					}
				}
			}()

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
	22: func(ctx context.Context, lastBlkc chan uint64, filter *DosproxyFilterer) (chan interface{}, chan error) {
		out := make(chan interface{})
		errc := make(chan error)
		go func() {
			blkNum := lastBlock(ctx, lastBlkc)

			ch := make(chan *DosproxyUpdateGroupMaturityPeriod)
			opt := &bind.WatchOpts{Context: ctx}
			sub, err := filter.WatchUpdateGroupMaturityPeriod(opt, ch)
			if err != nil {
				fmt.Println("WatchUpdateGroupMaturityPeriod err ", err)
				return
			}

			//2) get the historic data from proxy that start from lastBlkNum to latest
			logs, err := filter.FilterUpdateGroupMaturityPeriod(&bind.FilterOpts{Start: blkNum + 1, End: nil, Context: ctx})
			go func() {
				for logs.Next() {
					select {
					case <-ctx.Done():
					case out <- logs.Event:
					}
				}
			}()

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
	23: func(ctx context.Context, lastBlkc chan uint64, filter *DosproxyFilterer) (chan interface{}, chan error) {
		out := make(chan interface{})
		errc := make(chan error)
		go func() {
			blkNum := lastBlock(ctx, lastBlkc)

			ch := make(chan *DosproxyUpdateBootstrapCommitDuration)
			opt := &bind.WatchOpts{Context: ctx}
			sub, err := filter.WatchUpdateBootstrapCommitDuration(opt, ch)
			if err != nil {
				fmt.Println("WatchUpdateBootstrapCommitDuration err ", err)
				return
			}

			//2) get the historic data from proxy that start from lastBlkNum to latest
			logs, err := filter.FilterUpdateBootstrapCommitDuration(&bind.FilterOpts{Start: blkNum + 1, End: nil, Context: ctx})
			go func() {
				for logs.Next() {
					select {
					case <-ctx.Done():
					case out <- logs.Event:
					}
				}
			}()

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
	24: func(ctx context.Context, lastBlkc chan uint64, filter *DosproxyFilterer) (chan interface{}, chan error) {
		out := make(chan interface{})
		errc := make(chan error)
		go func() {
			blkNum := lastBlock(ctx, lastBlkc)

			ch := make(chan *DosproxyUpdateBootstrapRevealDuration)
			opt := &bind.WatchOpts{Context: ctx}
			sub, err := filter.WatchUpdateBootstrapRevealDuration(opt, ch)
			if err != nil {
				fmt.Println("WatchUpdateBootstrapRevealDuration err ", err)
				return
			}

			//2) get the historic data from proxy that start from lastBlkNum to latest
			logs, err := filter.FilterUpdateBootstrapRevealDuration(&bind.FilterOpts{Start: blkNum + 1, End: nil, Context: ctx})
			go func() {
				for logs.Next() {
					select {
					case <-ctx.Done():
					case out <- logs.Event:
					}
				}
			}()

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
	25: func(ctx context.Context, lastBlkc chan uint64, filter *DosproxyFilterer) (chan interface{}, chan error) {
		out := make(chan interface{})
		errc := make(chan error)
		go func() {
			blkNum := lastBlock(ctx, lastBlkc)

			ch := make(chan *DosproxyUpdatebootstrapStartThreshold)
			opt := &bind.WatchOpts{Context: ctx}
			sub, err := filter.WatchUpdatebootstrapStartThreshold(opt, ch)
			if err != nil {
				fmt.Println("WatchUpdatebootstrapStartThreshold err ", err)
				return
			}

			//2) get the historic data from proxy that start from lastBlkNum to latest
			logs, err := filter.FilterUpdatebootstrapStartThreshold(&bind.FilterOpts{Start: blkNum + 1, End: nil, Context: ctx})
			go func() {
				for logs.Next() {
					select {
					case <-ctx.Done():
					case out <- logs.Event:
					}
				}
			}()

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
	26: func(ctx context.Context, lastBlkc chan uint64, filter *DosproxyFilterer) (chan interface{}, chan error) {
		out := make(chan interface{})
		errc := make(chan error)
		go func() {
			blkNum := lastBlock(ctx, lastBlkc)

			ch := make(chan *DosproxyUpdatePendingGroupMaxLife)
			opt := &bind.WatchOpts{Context: ctx}
			sub, err := filter.WatchUpdatePendingGroupMaxLife(opt, ch)
			if err != nil {
				fmt.Println("WatchLogGrouping err ", err)
				return
			}

			//2) get the historic data from proxy that start from lastBlkNum to latest
			logs, err := filter.FilterUpdatePendingGroupMaxLife(&bind.FilterOpts{Start: blkNum + 1, End: nil, Context: ctx})
			go func() {
				for logs.Next() {
					select {
					case <-ctx.Done():
					case out <- logs.Event:
					}
				}
			}()

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
	27: func(ctx context.Context, lastBlkc chan uint64, filter *DosproxyFilterer) (chan interface{}, chan error) {
		out := make(chan interface{})
		errc := make(chan error)
		go func() {
			blkNum := lastBlock(ctx, lastBlkc)

			ch := make(chan *DosproxyGuardianReward)
			opt := &bind.WatchOpts{Context: ctx}
			sub, err := filter.WatchGuardianReward(opt, ch, []common.Address{})
			if err != nil {
				fmt.Println("WatchLogGrouping err ", err)
				return
			}

			//2) get the historic data from proxy that start from lastBlkNum to latest
			logs, err := filter.FilterGuardianReward(&bind.FilterOpts{Start: blkNum + 1, End: nil, Context: ctx}, []common.Address{})
			go func() {
				for logs.Next() {
					select {
					case <-ctx.Done():
					case out <- logs.Event:
					}
				}
			}()

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
}

func lastBlock(ctx context.Context, lastBlkc chan uint64) (blkNum uint64) {
	select {
	case <-ctx.Done():
		return
	case blkNum = <-lastBlkc:
	}
	return
}
func NewProxy(proxyAddr common.Address, client *ethclient.Client) (*DosproxySession, error) {
	p, err := NewDosproxy(proxyAddr, client)
	if err != nil {
		return nil, err
	}
	return &DosproxySession{Contract: p, CallOpts: bind.CallOpts{Context: context.Background()}}, nil
}

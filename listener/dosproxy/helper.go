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
		}()
		return out, errc
	},
	1: func(ctx context.Context, lastBlkc chan uint64, filter *DosproxyFilterer) (chan interface{}, chan error) {
		out := make(chan interface{})
		errc := make(chan error)
		go func() {
		}()
		return out, errc
	},
	2: func(ctx context.Context, lastBlkc chan uint64, filter *DosproxyFilterer) (chan interface{}, chan error) {
		out := make(chan interface{})
		errc := make(chan error)
		go func() {
		}()
		return out, errc
	},
	3: func(ctx context.Context, lastBlkc chan uint64, filter *DosproxyFilterer) (chan interface{}, chan error) {
		out := make(chan interface{})
		errc := make(chan error)
		go func() {
		}()
		return out, errc
	},
	4: func(ctx context.Context, lastBlkc chan uint64, filter *DosproxyFilterer) (chan interface{}, chan error) {
		out := make(chan interface{})
		errc := make(chan error)
		go func() {
		}()
		return out, errc
	},
	5: func(ctx context.Context, lastBlkc chan uint64, filter *DosproxyFilterer) (chan interface{}, chan error) {
		out := make(chan interface{})
		errc := make(chan error)
		go func() {
		}()
		return out, errc
	},
	6: func(ctx context.Context, lastBlkc chan uint64, filter *DosproxyFilterer) (chan interface{}, chan error) {
		out := make(chan interface{})
		errc := make(chan error)
		go func() {
		}()
		return out, errc
	},
	7: func(ctx context.Context, lastBlkc chan uint64, filter *DosproxyFilterer) (chan interface{}, chan error) {
		out := make(chan interface{})
		errc := make(chan error)
		go func() {
		}()
		return out, errc
	},
	8: func(ctx context.Context, lastBlkc chan uint64, filter *DosproxyFilterer) (chan interface{}, chan error) {
		out := make(chan interface{})
		errc := make(chan error)
		go func() {
		}()
		return out, errc
	},
	9: func(ctx context.Context, lastBlkc chan uint64, filter *DosproxyFilterer) (chan interface{}, chan error) {
		out := make(chan interface{})
		errc := make(chan error)
		go func() {
		}()
		return out, errc
	},
	10: func(ctx context.Context, lastBlkc chan uint64, filter *DosproxyFilterer) (chan interface{}, chan error) {
		out := make(chan interface{})
		errc := make(chan error)
		go func() {
			var blkNum uint64
			select {
			case <-ctx.Done():
				return
			case blkNum = <-lastBlkc:
			}
			fmt.Println("fetchTable ", blkNum)
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
		}()
		return out, errc
	},
	12: func(ctx context.Context, lastBlkc chan uint64, filter *DosproxyFilterer) (chan interface{}, chan error) {
		out := make(chan interface{})
		errc := make(chan error)
		go func() {
		}()
		return out, errc
	},
	13: func(ctx context.Context, lastBlkc chan uint64, filter *DosproxyFilterer) (chan interface{}, chan error) {
		out := make(chan interface{})
		errc := make(chan error)
		go func() {
		}()
		return out, errc
	},
	14: func(ctx context.Context, lastBlkc chan uint64, filter *DosproxyFilterer) (chan interface{}, chan error) {
		out := make(chan interface{})
		errc := make(chan error)
		go func() {
		}()
		return out, errc
	},
	15: func(ctx context.Context, lastBlkc chan uint64, filter *DosproxyFilterer) (chan interface{}, chan error) {
		out := make(chan interface{})
		errc := make(chan error)
		go func() {
		}()
		return out, errc
	},
	16: func(ctx context.Context, lastBlkc chan uint64, filter *DosproxyFilterer) (chan interface{}, chan error) {
		out := make(chan interface{})
		errc := make(chan error)
		go func() {
		}()
		return out, errc
	},
	17: func(ctx context.Context, lastBlkc chan uint64, filter *DosproxyFilterer) (chan interface{}, chan error) {
		out := make(chan interface{})
		errc := make(chan error)
		go func() {
		}()
		return out, errc
	},
	18: func(ctx context.Context, lastBlkc chan uint64, filter *DosproxyFilterer) (chan interface{}, chan error) {
		out := make(chan interface{})
		errc := make(chan error)
		go func() {
		}()
		return out, errc
	},
	19: func(ctx context.Context, lastBlkc chan uint64, filter *DosproxyFilterer) (chan interface{}, chan error) {
		out := make(chan interface{})
		errc := make(chan error)
		go func() {
		}()
		return out, errc
	},
	20: func(ctx context.Context, lastBlkc chan uint64, filter *DosproxyFilterer) (chan interface{}, chan error) {
		out := make(chan interface{})
		errc := make(chan error)
		go func() {
		}()
		return out, errc
	},
	21: func(ctx context.Context, lastBlkc chan uint64, filter *DosproxyFilterer) (chan interface{}, chan error) {
		out := make(chan interface{})
		errc := make(chan error)
		go func() {
		}()
		return out, errc
	},
	22: func(ctx context.Context, lastBlkc chan uint64, filter *DosproxyFilterer) (chan interface{}, chan error) {
		out := make(chan interface{})
		errc := make(chan error)
		go func() {
		}()
		return out, errc
	},
	23: func(ctx context.Context, lastBlkc chan uint64, filter *DosproxyFilterer) (chan interface{}, chan error) {
		out := make(chan interface{})
		errc := make(chan error)
		go func() {
		}()
		return out, errc
	},
	24: func(ctx context.Context, lastBlkc chan uint64, filter *DosproxyFilterer) (chan interface{}, chan error) {
		out := make(chan interface{})
		errc := make(chan error)
		go func() {
		}()
		return out, errc
	},
	25: func(ctx context.Context, lastBlkc chan uint64, filter *DosproxyFilterer) (chan interface{}, chan error) {
		out := make(chan interface{})
		errc := make(chan error)
		go func() {
		}()
		return out, errc
	},
	26: func(ctx context.Context, lastBlkc chan uint64, filter *DosproxyFilterer) (chan interface{}, chan error) {
		out := make(chan interface{})
		errc := make(chan error)
		go func() {
		}()
		return out, errc
	},
	27: func(ctx context.Context, lastBlkc chan uint64, filter *DosproxyFilterer) (chan interface{}, chan error) {
		out := make(chan interface{})
		errc := make(chan error)
		go func() {
		}()
		return out, errc
	},
}

func NewProxy(proxyAddr common.Address, client *ethclient.Client) (*DosproxySession, error) {
	p, err := NewDosproxy(proxyAddr, client)
	if err != nil {
		return nil, err
	}
	return &DosproxySession{Contract: p, CallOpts: bind.CallOpts{Context: context.Background()}}, nil
}

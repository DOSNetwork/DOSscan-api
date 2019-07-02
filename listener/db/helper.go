package db

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/DOSNetwork/explorer-Api/listener/dosproxy"

	"github.com/lib/pq"
)

/*
DROP TABLE IF EXISTS grouping;
CREATE TABLE grouping (
   id SERIAL NOT NULL,
   tx TEXT NOT NULL,
   blocknumber BIGINT NOT NULL,
   removed BOOLEAN NOT NULL,
   groupid TEXT,
   nodeId TEXT[],
   PRIMARY KEY (id)
);
*/
func LastBlk(ctx context.Context, event string, db *sql.DB) (chan uint64, chan error) {
	out := make(chan uint64)
	errc := make(chan error)
	go func() {
		var lastBlkNum uint64
		latestRecord := fmt.Sprintf("SELECT blocknumber FROM %s ORDER BY blocknumber DESC LIMIT 1;", "grouping")
		err := db.QueryRow(latestRecord).Scan(&lastBlkNum)
		if err != nil {
			lastBlkNum = 0
		}
		select {
		case <-ctx.Done():
		case out <- lastBlkNum:
		}
	}()
	return out, errc
}

var ProxyTable = []func(ctx context.Context, eventc chan interface{}, db *sql.DB) chan error{
	0: func(ctx context.Context, eventc chan interface{}, db *sql.DB) chan error {
		errc := make(chan error)
		go func() {
		}()
		return errc
	},
	1: func(ctx context.Context, eventc chan interface{}, db *sql.DB) chan error {
		errc := make(chan error)
		go func() {
		}()
		return errc
	},
	2: func(ctx context.Context, eventc chan interface{}, db *sql.DB) chan error {
		errc := make(chan error)
		go func() {
		}()
		return errc
	},
	3: func(ctx context.Context, eventc chan interface{}, db *sql.DB) chan error {
		errc := make(chan error)
		go func() {
		}()
		return errc
	},
	4: func(ctx context.Context, eventc chan interface{}, db *sql.DB) chan error {
		errc := make(chan error)
		go func() {
		}()
		return errc
	},
	5: func(ctx context.Context, eventc chan interface{}, db *sql.DB) chan error {
		errc := make(chan error)
		go func() {
		}()
		return errc
	},
	6: func(ctx context.Context, eventc chan interface{}, db *sql.DB) chan error {
		errc := make(chan error)
		go func() {
		}()
		return errc
	},
	7: func(ctx context.Context, eventc chan interface{}, db *sql.DB) chan error {
		errc := make(chan error)
		go func() {
		}()
		return errc
	},
	8: func(ctx context.Context, eventc chan interface{}, db *sql.DB) chan error {
		errc := make(chan error)
		go func() {
		}()
		return errc
	},
	9: func(ctx context.Context, eventc chan interface{}, db *sql.DB) chan error {
		errc := make(chan error)
		go func() {
		}()
		return errc
	},
	10: func(ctx context.Context, eventc chan interface{}, db *sql.DB) chan error {
		errc := make(chan error)
		go func() {
			for {
				select {
				case <-ctx.Done():
				case event := <-eventc:
					e, ok := event.(*dosproxy.DosproxyLogGrouping)
					if !ok {
						fmt.Println("saveTable get event !ok")
						return
					}
					var nodeid []string
					for _, node := range e.NodeId {
						nodeid = append(nodeid, fmt.Sprintf("%x", node.Bytes()))
					}
					var lastInsertId string
					err := db.QueryRow("INSERT INTO grouping(tx,blocknumber,removed,groupid,nodeId) VALUES($1,$2,$3,$4,$5) returning id;", fmt.Sprintf("%x", e.Raw.TxHash.Big()), e.Raw.BlockNumber, e.Raw.Removed, fmt.Sprintf("%x", e.GroupId), pq.Array(nodeid)).Scan(&lastInsertId)
					if err != nil {
						fmt.Println(":INSERT INTO grouping err ", err)
						select {
						case errc <- err:
						case <-ctx.Done():
						}
						return
					}
				}
			}
		}()
		return errc
	},
	11: func(ctx context.Context, eventc chan interface{}, db *sql.DB) chan error {
		errc := make(chan error)
		go func() {
		}()
		return errc
	},
	12: func(ctx context.Context, eventc chan interface{}, db *sql.DB) chan error {
		errc := make(chan error)
		go func() {
		}()
		return errc
	},
	13: func(ctx context.Context, eventc chan interface{}, db *sql.DB) chan error {
		errc := make(chan error)
		go func() {
		}()
		return errc
	},
	14: func(ctx context.Context, eventc chan interface{}, db *sql.DB) chan error {
		errc := make(chan error)
		go func() {
		}()
		return errc
	},
	15: func(ctx context.Context, eventc chan interface{}, db *sql.DB) chan error {
		errc := make(chan error)
		go func() {
		}()
		return errc
	},
	16: func(ctx context.Context, eventc chan interface{}, db *sql.DB) chan error {
		errc := make(chan error)
		go func() {
		}()
		return errc
	},
	17: func(ctx context.Context, eventc chan interface{}, db *sql.DB) chan error {
		errc := make(chan error)
		go func() {
		}()
		return errc
	},
	18: func(ctx context.Context, eventc chan interface{}, db *sql.DB) chan error {
		errc := make(chan error)
		go func() {
		}()
		return errc
	},
	19: func(ctx context.Context, eventc chan interface{}, db *sql.DB) chan error {
		errc := make(chan error)
		go func() {
		}()
		return errc
	},
	20: func(ctx context.Context, eventc chan interface{}, db *sql.DB) chan error {
		errc := make(chan error)
		go func() {
		}()
		return errc
	},
	21: func(ctx context.Context, eventc chan interface{}, db *sql.DB) chan error {
		errc := make(chan error)
		go func() {
		}()
		return errc
	},
	22: func(ctx context.Context, eventc chan interface{}, db *sql.DB) chan error {
		errc := make(chan error)
		go func() {
		}()
		return errc
	},
	23: func(ctx context.Context, eventc chan interface{}, db *sql.DB) chan error {
		errc := make(chan error)
		go func() {
		}()
		return errc
	},
	24: func(ctx context.Context, eventc chan interface{}, db *sql.DB) chan error {
		errc := make(chan error)
		go func() {
		}()
		return errc
	},
	25: func(ctx context.Context, eventc chan interface{}, db *sql.DB) chan error {
		errc := make(chan error)
		go func() {
		}()
		return errc
	},
	26: func(ctx context.Context, eventc chan interface{}, db *sql.DB) chan error {
		errc := make(chan error)
		go func() {
		}()
		return errc
	},
	27: func(ctx context.Context, eventc chan interface{}, db *sql.DB) chan error {
		errc := make(chan error)
		go func() {
		}()
		return errc
	},
}

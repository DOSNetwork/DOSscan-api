# Usage:
# $ 

CORE := ${GOPATH}/src/github.com/DOSNetwork/core

sync:
	cp ${CORE}/onchain/dosbridge/*.abi abi/
	cp ${CORE}/onchain/dospayment/*.abi abi/
	cp ${CORE}/onchain/dosstaking/*.abi abi/
	cp ${CORE}/onchain/dosproxy/*.abi abi/
	cp ${CORE}/onchain/commitreveal/*.abi abi/

gen: sync
	abigen --abi "abi/DOSAddressBridge.abi" --pkg dosbridge --out "models/dosbridge/DOSAddressBridge.go"
	abigen --abi "abi/DOSPayment.abi" --pkg dospayment --out "models/dospayment/DOSPayment.go"
	abigen --abi "abi/DOSProxy.abi" --pkg dosproxy --out "models/dosproxy/DOSProxy.go"
	abigen --abi "abi/Staking.abi" --pkg dosstaking --out "models/dosstaking/Staking.go"
	abigen --abi "abi/CommitReveal.abi" --pkg commitreveal --out "models/commitreveal/CommitReveal.go"

check-env:
ifeq ($(GETHURL),)
	echo "Please set GETHURL in environment variable"
	exit 1;
endif

build:
	go build -o subsciber subscriber/main.go
	go build -o server server/main.go

clean:
	dropdb dev
	createdb dev
	rm subsciber
	rm server

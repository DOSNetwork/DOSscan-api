.DEFAULT_GOAL := build

CORE := ${GOPATH}/src/github.com/DOSNetwork/core

sync:
	cp ${CORE}/onchain/dosbridge/*.abi abi/
	cp ${CORE}/onchain/dospayment/*.abi abi/
	cp ${CORE}/onchain/dosstaking/*.abi abi/
	cp ${CORE}/onchain/dosproxy/*.abi abi/
	cp ${CORE}/onchain/commitreveal/*.abi abi/
	cp ${CORE}/onchain/dosbridge/DOSAddressBridge.go models/dosbridge/DOSAddressBridge.go
	cp ${CORE}/onchain/dospayment/DOSPayment.go models/dospayment/DOSPayment.go
	cp ${CORE}/onchain/dosproxy/DOSProxy.go models/dosproxy/DOSProxy.go
	cp ${CORE}/onchain/dosstaking/Staking.go models/dosstaking/Staking.go
	cp ${CORE}/onchain/commitreveal/CommitReveal.go models/commitreveal/CommitReveal.go

check-env:
ifeq ($(GETHURL),)
	echo "Please set GETHURL in environment variable"
	exit 1;
endif

build: sync
	go build -o server/server server/main.go
	go build -o subscriber/subscriber subscriber/main.go

clean:
	rm -f subscriber/subscriber
	rm -f server/server

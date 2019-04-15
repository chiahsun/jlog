CUR_GOPATH=$(PWD)/../../../../../../../go

run:
	export GOPATH=$(CUR_GOPATH) && go run main/main.go

dep:
	export GOPATH=$(CUR_GOPATH) && dep ensure

test:
	export GOPATH=$(CUR_GOPATH) && go test --race

benchmark:
	export GOPATH=$(CUR_GOPATH) && go test -bench=.
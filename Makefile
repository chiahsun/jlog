all:
	echo "Put this makefile in your gopath root"

run:
	export GOPATH=$(PWD); go run src/192.168.12.16/Source/IM/Jello/jlog/main/main.go

dep:
	export GOPATH=$(PWD) && cd src/192.168.12.16/Source/IM/Jello/jlog && dep ensure

test:
	export GOPATH=$(PWD) && cd src/192.168.12.16/Source/IM/Jello/jlog && go test

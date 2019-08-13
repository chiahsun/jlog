.PHONY: all run test benchmark init

all:

run:
	go run main/main.go

test:
	go test --race

benchmark:
	go test -bench=.

init:
	go mod init 192.168.12.16/Source/IM/Jello/jlog

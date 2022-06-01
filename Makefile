LDFLAGS=-ldflags "-s -w -X main.Version=$(shell git describe --abbrev=0 --always --tags)"

build:
	#Linux
	GOOS=linux GOARCH=amd64 go build -o bin/shvtoanki ${LDFLAGS} ./cmd/shvToanki/main.go

	#Windows
	GOOS=windows GOARCH=amd64 go build -o bin/shvtoanki.exe ${LDFLAGS} ./cmd/shvToanki/main.go

clean:
	rm -f bin/*
	rm -f *.xml
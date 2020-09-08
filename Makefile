all: build

BIN_DIR = $(PWD)/bin

clean:
	rm -rf ${BIN_DIR}

dependencies:
	GO111MODULE=on go mod download

build: dependencies build-server

build-server:
	CGO_ENABLED=0 GO111MODULE=on go build -o ${BIN_DIR}/gan-server cmd/gan-server/main.go

ci: clean dependencies test

compile:
	# FreeBDS
	CGO_ENABLED=0 GO111MODULE=on GOOS=freebsd GOARCH=amd64 go build -o ${BIN_DIR}/gan-server-freebsd-amd64 cmd/gan-server/main.go
	# MacOS
	CGO_ENABLED=0 GO111MODULE=on GOOS=darwin GOARCH=amd64 go build -o ${BIN_DIR}/gan-server-darwin-amd64 cmd/gan-server/main.go
	# Linux
	CGO_ENABLED=0 GO111MODULE=on GOOS=linux GOARCH=amd64 go build -o ${BIN_DIR}/gan-server-linux-amd64 cmd/gan-server/main.go
	# Windows
	CGO_ENABLED=0 GO111MODULE=on GOOS=windows GOARCH=amd64 go build -o ${BIN_DIR}/gan-server-windows-amd64 cmd/gan-server/main.go

test:
	GO111MODULE=on go test -tags testing ./...
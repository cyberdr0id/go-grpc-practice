SHELL := powershell.exe
PACKAGE = $(shell (Get-Content go.mod -head 1).Split(" ")[1])
CHECK_DIR_CMD = if (!(Test-Path $@)) { $$e = [char]27; Write-Error "$$e[31mDirectory $@ doesn't exist$${e}[0m" }
SERVER_BIN = ${SERVER_DIR}.exe
CLIENT_BIN = ${CLIENT_DIR}.exe
BIN_DIR = bin
PROTO_DIR = proto
SERVER_DIR = server
CLIENT_DIR = client

.PHONY: unary

project := unary

all: $(project)

unary: $@

$(project):
	@${CHECK_DIR_CMD}
	protoc -I$@/${PROTO_DIR} --go_opt=module=${PACKAGE} --go_out=. --go-grpc_opt=module=${PACKAGE} --go-grpc_out=. $@/${PROTO_DIR}/*.proto
	go build -o ${BIN_DIR}/$@/${SERVER_BIN} ./$@/${SERVER_DIR}
	go build -o ${BIN_DIR}/$@/${CLIENT_BIN} ./$@/${CLIENT_DIR}

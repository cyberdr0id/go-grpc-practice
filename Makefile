SHELL := powershell.exe
PACKAGE = $(shell (Get-Content go.mod -head 1).Split(" ")[1])
SERVER_BIN = ${SERVER_DIR}.exe
CLIENT_BIN = ${CLIENT_DIR}.exe
BIN_DIR = bin
PROTO_DIR = proto
SERVER_DIR = server
CLIENT_DIR = client

.PHONY: unary sstream cstream

project := unary sstream cstream

all: $(project)

unary: $@
sstream: $@
cstream: $@

$(project):
	protoc -I$@/${PROTO_DIR} --go_opt=module=${PACKAGE} --go_out=. --go-grpc_opt=module=${PACKAGE} --go-grpc_out=. $@/${PROTO_DIR}/*.proto
	go build -o ${BIN_DIR}/$@/${SERVER_BIN} ./$@/${SERVER_DIR}
	go build -o ${BIN_DIR}/$@/${CLIENT_BIN} ./$@/${CLIENT_DIR}

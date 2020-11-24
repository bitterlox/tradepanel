.PHONY: dev dev-server

WD=$(pwd)
BIN_NAME=trade-panel
PROTO_SRC_DIR=server/remote
PROTO_DST_DIR=server/remote/proto

W_PID=build/dev/server.pid
S_PID=build/dev/wails.pid

#MAKE runs every line into a diff shell, so this if thing is not possible
#TODO: add a task that recompiles .proto and have it run before dev always
start-dev: build-proto
	@stty -echoctl

	@if [ -f $(W_PID) ] || [ -f $(S_PID) ]; then { echo "wails or server already running $$W_PID $$W_PID"; exit 1; } fi

	cd server && { go build && ./server 2>&1 >/dev/null & cd .. && echo $$! > $(S_PID); }
	cd client && { wails serve 2>&1 >/dev/null & cd .. && echo $$! > $(W_PID); }

	@echo "\nStarted wails bridge; started server component;"

stop-dev:
	kill `cat build/dev/server.pid` && rm build/dev/server.pid
	kill `cat build/dev/wails.pid` && rm build/dev/wails.pid #this may not be needed, when the binary is killed wails exits automatically
	pkill $(BIN_NAME)
	@rm -rf client/build server/server

build-proto:
	protoc -I=$(PROTO_SRC_DIR) --go_out=$(PROTO_DST_DIR) --go_opt=paths=source_relative --go-grpc_out=$(PROTO_DST_DIR) --go-grpc_opt=paths=source_relative $(PROTO_SRC_DIR)/rpc.proto
.PHONY: dev dev-server build-proto

WD=$(pwd)
BIN_NAME=trade-panel

PROTO_SRC_DIR=server/remote
PROTO_DST_DIR=server/remote/proto

BUILD_DIR=build

W_PID=$(BUILD_DIR)/dev/server.pid
S_PID=$(BUILD_DIR)/dev/wails.pid

DEV_DIR=$(BUILD_DIR)/dev
TMP_DIR=$(BUILD_DIR)/tmp
BIN_DIR=$(BUILD_DIR)/bin

#MAKE runs every line into a diff shell, so this if thing is not possible
#TODO: add a task that recompiles .proto and have it run before dev always
start-dev: build-proto build-dirs
	@if [ -d $(DEV_DIR) ]; then { mkdir $(DEV_DIR); } fi

	@if [ -f $(W_PID) ] || [ -f $(S_PID) ]; then { echo "wails or server already running $$W_PID $$W_PID"; exit 1; } fi

	cd cmd/server && { go build && ./server 1 >/dev/null & cd .. && echo $$! > $(S_PID); }
	cd cmd/client && { wails serve 1 >/dev/null & cd .. && echo $$! > $(W_PID); }

	@echo "\nStarted wails bridge; started server component;"

stop-dev:
	kill `cat build/dev/server.pid` && rm build/dev/server.pid
	kill `cat build/dev/wails.pid` && rm build/dev/wails.pid #this may not be needed, when the binary is killed wails exits automatically
	pkill $(BIN_NAME)
	@rm -rf cmd/client/build cmd/server/server

build-proto:
	protoc -I=$(PROTO_SRC_DIR) --go_out=$(PROTO_DST_DIR) --go_opt=paths=source_relative --go-grpc_out=$(PROTO_DST_DIR) --go-grpc_opt=paths=source_relative $(PROTO_SRC_DIR)/rpc.proto

trade-panel:
	@if [ ! -d $(TMP_DIR) ] ; then { mkdir  $(TMP_DIR); } fi
	@if [ ! -d $(BIN_DIR) ]; then { mkdir  $(BIN_DIR); } fi

	cp cmd/client/* $(TMP_DIR)
	cp go.mod go.sum $(TMP_DIR)

	cd $(TMP_DIR); wails build;

	rm -rf $(TMP_DIR)

build-dirs:
	@if [ ! -d $(BUILD_DIR) ] ; then { mkdir  $(BUILD_DIR); } fi
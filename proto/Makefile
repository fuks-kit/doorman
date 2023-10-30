## Shipping package
## Assumes user is on MacOS, if other OS, please change PROTO_ROOT_DIR to the path of protobuf installation
PROTO_ROOT_DIR = $(shell brew --prefix)/Cellar/protobuf/24.4/include
APP_DIR = /Users/patrick/Code/fuks_app/
APP_OUTPUT = $(APP_DIR)/lib/generated/doorman

## Dart requires you to manually ship all google provided proto files too.
dart:
	@mkdir -p $(APP_OUTPUT)
	@protoc -I=. --dart_out=grpc:$(APP_OUTPUT) *.proto
	@protoc -I$(PROTO_ROOT_DIR) --dart_out=$(APP_OUTPUT) $(PROTO_ROOT_DIR)/google/protobuf/*.proto

go:
	@protoc --go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		*.proto

gen: dart go
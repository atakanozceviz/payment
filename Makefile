API_PROTO_FILES=$(shell find api -name *.proto)

.PHONY: api
# generate api proto
api:
	protoc --proto_path ./api \
	       --proto_path ./third_party \
	       --go_out=paths=source_relative:./api \
	       --go-grpc_out=paths=source_relative:./api \
	       --validate_out=paths=source_relative,lang=go:./api \
	       $(API_PROTO_FILES)

APP_BIN =  gen/go/application/
APP_PROTO = proto/application/

.PHONY: gen
gen:
	rm -f $(APP_BIN)*
	#export PATH=$PATH:/path/to/protoc
	protoc -I proto $(APP_PROTO)*.proto --go_out=./gen/go --go_opt=paths=source_relative --go-grpc_out=./gen/go/ --go-grpc_opt=paths=source_relative --experimental_allow_proto3_optional


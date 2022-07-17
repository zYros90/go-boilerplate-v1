# for generating protos install 
```
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest 
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest 
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest 
go install github.com/google/gnostic/cmd/protoc-gen-openapi@latest 
go install github.com/envoyproxy/protoc-gen-validate@latest
go install github.com/chrusty/protoc-gen-jsonschema/cmd/protoc-gen-jsonschema@latest
```
* for dart proto:
https://github.com/google/protobuf.dart/tree/master/protoc_plugin#how-to-build-and-use=
```
dart pub global activate protoc_plugin
```

 # set protoc options in vscode
 preferences > settings > vscode-proto3 configuration
```
"protoc": {
    "options": [
        "--proto_path=./api/v1",
        "--proto_path=./api/third_party"
    ]
}
```

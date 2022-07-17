proto-jsonschema:
	$(eval THIRD_PARTY=./api/third_party)
	$(eval OUT_GO=./api/third_party/jsonschema)
	$(eval PROTO_PATH=./api/third_party/jsonschema)

	@mkdir -p $(OUT_GO)

	@protoc \
		--proto_path=$(PROTO_PATH) \
		--proto_path=$(THIRD_PARTY) \
		--proto_path=/usr/local/include \
		--go_out=paths=source_relative:$(OUT_GO) \
		$(PROTO_PATH)/*.proto

proto:
	$(eval PROTO_PATH=./api/v1)
	$(eval THIRD_PARTY=./api/third_party)
	$(eval OUT_GO=./api/v1/generated)
	$(eval OUT_SCHEMAS=./api/v1/generated_schemas)
	$(eval OUT_DART=./api/v1/generated_dart)

	@mkdir -p $(OUT_GO)
	@mkdir -p $(OUT_DART)
	@mkdir -p $(OUT_SCHEMAS)

	@protoc \
		--proto_path=$(PROTO_PATH) \
		--proto_path=$(THIRD_PARTY) \
		--proto_path=/usr/local/include \
		--go_out=paths=source_relative:$(OUT_GO) \
		--go-grpc_out=paths=source_relative:$(OUT_GO) \
		--validate_out=paths=source_relative,lang=go:$(OUT_GO) \
		--jsonschema_out=source_relative:$(OUT_SCHEMAS) \
		--dart_out=$(OUT_DART) \
		$(PROTO_PATH)/*.proto

docker:
	@echo 'using registry ${REGISTRY} with tag ${TAG}'

	docker build -t ${REGISTRY}:${TAG} .
	docker push ${REGISTRY}:${TAG} 

test:
	go mod download
	go test ./...
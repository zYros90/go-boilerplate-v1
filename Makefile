
proto:
	$(eval OUT_PATH=./api/v1/generated)
	$(eval THIRD_PARTY=./api/third_party)
	$(eval PROTO_PATH=./api/v1)
	$(eval SCHEMAS_PATH=./api/v1/generated/schemas)

	@mkdir -p $(OUT_PATH)
	@mkdir -p $(SCHEMAS_PATH)

	@protoc \
		--proto_path=$(PROTO_PATH) \
		--proto_path=$(THIRD_PARTY) \
		--go_out=paths=source_relative:$(OUT_PATH) \
		--go-grpc_out=paths=source_relative:$(OUT_PATH) \
		--go-errors_out=paths=source_relative:$(OUT_PATH) \
		--validate_out=paths=source_relative,lang=go:$(OUT_PATH) \
		--jsonschema_out=source_relative:$(SCHEMAS_PATH) \
		$(PROTO_PATH)/*.proto

docker:
	@echo 'using registry ${REGISTRY} with tag ${TAG}'

	docker build -t ${REGISTRY}:${TAG} .
	docker push ${REGISTRY}:${TAG} 

test:
	go mod download
	go test ./...
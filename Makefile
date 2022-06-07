
proto:
	$(eval OUT_PATH=./api/v1/generated)
	$(eval THIRD_PARTY=./api/third_party)
	$(eval PROTO_PATH=./api/v1)

	@mkdir -p $(OUT_PATH)

	@protoc \
		--proto_path=$(PROTO_PATH) \
		--proto_path=$(THIRD_PARTY) \
		--proto_path=$(PROTOC_COMMON_PATH) \
		--go_out=paths=source_relative:$(OUT_PATH) \
		--go-errors_out=paths=source_relative:$(OUT_PATH) \
		--validate_out=paths=source_relative,lang=go:$(OUT_PATH) \
		$(PROTO_PATH)/*.proto

docker:
	@echo 'using registry ${REGISTRY} with tag ${TAG}'

	docker build -t ${REGISTRY}:${TAG} .
	docker push ${REGISTRY}:${TAG} 

test:
	go mod download
	go test ./...
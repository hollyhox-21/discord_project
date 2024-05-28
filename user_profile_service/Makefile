LOCAL_BIN:=$(CURDIR)/bin

# Путь до завендореных protobuf файлов
VENDOR_PROTO_PATH := $(CURDIR)/vendor.protobuf

# Путь до сгенеренных .pb.go файлов
PKG_PROTO_PATH := $(CURDIR)/pkg/user_profile

.bin-deps:
	$(info Installing binary dependencies...)
	GOBIN=$(LOCAL_BIN) go install github.com/bufbuild/buf/cmd/buf@v1.31.0
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.34.1
	GOBIN=$(LOCAL_BIN) go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.3.0
	GOBIN=$(LOCAL_BIN) go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@v2.19.1
	GOBIN=$(LOCAL_BIN) go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@v2.19.1
	GOBIN=$(LOCAL_BIN) go install github.com/planetscale/vtprotobuf/cmd/protoc-gen-go-vtproto@v0.6.0
	GOBIN=$(LOCAL_BIN) go install github.com/swaggo/swag/cmd/swag@v1.16.3


# Генерация кода из protobuf
generate: .bin-deps .buf-generate .tidy

# go mod tidy
.tidy:
	GOBIN=$(LOCAL_BIN) go mod tidy

.buf-generate:
	rm -rf $(PKG_PROTO_PATH)
	$(LOCAL_BIN)/buf generate api

# vendor
vendor:	.vendor-reset .vendor-googleapis .vendor-google-protobuf .vendor-protovalidate .vendor-tidy .vendor-protoc-gen-openapiv2

.vendor-reset:
	rm -rf $(VENDOR_PROTO_PATH)
	mkdir -p $(VENDOR_PROTO_PATH)

.vendor-tidy:
	find $(VENDOR_PROTO_PATH) -type f ! -name "*.proto" -delete

# Устанавливаем proto описания google/protobuf
.vendor-google-protobuf:
	git clone -b main --single-branch -n --depth=1 --filter=tree:0 \
		https://github.com/protocolbuffers/protobuf $(VENDOR_PROTO_PATH)/protobuf &&\
	cd $(VENDOR_PROTO_PATH)/protobuf &&\
	git sparse-checkout set --no-cone src/google/protobuf &&\
	git checkout
	mkdir -p $(VENDOR_PROTO_PATH)/google
	mv $(VENDOR_PROTO_PATH)/protobuf/src/google/protobuf $(VENDOR_PROTO_PATH)/google
	rm -rf $(VENDOR_PROTO_PATH)/protobuf

# Устанавливаем proto описания validate
.vendor-protovalidate:
	git clone -b main --single-branch --depth=1 --filter=tree:0 \
		https://github.com/bufbuild/protovalidate $(VENDOR_PROTO_PATH)/protovalidate && \
	cd $(VENDOR_PROTO_PATH)/protovalidate
	git checkout
	mv $(VENDOR_PROTO_PATH)/protovalidate/proto/protovalidate/buf $(VENDOR_PROTO_PATH)
	rm -rf $(VENDOR_PROTO_PATH)/protovalidate

# Устанавливаем proto описания google/api
.vendor-googleapis:
	git clone -b master --single-branch -n --depth=1 --filter=tree:0 \
		https://github.com/googleapis/googleapis $(VENDOR_PROTO_PATH)/googleapis &&\
	cd $(VENDOR_PROTO_PATH)/googleapis &&\
	git checkout
	mv $(VENDOR_PROTO_PATH)/googleapis/google $(VENDOR_PROTO_PATH)
	rm -rf $(VENDOR_PROTO_PATH)/googleapis

# Устанавливаем proto описания protoc-gen-openapiv2/options
.vendor-protoc-gen-openapiv2:
	git clone -b main --single-branch -n --depth=1 --filter=tree:0 \
 		https://github.com/grpc-ecosystem/grpc-gateway $(VENDOR_PROTO_PATH)/grpc-gateway && \
 	cd $(VENDOR_PROTO_PATH)/grpc-gateway && \
	git sparse-checkout set --no-cone protoc-gen-openapiv2/options && \
	git checkout
	mkdir -p $(VENDOR_PROTO_PATH)/protoc-gen-openapiv2
	mv $(VENDOR_PROTO_PATH)/grpc-gateway/protoc-gen-openapiv2/options $(VENDOR_PROTO_PATH)/protoc-gen-openapiv2
	rm -rf $(VENDOR_PROTO_PATH)/grpc-gateway


# Билд приложения
build:
	go build -o $(LOCAL_BIN) ./cmd/notes/client
	go build -o $(LOCAL_BIN) ./cmd/notes/server

docs:
# либа к swag https://github.com/swaggo/swag
	$(LOCAL_BIN)/swag fmt
	$(LOCAL_BIN)/swag init --outputTypes "go,json" --parseInternal -g cmd/main.go

update-swagger:
	go get -u github.com/rakyll/statik

	rm -rf /tmp/swagger-ui
	git clone https://github.com/swagger-api/swagger-ui.git ./tmp/swagger-ui

	cd /tmp/swagger-ui; \
		mkdir ./html; \
		cat ./dist/index.html | perl -pe 's/https?:\/\/petstore.swagger.io\/v2\///g' > ./html/index.html; \
		cp ./dist/oauth2-redirect.html ./html; \
		cp ./dist/*.js ./html; \
		cp ./dist/*.css ./html; \
		cp ./dist/*.png ./html

	statik -dest ./ -f -p swagger -src ./tmp/swagger-ui/html -ns scratch

	#rm -fr /tmp/swagger-ui

# Объявляем, что текущие команды не являются файлами и
# интсрументируем Makefile не искать изменения в файловой системе
.PHONY: \
	.bin-deps \
	.buf-generate \
	.tidy \
	.vendor-google-protobuf \
	.vendor-protovalidate \
	.vendor-googleapis \
	vendor \
	generate \
	build
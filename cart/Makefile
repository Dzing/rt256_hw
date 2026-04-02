# Переменные
BIN_NAME=route-cart
PACKAGE=route/cart/cmd/app

BUILD_DIR=build
BUILD_DIR_F=$(CURDIR)/$(BUILD_DIR)

BIN_DIR=$(BUILD_DIR)/bin

PROTOC_PLUGINS_DIR_F=$(BUILD_DIR_F)/protoc-plugins

LINTER_DIR_F=$(BUILD_DIR_F)/lint
LINTER_NAME=golangci-lint
LINTER_BIN_F=$(LINTER_DIR_F)/$(LINTER_NAME)
LINTER_VERSION=v2.8.0

MAKE=make --no-print-directory

## prebuild: Подготавливает необходимое
.PHONY: prebuild
prebuild: 
	@echo "\n[i] Подготовка к сборке $(BIN_NAME) ..."
	@$(MAKE) .prep-protoc-plugins
	@$(MAKE) .protoc-gen
	@echo "[v] Подготовка к сборке $(BIN_NAME) завершена"

## build: Собирает проект
.PHONY: build
build: 
	@$(MAKE) prebuild
	@echo "\n[i] Сборка $(BIN_NAME) ..."
	@mkdir -p ${BUILD_DIR}
	go build -v -o ${BIN_DIR}/$(BIN_NAME) ${PACKAGE}
	@echo "[v] $(BIN_NAME) собран успешно!" 

## clean: Удаляет артефакты сборки
.PHONY: clean
clean:
	@echo "[i] Очистка..."
	@go clean
	rm -rf $(BIN_DIR)
	@echo "[v] Очистка завершена"

## precommit: Запуск проверок перед коммитом
.PHONY: precommit
precommit: 
	@$(MAKE) .lint
	@echo "[v] Подготовка к коммиту завершена!"

## lint: Запускает линтер (автоматически устанавливает при необходимости)
.lint:
	@if [ ! -f "$(LINTER_BIN_F)" ] \
	; then \
		echo "[!] Линтер не найден. Установка..."; \
		$(MAKE) .lint-install; \
	fi
	@echo "[i] Запуск линтера..."
	$(LINTER_BIN_F) run ./...

## lint-install: Устанавливает golangci-lint
.lint-install: 
	@echo "[i] Установка golangci-lint ($(LINTER_VERSION)) в $(LINTER_DIR_F)..."
	@rm -rf $(LINTER_DIR_F)
	@mkdir -p $(LINTER_DIR_F)
	@tmpdir=$$(mktemp -d); \
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b "$$tmpdir" $(LINTER_VERSION); \
	mv "$$tmpdir/$(LINTER_NAME)" $(LINTER_PATH); \
	rm -rf "$$tmpdir"
	@echo "[v] Установка завершена: $(LINTER_PATH)"



## prep-protoc-plugins: Устанавливает плагины для protoc
.prep-protoc-plugins:
	@if [ ! -f "$(LINTER_PATH)" ]; then \
		mkdir -p $(PROTOC_PLUGINS_DIR_F); \
	fi

	GOBIN=$(PROTOC_PLUGINS_DIR_F) go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	GOBIN=$(PROTOC_PLUGINS_DIR_F) go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	GOBIN=$(PROTOC_PLUGINS_DIR_F) go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
	GOBIN=$(PROTOC_PLUGINS_DIR_F) go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest
	@echo "[v] Плагины для protoc готовы"

## protoc-gen: Генерирует .go модули из .proto файлов 
.protoc-gen:
	@echo "[i] Генерация .go модулей из схем данных protobuf ..."
	@if [ ! -d "pkg" ] \
	; then \
		mkdir -p pkg \
	; fi
	protoc \
	-I . \
	-I ./../3rd-party/proto/googleapis \
	--plugin=protoc-gen-go=$(PROTOC_PLUGINS_DIR_F)/protoc-gen-go \
	--go_out ./pkg \
	--go_opt paths=source_relative \
	--plugin=protoc-gen-go-grpc=$(PROTOC_PLUGINS_DIR_F)/protoc-gen-go-grpc \
	--go-grpc_out ./pkg \
	--go-grpc_opt paths=source_relative \
	--plugin=protoc-gen-grpc-gateway=$(PROTOC_PLUGINS_DIR_F)/protoc-gen-grpc-gateway \
	--grpc-gateway_out ./pkg \
	--grpc-gateway_opt generate_unbound_methods=true \
	--grpc-gateway_opt paths=source_relative \
	api/v1/cart_service.proto
	@echo "[v] Генерация завершена"

clr:
	rm -rf pkg
	
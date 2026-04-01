# Переменные
BINARY_NAME=route-cart
BUILD_DIR=build
BIN_DIR=$(BUILD_DIR)/bin
LINTER_DIR=$(BUILD_DIR)/lint
LINTER_NAME=golangci-lint
LINTER_PATH=$(LINTER_DIR)/$(LINTER_NAME)
LINTER_VERSION=v2.8.0

PACKAGE=route/cart/cmd/app

## build: Собирает проект
.PHONY: build
build: 
	@echo "Сборка проекта ""$(BINARY_NAME)"" ..."
	@mkdir -p ${BUILD_DIR}
	go build -v -o ${BIN_DIR}/$(BINARY_NAME) ${PACKAGE}
	@echo "Сборка завершена! Артефакт: ${BIN_DIR}/$(BINARY_NAME)"

## clean: Удаляет артефакты сборки
.PHONY: clean
clean:
	@echo "Очистка..."
	@go clean
	@rm -rf $(BIN_DIR)
	@echo "Очистка завершена"

## precommit: Запуск проверок перед коммитом
.PHONY: precommit
precommit: lint
	@echo "Подготовка к коммиту завершена!"

## lint: Запускает линтер (автоматически устанавливает при необходимости)
.PHONY: lint
lint:
	@echo "Запуск линтера..."
	@if [ ! -f "$(LINTER_PATH)" ]; then \
		echo "Линтер не найден. Установка..."; \
		$(MAKE) lint-install; \
		echo "Линтер установлен. Повторный запуск..."; \
	fi
	$(LINTER_PATH) run ./...

## lint-install: Устанавливает golangci-lint в локальную папку bin/lint
.PHONY: lint-install
lint-install: 
	@echo "Установка golangci-lint ($(LINTER_VERSION)) в $(LINTER_DIR)..."
	@rm -rf $(LINTER_DIR)
	@mkdir -p $(LINTER_DIR)
	@tmpdir=$$(mktemp -d); \
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b "$$tmpdir" $(LINTER_VERSION); \
	mv "$$tmpdir/$(LINTER_NAME)" $(LINTER_PATH); \
	rm -rf "$$tmpdir"
	@echo "Установка завершена: $(LINTER_PATH)"

.PHONY: prep
prep: protoc-install

## lint-install: Устанавливает плагины для protoc
.PHONY: protoc-install
protoc-install:
	@echo "Установка protoc-gen-go@latest"
	GOBIN=$(BUILD_DIR_FULL) go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	GOBIN=$(BUILD_DIR_FULL) go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest


	
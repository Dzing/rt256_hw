# Переменные
BINARY_NAME=route-loms
BUILD_DIR=build
BIN_DIR=$(BUILD_DIR)/bin
LINTER_DIR=$(BUILD_DIR)/lint
LINTER_NAME=golangci-lint
LINTER_PATH=$(LINTER_DIR)/$(LINTER_NAME)
LINTER_VERSION=v2.8.0

PACKAGE=route/loms/cmd/app

.PHONY: build clean precommit lint lint-install

## build: Собирает проект
build: 
	@echo "Сборка проекта ""$(BINARY_NAME)"" ..."
	@mkdir -p ${BUILD_DIR}
	go build -v -o ${BIN_DIR}/$(BINARY_NAME) ${PACKAGE}
	@echo "Сборка завершена! Артефакт: ${BIN_DIR}/$(BINARY_NAME)"

## clean: Удаляет артефакты сборки
clean:
	@echo "Очистка..."
	@go clean
	@rm -rf $(BIN_DIR)
	@echo "Очистка завершена"

precommit: lint
	@echo "Подготовка к коммиту завершена!"

## lint: Запускает линтер (автоматически устанавливает при необходимости)
lint:
	@echo "Запуск линтера..."
	@if [ ! -f "$(LINTER_PATH)" ]; then \
		echo "Линтер не найден. Установка..."; \
		$(MAKE) lint-install; \
		echo "Линтер установлен. Повторный запуск..."; \
	fi
	$(LINTER_PATH) run ./...

## lint-install: Устанавливает golangci-lint в локальную папку bin/lint
lint-install: 
	@echo "Установка golangci-lint ($(LINTER_VERSION)) в $(LINTER_DIR)..."
	@mkdir -p $(LINTER_DIR)
	@tmpdir=$$(mktemp -d); \
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b "$$tmpdir" $(LINTER_VERSION); \
	mv "$$tmpdir/$(LINTER_NAME)" $(LINTER_PATH); \
	rm -rf "$$tmpdir"
	@echo "Установка завершена: $(LINTER_PATH)"

	
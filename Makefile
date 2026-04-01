3RD_PARTY_DIR=3rd-party
3RD_PARTY_PATH_F=$(CURDIR)/$(3RD_PARTY_DIR)

MAKE=make --no-print-directory

## prebuild: запускает подготовку инструментов к сборке
.PHONY: prebuild
prebuild:
	@echo "\n[i] Подготовка к сборке пакета ..."
	@$(MAKE) .3rd-party-prep
	@echo "[v] Подготовка к сборке пакета завершена"

## build: запускает сборку всех артефактов
.PHONY: build
build:
	@$(MAKE) prebuild
	@echo "\n[i] Сборка пакета артефактов ..."
	@cd loms && GOOS=linux GOARCH=amd64 $(MAKE) build
	@cd cart && GOOS=linux GOARCH=amd64 $(MAKE) build
	@echo "[v] Сборка пакета завершена"

## run: запускает сборку образов докера и последующий деплой
.PHONY: run
run: build
	docker compose up --force-recreate --build

## precommit: запускает проверки перед коммитом
.PHONY: precommit
precommit:
	cd loms && $(MAKE) precommit
	cd cart && $(MAKE) precommit

##################################################

.PHONY: clear
clear:
	rm -rf tmp
	rm -rf 3rd-party

.3rd-party-prep:
	@if [ ! -d "$(3RD_PARTY_PATH_F)" ]; then \
		mkdir -p $(3RD_PARTY_PATH_F); \
		echo "Создан каталог $(3RD_PARTY_PATH_F)"; \
	fi
	@$(MAKE) .git-proto/google/api

.git-proto/google/api:
	@$(eval TARGET_PATH=$(3RD_PARTY_PATH_F)/proto/googleapis/google/api)
	@if [ ! -d "$(TARGET_PATH)" ] \
	; then \
		echo "[!] proto/google/api не найден. Запрос из репозитория ..."; \
		git clone -b master --single-branch -n --depth=1 --filter=tree:0 \
			https://github.com/googleapis/googleapis tmp \
		&& cd tmp \
		&& git sparse-checkout set --no-cone google/api \
		&& git checkout \
		&& mkdir -p $(TARGET_PATH) \
		&& mv google/api/annotations.proto $(TARGET_PATH) \
		&& mv google/api/http.proto $(TARGET_PATH) \
		&& cd .. && rm -rf tmp \
		&& echo "[v] proto/google/api получен из репозитория" \
	; fi
	@echo "[v] proto/google/api готов"

	
 
	
ENV_PATH = .env.dev

include $(ENV_PATH)

COMPOSE_PATH = docker-compose.yml

DOCKER_COMPOSE := docker compose -f $(COMPOSE_PATH) --env-file $(ENV_PATH)

DB_URL = postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable
MIGRATIONS_PATH = migrations

.PHONY: test coverage-html clean docker-build docker-up docker-down docker-stop docker-logs generate

# === Вспомогательные команды ===

generate:
	@echo "==> Generating..."
	@go generate ./...


# === Тестирование ===

test:
	@echo "==> Запускаем тесты и генерируем отчет о покрытии..."
	@go test -coverprofile=coverage.out $(shell go list ./... | grep -v /mocks)
	@echo "\n==> Общее покрытие кода тестами:"
	@go tool cover -func=coverage.out | grep total

coverage-html: test
	@echo "==> Открываем HTML-отчет в браузере..."
	@go tool cover -html=coverage.out

clean:
	@echo "==> Очищаем сгенерированные файлы..."
	@rm -f coverage.out

# === Docker Compose ===

docker-build:
	@echo "==> Пересобираем и перезапускаем сервисы..."
	@$(DOCKER_COMPOSE) up -d --build

docker-up:
	@echo "==> Запускаем Docker Compose в фоновом режиме..."
	@$(DOCKER_COMPOSE) up -d

docker-down:
	@echo "==> Останавливаем и удаляем контейнеры..."
	@$(DOCKER_COMPOSE) down

docker-stop:
	@echo "==> Останавливаем сервисы Docker Compose..."
	@$(DOCKER_COMPOSE) stop

docker-logs:
	@echo "==> Просматриваем логи контейнеров..."
	@$(DOCKER_COMPOSE) logs -f


# === Migrations === #

migrate-down:
	@echo "==> Откатываем миграции..."
	@migrate -path $(MIGRATIONS_PATH) -database "$(DB_URL)" down


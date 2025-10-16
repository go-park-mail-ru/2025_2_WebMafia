ENV_PATH = .env
COMPOSE_PATH = docker-compose.yml

DOCKER_COMPOSE := docker compose -f $(COMPOSE_PATH) --env-file $(ENV_PATH)

DB_URL = postgres://$(DB_USER):$(DB_PASSWORD)@localhost:5432/$(DB_NAME)?sslmode=disable
MIGRATIONS_PATH = migrations

.PHONY: test coverage-html clean docker-build docker-up docker-down docker-stop docker-logs

# === Тестирование ===

test:
	@echo "==> Запускаем тесты и генерируем отчет о покрытии..."
	@go test -coverprofile=coverage.out ./...
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
	@echo "==> Собираем образы Docker..."
	@$(DOCKER_COMPOSE) build

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

migrate-up:
	@echo "==> Применяем миграции..."
	@migrate -path $(MIGRATIONS_PATH) -database "$(DB_URL)" up

migrate-down:
	@echo "==> Откатываем миграции..."
	@migrate -path $(MIGRATIONS_PATH) -database "$(DB_URL)" down
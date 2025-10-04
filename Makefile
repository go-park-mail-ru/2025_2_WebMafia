.PHONY: test coverage-html clean

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

up:
	@echo "==> Запускаем Docker Compose в фоновом режиме..."
	@sudo docker compose up --build -d

down:
	@echo "==> Останавливаем и удаляем контейнеры..."
	@sudo docker compose down

logs:
	@echo "==> Просматриваем логи контейнеров..."
	@sudo docker compose logs -f

stop:
	@echo "==> Останавливаем сервисы Docker Compose..."
	@sudo docker compose stop
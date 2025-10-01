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
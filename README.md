# База данных WaveMusic

В папке `internal/migrations/` лежат SQL-миграции для развёртывания схемы БД.  
Документация и диаграмма структуры — в `internal/normalized/`.


![Статус](https://img.shields.io/badge/статус-в_разработке-orange)

## Команда проекта

### Разработчики

- [Даниил Быстров](https://github.com/ska66696)
- [Артем Голубев](https://github.com/Xyzdat)
- [Дмитрий Дорофеев](https://github.com/poopaapopa)
- [Вероника Лемешкина](https://github.com/LemeshkinaVeronika)

### Менторы

- **Бэкенд:** [Илья Денисенко](https://github.com/MatiXxD)
- **Фронтенд:** [Алексей Зотов](https://github.com/let-robots-reign)
- **СУБД:** Алексей Фильчаков
- **UX/UI:** Анна Реутова

## Связанные проекты

- **Frontend-часть:** [Wave Music Frontend](https://github.com/frontend-park-mail-ru/2025_2_WebMafia/)

## Содержание

- [Технологии](#технологии)
- [Начало работы](#начало-работы)
- [Тестирование](#тестирование)
- [Deploy и CI/CD](#deploy-и-ci/cd)

## Технологии

- [Go v1.24.0.](https://tip.golang.org/doc/go1.24)
- [PostgreSQL 17](https://www.postgresql.org/docs/release/17.0/) **(в разработке)**
- [HTTP-роутинг Gorilla](github.com/gorilla/mux)

## Использование

API сервиса предоставляет RESTful endpoints для работы с данными.

**Пример запроса**

```sh
curl -X GET http://localhost:8080/
```

## Разработка

### Требования

Для установки и запуска проекта, необходим Go 1.24+.

### Установка зависимостей

Для установки зависимостей, выполните команду:

```sh
go mod tidy
```

### Запуск Development сервера

Чтобы запустить сервер для разработки, выполните команду:

```sh
go run server/main.go
```

### Создание билда

Чтобы выполнить production сборку, выполните команду:

```sh
go build -o wave-music server/main.go
```

## Тестирование

Инструменты тестирования находятся в процессе настройки. Планируемое покрытие 60% к концу разработки.

## Deploy и CI/CD

Настройки пайплайнов находятся в процессе разработки.

### Зачем мы разработали этот проект?

Это учебный проект для курса по веб-разработки от VK Education.

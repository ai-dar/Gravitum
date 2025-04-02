# MyApp

Это пример REST API для создания, редактирования и удаления пользователей, разработанный на языке Go с использованием Gin, GORM и PostgreSQL. Проект также включает модульное тестирование с использованием in-memory SQLite (через чисто Go‑драйвер) и контейнеризацию с Docker и docker-compose.

## Технологии

- **Go** – основной язык разработки.
- **Gin** – высокопроизводительный HTTP‑фреймворк.
- **GORM** – ORM для работы с PostgreSQL (и SQLite для тестирования).
- **PostgreSQL** – база данных для хранения пользователей.
- **Docker & docker-compose** – для контейнеризации приложения и базы данных.
- **Testify** – для написания тестов.

## Структура проекта
```bash
myapp/
├── Dockerfile
├── docker-compose.yml
├── go.mod
├── main.go
├── models/
│   └── user.go
├── repository/
│   └── user_repository.go
├── service/
│   └── user_service.go
├── handlers/
│   └── user_handler.go
├── routes/
│   └── routes.go
└── tests/
    └── user_handler_test.go
```

- **main.go** – точка входа в приложение, где происходит подключение к базе данных, миграция моделей и запуск сервера.
- **models/** – содержит определения структур данных (например, модель пользователя).
- **repository/** – слой доступа к данным (работа с PostgreSQL через GORM).
- **service/** – бизнес-логика работы с пользователями.
- **handlers/** – HTTP‑обработчики для реализации API.
- **routes/** – регистрация маршрутов API.
- **tests/** – тесты для проверки работы API с использованием in-memory базы данных SQLite.

## Как начать работу

### Локальный запуск

1. **Клонируйте репозиторий:**

   ```bash
   git clone https://github.com/ai-dar/Gravitum.git
   cd myapp
   ```
2. **Инициализируйте модуль и установите зависимости:**
   ```bash
   go mod init myapp
   go mod tidy
   ```
   
3. **Настройка подключения к базе данных**  
Убедитесь, что PostgreSQL запущен и создана база данных (например, `users_db`).  
Задайте переменную окружения для подключения. В CMD выполните:

  ```cmd
   set DATABASE_URL=host=localhost user=postgres password=postgres dbname=users_db port=5432 sslmode=disable
   ```

4. **Запуск приложения**  
В командной строке выполните:
  ```bash
  go run main.go
   ```

Приложение будет доступно по адресу: [http://localhost:8080](http://localhost:8080)

### 2. Работа с REST API

Приложение реализует следующие эндпоинты:

- **Создание пользователя**  
Метод: `POST`  
URL: `http://localhost:8080/users/`  
Пример JSON-данных:
  ```bash
  {"name": "John Doe", "email": "john@example.com"}
   ```

Пример запроса в CMD:
 ```bash
  curl -X POST http://localhost:8080/users/ -H "Content-Type: application/json" -d "{"name": "John Doe", "email": "john@example.com"}"
   ```

- **Получение пользователя по ID**  
Метод: `GET`  
URL: `http://localhost:8080/users/1`

- **Обновление данных пользователя**  
Метод: `PUT`  
URL: `http://localhost:8080/users/1`  
Пример JSON-данных:
```bash
  {"name": "Jane Doe", "email": "jane@example.com"}
   ```

Пример запроса:
 ```bash
curl -X PUT http://localhost:8080/users/1 -H "Content-Type: application/json" -d "{"name": "Jane Doe", "email": "jane@example.com"}"
```

- **Удаление пользователя**  
Метод: `DELETE`  
URL: `http://localhost:8080/users/1`  
Пример запроса:
```bash
curl -X DELETE http://localhost:8080/users/1
```

### 3. Запуск тестов

Тесты расположены в каталоге `tests`. Для запуска тестов откройте CMD в корневой директории и выполните:
```bash
go test ./...
```

Если тесты пройдут успешно, вы увидите сообщение вида:

```bash
ok      test/tests      0.296s
```

### 4. Контейнеризация с Docker

Проект содержит файлы `Dockerfile` и `docker-compose.yml` для контейнеризации.

1. **Сборка и запуск контейнеров**  
   В CMD в корневой директории выполните:
```bash
docker-compose up --build
```


2. **Доступ к сервисам**  
- Приложение будет доступно по адресу: [http://localhost:8080](http://localhost:8080)
- PostgreSQL будет доступен на указанном порту.

## Дополнительные сведения

- При запуске приложение автоматически выполняет миграцию модели пользователя с помощью `AutoMigrate`.
- Тесты используют in-memory базу данных с чисто Go‑драйвером (без cgo).
- Для управления изменениями используйте привычные команды Git:
- `git add .`
- `git commit -m "Описание изменений"`
- `git push origin main`





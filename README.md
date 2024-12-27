HTTP-калькулятор на Go

Калькулятор на Go

Простое приложение калькулятора, написанное на языке Go, которое поддерживает как командный интерфейс (CLI), так и HTTP API для выполнения математических вычислений.

Технологии
 • Go: основной язык разработки.
 • JSON: формат передачи данных.

Структура проекта

/cmd
  /calc_service
    main.go
/internal
  /calculator
    handler.go
    service.go
  /utils
    validator.go
/pkg
  /models
    expression.go
    result.go

/cmd/calc_service/main.go: точка входа в приложение, где инициализируется сервер и маршруты.

/internal/calculator/handler.go: обработчик HTTP-запросов, принимающий математические выражения и передающий их в сервис для вычисления.

/internal/calculator/service.go: сервис, выполняющий вычисления и возвращающий результаты.

/internal/utils/validator.go: утилита для валидации входных данных и выражений.

/pkg/models/expression.go: структура данных для представления математических выражений.

/pkg/models/result.go: структура данных для представления результатов вычислений.

Установка
 1. Клонирование репозитория:

go get https://github.com/gulovv/calc-go


 2. Запуск сервера:

go run cmd/calc_service/main.go

По умолчанию сервер запускается на порту 8080. Для изменения порта установите переменную окружения PORT.




Примеры использования
 1. Успешный расчет (HTTP 200):
Выполните следующий запрос для получения результата вычисления:

curl -X POST http://localhost:8080/api/v1/calculate \
     -H "Content-Type: application/json" \
     -d '{"expression": "3+5*2"}'


Ответ:

{
  "result": 13
}


 2. Некорректное выражение (HTTP 422):
Для запроса с ошибкой синтаксиса:

curl -X POST http://localhost:8080/api/v1/calculate \
     -H "Content-Type: application/json" \
     -d '{"expression": ""}'

Ответ:

{
"error":"Expression cannot be empty. Please provide a valid mathematical expression.
"}


 3. Внутренняя ошибка сервера (HTTP 500):
Для запроса, который вызывает внутреннюю ошибку сервера:

curl -X POST http://localhost:8080/api/v1/calculate \
     -H "Content-Type: application/json" \
     -d '{"expression": "Expression is not valid. Please check the syntax."}'

Ответ:

{
  "error": "Internal server error"
}

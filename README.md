# calc_service
Калькулятор арифметических выражений

# Установка и запуск
*все комманды ввводятся в git bash
(для удобства, комманды можно копировать и вставлять в git bash с помощью shift+insert)

```git clone https://github.com/horhhe/calc_service.git```          - клонирование репозитория

```cd calc_service```          - переход к репозиторию

```go run ./cmd/calc_service/...```          - запуск сервиса

# Примеры использования 
*для ввода запросов необходимо открыть второе окно git bash

Успешный запрос:
```curl --location 'http://localhost:8080/api/v1/calculate' --header 'Content-Type: application/json' --data '{"expression": "2+2*2"}'```

Ответ:
```{"result": 6}```

Ошибка 422:
```curl --location 'http://localhost:8080/api/v1/calculate' --header 'Content-Type: application/json' --data '{"expression": "2+2*a"}'```

Ответ:
```{"error": "недопустимый символ: a"}```

Ошибка 500:
```curl --location 'http://localhost:8080/api/v1/calculate' --header 'Content-Type: application/json' --data '{"expression": "2/0"}'```

Ответ:
```{"error": "деление на ноль"}```

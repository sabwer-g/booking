# 🌴 Booking service
Сервис отвечает за бронирования номеров в отелях, в котором реализована возможность забронировать свободный номер в отеле.
(тестовый прототип)

## 💻 Локальное окружение

Команды make
```
make up       # Поднятие локального окружения
make down     # Остановка локального окружения

make test     # Запуск go-тестов
```

## 🚀 Тестовый запуск

Для создания заказа необходимо выполнить запрос такого плана:
```
curl --location --request POST 'http://127.0.0.1:8080/api/booking/v1/orders/create' \
--header 'Content-Type: application/json' \
--data-raw '{
    "meta": {},
    "data": {
        "hotel_id": "reddison",
        "room_id": "lux",
        "email": "guest@mail.ru",
        "from": "2025-03-02T00:00:00Z",
        "to": "2025-03-04T00:00:00Z"
    }
}'
```
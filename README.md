# orders-service
Подключение и подписка на канал в nats-streaming
Полученные данные писать в Postgres
Так же полученные данные сохранить in memory в сервисе (Кеш)
В случае падения сервиса восстанавливать Кеш из Postgres
Поднять http сервер и выдавать данные по id из кеша
Сделать простейший интерфейс отображения полученных данных, для их запроса по id
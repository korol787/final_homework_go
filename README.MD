# Итоговое задание по "Программирование на Go"
## Запуск
Сначала необходимо скачать исходный код проекта и перейти в скачанную папку.
```
git clone https://github.com/korol787/users-balance-microservice

cd users-balance-microservice
```
Для запуска API сервера используется следующая команда:
```
docker-compose up
```
Для запуска тестов:
```
docker-compose -f docker-compose.test.yml up --build --abort-on-container-exit
```
По умолчанию сервер будет доступен по адресу http://localhost:8080/.

## Описание API

Детальное описание каждого endpoint'а с примерами открывается по клику:

- [Получить баланс пользователя](https://github.com/korol787/users-balance-microservice/blob/master/docs/balance.md)
  :`POST /v1/deposits/balance`
- [Изменить баланс пользователя](https://github.com/korol787/users-balance-microservice/blob/master/docs/update.md)
  :`POST /v1/deposits/update`
- [Перевести деньги между двумя пользователями](https://github.com/korol787/users-balance-microservice/blob/master/docs/transfer.md)
  :`POST /v1/deposits/transfer`
- [Получить историю операций пользователя](https://github.com/korol787/users-balance-microservice/blob/master/docs/history.md)
  :`POST /v1/deposits/history`

Также есть небольшая коллекция запросов для запуска в Postman, которая находится в файле [postman_examples.json](https://github.com/korol787/users-balance-microservice/blob/master/postman_examples.json).
Для получения ожидаемых ответов сервера рекомендуется отправлять запросы в исходном порядке.

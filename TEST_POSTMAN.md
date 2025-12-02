-- Удаляем старые роли
TRUNCATE TABLE roles RESTART IDENTITY CASCADE;

-- Создаем роли заново
INSERT INTO roles (role_name) VALUES ('admin');   -- id = 1
INSERT INTO roles (role_name) VALUES ('manager'); -- id = 2
INSERT INTO roles (role_name) VALUES ('viewer');  -- id = 3

После этого можешь проверить:

SELECT id, role_name FROM roles;

Ты должен увидеть:

 id | role_name
----+-----------
 1  | admin
 2  | manager
 3  | viewer


 Эндпоинт:
POST http://localhost:7777/auth/register

Тело запроса (JSON):

{
  "username": "admin",
  "password": "admin123",
  "email": "admin@example.com",
  "role_id": 1
}


Что проверяем:
Проверка валидности роли – role_id должен существовать в таблице roles.
Хэширование пароля – пароль не сохраняется в открытом виде, только password_hash.
Ответ сервера – успешная регистрация возвращает:
{
  "id": 1
}
Если ошибка роли:
{
  "error": "[service-user] invalid role_id"
}
 Эндпоинт:
POST /auth/login:
{
  "username": "admin",
  "password": "admin123"
}
Если всё настроено правильно, ответ будет примерно такой:
{
  "token": "<JWT_TOKEN>"
}
<JWT_TOKEN> нужно будет использовать в заголовке Authorization для всех последующих запросов к /items:
Authorization: Bearer <JWT_TOKEN>
Это обеспечит проверку роли и аутентификацию.
После успешного логина можно переходить к CRUD на /items.

POST http://localhost:7777/items
Authorization: Bearer <JWT_TOKEN>
Content-Type: application/json

Тело запроса:
{
  "sku": "TEST-001",
  "title": "Test Item",
  "quantity": 10,
  "price": 99.99
}
Если всё настроено верно, ответ должен быть примерно таким:
{
  "id": 1
}

можно добавить еще больше различных товаров для наглядности


GET http://localhost:7777/items
Authorization: Bearer <JWT_TOKEN>
Пример ответа:
[
    {
        "id": 1,
        "sku": "TEST-001",
        "title": "Test Item",
        "quantity": 10,
        "price": 99.99,
        "version": 1,
        "created_at": "2025-12-02T00:39:55.196307Z",
        "updated_at": "2025-12-02T00:39:55.196307Z"
    },
    {
        "id": 2,
        "sku": "TEST-002",
        "title": "Test Item2",
        "quantity": 101,
        "price": 919.99,
        "version": 1,
        "created_at": "2025-12-02T00:43:27.051492Z",
        "updated_at": "2025-12-02T00:43:27.051492Z"
    },
    {
        "id": 3,
        "sku": "TEST-003",
        "title": "Test Item3",
        "quantity": 1,
        "price": 19.9,
        "version": 1,
        "created_at": "2025-12-02T00:43:45.07245Z",
        "updated_at": "2025-12-02T00:43:45.07245Z"
    },
    {
        "id": 4,
        "sku": "TEST-004",
        "title": "Test Item4",
        "quantity": 5,
        "price": 200,
        "version": 1,
        "created_at": "2025-12-02T00:44:04.851308Z",
        "updated_at": "2025-12-02T00:44:04.851308Z"
    }
]

GET http://localhost:7777/items/3
Authorization: Bearer <JWT_TOKEN>
Пример ответа:
{
    "id": 3,
    "sku": "TEST-003",
    "title": "Test Item3",
    "quantity": 1,
    "price": 19.9,
    "version": 1,
    "created_at": "2025-12-02T00:43:45.07245Z",
    "updated_at": "2025-12-02T00:43:45.07245Z"
}


Обновление товара (PUT /items/1)

PUT http://localhost:7777/items/1
Authorization: Bearer <JWT_TOKEN>

Тело запроса:

{
  "sku": "TEST-001",
  "title": "Test Item Updated",
  "quantity": 15,
  "price": 109.99
}
Ожидаемый ответ:
{
    "status": "item updated"
}

DELETE http://localhost:7777/items/4
Authorization: Bearer <JWT_TOKEN>
Ожидаемый ответ:
{
  "status": "item deleted"
}
После удаления можно проверить:
Проверка, что товар исчез:
GET /items

 история
 GET http://localhost:7777/items/3/history

Ожидаемый ответ:
 [
    {
        "id": 2,
        "item_id": 2,
        "action_type": "insert",
        "old_value": null,
        "new_value": {
            "id": 2,
            "sku": "TEST-002",
            "price": 919.99,
            "title": "Test Item2",
            "version": 1,
            "quantity": 101,
            "created_at": "2025-12-02T00:43:27.051492",
            "updated_at": "2025-12-02T00:43:27.051492"
        },
        "changed_by": 1,
        "changed_at": "2025-12-02T00:43:27.051492Z"
    }
]

фильтрция историй
GET /items/:id/history/filter
Параметры передаются через query string:
user_id — фильтр по пользователю (целое число)
action_type — фильтр по типу действия (insert или update)
from — дата начала, например 2025-12-01T00:00:00Z
to — дата конца, например 2025-12-02T00:00:00Z

GET http://localhost:7777/items/1/history/filter?user_id=1&action_type=update&from=2025-12-01T00:00:00Z&to=2025-12-03T00:00:00Z
Authorization: Bearer <токен>

[
    {
        "id": 7,
        "item_id": 1,
        "action_type": "update",
        "old_value": {
            "id": 1,
            "sku": "TEST-001",
            "price": 99.99,
            "title": "Test Item",
            "version": 1,
            "quantity": 10,
            "created_at": "2025-12-02T00:39:55.196307",
            "updated_at": "2025-12-02T00:39:55.196307"
        },
        "new_value": {
            "id": 1,
            "sku": "TEST-001",
            "price": 494.99,
            "title": "Test 34",
            "version": 2,
            "quantity": 5655,
            "created_at": "2025-12-02T00:39:55.196307",
            "updated_at": "2025-12-02T13:26:53.339559"
        },
        "changed_by": 1,
        "changed_at": "2025-12-02T13:26:53.339559Z"
    }
]
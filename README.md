# Song Management API

## Описание проекта

Song Management API — это **RESTful API** для управления музыкальными треками. Реализованы функциональности для создания, чтения, обновления и удаления данных (CRUD) с поддержкой фильтрации, пагинации и работы с текстами песен.

-----

## Технологии

- Язык программирования: Go (Golang)
- База данных: PostgreSQL
- Web-фреймворк: Echo
- Документация: Swagger

## Установка
**Склонируйте репозиторий:**

```
git clone https://github.com/m1sol/musinfo.git
cd musinfo
```

Убедитесь, что у вас установлен Go версии 1.23 и PostgreSQL.
Скопируйте файл .env.example в .env и заполните переменные среды:

`cp .env.example .env`

Миграции базы данных находятся в директории */migrations* и устанавливаются автоматически при запуске проекта

## Использование

После запуска API будет доступно по адресу **http://localhost:8020**.

### Примеры запросов

**Получить список песен с фильтрацией и пагинацией:**

`GET /songs?page=1&limit=10&song=Muse`

**Ответ:**
```
Content-Type: application/json
{
  "status": 200,
  "data": [
        {
            "group": "Muse",
            "song": "Uprising",
            "releaseDate": "2009-09-07",
            "text": "The paranoia is in bloom...",
            "link": "https://example.com"
        }
        ...
    ]
}
```
**Добавить песню:**

`POST /songs`
```
{
    "group": "Muse",
    "song": "Uprising",
}
```

**Обновить информацию о песни:**

`PUT /songs/{id}`

```
{
    "group": "Muse",
    "song": "Uprising",
    "releaseDate": "2009-09-07",
    "text": "The paranoia is in bloom...",
    "link": "https://example.com"
}
```

**Получить информацию о песне по id с пагинацией по тексту песни**

`GET /songs/{id}?page=1&limit=10&`

**Ответ:**
```
Content-Type: application/json
{
  "status": 200,
  "data": {
        "group": "Muse",
        "song": "Uprising",
        "releaseDate": "2009-09-07",
        "text": "The paranoia is in bloom...",
        "link": "https://example.com"
    }
}
```

**Удалить песню:**

`DELETE /songs/{id}`

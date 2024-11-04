# Simple Books REST API


# Examples


### GET /books
```
curl --location --request GET 'localhost:8080/books' \
--header 'Content-Type: application/json' \
--data '{
    "title": "New Book",
    "author": "Arthur C",
    "price": 24.99
}'
```

Response:
```
{
  "data": [
    {
      "id": "1",
      "title": "Book One",
      "author": "Author A",
      "price": 19.99
    },
    {
      "id": "2",
      "title": "Book Two",
      "author": "Author B",
      "price": 29.99
    }
  ]
}
```


### GET /books/1
```
curl --location --request GET 'localhost:8080/books/1' \
--header 'Content-Type: application/json' \
--data '{
    "title": "Updated New Book",
    "author": "Arthur C",
    "price": 24.99
}'
```

Response:
```
{
  "data": {
    "id": "1",
    "title": "New Book",
    "author": "Arthur C",
    "price": 24.99
  }
}
```


### POST /books
```
curl --location 'localhost:8080/books' \
--header 'Content-Type: application/json' \
--data '{
    "title": "New Book",
    "author": "Arthur C",
    "price": 24.99
}'
```

Response:
```
{
  "data": {
    "id": "6",
    "title": "New Book",
    "author": "Arthur C",
    "price": 24.99
  }
}
```


### PUT /books/1
```
curl --location --request PUT 'localhost:8080/books/1' \
--header 'Content-Type: application/json' \
--data '{
    "title": "Updated New Book",
    "author": "Arthur C",
    "price": 24.99
}'
```

Response:
```
{
  "data": {
    "id": "1",
    "title": "Updated New Book",
    "author": "Arthur C",
    "price": 24.99
  }
}
```


# TODO
  - Add tests
  - Add OpenAPI 3 documentation

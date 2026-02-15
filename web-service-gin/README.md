## [Developing a RESTful API with Go and Gin](https://go.dev/doc/tutorial/web-service-gin)

This tutorial introduces the basics of writing a RESTful web service API with Go and the [Gin Web Framework](https://gin-gonic.com/en/docs/) (Gin).

Gin simplifies many coding tasks associated with building web applications, including web services. In this tutorial, you’ll use Gin to route requests, retrieve request details, and marshal JSON for responses.

This tutorial, we build a RESTful API server with two endpoints. Your example project will be a repository of data about vintage jazz records [db_access](https://github.com/konsbe/go-guide/tree/main/db_access).

The tutorial includes the following sections:

- Design API endpoints.
- Create a folder for your code.
- Create the data.
- Write a handler to return all items.
- Write a handler to add a new item.
- Write a handler to return a specific item.

We build an API that provides access to a store selling vintage recordings on vinyl. So we are going to provide endpoints through which a client can get and add albums for users.

When developing an API, we typically begin by designing the endpoints. The API’s users will have more success if the endpoints are easy to understand.

Here are the endpoints we’ll create in this tutorial.

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/albums` | Get a list of all albums, returned as JSON |
| POST | `/albums` | Add a new album from request data sent as JSON |
| GET | `/albums/:id` | Get an album by its ID, returning the album data as JSON |

From a new command line window, use curl to make a request to your running web service.

- GET request (/albums)
```bash
curl http://localhost:8080/albums
```
    [GIN] 2026/02/15 - 22:46:03 | 200 |     1.50721ms |       127.0.0.1 | GET      "/albums"

    [
        {
            "id": "1",
            "title": "Blue Train",
            "artist": "John Coltrane",
            "price": 56.99
        },
        {
            "id": "2",
            "title": "Jeru",
            "artist": "Gerry Mulligan",
            "price": 17.99
        },
        {
            "id": "3",
            "title": "Sarah Vaughan and Clifford Brown",
            "artist": "Sarah Vaughan",
            "price": 39.99
        },
        {
            "id": "4",
            "title": "The Modern Sound of Betty Carter",
            "artist": "Betty Carter",
            "price": 49.99
        }
    ]


- POST request (/albums)
```bash
curl http://localhost:8080/albums \
    --include \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '{"id": "4","title": "The Modern Sound of Betty Carter","artist": "Betty Carter","price": 49.99}'
```
    [GIN] 2026/02/15 - 22:43:49 | 201 |     132.685µs |       127.0.0.1 | POST     "/albums"

    HTTP/1.1 201 Created
    Content-Type: application/json; charset=utf-8
    Date: Wed, 02 Jun 2021 00:34:12 GMT
    Content-Length: 116

    {
        "id": "4",
        "title": "The Modern Sound of Betty Carter",
        "artist": "Betty Carter",
        "price": 49.99
    }

- GET request (/albums/:id)
```bash
curl http://localhost:8080/albums/2
```
    [GIN] 2026/02/15 - 22:53:48 | 200 |      92.638µs |       127.0.0.1 | GET      "/albums/2"

    {
        "id": "2",
        "title": "Jeru",
        "artist": "Gerry Mulligan",
        "price": 17.99
    }





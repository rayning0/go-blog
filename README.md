# Blog in Golang

Golang API for blog: users and posts

Run: `go run main.go`

`brew install httpie`. [HTTPie](https://httpie.org/) is a user-friendly command-line HTTP client. Easier form of `curl`.

POST: `http POST http://localhost:8080/books author="Raymond Gan" title="My First Book"`\
GET: `http http://localhost:8080/books`\
GET: `http http://localhost:8080/books/1`\
PATCH: `http PATCH http://localhost:8080/books/1 title="My Second Book"`\
DELETE: `http DELETE http://localhost:8080/books/1`

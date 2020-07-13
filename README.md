# Blog in Golang

Golang API for blog: users and posts

Run: `go run main.go`

`brew install httpie`. [HTTPie](https://httpie.org/) is a user-friendly command-line HTTP client. Easier form of `curl`.

POST: `http POST http://localhost:8080/users name="Mickey Mouse" email="mickey@mouse.com"`\
GET: `http http://localhost:8080/users`\
GET: `http http://localhost:8080/users/1`\
PATCH: `http PATCH http://localhost:8080/users/1 email="mickey2@mouse.com"`\
DELETE: `http DELETE http://localhost:8080/users/1`

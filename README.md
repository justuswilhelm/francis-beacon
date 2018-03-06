# Francis Beacon

Go + Postgres Beacon

## Quickstart

1. Set the environment variable `DATABASE_URL` to equal a Postgres database URL.
2. Ensure [glide](https://github.com/Masterminds/glide) is installed
3. Run `glide i`
4. Run `go run web.go`
5. In a separate window run `curl localhost:8080 -v -H 'Referer: Test'`

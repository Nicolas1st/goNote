# Simple Note API

An API project built with the idea of separting database and api code.

## Structure

cmd
- [main.go](./cmd/main.go)

  Launches the API

app
- api
    Contains endpoints handlers.
    The router for enpoints is built with dependency injection

- model
    Contains models
    Database query functions
    Migrations

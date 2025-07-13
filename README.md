# goNote

An application for note taking. Built in such a way so that
it'd possible to run it both on a server and in the browser only (no server part).

# Structure

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

# Golang REST API

## Requirements

- [`direnv`](https://direnv.net/) for loading environment variables.

## Goals

- [x] Controller-based dependency injection (inject dependency in the
      controller)
- [ ] Define external services (db, cache, logging and more) and inject
      accordingly to each service handler
- [ ] Custom validator for each controller connected to Echo
      in order to return appropriate status codes.

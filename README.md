# Golang REST API

## Requirements

- [`direnv`](https://direnv.net/) for loading environment variables.

## Goals

### Project

- [ ] Makefile for all commands
- [ ] Hot reloading with Air
- [ ] Docker-based

### Architecture

- [x] Dependency injection on each handler (inject service interfaces on
      handlers to abstract external service communication)
- [ ] Define external services communication (db, cache, logging and more) and
      inject accordingly to each service handler
- [ ] Custom validator for each controller connected to Echo
      in order to return appropriate status codes.

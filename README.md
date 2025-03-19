# Golang REST API

## Requirements

- [`Go`](https://go.dev/) (>= 1.23)
- Docker and Docker Compose
- [Swagger](https://github.com/swaggo/swag/cmd/swag) for API documentation.
- [`goose`](https://github.com/pressly/goose) for database migrations.
- [`air`](https://github.com/air-verse/air) for hot reloading
- [`direnv`](https://direnv.net/) for loading environment variables.
- [`Prettier`](https://prettier.io/) for formatting Markdown, JSON and related files.

## Installation

```
make reqs
sudo apt-get install direnv # Install this package according to your Linux flavor
```

## Usage

The command below list all possible commands that you might want to use:

```bash
make help
```

## Installing new drivers

```bash
./scripts/install_driver.sh
```

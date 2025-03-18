# Golang REST API

## Requirements

- [`Go`](https://go.dev/) (>= 1.23)
- [`goose`](https://github.com/pressly/goose) for database migrations.
- [`air`](https://github.com/air-verse/air) for hot reloading
- [`direnv`](https://direnv.net/) for loading environment variables.
- [`Prettier`](https://prettier.io/) for formatting Markdown, JSON and related files.

## Installation

- [Go](https://go.dev/doc/install)

- Goose

  ```
  go install github.com/pressly/goose/v3/cmd/goose@latest
  ```

- Air

  ```bash
  go install github.com/air-verse/air@latest
  ```

- direnv

  ```bash
  sudo apt-get install direnv # Debian
  ```

- Prettier

  ```bash
  npm install -g prettier
  ```

## Usage

The command below list all possible commands that you might want to use:

```bash
make help
```

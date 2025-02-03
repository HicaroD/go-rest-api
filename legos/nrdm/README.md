# NRDM - Non-Relational Database Manager

## Drivers supported

- [MongoDB](./mongo.go)

## How to install a new driver

Simply paste its respective file into this directory in your local environment
and you're good to go.

## Good practices

1. If you don't need some specific driver, just remove its respective file and
   clear all packages (tip: `go mod tidy` after removing the file)

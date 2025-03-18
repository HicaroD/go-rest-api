#!/bin/bash

install_driver() {
    case "$1" in
        mysql)
            go get github.com/go-sql-driver/mysql
            ;;
        postgresql|postgres)
            go get github.com/lib/pq
            ;;
        sqlite|sqlite3)
            go get github.com/mattn/go-sqlite3
            ;;
        *)
            echo "Unsupported database type. Choose from: mysql, postgresql, sqlite"
            exit 1
            ;;
    esac
}

if [ -z "$1" ]; then
    echo "Usage: $0 <mysql|postgresql|sqlite>"
    exit 1
fi

echo "Installing $1 driver..."
install_driver "$1"
echo "$1 driver installed successfully!"

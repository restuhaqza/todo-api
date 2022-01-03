#!/bin/bash

export MYSQL_HOST="tcp(127.0.0.1:3306)"
export MYSQL_PASSWORD="passw0rd"
export MYSQL_USER="user"
export MYSQL_DBNAME="todo_db"

go run main.go
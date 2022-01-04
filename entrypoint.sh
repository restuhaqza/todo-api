#!/bin/bash

export MYSQL_HOST="tcp(127.0.0.1:3306)"
export MYSQL_PASSWORD="passw0rd"
export MYSQL_USER="root"
export MYSQL_DBNAME="restu_haqqi_muzakir"

go run main.go
#!/bin/bash


function create_postgres() {
  docker run --name bank_db -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=123 -d postgres:12-alpine
}

function start_db() {
  docker start bank_db
}

function stop_db() {
  docker stop bank_db
}

function create_db() {
	docker exec bank_db createdb --username=root --owner=root simple_bank
}

function drop_db() {
	docker exec bank_db dropdb simple_bank
}

function migrate_up() {
	migrate -path db/migration -database "postgresql://root:123@localhost:5432/simple_bank?sslmode=disable" -verbose up
}

function migrate_down() {
	migrate -path db/migration -database "postgresql://root:123@localhost:5432/simple_bank?sslmode=disable" -verbose down
}

function sqlc_gen() {
  docker run --rm -v "%cd%:/src" -w /src sqlc/sqlc generate
#  docker run --rm -v "C:\programming\go\web-bank:/src" -w /src sqlc/sqlc generate
  read -n 1 -p "Press any key to exit" x
}

function test() {
    go test -v -cover ./...
}


name="${1?needs one argument}"
case $name in
        "create_postgres")
                create_postgres
                ;;
        "start_db")
                start_db
                ;;
        "stop_db")
                stop_db
                ;;
        "create_db")
                create_db
                ;;
        "drop_db")
                drop_db
                ;;
        "migrate_up")
                migrate_up
                ;;
        "migrate_down")
                migrate_down
                ;;
        "sqlc_gen")
                sqlc_gen
                ;;
        "test")
                test
                ;;
        *)
                echo "Not a valid argument"
                ;;
esac

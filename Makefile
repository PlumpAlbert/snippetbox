#!make
include .env
export $(shell sed 's/=.*//' .env)

dev:
	echo "./cmd/web -dsn ${DATABASE_URL}"
	fd -t f --absolute-path | entr -cr go run ./cmd/web -dsn ${DATABASE_URL}

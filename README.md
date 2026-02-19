# LeGo

I want to hit a btn to deploy an amazing product! This repo is still under developement! Don't use it :)

## Getting started 

### Database 

For dev env use this command for database setup:

```bash
docker compose up -d
```

### Migrations
I use [golang-migrate/migrate](https://github.com/golang-migrate/migrate) package for migrations.
You can run it with docker:
```bash
docker run -v $PWD/database/migrations:/migrations \
  --network host \
  migrate/migrate \
  -path=/migrations \
  -database "postgres://postgres:postgres@localhost:5432/go?sslmode=disable" \
  up
```

## JWT auth 

I wanna use JWT tokens for authentication:

```bash
go get github.com/golang-jwt/jwt/v5
```
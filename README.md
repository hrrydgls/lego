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
### Hot reloading 

I use [Air](https://github.com/air-verse/air) for live reloading. So you just need to run `air` in dev env to run it...

## Routing

I am using [chi](https://github.com/go-chi/chi) for routing. 


## JWT auth 

I wanna use JWT tokens for authentication:

```bash
go get github.com/golang-jwt/jwt/v5
```

## Todo 

- [x] Handle exceptions better
- [x] Refactor a bit
- [x] Use a package for routing 
- [ ] Create a middleware for auth
- [ ] Write some tests
- [ ] Use a package for env vars
- [ ] Encrypt passwords
- [ ] More info in the token, like name and email 
- [ ] Better logging
- [ ] Metrics
- [ ] Use a package for validation
- [ ] Rate limiter for login and register


## More general suggestion

### Go playground

If you wanna test sth on the fly, you can do it here: [https://go.dev/play/](https://go.dev/play/)
# sakaba-api
API repository for sakaba.link

## Build Docker image
```sh
docker build --no-cache -t sakaba/api .
```

&nbsp;

## Launch Docker container
```sh
docker run --rm -d \
    -e APP_CONFIG_JSON="{\"db.password\":\"DB_PWD\",\"db.host\":\"DB_HOST\",\"db.name\":\"DB_NAME\",\"db.user\":\"DB_USER\",\"aws.s3.id\":\"S3_ID\",\"aws.s3.secret\":\"S3_SECRET\",\"aws.s3.region\":\"S3_REGION\"}" \
    -p 8080:8080 \
    sakaba/api
```

&nbsp;

## Login
```sh
curl -X POST \
    -H 'Content-Type: application/json' \
    -d '{"email":"example@sakaba.link", "password":"xxxx"}' \
    http://localhost:8080/login
```

&nbsp;

## Refresh token
```sh
curl -X GET \
    -H 'Content-Type: application/json' \
    -H 'Authorization:Bearer xxxxxxxxx' \
    http://localhost:8080/auth/refresh_token
```

## Access
```sh
curl -X GET \
    -H 'Content-Type: application/json' \
    -H 'Authorization:Bearer xxxxxxxxx' \
    http://localhost:8080/auth/home
```

## Call add restaurant genre API
```sh
curl -X POST \
    -H 'Authorization:Bearer xxxxxxxxx' \
    -H 'Content-Type: application/json; charset=utf-8' \
    -d '{ "restaurant_id": "{RESTAURANT_ID}", "genre_id": "{GENRE_ID}" }' \
    http://localhost:8080/auth/restaurant-genre/
```

## Call add restaurant API
```sh
curl -X POST \
    -H 'Authorization:Bearer xxxxxxxxx' \
    -H 'Content-Type: application/json; charset=utf-8' \
    -d '{ "url": "{URL}", "name": "{NAME}", "genre": "{GENRE}", "tel": "{TEL}", "business_day_info": "{BUSINESS_DAY_INFO}", "address": "{ADDRESS}", "latitude": "{LATITUDE}", "longitude": "{LONGITUDE}", "area": "{AREA}" }' \
    http://localhost:8080/auth/restaurant/
```

&nbsp;

## Reference
- [Implementation of constructing minimal docker image of golang application](https://developpaper.com/implementation-of-constructing-minimal-docker-image-of-golang-application/)

&nbsp;

## Misc
### Install Go
```sh
brew install go
```

### Install `golint`
```sh
go install golang.org/x/lint/golint@latest
```

### Install shadow linter
```sh
go install golang.org/x/tools/go/analysis/passes/shadow/cmd/shadow@latest
```

### Linting and Vetting
```sh
golint ./...
```

```sh
go vet ./...
```

```sh
shadow ./...
```

### List environment variables for Go
```sh
go env
```

### Remove unused modules
```sh
go mod tidy -v
```

&nbsp;

## TODOs
- Read [Effective Go](https://go.dev/doc/effective_go)
- Use [golangci-lint](https://oreil.ly/O15u-)
- [Speed up Amazon ECS container deployments](https://nathanpeck.com/speeding-up-amazon-ecs-container-deployments/)

&nbsp;

## Links
- [Add JWT Authentication](https://github.com/appleboy/gin-jwt)
- [Delve debugger](https://oreil.ly/sosLu)
- [gopls](https://oreil.ly/TLapT)

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
    -e APP_CONFIG_JSON="{\"db.password\":\"DB_PWD\",\"db.host\":\"DB_HOST\",\"db.name\":\"DB_NAME\",\"db.user\":\"DB_USER\"}" \
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

## Call Admin API
```sh
curl -X POST \
    -H 'Authorization:Bearer xxxxxxxxx' \
    -H 'Content-Type: application/json; charset=utf-8' \
    -d '{ "restaurant_id": "{RESTAURANT_ID}", "genre_id": "{GENRE_ID}" }' \
    http://localhost:8080/auth/restaurant-genre/
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

### Linting and Vetting
```sh
golint ./...
```

```sh
go vet ./...
```

### List environment variables for Go
```sh
go env
```

### Remove unused modules
```sh
go mod tidy -v
```

## Deployment
### Deploy Lambda
```bash
cd xxx_xxx && ./deploy.sh
```

## Misc
### Create Lambda Function
```bash
aws lambda create-function \
    --function-name <FUNCTION_NAME> \
    --runtime python3.8 \
    --zip-file fileb://<ZIP_NAME> \
    --handler lambda_function.lambda_handler \
    --role arn:aws:iam::xxxxxxxxxxxx:role/lambda-vpc-role \
    --vpc-config SubnetIds=subnet-xxxxxxxxxxxxxxxxx,subnet-yyyyyyyyyyyyyyyyy,SecurityGroupIds=sg-xxxxxxxxxxxxxxxxx
```
### Remove Lambda Function
```bash
aws lambda delete-function --function-name <FUNCTION_NAME>
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

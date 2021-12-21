# sakaba-api
API repository for sakaba.link

## Build Docker image
```sh
docker build --no-cache -t sakaba/api .
```

## Launch Docker container
```sh
docker run --rm -d \
    -e SPRING_APPLICATION_JSON="{\"db.password\":\"DB_PWD\",\"db.host\":\"DB_HOST\",\"db.name\":\"DB_NAME\",\"db.user\":\"DB_USER\"}" \
    -p 8080:8080 \
    sakaba/api
```

## Access
```sh
open http://localhost:8080
```

## TODO
- Reduce Docker image size.
  - https://chemidy.medium.com/create-the-smallest-and-secured-golang-docker-image-based-on-scratch-4752223b7324
  - https://developpaper.com/implementation-of-constructing-minimal-docker-image-of-golang-application/

## Misc
### Remove unused modules
```sh
go mod tidy -v
```

## Deployment
### Deploy addMenu Lambda
```bash
cd add_menu && ./deploy.sh
```
### Deploy addPhotos Lambda
```bash
cd add_photos && ./deploy.sh
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



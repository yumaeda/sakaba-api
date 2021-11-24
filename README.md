# tokyo-takeout-api
API repository for sakaba.link

## Deployment
### Deploy getRestaurants Lambda
```bash
cd get_restaurants && ./deploy.sh
```
### Deploy getPhotos Lambda
```bash
cd get_photos && ./deploy.sh
```

## Test
### Test getRestaurants Lambda
```bash
aws lambda invoke --function-name getRestaurants output.json
```
### Test getPhotos Lambda
```bash
aws lambda invoke --function-name getPhotos output.json
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


# Go

## Build Docker image
```sh
docker build --no-cache -t sakaba/api .
```

## Create .env file
```sh
cp ./env.example ./.env
```

## Launch Docker container
```sh
docker run -d -p 8080:8080 sakaba/api
```

## Access
```sh
open http://localhost:8080
```

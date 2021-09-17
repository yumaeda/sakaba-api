#!/bin/sh

./clean.sh

cp ../common.py ./common.py

pip3 install --target ./package pymysql
cd package
zip -r9 ${OLDPWD}/get_restaurants.zip .
cd $OLDPWD
zip -g get_restaurants.zip lambda_function.py
zip -g get_restaurants.zip common.py

aws lambda update-function-code \
    --function-name getRestaurants \
    --zip-file fileb://get_restaurants.zip

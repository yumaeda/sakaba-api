#!/bin/sh

./clean.sh

cp ../common.py ./common.py

pip3 install --target ./package pymysql
cd package
zip -r9 ${OLDPWD}/get_restaurant_count.zip .
cd $OLDPWD
zip -g get_restaurant_count.zip lambda_function.py
zip -g get_restaurant_count.zip common.py

aws lambda update-function-code \
    --function-name getRestaurantCount \
    --zip-file fileb://get_restaurant_count.zip

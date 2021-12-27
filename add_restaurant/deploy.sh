#!/bin/sh

./clean.sh

cp ../common.py ./common.py

pip3 install --target ./package pymysql
cd package
zip -r9 ${OLDPWD}/add_restaurant.zip .
cd $OLDPWD
zip -g add_restaurant.zip common.py
zip -g add_restaurant.zip lambda_function.py

aws lambda update-function-code \
    --function-name addRestaurant \
    --zip-file fileb://add_restaurant.zip

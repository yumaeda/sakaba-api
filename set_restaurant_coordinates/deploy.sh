#!/bin/sh

./clean.sh

cp ../common.py ./common.py

pip3 install --target ./package pymysql requests
cd package
zip -r9 ${OLDPWD}/set_restaurant_coordinates.zip .
cd $OLDPWD
zip -g set_restaurant_coordinates.zip lambda_function.py
zip -g set_restaurant_coordinates.zip common.py

aws lambda update-function-code \
    --function-name setRestaurantCoordinates \
    --zip-file fileb://set_restaurant_coordinates.zip

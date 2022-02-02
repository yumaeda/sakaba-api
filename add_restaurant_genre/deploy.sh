#!/bin/sh

./clean.sh

cp ../common.py ./common.py

pip3 install --target ./package pymysql
cd package
zip -r9 ${OLDPWD}/add_restaurant_genre.zip .
cd $OLDPWD
zip -g add_restaurant_genre.zip common.py
zip -g add_restaurant_genre.zip lambda_function.py

aws lambda update-function-code \
    --function-name addRestaurantGenre \
    --zip-file fileb://add_restaurant_genre.zip

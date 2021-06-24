#!/bin/sh

./clean.sh

pip3 install --target ./package pymysql
cd package
zip -r9 ${OLDPWD}/get_wines.zip .
cd $OLDPWD
zip -g get_wines.zip lambda_function.py

aws lambda update-function-code \
    --function-name getWines \
    --zip-file fileb://get_wines.zip

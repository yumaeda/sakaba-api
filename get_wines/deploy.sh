#!/bin/sh

./clean.sh

cp ../common.py ./common.py

pip3 install --target ./package pymysql
cd package
zip -r9 ${OLDPWD}/get_wines.zip .
cd $OLDPWD
zip -g get_wines.zip lambda_function.py
zip -g get_wines.zip common.py

aws lambda update-function-code \
    --function-name getWines \
    --zip-file fileb://get_wines.zip

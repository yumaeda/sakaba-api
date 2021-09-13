#!/bin/sh

./clean.sh

pip3 install --target ./package pymysql
cd package
zip -r9 ${OLDPWD}/set_menu.zip .
cd $OLDPWD
zip -g set_menu.zip lambda_function.py

aws lambda update-function-code \
    --function-name setMenu \
    --zip-file fileb://set_menu.zip

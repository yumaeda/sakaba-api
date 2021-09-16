#!/bin/sh

./clean.sh

cp ../common.py ./common.py

pip3 install --target ./package pymysql
cd package
zip -r9 ${OLDPWD}/add_menu.zip .
cd $OLDPWD
zip -g add_menu.zip common.py
zip -g add_menu.zip lambda_function.py

aws lambda update-function-code \
    --function-name addMenu \
    --zip-file fileb://add_menu.zip

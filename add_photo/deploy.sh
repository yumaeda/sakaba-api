#!/bin/sh

./clean.sh

pip3 install --target ./package pymysql
cd package
zip -r9 ${OLDPWD}/add_photo.zip .
cd $OLDPWD
zip -g add_photo.zip lambda_function.py

aws lambda update-function-code \
    --function-name addPhoto \
    --zip-file fileb://add_photo.zip

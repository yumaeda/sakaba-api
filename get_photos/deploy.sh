#!/bin/sh

./clean.sh

pip3 install --target ./package pymysql
cd package
zip -r9 ${OLDPWD}/get_photos.zip .
cd $OLDPWD
zip -g get_photos.zip lambda_function.py

aws lambda update-function-code \
    --function-name getPhotos \
    --zip-file fileb://get_photos.zip


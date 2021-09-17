#!/bin/sh

./clean.sh

cp ../common.py ./common.py

pip3 install --target ./package pymysql
cd package
zip -r9 ${OLDPWD}/get_photos.zip .
cd $OLDPWD
zip -g get_photos.zip lambda_function.py
zip -g get_photos.zip common.py

aws lambda update-function-code \
    --function-name getPhotos \
    --zip-file fileb://get_photos.zip


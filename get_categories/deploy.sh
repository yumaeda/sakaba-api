#!/bin/sh

./clean.sh

cp ../common.py ./common.py

pip3 install --target ./package pymysql
cd package
zip -r9 ${OLDPWD}/get_categories.zip .
cd $OLDPWD
zip -g get_categories.zip lambda_function.py
zip -g get_categories.zip common.py

aws lambda update-function-code \
    --function-name getCategories \
    --zip-file fileb://get_categories.zip

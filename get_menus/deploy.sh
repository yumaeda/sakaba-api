#!/bin/sh

./clean.sh

cp ../common.py ./common.py

pip3 install --target ./package pymysql
cd package
zip -r9 ${OLDPWD}/get_menus.zip .
cd $OLDPWD
zip -g get_menus.zip lambda_function.py
zip -g get_menus.zip common.py

aws lambda update-function-code \
    --function-name getMenus \
    --zip-file fileb://get_menus.zip

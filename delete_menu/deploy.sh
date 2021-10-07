#!/bin/sh

./clean.sh

cp ../common.py ./common.py

pip3 install --target ./package pymysql
cd package
zip -r9 ${OLDPWD}/delete_menu.zip .
cd $OLDPWD
zip -g delete_menu.zip common.py
zip -g delete_menu.zip lambda_function.py

aws lambda update-function-code \
    --function-name deleteMenu \
    --zip-file fileb://delete_menu.zip

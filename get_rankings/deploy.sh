#!/bin/sh

./clean.sh

cp ../common.py ./common.py

pip3 install --target ./package pymysql
cd package
zip -r9 ${OLDPWD}/get_rankings.zip .
cd $OLDPWD
zip -g get_rankings.zip lambda_function.py
zip -g get_rankings.zip common.py

aws lambda update-function-code \
    --function-name getRankings \
    --zip-file fileb://get_rankings.zip

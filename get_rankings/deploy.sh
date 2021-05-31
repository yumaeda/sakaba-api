#!/bin/sh

./clean.sh

pip3 install --target ./package pymysql
cd package
zip -r9 ${OLDPWD}/get_rankings.zip .
cd $OLDPWD
zip -g get_rankings.zip lambda_function.py

aws lambda update-function-code \
    --function-name getRankings \
    --zip-file fileb://get_rankings.zip


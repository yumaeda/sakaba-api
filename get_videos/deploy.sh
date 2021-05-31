#!/bin/sh

./clean.sh

pip3 install --target ./package pymysql
cd package
zip -r9 ${OLDPWD}/get_videos.zip .
cd $OLDPWD
zip -g get_videos.zip lambda_function.py

aws lambda update-function-code \
    --function-name getVideos \
    --zip-file fileb://get_videos.zip

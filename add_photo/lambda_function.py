# System Module
import base64
import boto3
from botocore.config import Config
import logging
import os
import sys
import uuid

# 3rd Party Module
import pymysql

logger = logging.getLogger()
logger.setLevel(logging.INFO)

BUCKET_NAME = os.environ['S3_BUCKET_NAME']

try:
    conn = pymysql.connect(
        host=os.environ['DB_HOST'],
        user=os.environ['DB_USER'],
        passwd=os.environ['DB_PASSWORD'],
        db=os.environ['DB_NAME'],
        connect_timeout=10,
        charset='utf8mb4',
        cursorclass=pymysql.cursors.DictCursor
    )
except pymysql.MySQLError as e:
    logger.error('ERROR: Could not connect to MariaDB instance.')
    logger.error(e)
    sys.exit()

logger.info('SUCCESS: Connection to RDS MariaDB instance succeeded')

def insert_photo_metadata(restaurant_id: str) -> str:
    """
    Insert photo meta data to DB.
    """
    if restaurant_id is not None and restaurant_id != '':
        with conn.cursor() as cursor:
            try:
                file_name = str(uuid.uuid4())
                insert_sql = 'INSERT INTO photos(restaurant_id, name, type) VALUES (UuidToBin(\'{}\'), \'{}\', \'dish\')'.format(
                    restaurant_id,
                    file_name
                )
                cursor.execute(insert_sql)
                conn.commit()
                return file_name
            except Exception as ex:
                logging.exception(ex)
    return None

def lambda_handler(event, context):
    """
    This function fetches content from MySQL RDS instance
    """
    if 'restaurant_id' in event and 'file_content' in event:
        restaurant_id = event['restaurant_id']
        file_name = insert_photo_metadata(restaurant_id)
        if file_name is not None:
            file_content = base64.b64decode(event['file_content'])
            file_path = 'images/restaurants/{}/{}.jpeg'.format(restaurant_id, file_name)
            print('Store {} to S3'.format(file_path))
            s3_config = Config(
                region_name = 'ap-northeast-1',
                s3 = { 'addressing_style': 'path' }
            )
            try:
                s3 = boto3.client('s3', config=s3_config)
                s3.put_object(
                    Bucket=BUCKET_NAME,
                    Key=file_path,
                    Body=file_content
                )
                return { 'statusCode': 200, 'body': file_path }
            except Exception as ex:
                logging.exception(ex)
    return { 'statusCode': 400, 'body': 'Invalid parameters' }

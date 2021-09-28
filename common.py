# System Module
import logging
import os
import sys

# 3rd Party Module
import pymysql

logger = logging.getLogger()
logger.setLevel(logging.INFO)

def get_response(status_code: int, body: str):
    return {
        'statusCode': status_code,
        'body': body
    }

try:
    conn = pymysql.connect(
        host=os.environ['DB_HOST'],
        user=os.environ['DB_USER'],
        passwd=os.environ['DB_PASSWORD'],
        db=os.environ['DB_NAME'],
        connect_timeout=10,
        charset='utf8mb4',
        cursorclass=pymysql.cursors.DictCursor,
        autocommit=True
    )
except pymysql.MySQLError as e:
    logger.error('ERROR: Could not connect to MariaDB instance.')
    logger.error(e)
    sys.exit()

logger.info('SUCCESS: Connection to RDS MariaDB instance succeeded')

STATUS_CODE_OK = 200
STATUS_CODE_BAD_REQUEST = 400
STATUS_CODE_INTERNAL_SERVER_ERROR = 500
 
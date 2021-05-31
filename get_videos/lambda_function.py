# System Module
import json
import os
import logging
import sys

# 3rd Party Module
import pymysql

logger = logging.getLogger()
logger.setLevel(logging.INFO)

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

def lambda_handler(event, context):
    """
    This function fetches content from MySQL RDS instance
    """
    videos = []
    with conn.cursor() as cursor:
        sql = 'SELECT to_base64(UuidFromBin(r.id)) AS restaurant_id,' + \
              '       v.name AS name,' + \
              '       v.url AS url' + \
              '  FROM videos AS v' + \
              '  JOIN restaurants AS r' + \
              '    ON v.restaurant_id = r.id' + \
              ' ORDER BY v.name'

        if 'restaurant_id' in event:
            restaurant_id = event['restaurant_id']
            sql += ' WHERE r.id = UuidToBin(\'{}\')'.format(restaurant_id)
        cursor.execute(sql)
        videos = cursor.fetchall()
    return {
        'statusCode': 200,
        'body': json.dumps(videos)
    }

# System Module
import json
import logging
import os
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
    photos = []
    with conn.cursor() as cursor:
        sql = 'SELECT to_base64(UuidFromBin(r.id)) AS restaurant_id,' + \
              '       CONCAT(p.name, \'.jpg\') AS image,' + \
              '       CONCAT(p.name, \'.webp\') AS image_webp,' + \
              '       CONCAT(p.name, \'_thumbnail.jpg\') AS thumbnail,' + \
              '       CONCAT(p.name, \'_thumbnail.webp\') AS thumbnail_webp' + \
              '  FROM photos AS p' + \
              '  JOIN restaurants AS r' + \
              '    ON p.restaurant_id = r.id' + \
              ' ORDER BY p.create_time DESC'

        if 'restaurant_id' in event:
            restaurant_id = event['restaurant_id']
            sql += ' WHERE r.id = UuidToBin(\'{}\')'.format(restaurant_id)

        cursor.execute(sql)
        photos = cursor.fetchall()

    return {
        'statusCode': 200,
        'body': json.dumps(photos)
    }

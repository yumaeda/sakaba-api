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
    shops = []
    with conn.cursor() as cursor:
        sql = """
SELECT to_base64(UuidFromBin(r.id)) AS id,
       r.url,
       r.image_name,
       r.name,
       r.genre,
       r.tel,
       r.business_day_info,
       r.address,
       r.latitude,
       r.longitude,
       r.area,
       r.comment,
       r.takeout_available,
       COUNT(p.restaurant_id) AS photo_count
  FROM restaurants AS r
  LEFT JOIN photos AS p
    ON r.id = p.restaurant_id
 WHERE is_closed = 0
   AND JSON_EXTRACT(r.business_day_info, CONCAT('$.', DAYOFWEEK(CURDATE()))) IS NOT NULL
 GROUP BY r.id
 ORDER BY photo_count DESC
"""
        cursor.execute(sql)
        shops = cursor.fetchall()

    return {
        'statusCode': 200,
        'body': json.dumps(shops)
    }

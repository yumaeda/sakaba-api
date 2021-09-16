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
        cursorclass=pymysql.cursors.DictCursor,
        autocommit=True
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
    shop_counts = []
    with conn.cursor() as cursor:
        sql = """
SELECT area,
       COUNT(area) AS count
  FROM restaurants
 WHERE is_closed = 0
   AND REPLACE(JSON_EXTRACT(r.business_day_info, CONCAT('$.', DAYOFWEEK(CURDATE()), ".Start")), '"', '') <= DATE_FORMAT(CONVERT_TZ(NOW(), '+00:00', '+09:00'), '%H%i')
   AND REPLACE(JSON_EXTRACT(r.business_day_info, CONCAT('$.', DAYOFWEEK(CURDATE()), ".End")), '"', '') >= DATE_FORMAT(CONVERT_TZ(NOW(), '+00:00', '+09:00'), '%H%i')
 GROUP BY area
 ORDER BY COUNT(area) DESC
"""
        cursor.execute(sql)
        shop_counts = cursor.fetchall()

    return {
        'statusCode': 200,
        'body': json.dumps(shop_counts)
    }

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
    shop_counts = []
    with conn.cursor() as cursor:
        sql = """
SELECT area,
       COUNT(area) AS count
  FROM restaurants
 WHERE is_closed = 0
   AND JSON_EXTRACT(business_day_info, CONCAT('$.', DAYOFWEEK(CURDATE()))) IS NOT NULL
 GROUP BY area
 ORDER BY COUNT(area) DESC
"""
        cursor.execute(sql)
        shop_counts = cursor.fetchall()

    return {
        'statusCode': 200,
        'body': json.dumps(shop_counts)
    }

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
    rankings = []
    with conn.cursor() as cursor:
        sql = """
SELECT dishes.name AS 'dish',
       rankings.rank AS 'rank',
       restaurants.name AS 'restaurant',
       to_base64(UuidFromBin(restaurants.id)) AS 'restaurant_id',
       photos.name AS 'photo',
       restaurants.url AS 'restaurant_url'
  FROM rankings
  JOIN dishes
    ON rankings.dish_id = dishes.id
  JOIN photos
    ON rankings.photo_id = photos.id
  JOIN restaurants
    ON photos.restaurant_id = restaurants.id
 ORDER BY dishes.name ASC, rankings.rank ASC
"""
        cursor.execute(sql)
        rankings = cursor.fetchall()

    return {
        'statusCode': 200,
        'body': json.dumps(rankings)
    }

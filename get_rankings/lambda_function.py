# System Module
import json
from common import conn, get_response, STATUS_CODE_OK

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

    return get_response(STATUS_CODE_OK, json.dumps(rankings))

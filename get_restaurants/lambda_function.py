import json
from common import conn, get_response, STATUS_CODE_OK

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
   AND REPLACE(JSON_EXTRACT(r.business_day_info, CONCAT('$.', DAYOFWEEK(CURDATE()), ".Start")), '"', '') <= DATE_FORMAT(CONVERT_TZ(NOW(), '+00:00', '+09:00'), '%H%i')
   AND REPLACE(JSON_EXTRACT(r.business_day_info, CONCAT('$.', DAYOFWEEK(CURDATE()), ".End")), '"', '') >= DATE_FORMAT(CONVERT_TZ(NOW(), '+00:00', '+09:00'), '%H%i')
 GROUP BY r.id
 ORDER BY photo_count DESC
"""
        cursor.execute(sql)
        shops = cursor.fetchall()

    return get_response(STATUS_CODE_OK, json.dumps(shops))

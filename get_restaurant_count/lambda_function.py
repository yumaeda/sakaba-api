import json
from common import conn, get_response, STATUS_CODE_OK

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
   AND REPLACE(JSON_EXTRACT(business_day_info, CONCAT('$.', DAYOFWEEK(CURDATE()), ".Start")), '"', '') <= DATE_FORMAT(CONVERT_TZ(NOW(), '+00:00', '+09:00'), '%H%i')
   AND REPLACE(JSON_EXTRACT(business_day_info, CONCAT('$.', DAYOFWEEK(CURDATE()), ".End")), '"', '') >= DATE_FORMAT(CONVERT_TZ(NOW(), '+00:00', '+09:00'), '%H%i')
 GROUP BY area
 ORDER BY COUNT(area) DESC
"""
        cursor.execute(sql)
        shop_counts = cursor.fetchall()

    return get_response(STATUS_CODE_OK, json.dumps(shop_counts))

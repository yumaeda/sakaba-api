# System Module
import json
from common import conn, get_response, STATUS_CODE_OK

SQL_STMT = """
SELECT to_base64(UuidFromBin(id)) AS id,
       to_base64(UuidFromBin(restaurant_id)) AS restaurant_id,
       name,
       name_jpn,
       category,
       sub_category,
       region,
       price,
       is_min_price
  FROM menus
 ORDER BY category ASC, price ASC 
"""
 
def lambda_handler(event, context):
    """
    This function fetches content from MySQL RDS instance
    """
    entities = []
    with conn.cursor() as cursor:
        cursor.execute(SQL_STMT)
        entities = cursor.fetchall()
    return get_response(STATUS_CODE_OK, json.dumps(entities))

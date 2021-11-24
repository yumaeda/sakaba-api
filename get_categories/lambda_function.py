# System Module
import json
from common import conn, get_response, STATUS_CODE_OK, STATUS_CODE_BAD_REQUEST

SQL_STMT = """
SELECT id,
       parent_id,
       name
  FROM categories
 WHERE restaurant_id = UuidToBin('{restaurant_id}')
 ORDER BY parent_id ASC, id ASC
"""
 
def lambda_handler(event, context):
    """
    This function fetches content from MySQL RDS instance
    """
    if 'restaurant_id' not in event:
        return get_response(STATUS_CODE_BAD_REQUEST, 'Required columns are not specified.')
    restaurant_id = event['restaurant_id']
    entities = []
    with conn.cursor() as cursor:
        cursor.execute(SQL_STMT.format(restaurant_id=restaurant_id))
        entities = cursor.fetchall()
    return get_response(STATUS_CODE_OK, json.dumps(entities))
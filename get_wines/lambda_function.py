# System Module
import json
from common import conn, get_response, STATUS_CODE_OK

SQL_STMT = """
SELECT cepage,
       color,
       comment,
       country,
       name,
       name_jpn,
       producer,
       producer_jpn,
       region,
       region_jpn,
       vintage
  FROM wines
 ORDER BY color, producer, name
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

# System Module
import json
from common import conn, get_response, STATUS_CODE_OK

def lambda_handler(event, context):
    """
    This function fetches content from MySQL RDS instance
    """
    videos = []
    with conn.cursor() as cursor:
        sql = 'SELECT to_base64(UuidFromBin(r.id)) AS restaurant_id,' + \
              '       v.name AS name,' + \
              '       v.url AS url' + \
              '  FROM videos AS v' + \
              '  JOIN restaurants AS r' + \
              '    ON v.restaurant_id = r.id' + \
              ' ORDER BY v.name'

        if 'restaurant_id' in event:
            restaurant_id = event['restaurant_id']
            sql += ' WHERE r.id = UuidToBin(\'{}\')'.format(restaurant_id)
        cursor.execute(sql)
        videos = cursor.fetchall()
    return get_response(STATUS_CODE_OK, json.dumps(videos))

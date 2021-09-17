import json
from common import conn, get_response, STATUS_CODE_OK

def lambda_handler(event, context):
    """
    This function fetches content from MySQL RDS instance
    """
    photos = []
    with conn.cursor() as cursor:
        sql = 'SELECT to_base64(UuidFromBin(r.id)) AS restaurant_id,' + \
              '       CONCAT(p.name, \'.jpg\') AS image,' + \
              '       CONCAT(p.name, \'.webp\') AS image_webp,' + \
              '       CONCAT(p.name, \'_thumbnail.jpg\') AS thumbnail,' + \
              '       CONCAT(p.name, \'_thumbnail.webp\') AS thumbnail_webp' + \
              '  FROM photos AS p' + \
              '  JOIN restaurants AS r' + \
              '    ON p.restaurant_id = r.id' + \
              ' ORDER BY p.create_time DESC'

        if 'restaurant_id' in event:
            restaurant_id = event['restaurant_id']
            sql += ' WHERE r.id = UuidToBin(\'{}\')'.format(restaurant_id)
        cursor.execute(sql)
        photos = cursor.fetchall()

    return get_response(STATUS_CODE_OK, json.dumps(photos))

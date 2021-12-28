from common import conn, logging, get_response, STATUS_CODE_BAD_REQUEST, STATUS_CODE_INTERNAL_SERVER_ERROR, STATUS_CODE_OK

SQL_STMT = """
INSERT INTO restaurants(id, url, name, genre, tel, address, area, business_day_info, latitude, longitude)
     VALUES (
        UuidToBin(UUID()),
        '{url}',
        '{name}',
        '{genre}',
        '{tel}',
        '{address}',
        '{area}',
        '{business_day_info}',
        '{latitude}',
        '{longitude}'
    )
"""

def lambda_handler(event, context):
    """
    This function fetches content from MySQL RDS instance
    """
    if 'url' not in event or \
       'name' not in event or \
       'genre' not in event or \
       'tel' not in event or \
       'address' not in event or \
       'area' not in event or \
       'business_day_info' not in event or \
       'latitude' not in event or \
       'longitude' not in event:
        return get_response(STATUS_CODE_BAD_REQUEST, 'Required columns are not specified.')
    with conn.cursor() as cursor:
        try:
            insert_sql = SQL_STMT.format(
                url=event['url'],
                name=event['name'],
                genre=event['genre'],
                tel=event['tel'],
                address=event['address'],
                area=event['area'],
                business_day_info=event['business_day_info'],
                latitude=event['latitude'],
                longitude=event['longitude']
            )
            cursor.execute(insert_sql)
            conn.commit()
            return get_response(STATUS_CODE_OK, 'New restaurant [{}] is inserted.'.format(event['name']))
        except Exception as ex:
            logging.exception(ex)
    return get_response(STATUS_CODE_INTERNAL_SERVER_ERROR, 'Restaurant creation failed')

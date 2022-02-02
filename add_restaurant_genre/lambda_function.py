from common import conn, logging, get_response, STATUS_CODE_BAD_REQUEST, STATUS_CODE_INTERNAL_SERVER_ERROR, STATUS_CODE_OK

SQL_STMT = """
INSERT INTO restaurant_genres(restaurant_id, genre_id)
     VALUES (
        UuidToBin('{restaurant_id}'),
        '{genre_id}'
    )
"""

def lambda_handler(event, context):
    """
    This function fetches content from MySQL RDS instance
    """
    if 'restaurant_id' not in event or \
       'genre_id' not in event:
        return get_response(STATUS_CODE_BAD_REQUEST, 'Required columns are not specified.')
    with conn.cursor() as cursor:
        try:
            insert_sql = SQL_STMT.format(
                restaurant_id=event['restaurant_id'],
                genre_id=event['genre_id']
            )
            cursor.execute(insert_sql)
            conn.commit()
            return get_response(STATUS_CODE_OK, 'New restaurant genre is inserted.')
        except Exception as ex:
            logging.exception(ex)
    return get_response(STATUS_CODE_INTERNAL_SERVER_ERROR, 'Restaurant genre insertion failed')

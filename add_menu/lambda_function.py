from common import conn, logging, get_response, STATUS_CODE_BAD_REQUEST, STATUS_CODE_INTERNAL_SERVER_ERROR, STATUS_CODE_OK

SQL_STMT = """
INSERT INTO menus(id, restaurant_id, category, sub_category, region, name, name_jpn, price, is_min_price)
     VALUES (
         UuidToBin('{id}'),
         UuidToBin('{restaurant_id}'),
         {category},
         {sub_category},
         {region},
         '{name}',
         '{name_jpn}',
         {price},
         {is_min_price}
    )
"""

def lambda_handler(event, context):
    """
    This function fetches content from MySQL RDS instance
    """
    if 'id' not in event or \
       'restaurant_id' not in event or \
       'category' not in event or \
       'sub_category' not in event or \
       'region' not in event or \
       'name' not in event or \
       'name_jpn' not in event or \
       'price' not in event or \
       'is_min_price' not in event:
        return get_response(STATUS_CODE_BAD_REQUEST, 'Required columns are not specified.')
    with conn.cursor() as cursor:
        try:
            insert_sql = SQL_STMT.format(
                id=event['id'],
                restaurant_id=event['restaurant_id'],
                category=event['category'],
                sub_category=event['sub_category'],
                region=event['region'],
                name=event['name'],
                name_jpn=event['name_jpn'],
                price=event['price'],
                is_min_price=event['is_min_price']
            )
            cursor.execute(insert_sql)
            conn.commit()
            return get_response(STATUS_CODE_OK, 'New menu [{}] is inserted.'.format(event['id']))
        except Exception as ex:
            logging.exception(ex)
    return get_response(STATUS_CODE_INTERNAL_SERVER_ERROR, 'Menu creation failed')

from common import conn, logging, get_response, STATUS_CODE_OK, STATUS_CODE_BAD_REQUEST, STATUS_CODE_INTERNAL_SERVER_ERROR
import re

ID_KEY = 'id'
COLUMN_KEY = 'column'
VALUE_KEY = 'value'
COLUMNS = [
    'sort_order',
    'name',
    'name_jpn',
    'category',
    'sub_category',
    'region',
    'price',
    'is_min_price'
]
 
def lambda_handler(event, context):
    """
    This function fetches content from MySQL RDS instance
    """
    if ID_KEY not in event or COLUMN_KEY not in event or VALUE_KEY not in event:
        return get_response(STATUS_CODE_BAD_REQUEST, 'Required columns are not specified.')
    column = event[COLUMN_KEY]
    if column not in COLUMNS:
        return get_response(STATUS_CODE_BAD_REQUEST, 'Specified column is not supported.')
    id = event[ID_KEY]
    value = event[VALUE_KEY]
    with conn.cursor() as cursor:
        try:
            update_sql = 'UPDATE menus SET {column} = \'{value}\' WHERE id = UuidToBin(\'{id}\')'.format(
                column=column,
                value=conn.escape_string(value),
                id=id
            )
            cursor.execute(update_sql)
            conn.commit()
            return get_response(STATUS_CODE_OK, 'Menu [{}] is updated.'.format(id))
        except Exception as ex:
            logging.exception(ex)
    return get_response(STATUS_CODE_INTERNAL_SERVER_ERROR, 'Update failed')

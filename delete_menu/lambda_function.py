from common import conn, logging, get_response, STATUS_CODE_BAD_REQUEST, STATUS_CODE_INTERNAL_SERVER_ERROR, STATUS_CODE_OK

SQL_STMT = "DELETE FROM menus WHERE id = UuidToBin('{id}')"

def lambda_handler(event, context):
    """
    This function fetches content from MySQL RDS instance
    """
    if 'id' not in event:
        return get_response(STATUS_CODE_BAD_REQUEST, 'Required columns are not specified.')
    with conn.cursor() as cursor:
        menu_id = event['id']
        try:
            delete_sql = SQL_STMT.format(id=menu_id)
            cursor.execute(delete_sql)
            conn.commit()
            return get_response(STATUS_CODE_OK, 'The menu [{}] is deleted.'.format(menu_id))
        except Exception as ex:
            logging.exception(ex)
    return get_response(STATUS_CODE_INTERNAL_SERVER_ERROR, 'Menu deletion failed')

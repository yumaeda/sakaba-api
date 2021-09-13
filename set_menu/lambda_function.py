# System Module
import logging
import os
import sys

# 3rd Party Module
import pymysql

logger = logging.getLogger()
logger.setLevel(logging.INFO)

try:
    conn = pymysql.connect(
        host=os.environ['DB_HOST'],
        user=os.environ['DB_USER'],
        passwd=os.environ['DB_PASSWORD'],
        db=os.environ['DB_NAME'],
        connect_timeout=10,
        charset='utf8mb4',
        cursorclass=pymysql.cursors.DictCursor
    )
except pymysql.MySQLError as e:
    logger.error('ERROR: Could not connect to MariaDB instance.')
    logger.error(e)
    sys.exit()

logger.info('SUCCESS: Connection to RDS MariaDB instance succeeded')

STATUS_CODE_OK = 200
STATUS_CODE_BAD_REQUEST = 400
STATUS_CODE_INTERNAL_SERVER_ERROR = 500
ID_KEY = 'id'
COLUMN_KEY = 'column'
VALUE_KEY = 'value'
COLUMNS = [
   'name',
    'name_jpn',
    'category',
    'sub_category',
    'region',
    'price',
    'is_min_price'
]
 
def get_response(status_code: int, body: str):
    return {
        'statusCode': status_code,
        'body': body
    }

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
                value=value,
                id=id
            )
            cursor.execute(update_sql)
            conn.commit()
            return get_response(STATUS_CODE_OK, 'Menu [{}] is updated.'.format(id))
        except Exception as ex:
            logging.exception(ex)
    return get_response(STATUS_CODE_INTERNAL_SERVER_ERROR, 'Update failed')

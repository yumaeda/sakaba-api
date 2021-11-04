from common import conn, logging, get_response, STATUS_CODE_OK, STATUS_CODE_BAD_REQUEST, STATUS_CODE_INTERNAL_SERVER_ERROR
import requests
import json

ID_KEY = 'id'
GEOLOCATION_API_URI = 'https://msearch.gsi.go.jp/address-search/AddressSearch?q={}'
session = requests.Session()

def lambda_handler(event, context):
    """
    This function fetches content from MySQL RDS instance
    """
    if ID_KEY not in event:
        return get_response(STATUS_CODE_BAD_REQUEST, 'Required columns are not specified.')
    id = event[ID_KEY]

    with conn.cursor() as cursor:
        try:
            get_sql = 'SELECT address FROM restaurants WHERE id = UuidToBin(\'{id}\')'
            cursor.execute(get_sql.format(id=id))
            restaurant = cursor.fetchone()
            # Need to configure NAT Gateway to allow external API call.
            # https://aws.amazon.com/premiumsupport/knowledge-center/internet-access-lambda-function/
            response = session.get(GEOLOCATION_API_URI.format(restaurant['address']))
            json_data = json.loads(response.text)
            coordinates = json_data[0]['geometry']['coordinates'].split(',')
            update_sql = 'UPDATE restaurants SET latitude = \'{latitude}\', longitude = \'{longitude}\' WHERE id = UuidToBin(\'{id}\')'.format(
                latitude=coordinates[1],
                longitude=coordinates[0],
                id=id
            )
            cursor.execute(update_sql)
            conn.commit()
            return get_response(STATUS_CODE_OK, 'Coordinates of the restaurant [{}] is updated.'.format(id))
        except Exception as ex:
            logging.exception(ex)
    return get_response(STATUS_CODE_INTERNAL_SERVER_ERROR, 'Update failed')

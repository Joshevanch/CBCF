import requests
import argparse

parser = argparse.ArgumentParser(description='The parameter for the CBS message')
parser.add_argument('-id', '--messageId', type=int, help='The message ID', required = True)
parser.add_argument('-rs', '--ratSelector', type=str, help='Enumeration to choose between E-UTRA (4G) and NG (5G), the default is 5G', choices=['E-UTRA', 'NR'],  default='NR')
parser.add_argument('-n2', '--n2InformationClass', type=str, help='Enumeration to choose between class of the N2 message, it can be PWS, SM, RAN, NRPPa, PWS-BCAL, the default is PWS', choices=['PWS', 'SM', 'RAN', 'NRPPa', 'PWS-BCAL'],  default='PWS')
parser.add_argument('-t', '--tac', type=str, help='Parameter 1', default = '')
parser.add_argument('-mn', '--mnc', type=int, help='Mobile Network Code')
parser.add_argument('-mc', '--mcc', type=int, help='Mobile Country Code')
parser.add_argument('-r', '--repetitionPeriod', type=int, help='Repetition Period', required = True)
parser.add_argument('-n', '--numberOfBroadcastsRequested', type=int, help='Number of Broadcast Requested', required = True)
parser.add_argument('-m', '--warningMessageContents', type=str, help='Warning Message Contents', required = True)

args = parser.parse_args()

# Prepare the parameters
data = {
    'id': args.messageId,
    'ratSelector': args.ratSelector,
    'n2Information': args.n2InformationClass,
    'tac': args.tac,
    'mnc': args.mnc,
    'mcc': args.mcc,
    'repetitionPeriod': args.repetitionPeriod,
    'numberOfBroadcastsRequested': args.numberOfBroadcastsRequested,
    'warningMessageContents': args.warningMessageContents
}

# Send the HTTP request
response = requests.post('http://127.0.0.1:8080/', data=data)

# Check the response status code
if response.status_code == requests.codes.ok:
    print('Request was successful.')
    print('Response:', response.text)
else:
    print('Request failed with status code:', response.status_code)

import requests

def send_xml_data(url, xml_data):
    headers = {
        "Content-Type": "application/xml ;charset=utf-8",
    }
    
    try:
        response = requests.post(url, data=xml_data, headers=headers)
        response.raise_for_status()
        return response.content
    except requests.exceptions.RequestException as e:
        print(f"Error: {e}")
        return None

# Example XML data
xml_data = """
<?xml version="1.0" encoding="UTF-8"?> 
<alert xmlns="urn:oasis:names:tc:emergency:cap:1.1">
    <identifier>CWB-EQ112214</identifier> 
    <sender>cwb@scman.cwb.gov.tw</sender> 
    <sent>2023-07-27T00:08:00+08:00</sent> 
    <status>Actual</status> 
    <msgType>Alert</msgType>
    <source>CWB</source>
    <scope>Public</scope> 
    <info> 
        <language>zh-TW</language>
        <category>Met</category>
        <event>地震</event>
        <responseType>Shelter</responseType>
        <urgency>Immediate</urgency>
        <severity>Severe</severity>
        <certainty>Observed</certainty>
        <effective>2023-07-27T08:00:00+08:00</effective>
        <expires>2023-07-27T08:08:00+08:00</expires> 
        <senderName>中央氣象局</senderName> 
        <headline>地震報告</headline> 
        <description> 07/27-18:11 花蓮縣秀林鄉發生規模 5.3 有感地震，最大震度花蓮縣太魯閣、宜蘭縣南山、南投縣合歡山、臺中市德基 4 級。</description>
        <contact>123456</contact>  
        <area> 
            <areaDesc>最大震度 3 級地區</areaDesc> 
            <geoCode>10002</geoCode> 
        </area> 
        </info> 
    </alert>
"""

# Example URL to which you want to send the XML data
url = 'http://127.0.0.1:8080'
encoded_xml_data = xml_data.encode('utf-8')

# Sending the XML data and getting the response
response_content = send_xml_data(url, encoded_xml_data)
if response_content:
    print("Response:")
    print(response_content)
else:
    print("Failed to send XML data.")

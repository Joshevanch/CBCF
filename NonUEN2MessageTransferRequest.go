package main

import (
	"bytes"
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"github.com/free5gc/openapi/models"
)

func transfer(m map[string]string) {
	// Specify the URL you want to send the request to
	url := "http://127.0.0.18:8000/namf-comm/v1/non-ue-n2-messages/transfer/"
	// Create the request body
	message := models.NonUeN2MessageTransferRequest{}
	jsonString := []byte(`{
		"jsonData": {
		  "taiList": [
			{
			  "tac": "",
			  "plmnId": {
				"mnc": "",
				"mcc": ""
			  }
			}
		  ],
		  "ratSelector": "PWS",
		  "ecgiList": [
			{
			  "eutraCellId": "",
			  "plmnId": {
				"mnc": "",
				"mcc": ""
			  }
			}
		  ],
		  "ncgiList": [
			{
			  "nrCellId": "",
			  "plmnId": {
				"mnc": "",
				"mcc": ""
			  }
			}
		  ],
		  "globalRanNodeList": [
			{
			  "gNbId": {
				"bitLength": 24,
				"gNBValue": ""
			  },
			  "plmnId": {
				"mnc": "",
				"mcc": ""
			  },
			  "n3IwfId": "",
			  "ngeNbId": ""
			}
		  ],
		  "n2Information": {
			"n2InformationClass": "",
			"smInfo": {
			  "subjectToHo": false,
			  "pduSessionId": 29,
			  "n2InfoContent": {
				"ngapData": {
				  "contentId": "contentId"
				},
				"ngapIeType": "",
				"ngapMessageType": 32
			  },
			  "sNssai": {
				"sd": "sd",
				"sst": 32
			  }
			},
			"ranInfo": {
			  "n2InfoContent": {
				"ngapData": {
				  "contentId": "contentId"
				},
				"ngapIeType": "",
				"ngapMessageType": 32
			  }
			},
			"nrppaInfo": {
			  "nfId": "nfId",
			  "nrppaPdu": {
				"n2InfoContent": {
				  "ngapData": {
					"contentId": "contentId"
				  },
				  "ngapIeType": "",
				  "ngapMessageType": 32
				}
			  }
			},
			"pwsInfo": {
			  "messageIdentifier": 0,
			  "serialNumber": 0,
			  "pwsContainer": {
				"n2InfoContent": {
				  "ngapData": {
					"contentId": "contentId"
				  },
				  "ngapIeType": "",
				  "ngapMessageType": 51
				},
				"sendRanResponse": true,
				"omcId": true
			  }
			}
		  },
		  "supportedFeatures": ""
		}
	  }`)
	json.Unmarshal(jsonString, &message)
	if (m["ratSelector"] == "NR"){
		message.JsonData.RatSelector = models.RatSelector_NR
	}
	if (m["ratSelector"] == "E-UTRA"){
		message.JsonData.RatSelector = models.RatSelector_E_UTRA
	}
	(*message.JsonData.TaiList)[0].PlmnId.Mcc = m["mcc"]
	(*message.JsonData.TaiList)[0].PlmnId.Mnc = m["mnc"]
	(*message.JsonData.TaiList)[0].Tac = m["tac"]
	fmt.Printf("%+v", (*message.JsonData.TaiList)[0].PlmnId)
	jsonString, err := json.Marshal(message)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonString))
    req.Header.Set("X-Custom-Header", "myvalue")
    req.Header.Set("Content-Type", "application/json")
    client := &http.Client{}
    response, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error sending request: %s\n", err)
		return
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("Error reading response: %s\n", err)
		return
	}
	fmt.Println(string(body))
}

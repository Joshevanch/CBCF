package main

import (
	"bytes"
	"flag"
	"fmt"
	"net/http"
	"io/ioutil"
)

func main() {
	// Specify the URL you want to send the request to
	url := "http://127.0.0.18:8000/namf-comm/v1/non-ue-n2-messages/transfer/"
	ratSelector := flag.String("RAT", "", "enumeration to choose between E-UTRA (4G) and NG (5G)")
	tac := flag.String("TAC", "", "Tracking area code")
	mnc := flag.String("MNC", "", "Mobile Network Code")
	mcc := flag.String("MCC", "", "Mobile Country Code")
	flag.Parse()

	// Create the request body
	jsonString := []byte(`{
		"taiList": [
		  {
			"tac": "` + *tac + `",
			"plmnId": {
			  "mnc": "` + *mnc + `",
			  "mcc": "` + *mcc + `"
			}
		  }
		],
		"ratSelector": "` + *ratSelector + `",
		"ecgiList": [
		  {
			"eutraCellId": "",
			"plmnId": {
			  "mnc": "` + *mnc + `",
			  "mcc": "` + *mcc + `"
			}
		  }
		],
		"ncgiList": [
		  {
			"nrCellId": "",
			"plmnId": {
			  "mnc": "` + *mnc + `",
			  "mcc": "` + *mcc + `"
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
				"mnc": "` + *mnc + `",
				"mcc": "` + *mcc + `"
			},
			"n3IwfId": "",
			"ngeNbId": ""
		  }
		],
		"n2Information": {
		  "n2InformationClass": "PWS",
		  "smInfo": {
			"subjectToHo": true,
			"pduSessionId": 29,
			"n2InfoContent": {
			  "ngapData": {
				"contentId": "contentId"
			  },
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
			  "ngapMessageType": 32
			}
		  },
		  "nrppaInfo": {
			"nfId": "nfId",
			"nrppaPdu": {
			  "n2InfoContent": {
				"ngapData": {
				  "contentId": "contentId"
				}
			  }
			}
		  },
		  "pwsInfo": {
			"messageIdentifier": 600,
			"serialNumber": 32,
			"pwsContainer": {
			  "n2InfoContent": {
				"ngapData": {
				  "contentId": "contentId"
				},
				"sendRanResponse": true,
				"omcId": true
			  }
			}
		  }
		},
		"supportedFeatures": ""
	  }`)

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

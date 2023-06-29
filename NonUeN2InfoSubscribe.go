package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/free5gc/openapi/Namf_Communication"
	"github.com/free5gc/openapi/models"
)

func subscribe() {
	subscribe := models.NonUeN2InfoSubscriptionCreateData{}
	// Specify the URL you want to send the request to
	// Create the request body
	jsonString := []byte(`{
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
		"anTypeList":[

		],
		"n2InformationClass": "PWS",
		"n2NotifyCallbackUri": "127.0.0.1:8080/notify",
		"nfId": "",
		"supportedFeatures": ""
	  }`)
	  json.Unmarshal(jsonString, &subscribe)
	  namfConfiguration := Namf_Communication.NewConfiguration()
	  namfConfiguration.SetBasePath("http://127.0.0.18:8000")
	  fmt.Println(namfConfiguration.BasePath())
	  apiClient := Namf_Communication.NewAPIClient(namfConfiguration)
	  rep, res, err := apiClient.NonUEN2MessagesSubscriptionsCollectionDocumentApi.NonUeN2InfoSubscribe(context.TODO(), subscribe)
	  body, err := ioutil.ReadAll(res.Body)
	  fmt.Println(rep)
	if err != nil {
		fmt.Printf("Error reading response: %s\n", err)
		return
	}
	fmt.Println(string(body))
}

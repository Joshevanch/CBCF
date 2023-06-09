package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"github.com/free5gc/openapi/Namf_Communication"
	"github.com/free5gc/openapi/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func transfer(m map[string]string) {
	// Specify the URL you want to send the request to

	// Create the request body
	test := models.N2InformationTransferReqData{}
	message := models.NonUeN2MessageTransferRequest{}
	jsonString := []byte(`{
		"taiList": [
		  {
			"tac": "sdf",
			"plmnId": {
			  "mnc": "sdf",
			  "mcc": "sdf"
			}
		  }
		],
		"ratSelector": "PWS",
		"ecgiList": [
		  {
			"eutraCellId": "sdf",
			"plmnId": {
			  "mnc": "sdf",
			  "mcc": "sdf"
			}
		  }
		],
		"ncgiList": [
		  {
			"nrCellId": "sdf",
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
		  "n2InformationClass": "PWS",
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
			  "ngapData": {
				"contentId": "n2msg"
			  },
			  "ngapIeType": "",
			  "ngapMessageType": 51
			},
			"sendRanResponse": true,
			"omcId": true
		  }
		},
		"supportedFeatures": ""
	  }
	  `)
	  BinaryDataN2Information := []byte(`
	  {
		"messageType": "",
		"messageIdentifier": "",
		"serialNumber": "",
		"warningAreaList": "",
		"repetitionPeriod": "",
		"numberOfBroadcast": "",
		"warningType": "",
		"warningSecurityInformation": "",
		"dataCodingScheme": "",
		"warningMessageContents" : "",
		"concurrentWarningMessageIndicator": "",
		"warningAreaCoordinates": ""
	}`)
	json.Unmarshal(jsonString, &test)
	fmt.Println(&test)
	message.JsonData = &test
	if m["ratSelector"] == "NR" {
		message.JsonData.RatSelector = models.RatSelector_NR
	}
	if m["ratSelector"] == "E-UTRA" {
		message.JsonData.RatSelector = models.RatSelector_E_UTRA
	}
	id, err := strconv.ParseInt(m["id"], 10, 32)
	(*&message.JsonData.N2Information.PwsInfo.MessageIdentifier) = int32(id)
	(*message.JsonData.TaiList)[0].PlmnId.Mcc = m["mcc"]
	(*message.JsonData.TaiList)[0].PlmnId.Mnc = m["mnc"]
	(*message.JsonData.TaiList)[0].Tac = m["tac"]
	(*message.JsonData.EcgiList)[0].PlmnId.Mcc = m["mcc"]
	(*message.JsonData.EcgiList)[0].PlmnId.Mnc = m["mnc"]
	(*message.JsonData.NcgiList)[0].PlmnId.Mcc = m["mcc"]
	(*message.JsonData.NcgiList)[0].PlmnId.Mnc = m["mnc"]
	(*message.JsonData.GlobalRanNodeList)[0].PlmnId.Mcc = m["mcc"]
	(*message.JsonData.GlobalRanNodeList)[0].PlmnId.Mnc = m["mnc"]
	(*&message.BinaryDataN2Information) = BinaryDataN2Information
	fmt.Println(BinaryDataN2Information)
	
	jsonString, err = json.Marshal(message)
	namfConfiguration := Namf_Communication.NewConfiguration()
	namfConfiguration.SetBasePath("http://127.0.0.18:8000")
	fmt.Println(namfConfiguration.BasePath())
	apiClient := Namf_Communication.NewAPIClient(namfConfiguration)
	rep, res, err := apiClient.NonUEN2MessagesCollectionDocumentApi.NonUeN2MessageTransfer(context.TODO(), message)
	insertToDatabase(message)
	fmt.Println(rep)
	fmt.Println(res)
	if err != nil {
		log.Fatal(err)
	}
}

func insertToDatabase(message models.NonUeN2MessageTransferRequest) {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
	collection := client.Database("local").Collection("cbcf")
	insertResult, err := collection.InsertOne(context.TODO(), message)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Inserted document ID: %v\n", insertResult.InsertedID)
	sort := options.FindOne().SetSort(bson.D{{"_id", -1}})
	var result models.NonUeN2MessageTransferRequest
	err = collection.FindOne(context.TODO(), bson.D{}, sort).Decode(&result)
	var b []byte
	b, err = json.Marshal(result)
	fmt.Println(string(b))

}

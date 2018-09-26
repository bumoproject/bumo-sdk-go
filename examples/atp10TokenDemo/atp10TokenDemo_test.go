// atp10TokenDemo_test
package atp10TokenDemo_test

import (
	"encoding/json"
	"testing"

	"github.com/bumoproject/bumo-sdk-go/src/model"
	"github.com/bumoproject/bumo-sdk-go/src/sdk"
)

var testSdk sdk.Sdk

type Atp10Metadata struct {
	Code        string `json:"code"`
	Name        string `json:"name"`
	TotalSupply int64  `json:"total_supply"`
	Decimals    int64  `json:"decimals"`
	Description string `json:"description"`
	Version     string `json:"version"`
	Icon        string `json:"icon "`
}

//to initialize the SDK
func Test_Init(t *testing.T) {
	var reqData model.SDKInitRequest
	reqData.SetUrl("http://seed1.bumotest.io:26002")
	resData := testSdk.Init(reqData)
	if resData.ErrorCode != 0 {
		t.Errorf(resData.ErrorDesc)
	} else {
		t.Log("Test_NewSDK")
	}
}

/**
 * Issue the unlimited apt1.0 token successfully
 * Unlimited requirement: The totalSupply must be smaller than and equal to 0
 */
func Test_issueUnlimitedAtp10Token(t *testing.T) {
	// The account private key to issue atp1.0 token
	var issuerPrivateKey string = "privbtYzJ6miiFktK9BsDAMRNd3J4eKkuszfXqJ2huQ2h8DGUnRs9nuq"
	// The account address to send this transaction
	var issuerAddresss string = "buQVU86Jm4FeRW4JcQTD9Rx9NkUkHikYGp6z"

	// The apt token version
	var version string = "1.0"
	// The token code
	var code string = "TXT"
	// The token name
	var name string = "TXT"
	// The apt token icon
	var icon string = ""
	// The token total supply number
	var totalSupply int64 = 0
	// The token now supply number
	var nowSupply int64 = 1000000000
	// The token decimals
	var decimals int64 = 0

	// The token description
	var description string = "test unlimited issuance of apt1.0 token"
	// The operation notes
	var metadata string = "test the unlimited issuance of apt1.0 token"
	// The fixed write 1000L, the unit is MO
	var gasPrice int64 = 1000
	// Set up the maximum cost 0.01BU
	var feeLimit int64 = 5003000000
	// Transaction initiation account's Nonce + 1
	var nonce int64 = 10

	// If this is a atp 1.0 token, you must set transaction metadata like this
	var atp10Metadata Atp10Metadata
	atp10Metadata.Version = version
	atp10Metadata.Code = code
	atp10Metadata.Name = name
	atp10Metadata.Decimals = decimals
	atp10Metadata.TotalSupply = totalSupply
	atp10Metadata.Description = description
	atp10Metadata.Icon = icon
	atp10MetadataJson, err := json.Marshal(atp10Metadata)
	if err != nil {
		t.Errorf(err.Error())
	}
	// Build asset operation
	var reqDataIssue model.AssetIssueOperation
	reqDataIssue.Init()
	reqDataIssue.SetAmount(nowSupply)
	reqDataIssue.SetCode(code)
	reqDataIssue.SetSourceAddress(issuerAddresss)
	reqDataIssue.SetMetadata(metadata)

	var key string = "asset_property_" + code
	var value = string(atp10MetadataJson)
	var reqDataSetMetadata model.AccountSetMetadataOperation
	reqDataSetMetadata.Init()
	reqDataSetMetadata.SetSourceAddress(issuerAddresss)
	reqDataSetMetadata.SetKey(key)
	reqDataSetMetadata.SetValue(value)
	reqDataSetMetadata.SetMetadata(metadata)

	var operations []model.BaseOperation
	operations = append(operations, reqDataIssue)
	operations = append(operations, reqDataSetMetadata)

	// Record txhash for subsequent confirmation of the real result of the transaction.
	// After recommending five blocks, call again through txhash Get the transaction information
	// from the transaction Hash to confirm the final result of the transaction
	errorCode, errorDesc, hash := atp10BlobSubmit(testSdk, operations, issuerPrivateKey, issuerAddresss, nonce, gasPrice, feeLimit, string(atp10MetadataJson), metadata)
	if errorCode != 0 {
		t.Errorf(errorDesc)
	} else {
		t.Log("hash succeed", hash)
	}
}

/**
 * Issue the limited apt1.0 token successfully
 * Limited requirement: The totalSupply must be bigger than 0
 */
func Test_issuelimitedAtp10Token(t *testing.T) {
	// The account private key to issue atp1.0 token
	var issuerPrivateKey string = "privbtYzJ6miiFktK9BsDAMRNd3J4eKkuszfXqJ2huQ2h8DGUnRs9nuq"
	// The account address to send this transaction
	var issuerAddresss string = "buQVU86Jm4FeRW4JcQTD9Rx9NkUkHikYGp6z"

	// The apt token version
	var version string = "1.0"
	// The token code
	var code string = "TXT"
	// The token name
	var name string = "TXT"
	// The apt token icon
	var icon string = ""
	// The token total supply number
	var totalSupply int64 = 0
	// The token now supply number
	var nowSupply int64 = 1000000000
	// The token decimals
	var decimals int64 = 0
	// The token description
	var description string = "test unlimited issuance of apt1.0 token"
	// The operation notes
	var metadata string = "test the unlimited issuance of apt1.0 token"
	// The fixed write 1000L, the unit is MO
	var gasPrice int64 = 1000
	// Set up the maximum cost 0.01BU
	var feeLimit int64 = 5003000000
	// Transaction initiation account's Nonce + 1
	var nonce int64 = 10

	// If this is a atp 1.0 token, you must set transaction metadata like this
	var atp10Metadata Atp10Metadata
	atp10Metadata.Version = version
	atp10Metadata.Code = code
	atp10Metadata.Name = name
	atp10Metadata.Decimals = decimals
	atp10Metadata.TotalSupply = totalSupply
	atp10Metadata.Description = description
	atp10Metadata.Icon = icon
	atp10MetadataJson, err := json.Marshal(atp10Metadata)
	if err != nil {
		t.Errorf(err.Error())
	}
	// Build asset operation
	var reqDataIssue model.AssetIssueOperation
	reqDataIssue.Init()
	reqDataIssue.SetAmount(nowSupply)
	reqDataIssue.SetCode(code)
	reqDataIssue.SetSourceAddress(issuerAddresss)
	reqDataIssue.SetMetadata(metadata)

	var key string = "asset_property_" + code
	var value = string(atp10MetadataJson)
	var reqDataSetMetadata model.AccountSetMetadataOperation
	reqDataSetMetadata.Init()
	reqDataSetMetadata.SetSourceAddress(issuerAddresss)
	reqDataSetMetadata.SetKey(key)
	reqDataSetMetadata.SetValue(value)
	reqDataSetMetadata.SetMetadata(metadata)

	var operations []model.BaseOperation
	operations = append(operations, reqDataIssue)
	operations = append(operations, reqDataSetMetadata)

	// Record txhash for subsequent confirmation of the real result of the transaction.
	// After recommending five blocks, call again through txhash Get the transaction information
	// from the transaction Hash to confirm the final result of the transaction
	errorCode, errorDesc, hash := atp10BlobSubmit(testSdk, operations, issuerPrivateKey, issuerAddresss, nonce, gasPrice, feeLimit, string(atp10MetadataJson), metadata)
	if errorCode != 0 {
		t.Errorf(errorDesc)
	} else {
		t.Log("hash succeed", hash)
	}
}

/**
 * Send apt 1.0 token to other account
 */
func Test_sendAtp10Token(t *testing.T) {
	// The account private key to send atp1.0 token
	var senderPrivateKey string = "privbvCDPhjNmXdZD2p6RWfXhTC3qzpn8REtZtPSu64mMQDMxAJ3f1hu"
	// The account that issued the atp 1.0 token
	var issuerAddresss string = "buQVU86Jm4FeRW4JcQTD9Rx9NkUkHikYGp6z"
	// The account address to send this transaction
	var senderAddresss string = "buQtjhgK9SakQPYGzoZ3iHodfRvd8qTGoaYd"
	// The account to receive atp 1.0 token
	var destAddress string = "buQc77ZYKT2dYZ5pzdsfGdGjGMJGGR9ZVZ1p"
	// The token code
	var code string = "TXT"
	// The token amount to be sent
	var amount int64 = 100000
	// The operation notes
	var metadata string = "test one off issue apt1.0 token"
	// The fixed write 1000L, the unit is MO
	var gasPrice int64 = 1000
	// Set up the maximum cost 0.01BU
	var feeLimit int64 = 1000000
	// Transaction initiation account's Nonce + 1
	var nonce int64 = 10

	var operations []model.BaseOperation

	//  Check whether the destination account is activated
	var reqDataCheckActivated model.AccountCheckActivatedRequest
	reqDataCheckActivated.SetAddress(destAddress)
	if testSdk.Account.CheckActivated(reqDataCheckActivated).Result.IsActivated {
		t.Errorf("destAddress not activated")
		return
	}
	// Build asset operation
	var reqDataIssue model.AssetSendOperation
	reqDataIssue.Init()
	reqDataIssue.SetDestAddress(destAddress)
	reqDataIssue.SetAmount(amount)
	reqDataIssue.SetCode(code)
	reqDataIssue.SetIssuer(issuerAddresss)
	reqDataIssue.SetSourceAddress(senderAddresss)
	reqDataIssue.SetMetadata(metadata)

	operations = append(operations, reqDataIssue)

	// Record txhash for subsequent confirmation of the real result of the transaction.
	// After recommending five blocks, call again through txhash Get the transaction information
	// from the transaction Hash to confirm the final result of the transaction
	errorCode, errorDesc, hash := atp10BlobSubmit(testSdk, operations, senderPrivateKey, senderAddresss, nonce, gasPrice, feeLimit, "", metadata)
	if errorCode != 0 {
		t.Errorf(errorDesc)
	} else {
		t.Log("hash succeed", hash)
	}
}

func atp10BlobSubmit(testSdk sdk.Sdk, operations []model.BaseOperation, senderPrivateKey string, senderAddresss string, senderNonce int64, gasPrice int64, feeLimit int64, transMetadata string, metadata string) (errorCode int, errorDesc string, hash string) {
	//Blob
	var reqDataBlob model.TransactionBuildBlobRequest
	reqDataBlob.SetSourceAddress(senderAddresss)
	reqDataBlob.SetFeeLimit(feeLimit)
	reqDataBlob.SetGasPrice(gasPrice)
	reqDataBlob.SetNonce(senderNonce)
	reqDataBlob.SetMetadata(transMetadata)
	for i := range operations {
		reqDataBlob.AddOperation(operations[i])
	}

	resDataBlob := testSdk.Transaction.BuildBlob(reqDataBlob)
	if resDataBlob.ErrorCode != 0 {
		return resDataBlob.ErrorCode, resDataBlob.ErrorDesc, ""
	} else {
		//Sign
		PrivateKey := []string{senderPrivateKey}
		var reqData model.TransactionSignRequest
		reqData.SetBlob(resDataBlob.Result.Blob)
		reqData.SetPrivateKeys(PrivateKey)
		resDataSign := testSdk.Transaction.Sign(reqData)
		if resDataSign.ErrorCode != 0 {
			return resDataSign.ErrorCode, resDataSign.ErrorDesc, ""
		} else {
			//Submit
			var reqData model.TransactionSubmitRequest
			reqData.SetBlob(resDataBlob.Result.Blob)
			reqData.SetSignatures(resDataSign.Result.Signatures)
			resDataSubmit := testSdk.Transaction.Submit(reqData)
			if resDataSubmit.ErrorCode != 0 {
				return resDataSubmit.ErrorCode, resDataSubmit.ErrorDesc, ""
			} else {
				return resDataSubmit.ErrorCode, resDataSubmit.ErrorDesc, resDataBlob.Result.Blob
			}
		}
	}
}

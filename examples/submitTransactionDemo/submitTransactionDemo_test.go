// submitTransactionDemo
package submitTransactionDemo_test

import (
	"testing"

	"github.com/bumoproject/bumo-sdk-go/src/model"
	"github.com/bumoproject/bumo-sdk-go/src/sdk"
)

//take send BU, for example
func Test_submitTransactionDemo_BUSend(t *testing.T) {
	// The token amount to be sent
	var amount int64 = 100000
	// The account to receive
	var destAddress string = "buQVU86Jm4FeRW4JcQTD9Rx9NkUkHikYGp6z"
	var url string = "http://seed1.bumotest.io:26002"
	// The account that BU
	var sourceAddress string = "buQemmMwmRQY1JkcU7w3nhruoX5N3j6C29uo"
	// The fixed write 1000L, the unit is MO
	var gasPrice int64 = 1000
	// Set up the maximum cost 0.01BU
	var feeLimit int64 = 5003000000
	//Building SDK objects
	var testSdk sdk.Sdk
	var reqDataInit model.SDKInitRequest
	reqDataInit.SetUrl(url)
	resDataInit := testSdk.Init(reqDataInit)
	if resDataInit.ErrorCode != 0 {
		t.Errorf(resDataInit.ErrorDesc)
	}
	//Gets the latest Nonce
	var reqDataNonce model.AccountGetNonceRequest
	reqDataNonce.SetAddress(sourceAddress)
	resDataNonce := testSdk.Account.GetNonce(reqDataNonce)
	if resDataNonce.ErrorCode != 0 {
		t.Errorf(resDataNonce.ErrorDesc)
	}
	//Building Operation
	var reqDataOperation model.BUSendOperation
	reqDataOperation.Init()

	reqDataOperation.SetAmount(amount)
	reqDataOperation.SetDestAddress(destAddress)
	//Building Blob
	var reqDataBlob model.TransactionBuildBlobRequest
	reqDataBlob.SetSourceAddress(sourceAddress)
	reqDataBlob.SetFeeLimit(feeLimit)
	reqDataBlob.SetGasPrice(gasPrice)
	var nonce int64 = resDataNonce.Result.Nonce + 1
	reqDataBlob.SetNonce(nonce)
	reqDataBlob.SetOperation(reqDataOperation)
	resDataBlob := testSdk.Transaction.BuildBlob(reqDataBlob)
	if resDataBlob.ErrorCode != 0 {
		t.Errorf(resDataBlob.ErrorDesc)
	} else {
		//Sign
		PrivateKey := []string{"privbUPxs6QGkJaNdgWS2hisny6ytx1g833cD7V9C3YET9mJ25wdcq6h"}
		var reqData model.TransactionSignRequest
		reqData.SetBlob(resDataBlob.Result.Blob)
		reqData.SetPrivateKeys(PrivateKey)
		resDataSign := testSdk.Transaction.Sign(reqData)
		if resDataSign.ErrorCode != 0 {
			t.Errorf(resDataSign.ErrorDesc)
		} else {
			//Submit
			var reqData model.TransactionSubmitRequest
			reqData.SetBlob(resDataBlob.Result.Blob)
			reqData.SetSignatures(resDataSign.Result.Signatures)
			resDataSubmit := testSdk.Transaction.Submit(reqData)
			if resDataSubmit.ErrorCode != 0 {
				t.Errorf(resDataSubmit.ErrorDesc)
			} else {
				t.Log("Test_Transaction_BuildBlob_Sign_Submit succeed, Hash:", resDataSubmit.Result.Hash)
			}
		}
	}
}

//take Log create, for example
func Test_submitTransactionDemo_LogCreate(t *testing.T) {
	// The token amount to be sent
	var topic string = "log"
	// The account to receive
	var datas []string
	datas = append(datas, "txt", "doc")
	var url string = "http://seed1.bumotest.io:26002"
	// The account that BU
	var sourceAddress string = "buQemmMwmRQY1JkcU7w3nhruoX5N3j6C29uo"
	// The fixed write 1000L, the unit is MO
	var gasPrice int64 = 1000
	// Set up the maximum cost 0.01BU
	var feeLimit int64 = 103000000
	//Building SDK objects
	var testSdk sdk.Sdk
	var reqDataInit model.SDKInitRequest
	reqDataInit.SetUrl(url)
	resDataInit := testSdk.Init(reqDataInit)
	if resDataInit.ErrorCode != 0 {
		t.Errorf(resDataInit.ErrorDesc)
	}
	//Gets the latest Nonce
	var reqDataNonce model.AccountGetNonceRequest
	reqDataNonce.SetAddress(sourceAddress)
	resDataNonce := testSdk.Account.GetNonce(reqDataNonce)
	if resDataNonce.ErrorCode != 0 {
		t.Errorf(resDataNonce.ErrorDesc)
	}
	//Building Operation
	var reqDataOperation model.LogCreateOperation
	reqDataOperation.Init()
	reqDataOperation.SetDatas(datas)
	reqDataOperation.SetTopic(topic)
	//Building Blob
	var reqDataBlob model.TransactionBuildBlobRequest
	reqDataBlob.SetSourceAddress(sourceAddress)
	reqDataBlob.SetFeeLimit(feeLimit)
	reqDataBlob.SetGasPrice(gasPrice)
	var nonce int64 = resDataNonce.Result.Nonce + 1
	reqDataBlob.SetNonce(nonce)
	reqDataBlob.SetOperation(reqDataOperation)
	resDataBlob := testSdk.Transaction.BuildBlob(reqDataBlob)
	if resDataBlob.ErrorCode != 0 {
		t.Errorf(resDataBlob.ErrorDesc)
	} else {
		//Sign
		PrivateKey := []string{"privbUPxs6QGkJaNdgWS2hisny6ytx1g833cD7V9C3YET9mJ25wdcq6h"}
		var reqData model.TransactionSignRequest
		reqData.SetBlob(resDataBlob.Result.Blob)
		reqData.SetPrivateKeys(PrivateKey)
		resDataSign := testSdk.Transaction.Sign(reqData)
		if resDataSign.ErrorCode != 0 {
			t.Errorf(resDataSign.ErrorDesc)
		} else {
			//Submit
			var reqData model.TransactionSubmitRequest
			reqData.SetBlob(resDataBlob.Result.Blob)
			reqData.SetSignatures(resDataSign.Result.Signatures)
			resDataSubmit := testSdk.Transaction.Submit(reqData)
			if resDataSubmit.ErrorCode != 0 {
				t.Errorf(resDataSubmit.ErrorDesc)
			} else {
				t.Log("Test_Transaction_BuildBlob_Sign_Submit succeed, Hash:", resDataSubmit.Result.Hash)
			}
		}
	}
}

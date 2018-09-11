// exchangeDemo_test
package exchangeDemo_test

import (
	"encoding/json"
	"testing"

	"github.com/bumoproject/bumo-sdk-go/src/model"
	"github.com/bumoproject/bumo-sdk-go/src/sdk"
)

var testSdk sdk.Sdk

//to initialize the sdk
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

//check that the blocks are synchronized
func Test_Block_CheckStatus(t *testing.T) {
	resData := testSdk.Block.CheckStatus()
	if resData.ErrorCode != 0 {
		t.Errorf(resData.ErrorDesc)
	} else {
		t.Log("IsSynchronous:", resData.Result.IsSynchronous)
		t.Log("Test_Block_CheckStatus succeed", resData.Result)
	}

}

//create account
func Test_Account_Create(t *testing.T) {
	resData := testSdk.Account.Create()
	if resData.ErrorCode != 0 {
		t.Errorf(resData.ErrorDesc)
	} else {
		t.Log("Test_Account_Create", resData.Result)
	}
}

//verify account address
func Test_Account_checkValid(t *testing.T) {
	var reqData model.AccountCheckValidRequest
	var address string = "buQemmMwmRQY1JkcU7w3nhruoX5N3j6C29uo"
	reqData.SetAddress(address)
	resData := testSdk.Account.CheckValid(reqData)

	if resData.Result.IsValid {
		t.Log("Test_Account_CheckAddress succeed", resData.Result.IsValid)
	} else {
		t.Error("Test_Account_CheckAddress failured")
	}
}

//enquiry of account details
func Test_Account_GetInfo(t *testing.T) {
	var reqData model.AccountGetInfoRequest
	var address string = "buQemmMwmRQY1JkcU7w3nhruoX5N3j6C29uo"
	reqData.SetAddress(address)
	resData := testSdk.Account.GetInfo(reqData)
	if resData.ErrorCode != 0 {
		t.Errorf(resData.ErrorDesc)
	} else {
		data, _ := json.Marshal(resData.Result)
		t.Log("info:", string(data))
		t.Log("Test_Account_GetInfo succeed", resData.Result)
	}
}

//checking account balance
func Test_Account_GetBalance(t *testing.T) {
	var reqData model.AccountGetBalanceRequest
	var address string = "buQemmMwmRQY1JkcU7w3nhruoX5N3j6C29uo"
	reqData.SetAddress(address)
	resData := testSdk.Account.GetBalance(reqData)
	if resData.ErrorCode != 0 {
		t.Errorf(resData.ErrorDesc)
	} else {
		t.Log("Balance:", resData.Result.Balance)
		t.Log("Test_Account_GetBalance succeed", resData.Result)
	}
}

//check the account transaction serial number
func Test_Account_GetNonce(t *testing.T) {
	var reqData model.AccountGetNonceRequest
	var address string = "buQemmMwmRQY1JkcU7w3nhruoX5N3j6C29uo"
	reqData.SetAddress(address)
	resData := testSdk.Account.GetNonce(reqData)
	if resData.ErrorCode != 0 {
		t.Errorf(resData.ErrorDesc)
	} else {
		t.Log("Nonce:", resData.Result.Nonce)
		t.Log("Test_Account_GetNonce succeed", resData.Result)
	}
}

//submit and send bu transactions
func Test_submitTransaction(t *testing.T) {
	//Operation
	var reqDataOperation model.BUSendOperation
	reqDataOperation.Init()
	var amount int64 = 100
	var destAddress string = "buQVU86Jm4FeRW4JcQTD9Rx9NkUkHikYGp6z"
	reqDataOperation.SetAmount(amount)
	reqDataOperation.SetDestAddress(destAddress)
	//reqDataOperation.SetMetadata("send BU")
	//Blob
	var reqDataBlob model.TransactionBuildBlobRequest
	var sourceAddressBlob string = "buQemmMwmRQY1JkcU7w3nhruoX5N3j6C29uo"
	reqDataBlob.SetSourceAddress(sourceAddressBlob)
	var feeLimit int64 = 1000000000
	reqDataBlob.SetFeeLimit(feeLimit)
	var gasPrice int64 = 1000
	reqDataBlob.SetGasPrice(gasPrice)
	var nonce int64 = 97
	reqDataBlob.SetNonce(nonce)
	reqDataBlob.SetOperation(reqDataOperation)
	//reqDataBlob.SetMetadata("send BU")

	resDataBlob := testSdk.Transaction.BuildBlob(reqDataBlob)
	if resDataBlob.ErrorCode != 0 {
		t.Log(resDataBlob.ErrorDesc)
	} else {
		//Sign
		PrivateKey := []string{"privbUPxs6QGkJaNdgWS2hisny6ytx1g833cD7V9C3YET9mJ25wdcq6h"}
		var reqData model.TransactionSignRequest
		reqData.SetBlob(resDataBlob.Result.Blob)
		reqData.SetPrivateKeys(PrivateKey)

		resDataSign := testSdk.Transaction.Sign(reqData)
		if resDataSign.ErrorCode != 0 {
			t.Log(resDataSign.ErrorDesc)
		} else {
			//Submit
			var reqData model.TransactionSubmitRequest
			reqData.SetBlob(resDataBlob.Result.Blob)
			reqData.SetSignatures(resDataSign.Result.Signatures)
			resDataSubmit := testSdk.Transaction.Submit(reqData)

			if resDataSubmit.ErrorCode != 0 {
				t.Errorf(resDataSubmit.ErrorDesc)
			} else {
				t.Log("Hash:", resDataSubmit.Result.Hash)
				t.Log("Test_submitTransaction succeed", resDataSubmit.Result)
			}
		}
	}
}

//enquiry of transaction details
func Test_Transaction_GetInfo(t *testing.T) {
	var reqData model.TransactionGetInfoRequest
	var hash string = "c738fb80dc401d6aba2cf3802ec85ac07fbc23366c003537b64cd1a59ab307d8"
	reqData.SetHash(hash)
	resData := testSdk.Transaction.GetInfo(reqData)
	if resData.ErrorCode != 0 {
		t.Errorf(resData.ErrorDesc)
	} else {
		data, _ := json.Marshal(resData.Result)
		t.Log("info:", string(data))
		t.Log("Test_Transaction_GetInfo succeed", resData.Result)
	}
}

//get block height
func Test_Block_GetNumber(t *testing.T) {
	resData := testSdk.Block.GetNumber()
	if resData.ErrorCode != 0 {
		t.Errorf(resData.ErrorDesc)
	} else {
		t.Log("BlockNumber:", resData.Result.Header.BlockNumber)
		t.Log("Test_Block_GetNumber", resData.Result)
	}
}

//get block details
func Test_Block_GetInfo(t *testing.T) {
	var reqData model.BlockGetInfoRequest
	var blockNumber int64 = 581283
	reqData.SetBlockNumber(blockNumber)
	resData := testSdk.Block.GetInfo(reqData)
	if resData.ErrorCode != 0 {
		t.Errorf(resData.ErrorDesc)
	} else {
		data, _ := json.Marshal(resData.Result.Header)
		t.Log("Header:", string(data))
		t.Log("Test_Block_GetInfo succeed", resData.Result)
	}
}

//get the latest block information
func Test_Block_GetLatest(t *testing.T) {
	resData := testSdk.Block.GetLatest()
	if resData.ErrorCode != 0 {
		t.Errorf(resData.ErrorDesc)
	} else {
		data, _ := json.Marshal(resData.Result.Header)
		t.Log("Header:", string(data))
		t.Log("Test_Block_GetLatest succeed", resData.Result)
	}
}

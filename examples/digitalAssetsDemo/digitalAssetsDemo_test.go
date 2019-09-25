// digitalAssetsDemo
package digitalAssetsDemo_test

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

//get account assets
func Test_Account_GetAssets(t *testing.T) {
	var reqData model.AccountGetAssetsRequest
	var address string = "buQemmMwmRQY1JkcU7w3nhruoX5N3j6C29uo"
	reqData.SetAddress(address)
	resData := testSdk.Account.GetAssets(reqData)
	if resData.ErrorCode != 0 {
		t.Errorf(resData.ErrorDesc)
	} else {
		data, _ := json.Marshal(resData.Result.Assets)
		t.Log("Assets:", string(data))
		t.Log("Test_Account_GetAssets succeed", resData.Result)

	}
}

//get account metadata
func Test_Account_GetMetadata(t *testing.T) {
	var reqData model.AccountGetMetadataRequest
	var address string = "buQXoNR24p2pPqnXPyiDprmTWsU4SYLtBNCG"
	reqData.SetAddress(address)
	reqData.SetKey("global_attribute")
	resData := testSdk.Account.GetMetadata(reqData)
	if resData.ErrorCode != 0 {
		t.Errorf(resData.ErrorDesc)
	} else {
		data, _ := json.Marshal(resData.Result.Metadatas[0].Value)

		t.Log("Metadatas:", string(data))
		t.Log("Test_Account_GetMetadata succeed", resData.Result)
	}
}

//get asset info
func Test_Asset_GetInfo(t *testing.T) {
	var reqData model.AssetGetInfoRequest
	var address string = "buQemmMwmRQY1JkcU7w3nhruoX5N3j6C29uo"
	reqData.SetAddress(address)
	reqData.SetIssuer("buQnc3AGCo6ycWJCce516MDbPHKjK7ywwkuo")
	reqData.SetCode("HNC")
	resData := testSdk.Token.Asset.GetInfo(reqData)
	if resData.ErrorCode != 0 {
		t.Errorf(resData.ErrorDesc)
	} else {
		data, _ := json.Marshal(resData.Result.Assets)
		t.Log("Assets:", string(data))
		t.Log("Test_Asset_GetInfo succeed", resData.Result.Assets)
	}
}

//get contract info
func Test_Contract_GetInfo(t *testing.T) {
	var reqData model.ContractGetInfoRequest
	var address string = "buQXmYrmqt6ohcKtLFKgWFSZ5CjYKaSzaMjT"
	reqData.SetAddress(address)
	resData := testSdk.Contract.GetInfo(reqData)
	if resData.ErrorCode != 0 {
		t.Errorf(resData.ErrorDesc)
	} else {
		data, _ := json.Marshal(resData.Result.Contract)
		t.Log("Contract:", string(data))
		t.Log("Test_Contract_GetInfo succeed", resData.Result)
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

//evaluate fee
func Test_Transaction_EvaluateFee(t *testing.T) {
	var reqDataOperation model.BUSendOperation
	reqDataOperation.Init()
	var amount int64 = 100
	reqDataOperation.SetAmount(amount)
	reqDataOperation.SetMetadata("63")
	var destAddress string = "buQemmMwmRQY1JkcU7w3nhruoX5N3j6C29uo"
	reqDataOperation.SetDestAddress(destAddress)

	var reqDataEvaluate model.TransactionEvaluateFeeRequest
	var sourceAddress string = "buQVU86Jm4FeRW4JcQTD9Rx9NkUkHikYGp6z"
	reqDataEvaluate.SetSourceAddress(sourceAddress)
	var nonce int64 = 5
	reqDataEvaluate.SetNonce(nonce)
	var signatureNumber string = "3"
	reqDataEvaluate.SetSignatureNumber(signatureNumber)
	var SetCeilLedgerSeq int64 = 50
	reqDataEvaluate.SetCeilLedgerSeq(SetCeilLedgerSeq)
	reqDataEvaluate.SetMetadata("63")
	reqDataEvaluate.SetOperation(reqDataOperation)
	resDataEvaluate := testSdk.Transaction.EvaluateFee(reqDataEvaluate)
	if resDataEvaluate.ErrorCode != 0 {
		t.Log(resDataEvaluate)
		t.Errorf(resDataEvaluate.ErrorDesc)
	} else {
		data, _ := json.Marshal(resDataEvaluate.Result)
		t.Log("Evaluate:", string(data))
		t.Log("Test_EvaluateFee succeed", resDataEvaluate.Result)
	}
}

//Activate Account
func Test_activate_Account(t *testing.T) {
	// The account private key to activate a new account
	var activatePrivateKey string = "privbtZbKbz4CtsjtN9GnLsUgSG2vKGuMDri9ADnQ7AzYg5kHqcdCH4y"
	var activateAddress string = "buQXmYrmqt6ohcKtLFKgWFSZ5CjYKaSzaMjT"
	var initBalance int64 = 100000000000
	// The fixed write 1000L, the unit is MO
	var gasPrice int64 = 1000
	// Set up the maximum cost 0.01BU
	var feeLimit int64 = 100000000
	// Transaction initiation account's nonce + 1
	var nonce int64 = 8
	// The account to be activated
	var destAddress string = "buQVU86Jm4FeRW4JcQTD9Rx9NkUkHikYGp6z"
	//Operation
	var reqDataOperation model.AccountActivateOperation
	reqDataOperation.Init()
	reqDataOperation.SetDestAddress(destAddress)
	reqDataOperation.SetInitBalance(initBalance)

	errorCode, errorDesc, hash := submitTransaction(testSdk, reqDataOperation, activatePrivateKey, activateAddress, nonce, gasPrice, feeLimit)
	if errorCode != 0 {
		t.Log("errorDesc:", errorDesc)
	} else {
		t.Log("Test_BU_Send succeed", hash)
	}

}

//Asset Issue
func Test_Asset_Issue(t *testing.T) {
	// Init variable
	// The account private key to issue asset
	var issuePrivateKey string = "privbtZbKbz4CtsjtN9GnLsUgSG2vKGuMDri9ADnQ7AzYg5kHqcdCH4y"
	// The account address to send this transaction
	var issueAddress string = "buQXmYrmqt6ohcKtLFKgWFSZ5CjYKaSzaMjT"
	// Asset code
	var assetCode string = "TST"
	// Asset amount
	var assetAmount int64 = 10000000000000
	// metadata
	var metadata string = "issue TST"
	// The fixed write 1000L, the unit is MO
	var gasPrice int64 = 1000
	// Set up the maximum cost 50.01BU
	var feeLimit int64 = 10000000000000
	// Transaction initiation account's nonce + 1
	var nonce int64 = 7

	//Operation
	var reqDataOperation model.AssetIssueOperation
	reqDataOperation.Init()

	reqDataOperation.SetAmount(assetAmount)
	reqDataOperation.SetCode(assetCode)
	reqDataOperation.SetMetadata(metadata)
	errorCode, errorDesc, hash := submitTransaction(testSdk, reqDataOperation, issuePrivateKey, issueAddress, nonce, gasPrice, feeLimit)
	if errorCode != 0 {
		t.Log("errorDesc:", errorDesc)

	} else {
		t.Log("Test_BU_Send succeed", hash)
	}
}

//Build blob
func Test_Transaction_BuildBlob(t *testing.T) {
	// Init variable
	// The account address to send this transaction
	var senderAddress string = "buQXmYrmqt6ohcKtLFKgWFSZ5CjYKaSzaMjT"
	// The account to receive asset
	var destAddress string = "buQhapCK83xPPdjQeDuBLJtFNvXYZEKb6tKB"
	// Asset code
	var assetCode string = "TST"
	// The accout address of issuing asset
	var assetIssuer string = "buQcGP2a1PY45dauMfhk9QsFbn7a6BKKAM9x"
	// The asset amount to be sent
	var amount int64 = 1000000000000000
	// The fixed write 1000L, the unit is MO
	var gasPrice int64 = 0
	// Set up the maximum cost 0.01BU
	var feeLimit int64 = 0
	// Transaction initiation account's nonce + 1
	var nonce int64 = 2

	//Operation
	var reqDataOperation model.AssetSendOperation
	reqDataOperation.Init()
	reqDataOperation.SetAmount(amount)
	reqDataOperation.SetCode(assetCode)
	reqDataOperation.SetDestAddress(destAddress)
	reqDataOperation.SetIssuer(assetIssuer)

	var reqDataBlob model.TransactionBuildBlobRequest
	reqDataBlob.SetSourceAddress(senderAddress)
	reqDataBlob.SetFeeLimit(feeLimit)
	reqDataBlob.SetGasPrice(gasPrice)
	reqDataBlob.SetNonce(nonce)
	reqDataBlob.SetOperation(reqDataOperation)
	//reqDataBlob.SetMetadata("abc")

	resDataBlob := testSdk.Transaction.BuildBlob(reqDataBlob)
	if resDataBlob.ErrorCode != 0 {
		t.Log("errorDesc: ", resDataBlob.ErrorCode, resDataBlob.ErrorDesc, "")
	} else {
		t.Log("blob: " , resDataBlob.Result.Blob)
	}
}

//Asset Send
func Test_Asset_Send(t *testing.T) {
	// Init variable
	// The account private key to start this transaction
	var senderPrivateKey string = "privbtZbKbz4CtsjtN9GnLsUgSG2vKGuMDri9ADnQ7AzYg5kHqcdCH4y"
	// The account address to send this transaction
	var senderAddress string = "buQXmYrmqt6ohcKtLFKgWFSZ5CjYKaSzaMjT"
	// The account to receive asset
	var destAddress string = "buQhapCK83xPPdjQeDuBLJtFNvXYZEKb6tKB"
	// Asset code
	var assetCode string = "TST"
	// The accout address of issuing asset
	var assetIssuer string = "buQcGP2a1PY45dauMfhk9QsFbn7a6BKKAM9x"
	// The asset amount to be sent
	var amount int64 = 1000000000000000
	// The fixed write 1000L, the unit is MO
	var gasPrice int64 = 1000
	// Set up the maximum cost 0.01BU
	var feeLimit int64 = 1000000000000000
	// Transaction initiation account's nonce + 1
	var nonce int64 = 2

	//Operation
	var reqDataOperation model.AssetSendOperation
	reqDataOperation.Init()
	reqDataOperation.SetAmount(amount)
	reqDataOperation.SetCode(assetCode)
	reqDataOperation.SetDestAddress(destAddress)
	reqDataOperation.SetIssuer(assetIssuer)

	errorCode, errorDesc, hash := submitTransaction(testSdk, reqDataOperation, senderPrivateKey, senderAddress, nonce, gasPrice, feeLimit)
	if errorCode != 0 {
		t.Log("errorDesc:", errorDesc)
	} else {
		t.Log("Test_BU_Send succeed", hash)
	}
}

//BU Send
func Test_BU_Send(t *testing.T) {
	// Init variable
	// The account private key to start this transaction
	var senderPrivateKey string = "privbtZbKbz4CtsjtN9GnLsUgSG2vKGuMDri9ADnQ7AzYg5kHqcdCH4y"
	// The account address to send this transaction
	var senderAddress string = "buQXmYrmqt6ohcKtLFKgWFSZ5CjYKaSzaMjT"
	// The account address to receive bu
	var destAddress string = "buQsurH1M4rjLkfjzkxR9KXJ6jSu2r9xBNEw"
	// The amount to be sent
	var amount int64 = 1000000000000
	// The fixed write 1000L, the unit is MO
	var gasPrice int64 = 1000
	// Set up the maximum cost 0.01BU
	var feeLimit int64 = 10000000000
	// Transaction initiation account's nonce + 1
	var nonce int64 = 25

	//Operation
	var reqDataOperation model.BUSendOperation
	reqDataOperation.Init()
	reqDataOperation.SetAmount(amount)
	reqDataOperation.SetDestAddress(destAddress)
	errorCode, errorDesc, hash := submitTransaction(testSdk, reqDataOperation, senderPrivateKey, senderAddress, nonce, gasPrice, feeLimit)
	if errorCode != 0 {
		t.Log("errorDesc:", errorDesc)
	} else {
		t.Log("Test_BU_Send succeed", hash)
	}
}
func submitTransaction(testSdk sdk.Sdk, reqDataOperation model.BaseOperation, senderPrivateKey string, senderAddresss string, senderNonce int64, gasPrice int64, feeLimit int64) (errorCode int, errorDesc string, hash string) {
	//Blob
	var reqDataBlob model.TransactionBuildBlobRequest
	reqDataBlob.SetSourceAddress(senderAddresss)
	reqDataBlob.SetFeeLimit(feeLimit)
	reqDataBlob.SetGasPrice(gasPrice)
	reqDataBlob.SetNonce(senderNonce)
	reqDataBlob.SetOperation(reqDataOperation)
	//reqDataBlob.SetMetadata("abc")

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
			return resDataSubmit.ErrorCode, resDataSubmit.ErrorDesc, resDataSubmit.Result.Hash
		}
	}
}

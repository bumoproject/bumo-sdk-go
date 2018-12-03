// sdk_test
package sdk_test

import (
	"encoding/json"

	"testing"

	"github.com/bumoproject/bumo-sdk-go/src/model"
	"github.com/bumoproject/bumo-sdk-go/src/sdk"
)

var testSdk sdk.Sdk

//init
func Test_Init(t *testing.T) {
	var reqData model.SDKInitRequest
	reqData.SetUrl("http://127.0.0.1:36002")
	reqData.SetChainId(1)
	resData := testSdk.Init(reqData)
	if resData.ErrorCode != 0 {
		t.Errorf(resData.ErrorDesc)
	} else {
		t.Log("Test_NewSDK")
	}
}

//get block number
func Test_Block_GetNumber(t *testing.T) {

	resData := testSdk.Block.GetNumber()
	if resData.ErrorCode != 0 {
		t.Errorf(resData.ErrorDesc)
	} else {
		t.Log("BlockNumber:", resData.Result.Header.BlockNumber)
		t.Log("Test_Block_GetNumber", resData.Result)
	}
}

//check block status
func Test_Block_CheckStatus(t *testing.T) {
	resData := testSdk.Block.CheckStatus()
	if resData.ErrorCode != 0 {
		t.Errorf(resData.ErrorDesc)
	} else {
		t.Log("IsSynchronous:", resData.Result.IsSynchronous)
		t.Log("Test_Block_CheckStatus succeed", resData.Result)
	}

}

//get block transactions
func Test_Block_GetTransactions(t *testing.T) {
	var reqData model.BlockGetTransactionRequest
	var blockNumber int64 = 685714
	reqData.SetBlockNumber(blockNumber)
	resData := testSdk.Block.GetTransactions(reqData)
	if resData.ErrorCode != 0 {
		t.Errorf(resData.ErrorDesc)
	} else {
		data, _ := json.Marshal(resData.Result)
		t.Log("Result:", string(data))
		t.Log("Test_Block_GetTransactions succeed", resData.Result)
	}
}

//get block info
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

//get block latest
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

//get block validators
func Test_Block_GetValidators(t *testing.T) {
	var reqData model.BlockGetValidatorsRequest
	var blockNumber int64 = 581283
	reqData.SetBlockNumber(blockNumber)
	resData := testSdk.Block.GetValidators(reqData)
	if resData.ErrorCode != 0 {
		t.Errorf(resData.ErrorDesc)
	} else {
		data, _ := json.Marshal(resData.Result.Validators)
		t.Log("Validators:", string(data))
		t.Log("Test_Block_GetValidators succeed", resData.Result)
	}
}

//get block latest validators
func Test_Block_GetLatestValidators(t *testing.T) {
	resData := testSdk.Block.GetLatestValidators()
	if resData.ErrorCode != 0 {
		t.Errorf(resData.ErrorDesc)
	} else {
		data, _ := json.Marshal(resData.Result.Validators)
		t.Log("Validators:", string(data))
		t.Log("Test_Block_GetLatestValidators succeed", resData.Result)
	}
}

//get block reward
func Test_Block_GetReward(t *testing.T) {
	var reqData model.BlockGetRewardRequest
	var blockNumber int64 = 581283
	reqData.SetBlockNumber(blockNumber)
	resData := testSdk.Block.GetReward(reqData)
	if resData.ErrorCode != 0 {
		t.Errorf(resData.ErrorDesc)
	} else {
		t.Log("ValidatorsReward:", resData.Result.ValidatorsReward)
		t.Log("Test_Block_GetReward succeed", resData.Result)
	}
}

//get block latestreward
func Test_Block_GetLatestReward(t *testing.T) {
	resData := testSdk.Block.GetLatestReward()
	if resData.ErrorCode != 0 {
		t.Errorf(resData.ErrorDesc)
	} else {
		t.Log("ValidatorsReward:", resData.Result.ValidatorsReward)
		t.Log("Test_Block_GetLatestReward succeed", resData.Result)
	}
}

//get block fees
func Test_Block_GetFees(t *testing.T) {
	var reqData model.BlockGetFeesRequest
	var blockNumber int64 = 581283
	reqData.SetBlockNumber(blockNumber)
	resData := testSdk.Block.GetFees(reqData)
	if resData.ErrorCode != 0 {
		t.Errorf(resData.ErrorDesc)
	} else {
		data, _ := json.Marshal(resData.Result.Fees)
		t.Log("Fees:", string(data))
		t.Log("Test_Block_GetFees succeed", resData.Result)
	}
}

//get block latest fees
func Test_Block_GetLatestFees(t *testing.T) {
	resData := testSdk.Block.GetLatestFees()
	if resData.ErrorCode != 0 {
		t.Errorf(resData.ErrorDesc)
	} else {
		data, _ := json.Marshal(resData.Result.Fees)
		t.Log("Fees:", string(data))
		t.Log("Test_Block_GetLatestFees succeed", resData.Result)
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

//send bu
func Test_Transaction_BuildBlob_Sign_Submit(t *testing.T) {
	var reqDataOperation model.BUSendOperation
	reqDataOperation.Init()
	var amount int64 = 100
	var destAddress string = "buQVU86Jm4FeRW4JcQTD9Rx9NkUkHikYGp6z"
	reqDataOperation.SetAmount(amount)
	reqDataOperation.SetMetadata("63")
	reqDataOperation.SetDestAddress(destAddress)

	var reqDataBlob model.TransactionBuildBlobRequest
	var sourceAddressBlob string = "buQemmMwmRQY1JkcU7w3nhruoX5N3j6C29uo"
	reqDataBlob.SetSourceAddress(sourceAddressBlob)
	var feeLimit int64 = 1000000
	reqDataBlob.SetFeeLimit(feeLimit)
	var gasPrice int64 = 1000
	reqDataBlob.SetGasPrice(gasPrice)
	var nonce int64 = 109
	reqDataBlob.SetNonce(nonce)
	reqDataBlob.SetMetadata("63")
	var CeilLedgerSeq int64 = 50
	reqDataBlob.SetCeilLedgerSeq(CeilLedgerSeq)
	reqDataBlob.SetOperation(reqDataOperation)

	resDataBlob := testSdk.Transaction.BuildBlob(reqDataBlob)
	if resDataBlob.ErrorCode != 0 {
		t.Log(resDataBlob.ErrorDesc)
	} else {
		PrivateKeys := []string{"privbtYzJ6miiFktK9BsDAMRNd3J4eKkuszfXqJ2huQ2h8DGUnRs9nuq"}
		var reqData model.TransactionSignRequest
		reqData.SetBlob(resDataBlob.Result.Blob)
		reqData.SetPrivateKeys(PrivateKeys)
		resDataSign := testSdk.Transaction.Sign(reqData)
		if resDataSign.ErrorCode != 0 {
			t.Log(resDataSign.ErrorDesc)
		} else {
			var reqData model.TransactionSubmitRequest
			reqData.SetBlob(resDataBlob.Result.Blob)
			t.Log(resDataSign.Result.Signatures[0].SignData)
			reqData.SetSignatures(resDataSign.Result.Signatures)
			resDataSubmit := testSdk.Transaction.Submit(reqData)

			if resDataSubmit.ErrorCode != 0 {
				t.Errorf(resDataSubmit.ErrorDesc)
			} else {
				t.Log("Hash:", resDataSubmit.Result.Hash)
				t.Log("Test_Transaction_BuildBlob_Sign_Submit succeed", resDataSubmit.Result)
			}
		}
	}
}

//get transaction info
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

//checkvalid account
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

//create account
func Test_Account_Create(t *testing.T) {
	resData := testSdk.Account.Create()
	if resData.ErrorCode != 0 {
		t.Errorf(resData.ErrorDesc)
	} else {
		str, _ := json.Marshal(resData.Result)
		t.Log(string(str))
		t.Log("Test_Account_Create", resData.Result)
	}
}

//get account info
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

//get account nonce
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

//get account balance
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

//check account Activated
func Test_Account_CheckActivated(t *testing.T) {
	var reqData model.AccountCheckActivatedRequest
	var address string = "buQXoNR24p2pPqnXPyiDprmTWsU4SYLtBNCG"
	reqData.SetAddress(address)
	resData := testSdk.Account.CheckActivated(reqData)
	if resData.ErrorCode != 0 {
		t.Errorf(resData.ErrorDesc)
	} else {
		t.Log("Test_Account_CheckActivated succeed", resData.Result)
	}
}

//get asset info
func Test_Asset_GetInfo(t *testing.T) {
	var reqData model.AssetGetInfoRequest
	var address string = "buQemmMwmRQY1JkcU7w3nhruoX5N3j6C29uo"
	reqData.SetAddress(address)
	var issuer string = "buQnc3AGCo6ycWJCce516MDbPHKjK7ywwkuo"
	reqData.SetIssuer(issuer)
	var code string = "HNC"
	reqData.SetCode(code)
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

//check valid
func Test_Contract_CheckValid(t *testing.T) {
	var reqData model.ContractCheckValidRequest
	var address string = "buQXmYrmqt6ohcKtLFKgWFSZ5CjYKaSzaMjT"
	reqData.SetAddress(address)
	resData := testSdk.Contract.CheckValid(reqData)
	if resData.ErrorCode != 0 {
		t.Errorf(resData.ErrorDesc)
	} else {
		t.Log("Test_Contract_CheckValid succeed", resData.Result)
	}
}

//call
func Test_Contract_Call(t *testing.T) {
	var reqData model.ContractCallRequest
	var contractAddress string = "buQXmYrmqt6ohcKtLFKgWFSZ5CjYKaSzaMjT"
	reqData.SetContractAddress(contractAddress)
	var contractBalance string = "100000000000"
	reqData.SetContractBalance(contractBalance)
	var feeLimit int64 = 1000000
	reqData.SetFeeLimit(feeLimit)
	var gasPrice int64 = 1000
	reqData.SetGasPrice(gasPrice)
	var input string = "input"
	reqData.SetInput(input)
	var optType int64 = 2
	reqData.SetOptType(optType)
	var code string = "HNC"
	reqData.SetCode(code)
	resData := testSdk.Contract.Call(reqData)
	if resData.ErrorCode != 0 {
		t.Errorf(resData.ErrorDesc)
	} else {
		t.Log("Test_Contract_CheckValid succeed", resData.Result)
	}
}

//get address
func Test_Contract_GetAddress(t *testing.T) {
	var reqData model.ContractGetAddressRequest
	var hash string = "c738fb80dc401d6aba2cf3802ec85ac07fbc23366c003537b64cd1a59ab307d8"
	reqData.SetHash(hash)
	resData := testSdk.Contract.GetAddress(reqData)
	if resData.ErrorCode != 0 {
		t.Errorf(resData.ErrorDesc)
	} else {
		t.Log("Test_Contract_GetAddress succeed", resData.Result.ContractAddresInfos)
	}
}

//ctp10token allowance
func Test_Ctp10Token_Allowance(t *testing.T) {
	var reqData model.Ctp10TokenAllowanceRequest
	var contractAddress string = "buQXoNR24p2pPqnXPyiDprmTWsU4SYLtBNCG"
	reqData.SetContractAddress(contractAddress)
	var spender string = "buQW5p6gaCd331NerjxhD1cAHpmSGwxrt6e6"
	reqData.SetSpender(spender)
	var tokenOwner string = "buQnc3AGCo6ycWJCce516MDbPHKjK7ywwkuo"
	reqData.SetTokenOwner(tokenOwner)
	resData := testSdk.Token.Ctp10Token.Allowance(reqData)
	if resData.ErrorCode != 0 {
		t.Errorf(resData.ErrorDesc)
	} else {
		t.Log("Allowance:", resData.Result.Allowance)
		t.Log("Test_Ctp10Token_Allowance succeed", resData.Result)
	}
}

//get ctp10token info
func Test_Ctp10Token_GetInfo(t *testing.T) {
	var reqData model.Ctp10TokenGetInfoRequest
	var contractAddress string = "buQXoNR24p2pPqnXPyiDprmTWsU4SYLtBNCG"
	reqData.SetContractAddress(contractAddress)
	resData := testSdk.Token.Ctp10Token.GetInfo(reqData)
	if resData.ErrorCode != 0 {
		t.Log(resData)
		t.Errorf(resData.ErrorDesc)
	} else {
		data, _ := json.Marshal(resData.Result)
		t.Log("info:", string(data))
		t.Log("Test_Ctp10Token_GetInfo succeed", resData.Result)
	}
}

//get ctp10token name
func Test_Ctp10Token_GetName(t *testing.T) {
	var reqData model.Ctp10TokenGetNameRequest
	var contractAddress string = "buQXoNR24p2pPqnXPyiDprmTWsU4SYLtBNCG"
	reqData.SetContractAddress(contractAddress)
	resData := testSdk.Token.Ctp10Token.GetName(reqData)
	if resData.ErrorCode != 0 {
		t.Errorf(resData.ErrorDesc)
	} else {
		t.Log("Name:", resData.Result.Name)
		t.Log("Test_Ctp10Token_GetName succeed", resData.Result)
	}
}

//get ctp10token symbol
func Test_Ctp10Token_GetSymbol(t *testing.T) {
	var reqData model.Ctp10TokenGetSymbolRequest
	var contractAddress string = "buQXoNR24p2pPqnXPyiDprmTWsU4SYLtBNCG"
	reqData.SetContractAddress(contractAddress)
	resData := testSdk.Token.Ctp10Token.GetSymbol(reqData)
	if resData.ErrorCode != 0 {
		t.Errorf(resData.ErrorDesc)
	} else {
		t.Log("Symbol:", resData.Result.Symbol)
		t.Log("Test_Ctp10Token_GetSymbol succeed", resData.Result)
	}
}

//get ctp10token decimals
func Test_Ctp10Token_GetDecimals(t *testing.T) {
	var reqData model.Ctp10TokenGetDecimalsRequest
	var contractAddress string = "buQXoNR24p2pPqnXPyiDprmTWsU4SYLtBNCG"
	reqData.SetContractAddress(contractAddress)
	resData := testSdk.Token.Ctp10Token.GetDecimals(reqData)
	if resData.ErrorCode != 0 {
		t.Errorf(resData.ErrorDesc)
	} else {
		t.Log("Decimals:", resData.Result.Decimals)
		t.Log("Test_Ctp10Token_GetDecimals succeed", resData.Result)
	}
}

//get ctp10token totalsupply
func Test_Ctp10Token_GetTotalSupply(t *testing.T) {
	var reqData model.Ctp10TokenGetTotalSupplyRequest
	var contractAddress string = "buQXoNR24p2pPqnXPyiDprmTWsU4SYLtBNCG"
	reqData.SetContractAddress(contractAddress)
	resData := testSdk.Token.Ctp10Token.GetTotalSupply(reqData)
	if resData.ErrorCode != 0 {
		t.Errorf(resData.ErrorDesc)
	} else {
		t.Log("TotalSupply:", resData.Result.TotalSupply)
		t.Log("Test_Ctp10Token_GetTotalSupply succeed", resData.Result)
	}
}

//get ctp10token balance
func Test_Ctp10Token_GetBalance(t *testing.T) {
	var reqData model.Ctp10TokenGetBalanceRequest
	var contractAddress string = "buQXoNR24p2pPqnXPyiDprmTWsU4SYLtBNCG"
	reqData.SetContractAddress(contractAddress)
	var tokenOwner string = "buQW5p6gaCd331NerjxhD1cAHpmSGwxrt6e6"
	reqData.SetTokenOwner(tokenOwner)
	resData := testSdk.Token.Ctp10Token.GetBalance(reqData)
	if resData.ErrorCode != 0 {
		t.Log(resData)
		t.Errorf(resData.ErrorDesc)
	} else {
		t.Log("Balance:", resData.Result.Balance)
		t.Log("Test_Ctp10Token_GetBalance succeed", resData.Result)
	}
}

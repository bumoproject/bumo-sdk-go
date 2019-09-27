// transaction
package blockchain

import (
	"encoding/hex"
	"encoding/json"
	"math"
	"strconv"

	"github.com/bumoproject/bumo-sdk-go/src/common"
	"github.com/bumoproject/bumo-sdk-go/src/crypto/keypair"
	"github.com/bumoproject/bumo-sdk-go/src/crypto/protocol"
	"github.com/bumoproject/bumo-sdk-go/src/crypto/signature"
	"github.com/bumoproject/bumo-sdk-go/src/exception"
	"github.com/bumoproject/bumo-sdk-go/src/model"
	"github.com/golang/protobuf/proto"
)

type TransactionOperation struct {
	Url string
}

// build blob
func (transaction *TransactionOperation) BuildBlob(reqData model.TransactionBuildBlobRequest) model.TransactionBuildBlobResponse {
	var resData model.TransactionBuildBlobResponse
	if !keypair.CheckAddress(reqData.GetSourceAddress()) {
		SDKRes := exception.GetSDKRes(exception.INVALID_SOURCEADDRESS_ERROR)
		resData.ErrorCode = SDKRes.ErrorCode
		resData.ErrorDesc = SDKRes.ErrorDesc
		return resData
	}
	if reqData.GetNonce() <= 0 {
		SDKRes := exception.GetSDKRes(exception.INVALID_NONCE_ERROR)
		resData.ErrorCode = SDKRes.ErrorCode
		resData.ErrorDesc = SDKRes.ErrorDesc
		return resData
	}
	if reqData.GetCeilLedgerSeq() < 0 {
		SDKRes := exception.GetSDKRes(exception.INVALID_CEILLEDGERSEQ_ERROR)
		resData.ErrorCode = SDKRes.ErrorCode
		resData.ErrorDesc = SDKRes.ErrorDesc
		return resData
	}
	if reqData.GetGasPrice() < 0 {
		SDKRes := exception.GetSDKRes(exception.INVALID_GASPRICE_ERROR)
		resData.ErrorCode = SDKRes.ErrorCode
		resData.ErrorDesc = SDKRes.ErrorDesc
		return resData
	}
	if reqData.GetFeeLimit() < 0 {
		SDKRes := exception.GetSDKRes(exception.INVALID_FEELIMIT_ERROR)
		resData.ErrorCode = SDKRes.ErrorCode
		resData.ErrorDesc = SDKRes.ErrorDesc
		return resData
	}
	operationsData := reqData.GetOperations()
	if operationsData.Len() == 0 {
		SDKRes := exception.GetSDKRes(exception.OPERATIONS_EMPTY_ERROR)
		resData.ErrorCode = SDKRes.ErrorCode
		resData.ErrorDesc = SDKRes.ErrorDesc
		return resData
	}
	operations, SDKRes := common.GetOperations(operationsData, transaction.Url, reqData.GetSourceAddress())
	if SDKRes.ErrorCode != 0 {
		resData.ErrorCode = SDKRes.ErrorCode
		resData.ErrorDesc = SDKRes.ErrorDesc
		return resData
	}
	for i := range operations {
		switch operations[i].Type {
		case protocol.Operation_CREATE_ACCOUNT:
			if operations[i].GetCreateAccount().GetDestAddress() == reqData.GetSourceAddress() {
				resData.ErrorCode = exception.SOURCEADDRESS_EQUAL_DESTADDRESS_ERROR
				resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
				return resData
			}
		case protocol.Operation_PAY_ASSET:
			if operations[i].GetPayAsset().GetDestAddress() == reqData.GetSourceAddress() {
				resData.ErrorCode = exception.SOURCEADDRESS_EQUAL_CONTRACTADDRESS_ERROR
				resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
				return resData
			}
		case protocol.Operation_PAY_COIN:
			if operations[i].GetPayCoin().GetDestAddress() == reqData.GetSourceAddress() {
				resData.ErrorCode = exception.SOURCEADDRESS_EQUAL_CONTRACTADDRESS_ERROR
				resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
				return resData
			}
		default:
		}
	}
	if SDKRes.ErrorCode != 0 {
		resData.ErrorCode = SDKRes.ErrorCode
		resData.ErrorDesc = SDKRes.ErrorDesc
		return resData
	}
	if reqData.GetCeilLedgerSeq() < 0 {
		resData.ErrorCode = exception.INVALID_CEILLEDGERSEQ_ERROR
		resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
		return resData
	}
	var seq int64 = 0
	if reqData.GetCeilLedgerSeq() > 0 {
		var Block BlockOperation
		Block.Url = transaction.Url
		resDataNumber := Block.GetNumber()
		seq = reqData.GetCeilLedgerSeq() + resDataNumber.Result.Header.BlockNumber
	}
	Transaction := protocol.Transaction{
		SourceAddress: reqData.GetSourceAddress(),
		Nonce:         reqData.GetNonce(),
		CeilLedgerSeq: seq,
		FeeLimit:      reqData.GetFeeLimit(),
		GasPrice:      reqData.GetGasPrice(),
		Metadata:      []byte(reqData.GetMetadata()),
		Operations:    operations,
	}
	data, err := proto.Marshal(&Transaction)
	if err != nil {
		SDKRes := exception.GetSDKRes(exception.SYSTEM_ERROR)
		resData.ErrorCode = SDKRes.ErrorCode
		resData.ErrorDesc = SDKRes.ErrorDesc
		return resData
	}
	dataStr := hex.EncodeToString(data)
	resData.Result.Blob = dataStr
	resData.ErrorCode = exception.SUCCESS
	return resData
}

// evaluate fee
func (transaction *TransactionOperation) EvaluateFee(reqData model.TransactionEvaluateFeeRequest) model.TransactionEvaluateFeeResponse {
	var resDataD model.TransactionEvaluateFeeData
	var resData model.TransactionEvaluateFeeResponse
	if !keypair.CheckAddress(reqData.GetSourceAddress()) {
		SDKRes := exception.GetSDKRes(exception.INVALID_SOURCEADDRESS_ERROR)
		resData.ErrorCode = SDKRes.ErrorCode
		resData.ErrorDesc = SDKRes.ErrorDesc
		return resData
	}
	if reqData.GetNonce() <= 0 {
		SDKRes := exception.GetSDKRes(exception.INVALID_NONCE_ERROR)
		resData.ErrorCode = SDKRes.ErrorCode
		resData.ErrorDesc = SDKRes.ErrorDesc
		return resData
	}
	operationsData := reqData.GetOperations()
	if operationsData.Len() == 0 {
		SDKRes := exception.GetSDKRes(exception.OPERATIONS_EMPTY_ERROR)
		resData.ErrorCode = SDKRes.ErrorCode
		resData.ErrorDesc = SDKRes.ErrorDesc
		return resData
	}
	if reqData.GetCeilLedgerSeq() < 0 {
		resData.ErrorCode = exception.INVALID_CEILLEDGERSEQ_ERROR
		resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
		return resData
	}
	var SignatureNumber int64 = 1
	if len(reqData.GetSignatureNumber()) != 0 {
		var err error
		SignatureNumber, err = strconv.ParseInt(reqData.GetSignatureNumber(), 10, 64)
		if err != nil || SignatureNumber <= 0 || SignatureNumber > math.MaxInt32 {
			SDKRes := exception.GetSDKRes(exception.INVALID_SIGNATURENUMBER_ERROR)
			resData.ErrorCode = SDKRes.ErrorCode
			resData.ErrorDesc = SDKRes.ErrorDesc
			return resData
		}
	}
	operations, SDKRes := common.GetOperations(operationsData, transaction.Url, reqData.GetSourceAddress())
	if SDKRes.ErrorCode != 0 {
		resData.ErrorCode = SDKRes.ErrorCode
		resData.ErrorDesc = SDKRes.ErrorDesc
		return resData
	}
	Operations := make([]model.OperationEvaluat, len(operations))
	for i := range operations {
		Operations[i].SourceAddress = operations[i].GetSourceAddress()
		Operations[i].Metadata = string(operations[i].GetMetadata())
		Operations[i].Type = operations[i].GetType()
		Operations[i].CreateAccount = operations[i].CreateAccount
		Operations[i].IssueAsset = operations[i].IssueAsset
		Operations[i].Log = operations[i].Log
		Operations[i].PayAsset = operations[i].PayAsset
		Operations[i].PayCoin = operations[i].PayCoin
		Operations[i].SetMetadata = operations[i].SetMetadata
		Operations[i].SetPrivilege = operations[i].SetPrivilege
		Operations[i].SetSignerWeight = operations[i].SetSignerWeight
		Operations[i].SetThreshold = operations[i].SetThreshold
	}
	var seq int64 = 0
	if reqData.GetCeilLedgerSeq() > 0 {
		var Block BlockOperation
		Block.Url = transaction.Url
		resDataNumber := Block.GetNumber()
		seq = reqData.GetCeilLedgerSeq() + resDataNumber.Result.Header.BlockNumber
	}
	request := &model.WebTransactionEvaluateFeeResponse{
		Items: []model.Item{
			{
				TransactionJson: model.TransactionJson{
					SourceAddress: reqData.GetSourceAddress(),
					Metadata:      reqData.GetMetadata(),
					Nonce:         reqData.GetNonce(),
					CeilLedgerSeq: seq,
					Operations:    Operations,
				},
				SignatureNumber: SignatureNumber,
			},
		},
	}
	requestJson, err := json.Marshal(request)
	if err != nil {
		SDKRes := exception.GetSDKRes(exception.SYSTEM_ERROR)
		resData.ErrorCode = SDKRes.ErrorCode
		resData.ErrorDesc = SDKRes.ErrorDesc
		return resData
	}
	response, SDKRes := common.PostRequest(transaction.Url, "/testTransaction", requestJson)
	if SDKRes.ErrorCode != 0 {
		resData.ErrorCode = SDKRes.ErrorCode
		resData.ErrorDesc = SDKRes.ErrorDesc
		return resData
	}
	defer response.Body.Close()
	if response.StatusCode == 200 {
		decoder := json.NewDecoder(response.Body)
		decoder.UseNumber()
		err = decoder.Decode(&resDataD)
		if err != nil {
			SDKRes := exception.GetSDKRes(exception.SYSTEM_ERROR)
			resData.ErrorCode = SDKRes.ErrorCode
			resData.ErrorDesc = SDKRes.ErrorDesc
			return resData
		}
		if resDataD.ErrorCode == 0 {
			if resDataD.Result.Txs == nil {
				resData.ErrorCode = exception.THE_QUERY_FAILED
				resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
				return resData
			}
			resData.ErrorCode = exception.SUCCESS
			resData.Result.FeeLimit = resDataD.Result.Txs[0].TransactionEnv.Transaction.FeeLimit
			resData.Result.GasPrice = resDataD.Result.Txs[0].TransactionEnv.Transaction.GasPrice
			return resData
		} else {
			resData.ErrorCode = resDataD.ErrorCode
			resData.ErrorDesc = resDataD.ErrorDesc
			return resData
		}
	} else {
		SDKRes := exception.GetSDKRes(exception.CONNECTNETWORK_ERROR)
		resData.ErrorCode = SDKRes.ErrorCode
		resData.ErrorDesc = SDKRes.ErrorDesc
		return resData
	}
}

// sign
func (transaction *TransactionOperation) Sign(reqData model.TransactionSignRequest) model.TransactionSignResponse {
	var resData model.TransactionSignResponse
	if reqData.GetBlob() == "" {
		SDKRes := exception.GetSDKRes(exception.INVALID_BLOB_ERROR)
		resData.ErrorCode = SDKRes.ErrorCode
		resData.ErrorDesc = SDKRes.ErrorDesc
		return resData
	}
	if reqData.GetPrivateKeys() == nil {
		SDKRes := exception.GetSDKRes(exception.PRIVATEKEY_NULL_ERROR)
		resData.ErrorCode = SDKRes.ErrorCode
		resData.ErrorDesc = SDKRes.ErrorDesc
		return resData
	}
	for i := range reqData.GetPrivateKeys() {
		if !keypair.CheckPrivateKey(reqData.GetPrivateKeys()[i]) {
			SDKRes := exception.GetSDKRes(exception.PRIVATEKEY_ONE_ERROR)
			resData.ErrorCode = SDKRes.ErrorCode
			resData.ErrorDesc = SDKRes.ErrorDesc
			return resData
		}
	}
	signatures := make([]model.Signature, len(reqData.GetPrivateKeys()))
	var err error
	for i := range reqData.GetPrivateKeys() {
		signatures[i].PublicKey, err = keypair.GetEncPublicKey(reqData.GetPrivateKeys()[i])
		if err != nil {
			SDKRes := exception.GetSDKRes(exception.GET_ENCPUBLICKEY_ERROR)
			resData.ErrorCode = SDKRes.ErrorCode
			resData.ErrorDesc = SDKRes.ErrorDesc
			return resData
		}
	}
	TransactionBlob, err := hex.DecodeString(reqData.GetBlob())
	if err != nil {
		SDKRes := exception.GetSDKRes(exception.INVALID_BLOB_ERROR)
		resData.ErrorCode = SDKRes.ErrorCode
		resData.ErrorDesc = SDKRes.ErrorDesc
		return resData
	}
	var transactionBlob protocol.Transaction
	err = proto.Unmarshal(TransactionBlob, &transactionBlob)
	if err != nil {
		resData.ErrorCode = exception.INVALID_BLOB_ERROR
		resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
		return resData
	}
	for i := range reqData.GetPrivateKeys() {
		signatures[i].SignData, err = signature.Sign(reqData.GetPrivateKeys()[i], TransactionBlob)
		if err != nil {
			SDKRes := exception.GetSDKRes(exception.SIGN_ERROR)
			resData.ErrorCode = SDKRes.ErrorCode
			resData.ErrorDesc = SDKRes.ErrorDesc
			return resData
		}
	}
	resData.Result.Signatures = signatures
	resData.ErrorCode = exception.SUCCESS
	return resData
}

// submit
func (transaction *TransactionOperation) Submit(reqData model.TransactionSubmitRequest) model.TransactionSubmitResponse {
	var resDatas model.TransactionSubmitData
	var resData model.TransactionSubmitResponse
	var reqDatas model.TransactionSubmitRequests
	reqDatas.Items = make([]model.TransactionSubmitRequest, 1)
	reqDatas.Items[0] = reqData
	if reqData.GetBlob() == "" {
		SDKRes := exception.GetSDKRes(exception.INVALID_BLOB_ERROR)
		resData.ErrorCode = SDKRes.ErrorCode
		resData.ErrorDesc = SDKRes.ErrorDesc
		return resData
	}
	TransactionBlob, err := hex.DecodeString(reqData.GetBlob())
	if err != nil {
		SDKRes := exception.GetSDKRes(exception.INVALID_BLOB_ERROR)
		resData.ErrorCode = SDKRes.ErrorCode
		resData.ErrorDesc = SDKRes.ErrorDesc
		return resData
	}
	var transactionBlob protocol.Transaction
	err = proto.Unmarshal(TransactionBlob, &transactionBlob)
	if err != nil {
		resData.ErrorCode = exception.INVALID_BLOB_ERROR
		resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
		return resData
	}
	for i := range reqDatas.Items {
		if reqDatas.Items[i].GetSignatures() == nil {
			resData.ErrorCode = exception.SIGNATURE_EMPTY_ERROR
			resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
			return resData
		}
		for j := range reqDatas.Items[i].GetSignatures() {
			if !keypair.CheckPublicKey(reqDatas.Items[i].GetSignatures()[j].PublicKey) || reqDatas.Items[i].GetSignatures()[j].SignData == "" {
				SDKRes := exception.GetSDKRes(exception.INVALID_BLOB_ERROR)
				resData.ErrorCode = SDKRes.ErrorCode
				resData.ErrorDesc = SDKRes.ErrorDesc
				return resData
			}
		}
	}
	requestJson, SDKRes := common.GetRequestJson(reqDatas)
	if SDKRes.ErrorCode != 0 {
		resData.ErrorCode = SDKRes.ErrorCode
		resData.ErrorDesc = SDKRes.ErrorDesc
		return resData
	}
	response, SDKRes := common.PostRequest(transaction.Url, "/submitTransaction", requestJson)
	if SDKRes.ErrorCode != 0 {
		resData.ErrorCode = SDKRes.ErrorCode
		resData.ErrorDesc = SDKRes.ErrorDesc
		return resData
	}
	defer response.Body.Close()
	if response.StatusCode == 200 {
		decoder := json.NewDecoder(response.Body)
		decoder.UseNumber()
		err := decoder.Decode(&resDatas)
		if err != nil {
			SDKRes := exception.GetSDKRes(exception.SYSTEM_ERROR)
			resData.ErrorCode = SDKRes.ErrorCode
			resData.ErrorDesc = SDKRes.ErrorDesc
			return resData
		}
		if resDatas.Results[0].ErrorCode == 0 {
			resData.Result.Hash = resDatas.Results[0].Hash
			return resData
		} else {
			resData.ErrorCode = resDatas.Results[0].ErrorCode
			resData.ErrorDesc = resDatas.Results[0].ErrorDesc
			return resData
		}
	} else {
		SDKRes := exception.GetSDKRes(exception.CONNECTNETWORK_ERROR)
		resData.ErrorCode = SDKRes.ErrorCode
		resData.ErrorDesc = SDKRes.ErrorDesc
		return resData
	}
}

// get info
func (transaction *TransactionOperation) GetInfo(reqData model.TransactionGetInfoRequest) model.TransactionGetInfoResponse {
	var resData model.TransactionGetInfoResponse
	if len(reqData.GetHash()) != 64 {
		SDKRes := exception.GetSDKRes(exception.INVALID_HASH_ERROR)
		resData.ErrorCode = SDKRes.ErrorCode
		resData.ErrorDesc = SDKRes.ErrorDesc
		return resData

	}
	get := "/getTransactionHistory?hash="
	response, SDKRes := common.GetRequest(transaction.Url, get, reqData.GetHash())
	if SDKRes.ErrorCode != 0 {
		resData.ErrorCode = SDKRes.ErrorCode
		resData.ErrorDesc = SDKRes.ErrorDesc
		return resData
	}
	defer response.Body.Close()
	if response.StatusCode == 200 {
		decoder := json.NewDecoder(response.Body)
		decoder.UseNumber()
		err := decoder.Decode(&resData)
		if err != nil {
			SDKRes := exception.GetSDKRes(exception.SYSTEM_ERROR)
			resData.ErrorCode = SDKRes.ErrorCode
			resData.ErrorDesc = SDKRes.ErrorDesc
			return resData
		}
		if resData.ErrorCode == 0 {
			for i := range resData.Result.Transactions {
				data, err := hex.DecodeString(resData.Result.Transactions[i].Transaction.Metadata)
				if err != nil {
					SDKRes := exception.GetSDKRes(exception.SYSTEM_ERROR)
					resData.ErrorCode = SDKRes.ErrorCode
					resData.ErrorDesc = SDKRes.ErrorDesc
					return resData
				}
				resData.Result.Transactions[i].Transaction.Metadata = string(data)
				for j := range resData.Result.Transactions[i].Transaction.Operations {
					data, err := hex.DecodeString(resData.Result.Transactions[i].Transaction.Operations[j].Metadata)
					if err != nil {
						SDKRes := exception.GetSDKRes(exception.SYSTEM_ERROR)
						resData.ErrorCode = SDKRes.ErrorCode
						resData.ErrorDesc = SDKRes.ErrorDesc
						return resData
					}
					resData.Result.Transactions[i].Transaction.Operations[j].Metadata = string(data)
				}
			}
			return resData
		} else {
			if resData.ErrorCode == 4 {
				resData.ErrorDesc = "Get Transaction failed"
				return resData
			}
			return resData

		}
	} else {
		SDKRes := exception.GetSDKRes(exception.CONNECTNETWORK_ERROR)
		resData.ErrorCode = SDKRes.ErrorCode
		resData.ErrorDesc = SDKRes.ErrorDesc
		return resData

	}
}

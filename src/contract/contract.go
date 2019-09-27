// contract
package contract

import (
	"encoding/json"
	"strings"

	"github.com/bumoproject/bumo-sdk-go/src/account"
	"github.com/bumoproject/bumo-sdk-go/src/blockchain"
	"github.com/bumoproject/bumo-sdk-go/src/common"
	"github.com/bumoproject/bumo-sdk-go/src/crypto/keypair"
	"github.com/bumoproject/bumo-sdk-go/src/exception"
	"github.com/bumoproject/bumo-sdk-go/src/model"
)

type ContractOperation struct {
	Url string
}

//Check Valid
func (contract *ContractOperation) CheckValid(reqData model.ContractCheckValidRequest) model.ContractCheckValidResponse {
	var Account account.AccountOperation
	Account.Url = contract.Url
	var reqDataAcc model.AccountGetInfoRequest
	var resData model.ContractCheckValidResponse
	resData.Result.IsValid = false
	reqDataAcc.SetAddress(reqData.GetAddress())
	if !keypair.CheckAddress(reqData.GetAddress()) {
		resData.ErrorCode = exception.INVALID_CONTRACTADDRESS_ERROR
		resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
		return resData
	}
	resDataAcc := Account.GetInfo(reqDataAcc)
	if resDataAcc.ErrorCode != 0 {
		resData.ErrorCode = resDataAcc.ErrorCode
		resData.ErrorDesc = resDataAcc.ErrorDesc
		return resData
	}
	if resDataAcc.Result.Priv.MasterWeight == 0 && resDataAcc.Result.Priv.Thresholds.TxThreshold == 1 && len(resDataAcc.Result.Contract.Payload) != 0 {
		resData.Result.IsValid = true
		return resData
	} else {
		resData.Result.IsValid = false
		return resData
	}
}

//Get Info
func (contract *ContractOperation) GetInfo(reqData model.ContractGetInfoRequest) model.ContractGetInfoResponse {
	var resData model.ContractGetInfoResponse
	var reqDataCheck model.ContractCheckValidRequest
	reqDataCheck.SetAddress(reqData.GetAddress())
	resDataCheck := contract.CheckValid(reqDataCheck)
	if resDataCheck.ErrorCode != 0 {
		resData.ErrorCode = resDataCheck.ErrorCode
		resData.ErrorDesc = resDataCheck.ErrorDesc
		return resData
	}
	if resDataCheck.Result.IsValid == false {
		resData.ErrorCode = exception.INVALID_CONTRACTADDRESS_ERROR
		resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
		return resData
	}
	get := "/getAccount?address="
	response, SDKRes := common.GetRequest(contract.Url, get, reqData.GetAddress())
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
			resData.ErrorCode = exception.SYSTEM_ERROR
			resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
			return resData
		}
		if resData.ErrorCode == 0 {
			resData.ErrorCode = exception.SUCCESS
			return resData
		} else {
			if resData.ErrorCode == 4 {
				resData.ErrorDesc = "Get account failed"
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

//Call
func (contract *ContractOperation) Call(reqData model.ContractCallRequest) model.ContractCallResponse {
	var resData model.ContractCallResponse
	if reqData.GetContractAddress() == "" && reqData.GetCode() == "" {
		resData.ErrorCode = exception.CONTRACTADDRESS_CODE_BOTH_NULL_ERROR
		resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
		return resData
	}
	if reqData.GetSourceAddress() != "" {
		if !keypair.CheckAddress(reqData.GetSourceAddress()) {
			resData.ErrorCode = exception.INVALID_SOURCEADDRESS_ERROR
			resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
			return resData
		}
	}
	if reqData.GetContractAddress() != "" {
		if !keypair.CheckAddress(reqData.GetContractAddress()) {
			resData.ErrorCode = exception.INVALID_CONTRACTADDRESS_ERROR
			resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
			return resData
		}
	}
	if reqData.GetOptType() < 0 || reqData.GetOptType() > 2 {
		resData.ErrorCode = exception.INVALID_OPTTYPE_ERROR
		resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
		return resData
	}
	callData := model.CallContractRequest{
		ContractAddress: reqData.GetContractAddress(),
		Code:            reqData.GetCode(),
		Input:           reqData.GetInput(),
		ContractBalance: reqData.GetContractBalance(),
		FeeLimit:        reqData.GetFeeLimit(),
		GasPrice:        reqData.GetGasPrice(),
		OptType:         reqData.GetOptType(),
		SourceAddress:   reqData.GetSourceAddress(),
	}
	reqDataByte, err := json.Marshal(callData)
	if err != nil {
		resData.ErrorCode = exception.SYSTEM_ERROR
		resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
		return resData
	}
	response, SDKRes := common.PostRequest(contract.Url, "/callContract", reqDataByte)
	if SDKRes.ErrorCode != 0 {
		resData.ErrorCode = exception.SYSTEM_ERROR
		resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
		return resData
	}
	defer response.Body.Close()
	if response.StatusCode == 200 {
		decoder := json.NewDecoder(response.Body)
		decoder.UseNumber()
		err := decoder.Decode(&resData)
		if err != nil {
			resData.ErrorCode = exception.SYSTEM_ERROR
			resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
			return resData
		}
		if resData.ErrorCode == 0 {
			return resData
		} else {
			resData.ErrorCode = resData.ErrorCode
			resData.ErrorDesc = resData.ErrorDesc
			return resData
		}
	} else {
		resData.ErrorCode = exception.CONNECTNETWORK_ERROR
		resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
		return resData
	}
}

//Get Address
func (contract *ContractOperation) GetAddress(reqData model.ContractGetAddressRequest) model.ContractGetAddressResponse {
	var resData model.ContractGetAddressResponse
	var Transaction blockchain.TransactionOperation
	Transaction.Url = contract.Url
	var reqDataInfo model.TransactionGetInfoRequest
	reqDataInfo.SetHash(reqData.GetHash())
	resDataInfo := Transaction.GetInfo(reqDataInfo)
	if resDataInfo.ErrorCode != 0 {
		resData.ErrorCode = resDataInfo.ErrorCode
		resData.ErrorDesc = resDataInfo.ErrorDesc
		return resData
	} else {
		if resDataInfo.Result.Transactions[0].ErrorCode != 0 {
			resData.ErrorCode = int(resDataInfo.Result.Transactions[0].ErrorCode)
			resData.ErrorDesc = resDataInfo.Result.Transactions[0].ErrorDesc
			return resData
		} else {
			if resDataInfo.Result.Transactions[0].ErrorDesc == "" {
				resData.ErrorCode = exception.QUERY_NO_RESULTS
				resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
				return resData
			}
			NewReader := strings.NewReader(resDataInfo.Result.Transactions[0].ErrorDesc)
			decoder := json.NewDecoder(NewReader)
			decoder.UseNumber()
			err := decoder.Decode(&resData.Result.ContractAddresInfos)
			if err != nil {
				resData.ErrorCode = exception.SYSTEM_ERROR
				resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
				return resData
			}
			return resData
		}
	}
}

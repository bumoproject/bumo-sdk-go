// Ctp10Token
package token

import (
	"encoding/json"
	"strconv"
	"strings"

	"github.com/bumoproject/bumo-sdk-go/src/account"
	"github.com/bumoproject/bumo-sdk-go/src/contract"
	"github.com/bumoproject/bumo-sdk-go/src/crypto/keypair"
	"github.com/bumoproject/bumo-sdk-go/src/exception"
	"github.com/bumoproject/bumo-sdk-go/src/model"
)

type Ctp10TokenOperation struct {
	Url string
}

//Check Valid
func (Ctp10Token *Ctp10TokenOperation) CheckValid(reqData model.Ctp10TokenCheckValidRequest) model.Ctp10TokenCheckValidResponse {
	var Contract contract.ContractOperation
	Contract.Url = Ctp10Token.Url
	var Account account.AccountOperation
	Account.Url = Ctp10Token.Url
	var resData model.Ctp10TokenCheckValidResponse
	resData.Result.IsValid = false
	var raqDataCheck model.ContractCheckValidRequest
	raqDataCheck.SetAddress(reqData.GetContractAddress())
	resDataCheck := Contract.CheckValid(raqDataCheck)
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
	var raqDataMetadata model.AccountGetMetadataRequest
	raqDataMetadata.SetAddress(reqData.GetContractAddress())
	raqDataMetadata.SetKey("global_attribute")
	rasDataMetadata := Account.GetMetadata(raqDataMetadata)
	if rasDataMetadata.ErrorCode == 0 {
		var data model.Params
		strReader := strings.NewReader(rasDataMetadata.Result.Metadatas[0].Value)
		decoder := json.NewDecoder(strReader)
		decoder.UseNumber()
		err := decoder.Decode(&data)
		if err != nil {
			SDKRes := exception.GetSDKRes(exception.SYSTEM_ERROR)
			resData.ErrorCode = SDKRes.ErrorCode
			resData.ErrorDesc = SDKRes.ErrorDesc
			return resData
		}
		balance, err := strconv.ParseInt(data.Balance, 10, 64)
		if err != nil {
			resData.ErrorCode = exception.INVALID_CONTRACTADDRESS_ERROR
			resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
			return resData
		}
		if balance <= 0 {
			resData.ErrorCode = exception.INVALID_CONTRACTADDRESS_ERROR
			resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
			return resData
		}
		if data.Decimals < 0 || data.Decimals > 8 {
			resData.ErrorCode = exception.INVALID_CONTRACTADDRESS_ERROR
			resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
			return resData
		}
		if len(data.Name) > 1024 {
			resData.ErrorCode = exception.INVALID_CONTRACTADDRESS_ERROR
			resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
			return resData
		}
		if len(data.Symbol) > 1024 {
			resData.ErrorCode = exception.INVALID_TOKEN_SIMBOL_ERROR
			resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
			return resData
		}
		totalSupply, err := strconv.ParseInt(data.TotalSupply, 10, 64)
		if err != nil {
			resData.ErrorCode = exception.INVALID_CONTRACTADDRESS_ERROR
			resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
			return resData
		}
		if totalSupply < 0 {
			resData.ErrorCode = exception.INVALID_CONTRACTADDRESS_ERROR
			resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
			return resData
		}
		if !keypair.CheckAddress(data.Ctp10TokenOwner) {
			resData.ErrorCode = exception.INVALID_CONTRACTADDRESS_ERROR
			resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
			return resData
		}
		if data.Ctp != "1.0" {
			resData.ErrorCode = exception.INVALID_CONTRACTADDRESS_ERROR
			resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
			return resData
		}
		resData.Result.IsValid = true
		return resData
	} else if rasDataMetadata.ErrorCode == exception.NO_METADATA_ERROR {
		resData.ErrorCode = exception.INVALID_CONTRACTADDRESS_ERROR
		resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
		return resData
	} else {
		resData.ErrorCode = rasDataMetadata.ErrorCode
		resData.ErrorDesc = rasDataMetadata.ErrorDesc
		return resData
	}
}

//Allowance
func (Ctp10Token *Ctp10TokenOperation) Allowance(reqData model.Ctp10TokenAllowanceRequest) model.Ctp10TokenAllowanceResponse {
	var resData model.Ctp10TokenAllowanceResponse
	var reqDataCheck model.Ctp10TokenCheckValidRequest
	reqDataCheck.SetContractAddress(reqData.GetContractAddress())
	resDataCheck := Ctp10Token.CheckValid(reqDataCheck)
	if resDataCheck.ErrorCode != 0 {
		resData.ErrorCode = exception.INVALID_CONTRACTADDRESS_ERROR
		resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
		return resData
	}
	if !keypair.CheckAddress(reqData.GetTokenOwner()) {
		resData.ErrorCode = exception.INVALID_TOKENOWNER_ERROR
		resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
		return resData
	}
	if !keypair.CheckAddress(reqData.GetSpender()) {
		resData.ErrorCode = exception.INVALID_SPENDER_ERROR
		resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
		return resData
	}
	var Contract contract.ContractOperation
	Contract.Url = Ctp10Token.Url
	var reqDataCall model.ContractCallRequest
	reqDataCall.SetContractAddress("buQXoNR24p2pPqnXPyiDprmTWsU4SYLtBNCG")
	reqDataCall.SetOptType(2)
	var Input model.Input
	Input.Method = "allowance"
	Input.Params.Owner = reqData.GetTokenOwner()
	Input.Params.Spender = reqData.GetSpender()
	InputByte, err := json.Marshal(Input)
	if err != nil {
		resData.ErrorCode = exception.SYSTEM_ERROR
		resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
		return resData
	}
	reqDataCall.SetInput(string(InputByte))
	resDataCall := Contract.Call(reqDataCall)
	if resDataCall.ErrorCode != 0 {
		resData.ErrorCode = resDataCall.ErrorCode
		resData.ErrorDesc = resDataCall.ErrorDesc
		return resData
	} else {
		if resDataCall.Result.QueryRets[0].Error.Data.Exception != "" {
			resData.ErrorCode = exception.GET_ALLOWANCE_ERROR
			resData.ErrorDesc = resDataCall.Result.QueryRets[0].Error.Data.Exception
			return resData
		}
		strReader := strings.NewReader(resDataCall.Result.QueryRets[0].Result.Value)
		decoder := json.NewDecoder(strReader)
		decoder.UseNumber()
		dataStr := make(map[string]string)
		err = decoder.Decode(&dataStr)
		if err != nil {
			resData.ErrorCode = exception.SYSTEM_ERROR
			resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
			return resData
		}
		resData.Result.Allowance, err = strconv.ParseInt(dataStr["allowance"], 10, 64)
		if err != nil {
			resData.ErrorCode = exception.SYSTEM_ERROR
			resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
			return resData
		}
		resData.ErrorCode = exception.SUCCESS
		return resData
	}
}

//Get Info
func (Ctp10Token *Ctp10TokenOperation) GetInfo(reqData model.Ctp10TokenGetInfoRequest) model.Ctp10TokenGetInfoResponse {
	var resData model.Ctp10TokenGetInfoResponse
	var reqDataCheck model.Ctp10TokenCheckValidRequest
	reqDataCheck.SetContractAddress(reqData.GetContractAddress())
	resDataCheck := Ctp10Token.CheckValid(reqDataCheck)
	if resDataCheck.ErrorCode != 0 {
		resData.ErrorCode = exception.INVALID_CONTRACTADDRESS_ERROR
		resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
		return resData
	}
	var Contract contract.ContractOperation
	Contract.Url = Ctp10Token.Url
	var reqDataCall model.ContractCallRequest
	reqDataCall.SetContractAddress("buQXoNR24p2pPqnXPyiDprmTWsU4SYLtBNCG")
	reqDataCall.SetOptType(2)
	var Input model.Input
	Input.Method = "contractInfo"
	InputByte, err := json.Marshal(Input)
	if err != nil {
		resData.ErrorCode = exception.SYSTEM_ERROR
		resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
		return resData
	}
	reqDataCall.SetInput(string(InputByte))
	resDataCall := Contract.Call(reqDataCall)
	if resDataCall.ErrorCode != 0 {
		resData.ErrorCode = resDataCall.ErrorCode
		resData.ErrorDesc = resDataCall.ErrorDesc
		return resData
	} else {
		if resDataCall.Result.QueryRets[0].Error.Data.Exception != "" {
			resData.ErrorCode = exception.GET_TOKEN_INFO_ERROR
			resData.ErrorDesc = resDataCall.Result.QueryRets[0].Error.Data.Exception
			return resData
		}
		strReader := strings.NewReader(resDataCall.Result.QueryRets[0].Result.Value)
		decoder := json.NewDecoder(strReader)
		decoder.UseNumber()
		var valueData model.Value
		err = decoder.Decode(&valueData)
		if err != nil {
			resData.ErrorCode = exception.SYSTEM_ERROR
			resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
			return resData
		}
		resData.Result.Decimals = valueData.ContractInfo.Decimals
		resData.Result.Name = valueData.ContractInfo.Name
		resData.Result.Symbol = valueData.ContractInfo.Symbol
		resData.Result.TotalSupply, err = strconv.ParseInt(valueData.ContractInfo.TotalSupply, 10, 64)
		if err != nil {
			resData.ErrorCode = exception.SYSTEM_ERROR
			resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
			return resData
		}
		resData.ErrorCode = exception.SUCCESS
		return resData
	}
}

//Get Name
func (Ctp10Token *Ctp10TokenOperation) GetName(reqData model.Ctp10TokenGetNameRequest) model.Ctp10TokenGetNameResponse {
	var resData model.Ctp10TokenGetNameResponse
	var reqDataCheck model.Ctp10TokenCheckValidRequest
	reqDataCheck.SetContractAddress(reqData.GetContractAddress())
	resDataCheck := Ctp10Token.CheckValid(reqDataCheck)
	if resDataCheck.ErrorCode != 0 {
		resData.ErrorCode = exception.INVALID_CONTRACTADDRESS_ERROR
		resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
		return resData
	}
	var Contract contract.ContractOperation
	Contract.Url = Ctp10Token.Url
	var reqDataCall model.ContractCallRequest
	reqDataCall.SetContractAddress("buQXoNR24p2pPqnXPyiDprmTWsU4SYLtBNCG")
	reqDataCall.SetOptType(2)
	var Input model.Input
	Input.Method = "name"
	InputByte, err := json.Marshal(Input)
	if err != nil {
		resData.ErrorCode = exception.SYSTEM_ERROR
		resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
		return resData
	}
	reqDataCall.SetInput(string(InputByte))
	resDataCall := Contract.Call(reqDataCall)
	if resDataCall.ErrorCode != 0 {
		resData.ErrorCode = resDataCall.ErrorCode
		resData.ErrorDesc = resDataCall.ErrorDesc
		return resData
	} else {
		if resDataCall.Result.QueryRets[0].Error.Data.Exception != "" {
			resData.ErrorCode = exception.GET_TOKEN_INFO_ERROR
			resData.ErrorDesc = resDataCall.Result.QueryRets[0].Error.Data.Exception
			return resData
		}
		strReader := strings.NewReader(resDataCall.Result.QueryRets[0].Result.Value)
		decoder := json.NewDecoder(strReader)
		decoder.UseNumber()
		valueData := make(map[string]string)
		err = decoder.Decode(&valueData)
		if err != nil {
			resData.ErrorCode = exception.SYSTEM_ERROR
			resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
			return resData
		}
		resData.Result.Name = valueData["name"]
		resData.ErrorCode = exception.SUCCESS
		return resData
	}
}

//Get Symbol
func (Ctp10Token *Ctp10TokenOperation) GetSymbol(reqData model.Ctp10TokenGetSymbolRequest) model.Ctp10TokenGetSymbolResponse {

	var resData model.Ctp10TokenGetSymbolResponse
	var reqDataCheck model.Ctp10TokenCheckValidRequest
	reqDataCheck.SetContractAddress(reqData.GetContractAddress())
	resDataCheck := Ctp10Token.CheckValid(reqDataCheck)
	if resDataCheck.ErrorCode != 0 {
		resData.ErrorCode = exception.INVALID_CONTRACTADDRESS_ERROR
		resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
		return resData
	}
	var Contract contract.ContractOperation
	Contract.Url = Ctp10Token.Url
	var reqDataCall model.ContractCallRequest
	reqDataCall.SetContractAddress("buQXoNR24p2pPqnXPyiDprmTWsU4SYLtBNCG")
	reqDataCall.SetOptType(2)
	var Input model.Input
	Input.Method = "symbol"
	InputByte, err := json.Marshal(Input)
	if err != nil {
		resData.ErrorCode = exception.SYSTEM_ERROR
		resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
		return resData
	}
	reqDataCall.SetInput(string(InputByte))
	resDataCall := Contract.Call(reqDataCall)
	if resDataCall.ErrorCode != 0 {
		resData.ErrorCode = resDataCall.ErrorCode
		resData.ErrorDesc = resDataCall.ErrorDesc
		return resData
	} else {
		if resDataCall.Result.QueryRets[0].Error.Data.Exception != "" {
			resData.ErrorCode = exception.GET_TOKEN_INFO_ERROR
			resData.ErrorDesc = resDataCall.Result.QueryRets[0].Error.Data.Exception
			return resData
		}
		strReader := strings.NewReader(resDataCall.Result.QueryRets[0].Result.Value)
		decoder := json.NewDecoder(strReader)
		decoder.UseNumber()
		valueData := make(map[string]string)
		err = decoder.Decode(&valueData)
		if err != nil {
			resData.ErrorCode = exception.SYSTEM_ERROR
			resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
			return resData
		}
		resData.Result.Symbol = valueData["symbol"]
		resData.ErrorCode = exception.SUCCESS
		return resData
	}
}

//Get Decimals
func (Ctp10Token *Ctp10TokenOperation) GetDecimals(reqData model.Ctp10TokenGetDecimalsRequest) model.Ctp10TokenGetDecimalsResponse {
	var resData model.Ctp10TokenGetDecimalsResponse
	var reqDataCheck model.Ctp10TokenCheckValidRequest
	reqDataCheck.SetContractAddress(reqData.GetContractAddress())
	resDataCheck := Ctp10Token.CheckValid(reqDataCheck)
	if resDataCheck.ErrorCode != 0 {
		resData.ErrorCode = exception.INVALID_CONTRACTADDRESS_ERROR
		resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
		return resData
	}
	var Contract contract.ContractOperation
	Contract.Url = Ctp10Token.Url
	var reqDataCall model.ContractCallRequest
	reqDataCall.SetContractAddress("buQXoNR24p2pPqnXPyiDprmTWsU4SYLtBNCG")
	reqDataCall.SetOptType(2)
	var Input model.Input
	Input.Method = "decimals"
	InputByte, err := json.Marshal(Input)
	if err != nil {
		resData.ErrorCode = exception.SYSTEM_ERROR
		resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
		return resData
	}
	reqDataCall.SetInput(string(InputByte))
	resDataCall := Contract.Call(reqDataCall)
	if resDataCall.ErrorCode != 0 {
		resData.ErrorCode = resDataCall.ErrorCode
		resData.ErrorDesc = resDataCall.ErrorDesc
		return resData
	} else {
		if resDataCall.Result.QueryRets[0].Error.Data.Exception != "" {
			resData.ErrorCode = exception.GET_TOKEN_INFO_ERROR
			resData.ErrorDesc = resDataCall.Result.QueryRets[0].Error.Data.Exception
			return resData
		}
		strReader := strings.NewReader(resDataCall.Result.QueryRets[0].Result.Value)
		decoder := json.NewDecoder(strReader)
		decoder.UseNumber()
		valueData := make(map[string]interface{})
		err = decoder.Decode(&valueData)
		if err != nil {
			resData.ErrorCode = exception.SYSTEM_ERROR
			resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
			return resData
		}
		str := valueData["decimals"].(json.Number)
		resData.Result.Decimals, err = strconv.ParseInt(string(str), 10, 64)
		if err != nil {
			resData.ErrorCode = exception.SYSTEM_ERROR
			resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
			return resData
		}
		resData.ErrorCode = exception.SUCCESS
		return resData
	}
}

//Get TotalSupply
func (Ctp10Token *Ctp10TokenOperation) GetTotalSupply(reqData model.Ctp10TokenGetTotalSupplyRequest) model.Ctp10TokenGetTotalSupplyResponse {
	var resData model.Ctp10TokenGetTotalSupplyResponse
	var reqDataCheck model.Ctp10TokenCheckValidRequest
	reqDataCheck.SetContractAddress(reqData.GetContractAddress())
	resDataCheck := Ctp10Token.CheckValid(reqDataCheck)
	if resDataCheck.ErrorCode != 0 {
		resData.ErrorCode = exception.INVALID_CONTRACTADDRESS_ERROR
		resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
		return resData
	}
	var Contract contract.ContractOperation
	Contract.Url = Ctp10Token.Url
	var reqDataCall model.ContractCallRequest
	reqDataCall.SetContractAddress("buQXoNR24p2pPqnXPyiDprmTWsU4SYLtBNCG")
	reqDataCall.SetOptType(2)
	var Input model.Input
	Input.Method = "totalSupply"
	InputByte, err := json.Marshal(Input)
	if err != nil {
		resData.ErrorCode = exception.SYSTEM_ERROR
		resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
		return resData
	}
	reqDataCall.SetInput(string(InputByte))
	resDataCall := Contract.Call(reqDataCall)
	if resDataCall.ErrorCode != 0 {
		resData.ErrorCode = resDataCall.ErrorCode
		resData.ErrorDesc = resDataCall.ErrorDesc
		return resData
	} else {
		if resDataCall.Result.QueryRets[0].Error.Data.Exception != "" {
			resData.ErrorCode = exception.GET_TOKEN_INFO_ERROR
			resData.ErrorDesc = resDataCall.Result.QueryRets[0].Error.Data.Exception
			return resData
		}
		strReader := strings.NewReader(resDataCall.Result.QueryRets[0].Result.Value)
		decoder := json.NewDecoder(strReader)
		decoder.UseNumber()
		valueData := make(map[string]interface{})
		err = decoder.Decode(&valueData)
		if err != nil {
			resData.ErrorCode = exception.SYSTEM_ERROR
			resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
			return resData
		}
		str := valueData["totalSupply"].(string)
		resData.Result.TotalSupply, err = strconv.ParseInt(str, 10, 64)
		if err != nil {
			resData.ErrorCode = exception.SYSTEM_ERROR
			resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
			return resData
		}
		resData.ErrorCode = exception.SUCCESS
		return resData
	}
}

//Get Balance
func (Ctp10Token *Ctp10TokenOperation) GetBalance(reqData model.Ctp10TokenGetBalanceRequest) model.Ctp10TokenGetBalanceResponse {
	var resData model.Ctp10TokenGetBalanceResponse
	var reqDataCheck model.Ctp10TokenCheckValidRequest
	reqDataCheck.SetContractAddress(reqData.GetContractAddress())
	resDataCheck := Ctp10Token.CheckValid(reqDataCheck)
	if resDataCheck.ErrorCode != 0 {
		resData.ErrorCode = exception.INVALID_CONTRACTADDRESS_ERROR
		resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
		return resData
	}
	if !keypair.CheckAddress(reqData.GetTokenOwner()) {
		resData.ErrorCode = exception.INVALID_TOKENOWNER_ERROR
		resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
		return resData
	}
	var Contract contract.ContractOperation
	Contract.Url = Ctp10Token.Url
	var reqDataCall model.ContractCallRequest
	reqDataCall.SetContractAddress("buQXoNR24p2pPqnXPyiDprmTWsU4SYLtBNCG")
	reqDataCall.SetOptType(2)
	var Input model.Input
	Input.Method = "balanceOf"
	Input.Params.Address = reqData.GetTokenOwner()
	InputByte, err := json.Marshal(Input)
	if err != nil {
		resData.ErrorCode = exception.SYSTEM_ERROR
		resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
		return resData
	}
	reqDataCall.SetInput(string(InputByte))
	resDataCall := Contract.Call(reqDataCall)
	if resDataCall.ErrorCode != 0 {
		resData.ErrorCode = resDataCall.ErrorCode
		resData.ErrorDesc = resDataCall.ErrorDesc
		return resData
	} else {
		if resDataCall.Result.QueryRets[0].Error.Data.Exception != "" {
			resData.ErrorCode = exception.GET_TOKEN_INFO_ERROR
			resData.ErrorDesc = resDataCall.Result.QueryRets[0].Error.Data.Exception
			return resData
		}
		strReader := strings.NewReader(resDataCall.Result.QueryRets[0].Result.Value)
		decoder := json.NewDecoder(strReader)
		decoder.UseNumber()
		valueData := make(map[string]interface{})
		err = decoder.Decode(&valueData)
		if err != nil {
			resData.ErrorCode = exception.SYSTEM_ERROR
			resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
			return resData
		}
		str := valueData["balance"].(string)
		resData.Result.Balance, err = strconv.ParseInt(str, 10, 64)
		if err != nil {
			resData.ErrorCode = exception.SYSTEM_ERROR
			resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
			return resData
		}
		resData.ErrorCode = exception.SUCCESS
		return resData
	}
}

// account
package account

import (
	"bytes"
	"encoding/json"

	"github.com/bumoproject/bumo-sdk-go/src/common"
	"github.com/bumoproject/bumo-sdk-go/src/crypto/keypair"
	"github.com/bumoproject/bumo-sdk-go/src/exception"
	"github.com/bumoproject/bumo-sdk-go/src/model"
)

type AccountOperation struct {
	Url string
}

// Check the validity of the address
func (account *AccountOperation) CheckValid(reqData model.AccountCheckValidRequest) model.AccountCheckValidResponse {
	var resData model.AccountCheckValidResponse
	resData.Result.IsValid = keypair.CheckAddress(reqData.GetAddress())
	resData.ErrorCode = exception.SUCCESS
	return resData
}

// Create public and private key pairs
func (account *AccountOperation) Create() model.AccountCreateResponse {
	var resData model.AccountCreateResponse
	var err error
	resData.Result.PublicKey, resData.Result.PrivateKey, resData.Result.Address, err = keypair.Create()
	if err != nil {
		resData.ErrorCode = exception.ACCOUNT_CREATE_ERROR
		resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
		return resData
	}
	resData.ErrorCode = exception.SUCCESS
	return resData
}

// Get account info
func (account *AccountOperation) GetInfo(reqData model.AccountGetInfoRequest) model.AccountGetInfoResponse {
	var resData model.AccountGetInfoResponse
	if !keypair.CheckAddress(reqData.GetAddress()) {
		resData.ErrorCode = exception.INVALID_ADDRESS_ERROR
		resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
		return resData
	}
	response, SDKRes := common.GetRequest(account.Url, "/getAccount?address=", reqData.GetAddress())
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
		resData.ErrorCode = exception.CONNECTNETWORK_ERROR
		resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
		return resData
	}
}

// Get Nonce
func (account *AccountOperation) GetNonce(reqData model.AccountGetNonceRequest) model.AccountGetNonceResponse {
	var resData model.AccountGetNonceResponse
	if !keypair.CheckAddress(reqData.GetAddress()) {
		SDKRes := exception.GetSDKRes(exception.INVALID_ADDRESS_ERROR)
		resData.ErrorCode = SDKRes.ErrorCode
		resData.ErrorDesc = SDKRes.ErrorDesc
		return resData
	}
	response, SDKRes := common.GetRequest(account.Url, "/getAccount?address=", reqData.GetAddress())
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

// Get Balance
func (account *AccountOperation) GetBalance(reqData model.AccountGetBalanceRequest) model.AccountGetBalanceResponse {
	var resData model.AccountGetBalanceResponse
	if !keypair.CheckAddress(reqData.GetAddress()) {
		SDKRes := exception.GetSDKRes(exception.INVALID_ADDRESS_ERROR)
		resData.ErrorCode = SDKRes.ErrorCode
		resData.ErrorDesc = SDKRes.ErrorDesc
		return resData
	}
	get := "/getAccount?address="
	response, SDKRes := common.GetRequest(account.Url, get, reqData.GetAddress())
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

// Get Assets
func (account *AccountOperation) GetAssets(reqData model.AccountGetAssetsRequest) model.AccountGetAssetsResponse {
	var resData model.AccountGetAssetsResponse
	if !keypair.CheckAddress(reqData.GetAddress()) {
		resData.ErrorCode = exception.INVALID_ADDRESS_ERROR
		resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
		return resData
	}
	get := "/getAccount?address="
	response, SDKRes := common.GetRequest(account.Url, get, reqData.GetAddress())
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
			if resData.Result.Assets == nil {
				resData.ErrorCode = exception.NO_ASSET_ERROR
				resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
				return resData
			}
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
		resData.ErrorCode = exception.CONNECTNETWORK_ERROR
		resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
		return resData

	}
}

// Get Metadata
func (account *AccountOperation) GetMetadata(reqData model.AccountGetMetadataRequest) model.AccountGetMetadataResponse {
	var resData model.AccountGetMetadataResponse
	if !keypair.CheckAddress(reqData.GetAddress()) {
		SDKRes := exception.GetSDKRes(exception.INVALID_ADDRESS_ERROR)
		resData.ErrorCode = SDKRes.ErrorCode
		resData.ErrorDesc = SDKRes.ErrorDesc
		return resData
	}
	if len(reqData.GetKey()) > 1024 {
		resData.ErrorCode = exception.INVALID_DATAKEY_ERROR
		resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
		return resData
	}
	get := "/getAccount?address="
	var buf bytes.Buffer
	buf.WriteString(reqData.GetAddress())
	buf.WriteString("&key=")
	buf.WriteString(reqData.GetKey())
	str := buf.String()
	response, SDKRes := common.GetRequest(account.Url, get, str)
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
			if resData.Result.Metadatas == nil {
				resData.ErrorCode = exception.NO_METADATA_ERROR
				resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
				return resData
			}
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

// Check Activated
func (account *AccountOperation) CheckActivated(reqData model.AccountCheckActivatedRequest) model.AccountCheckActivatedResponse {
	var resData model.AccountCheckActivatedResponse
	resData.Result.IsActivated = false
	var reqDataInfo model.AccountGetInfoRequest
	reqDataInfo.SetAddress(reqData.GetAddress())
	resDataInfo := account.GetInfo(reqDataInfo)
	if resDataInfo.ErrorCode == 0 {
		resData.Result.IsActivated = true
		return resData
	} else if resDataInfo.ErrorCode == 4 {
		return resData
	} else {
		resData.ErrorCode = resDataInfo.ErrorCode
		resData.ErrorDesc = resDataInfo.ErrorDesc
		return resData
	}
}

// asset
package token

import (
	"bytes"
	"encoding/json"

	"github.com/bumoproject/bumo-sdk-go/src/common"
	"github.com/bumoproject/bumo-sdk-go/src/crypto/keypair"
	"github.com/bumoproject/bumo-sdk-go/src/exception"
	"github.com/bumoproject/bumo-sdk-go/src/model"
)

type TokenOperation struct {
	Asset      AssetOperation
	Ctp10Token Ctp10TokenOperation
}
type AssetOperation struct {
	Url string
}

//获取账户指定资产数量
func (asset *AssetOperation) GetInfo(reqData model.AssetGetInfoRequest) model.AssetGetInfoResponse {
	var resData model.AssetGetInfoResponse
	if !keypair.CheckAddress(reqData.GetAddress()) {
		resData.ErrorCode = exception.INVALID_ADDRESS_ERROR
		resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
		return resData
	}
	if len(reqData.GetCode()) > 64 || len(reqData.GetCode()) == 0 {
		resData.ErrorCode = exception.INVALID_ASSET_CODE_ERROR
		resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
		return resData
	}
	if !keypair.CheckAddress(reqData.GetIssuer()) {
		resData.ErrorCode = exception.INVALID_ISSUER_ADDRESS_ERROR
		resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
		return resData
	}
	get := "/getAccount?address="
	var buf bytes.Buffer
	buf.WriteString(reqData.GetAddress())
	buf.WriteString("&code=")
	buf.WriteString(reqData.GetCode())
	buf.WriteString("&issuer=")
	buf.WriteString(reqData.GetIssuer())
	str := buf.String()
	response, SDKRes := common.GetRequest(asset.Url, get, str)
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

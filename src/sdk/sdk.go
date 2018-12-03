// sdk
package sdk

import (
	_ "sync"

	"github.com/bumoproject/bumo-sdk-go/src/account"
	"github.com/bumoproject/bumo-sdk-go/src/blockchain"
	"github.com/bumoproject/bumo-sdk-go/src/common"
	"github.com/bumoproject/bumo-sdk-go/src/contract"
	"github.com/bumoproject/bumo-sdk-go/src/exception"
	"github.com/bumoproject/bumo-sdk-go/src/model"
	"github.com/bumoproject/bumo-sdk-go/src/token"
)

type Sdk struct {
	Account     account.AccountOperation
	Contract    contract.ContractOperation
	Transaction blockchain.TransactionOperation
	Block       blockchain.BlockOperation
	Token       token.TokenOperation
}

//Init
func (sdk *Sdk) Init(reqData model.SDKInitRequest) model.SDKInitResponse {
	var resData model.SDKInitResponse
	if reqData.GetUrl() == "" {
		resData.ErrorCode = exception.URL_EMPTY_ERROR
		resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
		return resData
	}
	get := "/hello"
	commonStruct := common.GetIns()
	commonStruct.Url = reqData.GetUrl()
	commonStruct.ConnectTimeout = reqData.GetConnectTimeout()
	commonStruct.ReadWriteTimeout = reqData.GetReadWriteTimeout()
	commonStruct.ChainId = reqData.GetChainId()
	response, SDKRes := common.GetRequest(get, "")
	if SDKRes.ErrorCode != 0 {
		resData.ErrorCode = SDKRes.ErrorCode
		resData.ErrorDesc = SDKRes.ErrorDesc
		return resData
	}
	defer response.Body.Close()
	if response.StatusCode != 200 {
		resData.ErrorCode = exception.URL_EMPTY_ERROR
		resData.ErrorDesc = exception.GetErrDesc(resData.ErrorCode)
		return resData
	}
	resData.ErrorCode = exception.SUCCESS
	return resData
}

//type configSdk struct {
//	Account     account.AccountOperation
//	Contract    contract.ContractOperation
//	Transaction blockchain.TransactionOperation
//	Block       blockchain.BlockOperation
//	Token       token.TokenOperation
//}

//var insSdk *configSdk
//var once sync.Once

//func GetInsSdk(reqData model.SDKInitRequest) *configSdk {
//	once.Do(func() {
//		insSdk = &configSdk{}
//		insSdk.Init(reqData)
//	})
//	return insSdk
//}

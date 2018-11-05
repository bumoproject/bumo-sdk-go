# Bumo Go SDK

## 概述
本文档简要概述Bumo Go SDK常用接口文档，方便开发者更方便地写入和查询BU区块链。

- [配置](#配置)
	- [包引用](#包引用)
- [名词解析](#名词解析)
- [参数与响应](#参数与响应)
    - [请求参数](#请求参数)
    - [响应结果](#响应结果)
- [使用方法](#使用方法)
    - [包导入](#包导入)
    - [生成SDK实例](#生成sdk实例)
    - [生成公私钥地址](#生成公私钥地址)
    - [有效性校验](#有效性校验)
    - [查询](#查询)
	- [提交交易](#提交交易)
		- [获取账户nonce值](#获取账户nonce值)
		- [构建操作](#构建操作)
		- [构建交易Blob](#构建交易blob)
		- [签名交易](#签名交易)
		- [广播交易](#广播交易)
- [账户服务](#账户服务)
    - [CheckValid](#checkvalid)
    - [Create](#create)
    - [GetInfo](#getinfo-account)
    - [GetNonce](#getnonce)
    - [GetBalance](#getbalance-account)
    - [GetAssets](#getassets)
    - [GetMetadata](#getmetadata)
- [资产服务](#资产服务)
	- [GetInfo](#getinfo-asset)
- [合约服务](#合约服务)
    - [GetInfo](#getinfo-contract)
- [交易服务](#交易服务)
    - [操作说明](#操作说明)
    - [BuildBlob](#buildblob)
    - [EvaluateFee](#evaluatefee)
    - [Sign](#sign)
    - [Submit](#submit)
    - [GetInfo](#getinfo-transaction)
- [区块服务](#区块服务)
    - [GetNumber](#getnumber)
    - [CheckStatus](#checkstatus)
    - [GetTransactions](#gettransactions)
    - [GetInfo](#getinfo-block)
    - [GetLatest](#getlatest)
    - [GetValidators](#getvalidators)
    - [GetLatestValidators](#getlatestvalidators)
    - [GetReward](#getreward)
    - [GetLatestReward](#getlatestreward)
    - [GetFees](#getfees)
    - [GetLatestFees](#getlatestfees)
- [错误码](#错误码)

## 配置

### 包引用
所依赖的golang包在src文件夹中寻找，依赖的golang包如下：

```
	//获取包
	go get github.com/bumoproject/bumo-sdk-go
```
## 名词解析

操作BU区块链： 向BU区块链写入或修改数据

提交交易： 向BU区块链发送修改的内容

查询BU区块链： 查询BU区块链中的数据

账户服务： 提供账户相关的有效性校验与查询接口

资产服务： 提供资产相关的查询接口

Ctp10Token服务： 提供合约资产相关的有效性校验与查询接口

合约服务： 提供合约相关的有效性校验与查询接口

交易服务： 提供交易相关的提交与查询接口

区块服务： 提供区块的查询接口

账户nonce值： 每个账户都维护一个序列号，用于用户提交交易时标识交易执行顺序的
## 参数与响应
#### 请求参数
> 请求参数的格式，是[类名][方法名]Request，比如Account.GetInfo()的请求参数是AccountGetInfoRequest。
请求参数的成员，是各个方法的入参的成员变量名。

例如：Account.GetInfo()的入参成员是address，那么AccountGetInfoRequest的结构如下：

```
type AccountGetInfoRequest struct {
	address string
}
```

#### 响应结果
响应结果的格式，包含错误码，错误描述和result，格式是[类名][方法名]Response。

例如Account.GetInfo()的结构体名是AccountGetInfoResponse：

```
type AccountGetInfoResponse struct {
	ErrorCode int
	ErrorDesc string
	Result	AccountGetInfoResult
}
```
说明：

(1) ErrorCode: 0表示无错误，大于0表示有错误

(2) ErrorDesc: 空表示无错误，有内容表示有错误

(3) Result: 返回结果的结构体，其中结构体的名称，格式是[类名][方法名]Result。
例如Account.GetNonce()的结构体名是AccountGetNonceResult：

```
type AccountGetNonceResult struct {
	Nonce int64
}
```


## 使用方法

> 这里介绍SDK的使用流程，首先需要生成SDK实例，然后调用相应服务的接口，其中服务包括账户服务、资产服务、合约服务、交易服务、区块服务，接口按使用分类分为生成公私钥地址接口、有效性校验接口、查询接口、提交交易相关接口.

### 包导入
>导入使用的包

```
import (

	"github.com/bumoproject/bumo-sdk-go/src/model"
	"github.com/bumoproject/bumo-sdk-go/src/sdk"
)
```

### 生成SDK实例
>初始化Sdk结构体

```
var testSdk sdk.Sdk
```
>调用SDK的接口Init

```
url := "http://seed1.bumotest.io:26002"
var reqData model.SDKInitRequest
reqData.SetUrl(url)
resData := testSdk.Init(reqData)
```
### 生成公私钥地址
>通过调用Account的Create生成账户，例如：

```
resData := testSdk.Account.Create()
```
### 有效性校验
此接口用于校验信息的有效性的，直接调用相应的接口即可，比如，校验账户地址有效性，调用如下：

```
//初始化传入参数
var reqData model.AccountCheckValidRequest
var address string = "buQtfFxpQP9JCFgmu4WBojBbEnVyQGaJDgGn"
reqData.SetAddress(address)
//调用接口检查
resData := testSdk.Account.CheckValid(reqData)
```
### 查询
调用相应的接口，例如：查询账户信息

```
//初始化传入参数
var reqData model.AccountGetInfoRequest
var address string = "buQtfFxpQP9JCFgmu4WBojBbEnVyQGaJDgGn"
reqData.SetAddress(address)
//调用接口查询
resData := testSdk.Account.GetInfo(reqData)
```

### 提交交易
> 提交交易的过程包括以下几步：获取账户nonce值，构建操作，构建交易Blob，签名交易和广播交易

#### 获取账户nonce值

开发者可自己维护各个账户nonce，在提交完一个交易后，自动递增1，这样可以在短时间内发送多笔交易，否则，必须等上一个交易执行完成后，账户的nonce值才会加1。接口调用如下：

```
// 初始化请求参数
var reqData model.AccountGetNonceRequest
var address string = "buQtfFxpQP9JCFgmu4WBojBbEnVyQGaJDgGn"
reqData.SetAddress(address)
resData := testSdk.Account.GetNonce(reqData)
// 调用GetNonce接口
resData := testSdk.Account.GetNonce(reqData)
```
#### 构建操作

> 这里的操作是指在交易中做的一些动作。 例如：构建发送BU操作BUSendOperation，调用如下：

```
var buSendOperation model.BUSendOperation
buSendOperation.Init()
var amount int64 = 100
var address string = "buQVU86Jm4FeRW4JcQTD9Rx9NkUkHikYGp6z"
buSendOperation.SetAmount(amount)
buSendOperation.SetDestAddress(address)
```

#### 构建交易Blob

> 该接口用于生成交易Blob串，接口调用如下：
> 注意：gasPrice和feeLimit的单位是MO，且 1 BU = 10^8 MO

```
//初始化传入参数
var reqDataBlob model.TransactionBuildBlobRequest
reqDataBlob.SetSourceAddress(surceAddress)
reqDataBlob.SetFeeLimit(feeLimit)
reqDataBlob.SetGasPrice(gasPrice)
reqDataBlob.SetNonce(senderNonce)
reqDataBlob.SetOperation(buSendOperation)
//调用BuildBlob接口
resDataBlob := testSdk.Transaction.BuildBlob(reqDataBlob)
```

#### 签名交易

> 该接口用于交易发起者使用私钥对交易进行签名。接口调用如下：

```
//初始化传入参数
PrivateKey := []string{"privbUPxs6QGkJaNdgWS2hisny6ytx1g833cD7V9C3YET9mJ25wdcq6h"}
var reqData model.TransactionSignRequest
reqData.SetBlob(resDataBlob.Result.Blob)
reqData.SetPrivateKeys(PrivateKey)
//调用Sign接口
resDataSign := testSdk.Transaction.Sign(reqData)
```
#### 广播交易

> 该接口用于向BU区块链发送交易，触发交易的执行。接口调用如下：

```
//初始化传入参数
var reqData model.TransactionSubmitRequest
reqData.SetBlob(resDataBlob.Result.Blob)
reqData.SetSignatures(resDataSign.Result.Signatures)
//调用Submit接口
resDataSubmit := testSdk.Transaction.Submit(reqData)
```

## 账户服务

账户服务主要是账户相关的接口，包括5个接口：CheckValid, GetInfo, GetNonce, GetBalance, GetAssets, GetMetadata。

#### CheckValid
> 接口说明

该接口用于检测账户地址的有效性
> 调用方法

CheckValid(model.AccountCheckValidRequest) model.AccountCheckValidResponse

> 请求参数

参数	|	 类型	 |	描述
----------- |------------|----------------
address	 |string	 |待检测的账户地址

> 响应数据

参数	|	类型	|	描述
----------- |------------|----------------
IsValid	 |string	 |账户地址是否有效

> 错误码

异常	|	 错误码|描述
-----------|-----------|--------
SYSTEM_ERROR |20000	 |System error

> 示例

```
var reqData model.AccountCheckValidRequest
address := "buQtfFxpQP9JCFgmu4WBojBbEnVyQGaJDgGn"
reqData.SetAddress(address)
resData := testSdk.Account.CheckValid(reqData)
if resData.ErrorCode == 0 {
	fmt.Println(resData.Result.IsValid)
}
```

#### Create
> 接口说明

生成公私钥对
> 调用方法

Create() model.AccountCreateResponse

> 响应数据

参数	|	类型	|	描述
----------- |------------|----------------
PrivateKey	 |string	 |私钥
PublicKey	 |string	 |公钥
Address	 |string	 |地址

> 示例

```
resData := testSdk.Account.Create()
if resData.ErrorCode == 0 {
	fmt.Println("Address:",resData.Result.Address)
	fmt.Println("PrivateKey:",resData.Result.PrivateKey)
	fmt.Println("PublicKey:",resData.Result.PublicKey)
}
```

#### GetInfo-Account

> 接口说明

查询账户信息

> 调用方法

GetInfo(model.AccountGetInfoRequest) model.AccountGetInfoResponse

> 请求参数

参数	|	类型	|	描述
----------- |------------|----------------
address	 |string	 |待检测的账户地址

> 响应数据

参数	|	 类型	|	描述
--------- |------------- |----------------
Address	|	string	 |	账户地址
Balance	|	int64	|	账户余额
Nonce	|	int64	|	账户交易序列号
Priv	|[Priv](#priv) |	账户权限

#### Priv
参数	|	类型	|	描述
-----------|------------ |----------------
MasterWeight |	 int64	|账户自身权重
Signers	 |[] [Signer](#signer)|签名者权重
Thresholds	 |[Threshold](#threshold)|	门限

#### Signer
参数	|	类型	|	描述
-----------|------------ |----------------
Address	 |string	|签名账户地址
Weight	 |int64	|签名账户权重

#### Threshold
参数	|	类型	|	描述
-----------|------------ |----------------
TxThreshold	 |	string	|交易默认门限
TypeThresholds|[TypeThreshold](#typethreshold)|不同类型交易的门限

#### TypeThreshold
参数	|	类型	|	描述
-----------|------------ |----------------
Type	 |	int64	|	操作类型
Threshold	|	int64	|	门限

> 错误码

异常	|	 错误码|描述
-----------|-----------|--------
INVALID_ADDRESS_ERROR|11006 |Invalid address
CONNECTNETWORK_ERROR|11007|Fail to Connect network
SYSTEM_ERROR |20000	 |System error

> 示例

```
var reqData model.AccountGetInfoRequest
var address string = "buQtfFxpQP9JCFgmu4WBojBbEnVyQGaJDgGn"
reqData.SetAddress(address)
resData := testSdk.Account.GetInfo(reqData)
if resData.ErrorCode == 0 {
	data, _ := json.Marshal(resData.Result)
	fmt.Println("Info:", string(data))
}
```
#### GetNonce
> 接口说明

该接口用于获取账户的nonce

> 调用方法

GetNonce(model.AccountGetNonceRequest) model.AccountGetNonceResponse

> 请求参数

参数	|	类型	|	描述
----------- |------------|----------------
address	 |string	 |待检测的账户地址

> 响应数据

参数	|	类型	|	描述
----------- |------------|----------------
Nonce	|int64	 |该账户的交易序列号

> 错误码

异常	|	 错误码|描述
-----------|-----------|--------
INVALID_ADDRESS_ERROR	|	 11006	 |	 Invalid address
CONNECTNETWORK_ERROR	|	 11007	|	 Fail to Connect network
SYSTEM_ERROR |20000	 |System error

> 示例

```
var reqData model.AccountGetNonceRequest
var address string = "buQtfFxpQP9JCFgmu4WBojBbEnVyQGaJDgGn"
reqData.SetAddress(address)
resData := testSdk.Account.GetNonce(reqData)
if resData.ErrorCode == 0 {
	fmt.Println(resData.Result.Nonce)
}
```


#### GetBalance-Account
> 接口说明

该接口用于获取账户的Balance

> 调用方法

GetBalance(model.AccountGetBalanceRequest) model.AccountGetBalanceResponse

> 请求参数

参数	|	类型	|	描述
----------- |------------|----------------
address	 |string	 |待检测的账户地址

> 响应数据

参数	|	类型	|	描述
----------- |------------|----------------
Balance	 |int64	 |该账户的余额

> 错误码

异常	|	 错误码|描述
-----------	|	 -----------	 |	 --------
INVALID_ADDRESS_ERROR	|	 11006	 |	 Invalid address
CONNECTNETWORK_ERROR	|	 11007	|	 Fail to Connect network
SYSTEM_ERROR |20000	 |System error


> 示例

```
var reqData model.AccountGetBalanceRequest
var address string = "buQtfFxpQP9JCFgmu4WBojBbEnVyQGaJDgGn"
reqData.SetAddress(address)
resData := testSdk.Account.GetBalance(reqData)
if resData.ErrorCode == 0 {
	fmt.Println("Balance", resData.Result.Balance)
}
```

#### GetAssets

> 接口说明

该接口用于获取账户的nonce

> 调用方法

GetAssets(model.AccountGetAssetsRequest) model.AccountGetAssetsResponse

> 请求参数

参数	|	类型	|	描述
----------- |------------|----------------
address	 |string	 |待检测的账户地址

> 响应数据

参数	|	类型	|	描述
----------- |------------|----------------
Assets	|[] [Asset](#asset) |账户资产

#### Asset

参数	|	类型	|	描述
----------- |------------|----------------
Key	|[Key](#key)|资产惟一标识
Amount	|int64	|资产数量

 #### Key

参数|	 类型	|	 描述
-------- |----------- |	-----------
Code	 |	string	|	资产编码，长度[1 64]
Issuer	|	string	|	资产发行账户地址

> 错误码

异常	|	 错误码|描述
-----------|-----------|--------
INVALID_ADDRESS_ERROR	|	 11006	 |	 Invalid address
CONNECTNETWORK_ERROR	|	 11007	|	 Fail to Connect network
SYSTEM_ERROR	 |	20000	 |	System error


> 示例

```
var reqData model.AccountGetAssetsRequest
var address string = "buQtfFxpQP9JCFgmu4WBojBbEnVyQGaJDgGn"
reqData.SetAddress(address)
resData := testSdk.Account.GetAssets(reqData)
if resData.ErrorCode == 0 {
	data, _ := json.Marshal(resData.Result.Assets)
	fmt.Println("Assets:", string(data))
}
```


#### GetMetadata
> 接口说明

获取账户的metadata信息

> 调用方法

GetMetadata(model.AccountGetMetadataRequest) model.AccountGetMetadataResponse

> 请求参数

参数|类型|	描述
-------- |-------- |----------------
address|string|待检测的账户地址
key	|string|选填，metadata关键字，长度[1, 1024]

> 响应数据

参数	|	 类型	|	描述
----------- |----------- |----------------
Metadatas	|[] [Metadata](#metadata)|账户

#### Metadata
参数	|	 类型	|	描述
----------- |----------- |----------------
Key	 |string	 |metadata的关键词
Value	|string	 |metadata的内容
Version	 |int64	|metadata的版本


> 错误码

异常	|	 错误码|描述|
-----------|-----------|--------
INVALID_ADDRESS_ERROR	 |	 11006	 |	 Invalid address
CONNECTNETWORK_ERROR	 |	 11007	 |	 Fail to Connect network
INVALID_DATAKEY_ERROR	 |	 11011	 |	 The length of key must be between 1 and 1024
SYSTEM_ERROR	 |	 20000	|	 System error

> 示例

```
var reqData model.AccountGetMetadataRequest
var address string = "buQemmMwmRQY1JkcU7w3nhruoX5N3j6C29uo"
reqData.SetAddress(address)
resData := testSdk.Account.GetMetadata(reqData)
if resData.ErrorCode == 0 {
	data, _ := json.Marshal(resData.Result.Metadatas)
	fmt.Println("Metadatas:", string(data))
}
```



## 资产服务

资产服务主要是资产相关的接口，目前有1个接口：GetInfo

#### GetInfo-Asset

> 接口说明

获取账户指定资产数量

> 调用方法

GetInfo(model.AssetGetInfoRequest) model.AssetGetInfoResponse

> 请求参数

参数	|	 类型	|	描述
-----------|------------|----------------
address	|string	|必填，待查询的账户地址
code	|string	|必填，资产编码，长度[1, 64]
issuer	|string	|必填，资产发行账户地址

> 响应数据

参数	|	 类型	|	描述
-----------|------------|----------------
Assets	|[] [Asset](#asset)|账户资产

> 错误码

异常	|	 错误码|描述|
-----------|-----------|--------
INVALID_ADDRESS_ERROR	|	11006	|	Invalid address
CONNECTNETWORK_ERROR	|	11007	|	Fail to Connect network
INVALID_ASSET_CODE_ERROR	|	11023	|	The length of asset code must be between 1 and 1024
INVALID_ISSUER_ADDRESS_ERROR	|	11027	|	Invalid issuer address
SYSTEM_ERROR	|	20000	|	System error

> 示例

```
var reqData model.AssetGetInfoRequest
var address string = "buQemmMwmRQY1JkcU7w3nhruoX5N3j6C29uo"
reqData.SetAddress(address)
reqData.SetIssuer("buQnc3AGCo6ycWJCce516MDbPHKjK7ywwkuo")
reqData.SetCode("HNC")
resData := testSdk.Token.Asset.GetInfo(reqData)
if resData.ErrorCode == 0 {
	data, _ := json.Marshal(resData.Result.Assets)
	fmt.Println("Assets:", string(data))
}
```

## 合约服务

合约服务主要是合约相关的接口,目前有1个接口:GetInfo

#### GetInfo-contract
> 接口说明

获取合约信息

> 调用方法

GetInfo(model.ContractGetInfoRequest) model.ContractGetInfoResponse

> 请求参数

参数|类型|描述|
----|------|-----|
contractAddress	|	string	|	必填，合约账户地址	|



> 响应数据

参数		|		 类型			|	描述	|
--------	|---------------|------------|
Type	|	int64	|	合约类型，默认0	|
Payload	|	string	|	合约代码	|


> 错误码

异常	|	 错误码|描述|
-----------|-----------|--------
INVALID_CONTRACTADDRESS_ERROR	|	11037	|	Invalid contract address	|
CONTRACTADDRESS_NOT_CONTRACTACCOUNT_ERROR	|	11038	|	contractAddress is not a contract account	|
CONNECTNETWORK_ERROR	|	11007	|	Fail to Connect network	|
SYSTEM_ERROR	|	20000	|	System error	|

> 示例

```
var reqData model.ContractGetInfoRequest
var address string = "buQfnVYgXuMo3rvCEpKA6SfRrDpaz8D8A9Ea"
reqData.SetAddress(address)
resData := testSdk.Contract.GetInfo(reqData)
if resData.ErrorCode == 0 {
	data, _ := json.Marshal(resData.Result.Contract)
	fmt.Println("Contract:", string(data))
}
```

## 交易服务

交易服务主要是交易相关的接口，目前有5个接口：BuildBlob, EvaluateFee, Sign, Submit, GetInfo。

其中调用BuildBlob之前需要构建一些操作，目前操作有16种，分别包括AccountActivateOperation，AccountSetMetadataOperation, AccountSetPrivilegeOperation, AssetIssueOperation, AssetSendOperation, BUSendOperation, Ctp10TokenIssueOperation, Ctp10TokenTransferOperation, Ctp10TokenTransferFromOperation, Ctp10TokenApproveOperation, Ctp10TokenAssignOperation, Ctp10TokenChangeOwnerOperation, ContractInvokeByAssetOperation, ContractInvokeByBUOperation, LogCreateOperation,ContractCreateOperation

### 操作说明

#### BaseOperation

操作对象，根据不同的操作生成，详情如下：

>AccountActivateOperation



成员变量	|	 类型|	描述
-------------|--------|--------------
sourceAddress|string|选填，操作源账户
destAddress|string|必填，目标账户地址
initBalance|int64|必填，初始化资产，大小[1, max(int64)]
metadata	|	string	|	选填，备注

> AccountSetMetadataOperation



成员变量	|	 类型|	描述
-------------|---------|-----------
sourceAddress|string|选填，操作源账户
key	|string|必填，metadata的关键词，长度[1, 1024]
value	|string|选填，metadata的内容，长度[0, 256K]
version	|int64	|选填，metadata的版本
deleteFlag	|bool	|选填，是否删除metadata
metadata|string|选填，备注

> AccountSetPrivilegeOperation



成员变量	|	 类型|	描述
-------------|---------|--------------------------
sourceAddress|string|选填，操作源账户
masterWeight	|	string	|	选填，账户自身权重，大小[0, max(uint32)]
signers	|	[] [Signer](#signer)	|	选填，签名者权重列表
txThreshold	|	string	|	选填，交易门限，大小[0, max(int64)]
typeThreshold	|	[TypeThreshold](#typethreshold)	|	选填，指定类型交易门限
metadata	|	string	|	选填，备注

> AssetIssueOperation



成员变量	|	 类型|	描述
-------------|---------|------------------------
sourceAddress	|	string	|	选填，发起该操作的源账户地址
code	|	string	|	必填，资产编码，长度[1 64]
amount	|	int64	|	必填，资产发行数量，大小[1, max(int64)]
metadata	|	string	|	选填，备注


> AssetSendOperation

成员变量	|	 类型|	描述
-------------|---------|----------------------
sourceAddress	|	string	|	选填，发起该操作的源账户地址
destAddress	|	string	|	必填，目标账户地址
code	|	string	|	必填，资产编码，长度[1 64]
issuer	|	string	|	必填，资产发行账户地址
amount	|	int64	|	必填，资产数量，大小[ 0, max(int64)]
metadata	|	string	|	选填，备注

> BUSendOperation

成员变量	|	 类型|	描述
-------------|---------|---------------------
sourceAddress	|	string	|	选填，发起该操作的源账户地址
destAddress	|	string	|	必填，目标账户地址
amount	|	int64	|	必填，资产发行数量，大小[0, max(int64)]
metadata	|	string	|	选填，备注

> Ctp10TokenIssueOperation

成员变量	|	 类型|	描述
-------------|---------|---------------------
sourceAddress	|	string	|	选填，发起该操作的源账户地址
initBalance	|	int64	|	必填，给合约账户的初始化资产，大小[1, max(int64)]
name	|	string	|	必填，token名称，长度[1, 1024]
symbol	|	string	|	必填，token符号，长度[1, 1024]
decimals	|	int64	|	必填，token数量的精度，大小[0, 8]
supply	|	int64	|	必填，token发行的总供应量，大小[1, max(int64)]
metadata	|	string	|	选填，备注

> Ctp10TokenTransferOperation

成员变量	|	 类型|	描述
-------------|---------|---------------------
sourceAddress	|	string	|	选填，发起该操作的源账户地址
contractAddress	|	string	|	必填，合约账户地址
destAddress	|	string	|	必填，待转移的目标账户地址
amount	|	int64	|	必填，待转移的token数量，大小[1, max(int64)]
metadata	|	string	|	选填，备注

> Ctp10TokenTransferFromOperation

成员变量	|	 类型|	描述
-------------|---------|---------------------
sourceAddress	|	string	|	选填，发起该操作的源账户地址
contractAddress	|	string	|	必填，合约账户地址
fromAddress	|	string	|	必填，待转移的源账户地址
destAddress	|	string	|	必填，待转移的目标账户地址
amount	|	int64	|	必填，待转移的token数量，大小[1, max(int64)]
metadata	|	string	|	选填，备注

> Ctp10TokenApproveOperation

成员变量	|	 类型|	描述
-------------|---------|---------------------
sourceAddress	|	string	|	选填，发起该操作的源账户地址
contractAddress	|	string	|	必填，合约账户地址
spender	|	string	|	必填，授权的账户地址
amount	|	int64	|	必填，被授权的待转移的token数量，大小[1, max(int64)]
metadata	|	string	|	选填，备注

> Ctp10TokenAssignOperation

成员变量	|	 类型|	描述
-------------|---------|---------------------
sourceAddress	|	string	|	选填，发起该操作的源账户地址
contractAddress	|	string	|	必填，合约账户地址
destAddress	|	string	|	必填，待分配的目标账户地址
amount	|	int64	|	必填，待分配的token数量，大小[1, max(int64)]
metadata	|	string	|	选填，备注

> Ctp10TokenChangeOwnerOperation

成员变量	|	 类型|	描述
-------------|---------|---------------------
sourceAddress	|	string	|	选填，发起该操作的源账户地址
contractAddress	|	string	|	必填，合约账户地址
tokenOwner	|	string	|	必填，待分配的目标账户地址
metadata	|	string	|	选填，备注

> ContractCreateOperation

成员变量	|	 类型|	描述
-------------|---------|---------------------
sourceAddress	|	string	|	选填，发起该操作的源账户地址
initBalance	|	int64	|	必填，给合约账户的初始化资产，大小[1, max(int64)]
initInput	|	string	|	选填，对应的合约初始化参数
payload	|	string	|	必填，对应的合约代码
metadata	|	string	|	选填，备注

> ContractInvokeByAssetOperation

成员变量	|	 类型|	描述
-------------|---------|---------------------
sourceAddress	|	string	|	选填，发起该操作的源账户地址
contractAddress	|	string	|	必填，合约账户地址
code	|	string	|	选填，资产编码，长度[0, 64]，当为null时，仅触发合约
issuer	|	string	|	选填，资产发行账户地址，当为null时，仅触发合约
amount	|	int64	|	选填，资产数量，大小[0, max(int64)]，当是0时，仅触发合约
input	|	string	|	选填，待触发的合约的main()入参
metadata	|	string	|	选填，备注

> ContractInvokeByBUOperation

成员变量	|	 类型|	描述
-------------|---------|---------------------
sourceAddress	|	string	|	选填，发起该操作的源账户地址
contractAddress	|	string	|	必填，合约账户地址
amount	|	int64	|	选填，资产发行数量，大小[0, max(int64)]，当0时仅触发合约
input	|	string	|	选填，待触发的合约的main()入参
metadata	|	string	|	选填，备注

> LogCreateOperation

成员变量	|	 类型|	描述
-------------|---------|---------------------
sourceAddress	|	string	|	选填，发起该操作的源账户地址
topic	|	string	|	必填，日志主题，长度[1, 128]
data	|	[]string	|	必填，日志内容，每个字符串长度[1, 1024]
metadata	|	string	|	选填，备注


#### EvaluateFee
> 接口说明

评估费用

> 调用方法

EvaluateFee(model.TransactionEvaluateFeeRequest) model.TransactionEvaluateFeeResponse

> 请求参数

参数	|	类型	|	描述
-----------|------------|----------------
sourceAddress	|	string	|	必填，发起该操作的源账户地址
nonce	|	int64	|	必填，待发起的交易序列号，大小[1, max(int64)]
operations	|	list.List	|	必填，待提交的操作列表，不能为空
signtureNumber	|	string	|	选填，待签名者的数量，默认是1，大小[1, max(int32)]
metadata	|	string	|	选填，备注
ceilLedgerSeq	|	int64	|	选填，距离当前区块高度指定差值的区块内执行的限制，当区块超出当时区块高度与所设差值的和后，交易执行失败。必须大于等于0，是0时不限制

> 响应数据

成员变量	|	 类型	|	描述	|
-----------|------------|----------------
FeeLimit	|	int64	|	交易费用
GasPrice	|	int64	|	打包费用

> 错误码

异常	|	 错误码|描述|
-----------|-----------|--------
INVALID_SOURCEADDRESS_ERROR	|	11002	|	Invalid sourceAddress
INVALID_NONCE_ERROR	|	11048	|	Nonce must be between 1 and max(int64)
INVALID_OPERATIONS_ERROR	|	11051	|	Operations cannot be resolved
OPERATIONS_ONE_ERROR	|	11053	|	One of operations error
INVALID_SIGNATURENUMBER_ERROR	|	11054	|	SignatureNumber must be between 1 and max(int32)
SYSTEM_ERROR	|	20000	|	System error

> 示例

```
var reqDataOperation model.BUSendOperation
reqDataOperation.Init()
var amount int64 = 100
reqDataOperation.SetAmount(amount)
var destAddress string = "buQVU86Jm4FeRW4JcQTD9Rx9NkUkHikYGp6z"
reqDataOperation.SetDestAddress(destAddress)

var reqDataEvaluate model.TransactionEvaluateFeeRequest
var sourceAddress string = "buQVU86Jm4FeRW4JcQTD9Rx9NkUkHikYGp6z"
reqDataEvaluate.SetSourceAddress(sourceAddress)
var nonce int64 = 88
reqDataEvaluate.SetNonce(nonce)
var signatureNumber string = "3"
reqDataEvaluate.SetSignatureNumber(signatureNumber)
var SetCeilLedgerSeq int64 = 50
reqDataEvaluate.SetCeilLedgerSeq(SetCeilLedgerSeq)
reqDataEvaluate.SetOperation(reqDataOperation)
resDataEvaluate := testSdk.Transaction.EvaluateFee(reqDataEvaluate)
if resDataEvaluate.ErrorCode == 0 {
	data, _ := json.Marshal(resDataEvaluate.Result)
	fmt.Println("Evaluate:", string(data))
}
```


#### BuildBlob
> 接口说明

该接口用于序列化交易，生成交易Blob串，便于网络传输

> 调用方法

BuildBlob(model.TransactionBuildBlobRequest) model.TransactionBuildBlobResponse

> 请求参数

参数	|	 类型	|	描述
-----------|------------|----------------
sourceAddress	|	string	|	必填，发起该操作的源账户地址
nonce	|	int64	|	必填，待发起的交易序列号，函数里+1，大小[1, max(int64)]
gasPrice	|	int64	|	必填，交易打包费用，单位MO，1 BU = 10^8 MO，大小[1000, max(int64)]
feeLimit	|	int64	|	必填，交易手续费，单位MO，1 BU = 10^8 MO，大小[1, max(int64)]
operations	|	list.List	|	必填，待提交的操作列表，不能为空
ceilLedgerSeq	|	int64	|	选填，距离当前区块高度指定差值的区块内执行的限制，当区块超出当时区块高度与所设差值的和后，交易执行失败。必须大于等于0，是0时不限制
metadata	|	string	|	选填，备注


> 响应数据

参数	|	类型	|	描述
-----------|------------|----------------
TransactionBlob	|	string	|	Transaction序列化后的16进制字符串

> 错误码

异常	|	 错误码|描述|
-----------|-----------|--------
INVALID_SOURCEADDRESS_ERROR	|	11002	|	Invalid sourceAddress
INVALID_NONCE_ERROR	|	11048	|	Nonce must be between 1 and max(int64)
INVALID_DESTADDRESS_ERROR	|	11003	|	Invalid destAddress
INVALID_INITBALANCE_ERROR	|	11004	|	InitBalance must be between 1 and max(int64)
SOURCEADDRESS_EQUAL_DESTADDRESS_ERROR	|	11005	|	SourceAddress cannot be equal to destAddress
INVALID_ISSUE_AMMOUNT_ERROR	|	11008	|	AssetAmount this will be issued must be between 1 and max(int64)
INVALID_DATAKEY_ERROR	|	11011	|	The length of key must be between 1 and 1024
INVALID_DATAVALUE_ERROR	|	11012	|	The length of value must be between 0 and 256000
INVALID_DATAVERSION_ERROR	|	11013	|	The version must be equal or bigger than 0
INVALID_MASTERWEIGHT _ERROR	|	11015	|	MasterWeight must be between 0 and max(uint32)
INVALID_SIGNER_ADDRESS_ERROR	|	11016	|	Invalid signer address
INVALID_SIGNER_WEIGHT _ERROR	|	11017	|	Signer weight must be between 0 and max(uint32)
INVALID_TX_THRESHOLD_ERROR	|	11018	|	TxThreshold must be between 0 and max(int64)
INVALID_OPERATION_TYPE_ERROR	|	11019	|	Operation type must be between 1 and 100
INVALID_TYPE_THRESHOLD_ERROR	|	11020	|	TypeThreshold must be between 0 and max(int64)
INVALID_ASSET_CODE _ERROR	|	11023	|	The length of key must be between 1 and 64
INVALID_ASSET_AMOUNT_ERROR	|	11024	|	AssetAmount must be between 0 and max(int64)
INVALID_BU_AMOUNT_ERROR	|	11026	|	BuAmount must be between 0 and max(int64)
INVALID_ISSUER_ADDRESS_ERROR	|	11027	|	Invalid issuer address
NO_SUCH_TOKEN_ERROR	|	11030	|	No such token
INVALID_TOKEN_NAME_ERROR	|	11031	|	The length of token name must be between 1 and 1024
INVALID_TOKEN_SYMBOL_ERROR	|	11032	|	The length of symbol must be between 1 and 1024
INVALID_TOKEN_DECIMALS_ERROR	|	11033	|	Decimals must be between 0 and 8
INVALID_TOKEN_TOTALSUPPLY_ERROR	|	11034	|	TotalSupply must be between 1 and max(int64)
INVALID_TOKENOWNER_ERRPR	|	11035	|	Invalid token owner
INVALID_CONTRACTADDRESS_ERROR	|	11037	|	Invalid contract address
CONTRACTADDRESS_NOT_CONTRACTACCOUNT_ERROR	|	11038	|	ContractAddress is not a contract account
INVALID_TOKEN_AMOUNT_ERROR	|	11039	|	Token amount must be between 1 and max(int64)
SOURCEADDRESS_EQUAL_CONTRACTADDRESS_ERROR	|	11040	|	SourceAddress cannot be equal to contractAddress
INVALID_FROMADDRESS_ERROR	|	11041	|	Invalid fromAddress
FROMADDRESS_EQUAL_DESTADDRESS_ERROR	|	11042	|	FromAddress cannot be equal to destAddress
INVALID_SPENDER_ERROR	|	11043	|	Invalid spender
PAYLOAD_EMPTY_ERROR	|	11044	|	Payload cannot be empty
INVALID_LOG_TOPIC _ERROR	|	11045	|	The length of key must be between 1 and 128
INVALID_LOG_DATA _ERROR	|	11046	|	The length of value must be between 1 and 1024
INVALID_CONTRACT_TYPE_ERROR	|	11047	|	Type must be equal or bigger than 0
INVALID_NONCE_ERROR	|	11048	|	Nonce must be between 1 and max(int64)
INVALID_ GASPRICE_ERROR	|	11049	|	GasPrice must be between 1000 and max(int64)
INVALID_FEELIMIT_ERROR	|	11050	|	FeeLimit must be between 1 and max(int64)
OPERATIONS_EMPTY_ERROR	|	11051	|	Operations cannot be empty
INVALID_CEILLEDGERSEQ_ERROR	|	11052	|	CeilLedgerSeq must be equal or bigger than 0
OPERATIONS_ONE_ERROR	|	11053	|	One of operations cannot be resolved
SYSTEM_ERROR	|	20000	|	System error


> 示例

```
var reqDataOperation model.BUSendOperation
reqDataOperation.Init()
var amount int64 = 100
var destAddress string = "buQVU86Jm4FeRW4JcQTD9Rx9NkUkHikYGp6z"
reqDataOperation.SetAmount(amount)
reqDataOperation.SetDestAddress(destAddress)

var reqDataBlob model.TransactionBuildBlobRequest
var sourceAddressBlob string = "buQemmMwmRQY1JkcU7w3nhruoX5N3j6C29uo"
reqDataBlob.SetSourceAddress(sourceAddressBlob)
var feeLimit int64 = 1000000000
reqDataBlob.SetFeeLimit(feeLimit)
var gasPrice int64 = 1000
reqDataBlob.SetGasPrice(gasPrice)
var nonce int64 = 88
reqDataBlob.SetNonce(nonce)
reqDataBlob.SetOperation(reqDataOperation)

resDataBlob := testSdk.Transaction.BuildBlob(reqDataBlob)
if resDataBlob.ErrorCode == 0 {
	fmt.Println("Blob:", resDataBlob.Result)
}
```

#### Sign
> 接口说明

该接口用于实现交易的签名

> 调用方法

Sign(model.TransactionSignRequest) model.TransactionSignResponse

> 请求参数

参数		|		类型	|	描述
-----------|------------|----------------
blob	|	string	|	必填，待签名的交易Blob
privateKeys	|	[] string	|	必填，私钥列表


> 响应数据

参数	|	类型	|	描述
-----------|------------|----------------
Signatures	|	[] [Signature](#signature)	|签名后的数据列表

#### Signature

成员变量	|	 类型	|	描述	|
-----------|------------|----------------
signData	|	int64	|	签名后数据
publicKey	|	int64	|	公钥

> 错误码

异常	|	 错误码|描述|
-----------|-----------|--------
INVALID_BLOB_ERROR	|	11056	|	Invalid blob
PRIVATEKEY_NULL_ERROR	|	11057	|	PrivateKeys cannot be empty
PRIVATEKEY_ONE_ERROR	|	11058	|	One of privateKeys error
GET_ENCPUBLICKEY_ERROR	|	14000	|	The function 'GetEncPublicKey' failed
SIGN_ERROR	|	14001	|	The function 'Sign' failed
SYSTEM_ERROR	|	20000	|	System error

> 示例

```
PrivateKey := []string{"privbUPxs6QGkJaNdgWS2hisny6ytx1g833cD7V9C3YET9mJ25wdcq6h"}
var reqData model.TransactionSignRequest
reqData.SetBlob(resDataBlob.Result.Blob)
reqData.SetPrivateKeys(PrivateKey)
resDataSign := testSdk.Transaction.Sign(reqData)
if resDataSign.ErrorCode == 0 {
	fmt.Println("Sign:", resDataSign.Result)
}
```

#### Submit
> 接口说明

提交

> 调用方法

Submit(model.TransactionSubmitRequest) model.TransactionSubmitResponse

> 请求参数

参数	|	类型	|	描述
-----------|------------|----------------
blob	|	string	|	必填，交易blob
signature	|	[] [Signature](#signature)	|	必填，签名列表

> 响应数据

参数	|	类型	|	描述
-----------|------------|----------------
Hash	|	string	|	交易hash

> 错误码

异常	|	 错误码|描述|
-----------|-----------|--------
INVALID_BLOB_ERROR	|	11052	|	Invalid blob
SYSTEM_ERROR	|	20000	|	System error

> 示例

```
var reqData model.TransactionSubmitRequest
reqData.SetBlob(resDataBlob.Result.Blob)
reqData.SetSignatures(resDataSign.Result.Signatures)
resDataSubmit := testSdk.Transaction.Submit(reqData.Result)
if resDataSubmit.ErrorCode == 0 {
	fmt.Println("Hash:", resDataSubmit.Result.Hash)
}
```

#### GetInfo-transaction
> 接口说明

根据hash查询交易

> 调用方法

GetInfo(model.TransactionGetInfoRequest) model.TransactionGetInfoResponse

> 请求参数

参数	|	类型	|	描述
-----------|------------|----------------
hash	|	string	|	交易hash

> 响应数据

参数	|	类型	|	描述
-----------|------------|----------------
TotalCount	|	int64	|	返回的总交易数
Transactions	|	[] [TransactionHistory](#transactionhistory)	|	交易内容

#### TransactionHistory

成员变量|	类型	|	描述
-----------|------------|----------------
ActualFee	|	string	|	交易实际费用
CloseTime	|	int64	|	交易关闭时间
ErrorCode	|	int64	|	交易错误码
ErrorDesc	|	string	|	交易描述
Hash	|	string	|	交易hash
LedgerSeq	|	int64	|	区块序列号
Transactions	|	[Transaction](#transaction)	|	交易内容列表
Signatures	|	[] [Signature](#signature)	|	签名列表
TxSize	|	int64	|	交易大小

#### Transaction

成员	|	 类型	|	描述	|
-----------|------------|----------------
SourceAddress	|	string	|	交易发起的源账户地址
FeeLimit	|	int64	|	交易费用
GasPrice	|	int64	|	交易打包费用
Nonce	|	int64	|	交易序列号
Operations	|	[] [Operation](#operation)	|	操作列表

#### ContractTrigger
成员	|	 类型	|	描述	|
-----------|------------|----------------
Transaction	|	[TriggerTransaction](#triggertransaction)	|	触发交易

#### Operation

成员	|	 类型	|	描述	|
-----------|------------|----------------
Type	|	int64	|	操作类型
SourceAddress	|	string	|	操作发起源账户地址
Metadata	|	string	|	备注
CreateAccount	|	[CreateAccount](#createaccount)	|	创建账户操作
IssueAsset	|	[IssueAsset](#issueasset)	|	发行资产操作
PayAsset	|	[PayAsset](#payasset)	|	转移资产操作
PayCoin	|	[PayCoin](#paycoin)	|	发送BU操作
SetMetadata	|	[SetMetadata](#setmetadata)	|	设置metadata操作
SetPrivilege	|	[SetPrivilege](#setprivilege)	|	设置账户权限操作
Log	|	[Log](#log)	|	记录日志

#### TriggerTransaction

成员	|	 类型	|	描述	|
-----------|------------|----------------
hash	|	string	|	交易hash

#### CreateAccount

成员	|	 类型	|	描述	|
-----------|------------|----------------
DestAddress	|	string	|	目标账户地址
Contract	|	[Contract](#contract)	|	合约信息
Priv	|	[Priv](#priv)	|	账户权限
Metadata	|	[] [Metadata](#metadata)	|	账户
InitBalance	|	int64	|	账户资产
InitInput	|	string	|	合约init函数的入参

#### Contract

成员	|	 类型	|	描述	|
-----------|------------|----------------
Type	|	int64	|	合约的语种，默认不赋值
Payload	|	string	|	对应语种的合约代码

#### Metadata

成员	|	 类型	|	描述	|
-----------|------------|----------------
Key	|	string	|	metadata的关键词
Value	|	string	|	metadata的内容
Version	|	int64	|	metadata的版本

#### IssueAsset

成员	|	 类型	|	描述	|
-----------|------------|----------------
Code	|	string	|	资产编码，长度[1 64]
Amount	|	int64	|	资产数量

#### PayAsset

成员	|	 类型	|	描述	|
-----------|------------|----------------
DestAddress	|	string	|	待转移的目标账户地址
Asset	|	[Asset](#asset)	|	账户资产
Input	|	string	|	合约main函数入参

#### PayCoin

成员	|	 类型	|	描述	|
-----------	|	 ------------	|	 ----------------	|
DestAddress	|	string	|	待转移的目标账户地址
Amount	|	int64	|	待转移的BU数量
Input	|	string	|	合约main函数入参

#### SetMetadata

成员	|	 类型	|	描述	|
-----------|------------|----------------
Key	|	string	|	metadata的关键词
Value	|	string	|	metadata的内容
Version	|	int64	|	metadata的版本
DeleteFlag	|	bool	|	是否删除metadata

#### SetPrivilege

成员	|	 类型	|	描述	|
-----------|------------|----------------
MasterWeight	|	string	|	账户自身权重，大小[0, max(uint32)]
Signers	|	[] [Signer](#signer)	|	签名者权重列表
TxThreshold	|	string	|	交易门限，大小[0, max(int64)]
TypeThreshold	|	[TypeThreshold](#typethreshold)	|	指定类型交易门限

#### Log

成员	|	 类型	|	描述	|
-----------|------------|----------------
Topic	|	string	|	日志主题
Data	|	[]string	|	日志内容

> 示例

```
var reqData model.TransactionGetInfoRequest
var hash string = "cd33ad1e033d6dfe3db3a1d29a55e190935d9d1ff40a138d777e9406ebe0fdb1"
reqData.SetHash(hash)
resData := testSdk.Transaction.GetInfo(reqData)
if resData.ErrorCode == 0 {
	data, _ := json.Marshal(resData.Result)
	fmt.Println("info:", string(data)
}
```


## 区块服务

区块服务主要是区块相关的接口，目前有11个接口：GetNumber, CheckStatus, GetTransactions , GetInfo, GetLatestInfo, GetValidators, GetLatestValidators, GetReward, GetLatestReward, GetFees, GetLatestFees。

#### GetNumber
> 接口说明

获取区块高度

> 调用方法

GetNumber() model.BlockGetNumberResponse
> 响应数据

参数		|		 类型			|	描述	|
--------	|---------------|------------|
BlockNumber	|	int64	|	最新的区块高度，对应底层字段seq	|


> 错误码

异常	|	 错误码|描述|
-----------|-----------|--------
CONNECTNETWORK_ERROR	|	11007	|	Fail to Connect network	|
SYSTEM_ERROR	|	20000	|	System error	|

> 示例

```
resData := testSdk.Block.GetNumber()
if resData.ErrorCode == 0 {
	fmt.Println("BlockNumber:", resData.Result.BlockNumber)
}
```


#### CheckStatus
> 接口说明

检查区块同步

> 调用方法

CheckStatus() model.BlockCheckStatusResponse


> 响应数据

参数		|		 类型			|	描述	|
--------	|---------------|------------|
IsSynchronous|bool|区块是否同步|


> 错误码

异常	|	 错误码|描述|
-----------|-----------|--------
CONNECTNETWORK_ERROR	|	11007	|	Fail to Connect network	|
SYSTEM_ERROR	|	20000	|	System error	|

> 示例

```
resData := testSdk.Block.CheckStatus()
if resData.ErrorCode == 0 {
	fmt.Println("IsSynchronous:", resData.Result.IsSynchronous)
}
```


#### GetTransactions
> 接口说明

根据高度查询交易

> 调用方法

GetTransactions(model.BlockGetTransactionRequest) model.BlockGetTransactionResponse

> 请求参数


参数	|	类型	|	描述
-----------|------------|----------------
blockNumber	|	int64	|	必填，待查询的区块高度

> 响应数据

参数	|	类型	|	描述
-----------|------------|----------------
TotalCount	|	int64	|	返回的总交易数
Transactions	|	[] [TransactionHistory](#transactionhistory)	|	交易内容


> 错误码

异常	|	 错误码|描述|
-----------|-----------|--------
INVALID_BLOCKNUMBER_ERROR	|	11060	|	BlockNumber must bigger than 0
CONNECTNETWORK_ERROR	|	11007	|	Fail to Connect network
SYSTEM_ERROR	|	20000	|	System error

> 示例

```
var reqData model.BlockGetTransactionRequest
var blockNumber int64 = 581283
reqData.SetBlockNumber(blockNumber)
resData := testSdk.Block.GetTransactions(reqData)
if resData.ErrorCode == 0 {
	data, _ := json.Marshal(resData.Result.Transactions)
	fmt.Println("Transactions:", string(data))
}
```


#### GetInfo-block
> 接口说明

获取区块信息

> 调用方法

GetInfo(model.BlockGetInfoRequest) model.BlockGetInfoResponse

> 请求参数

参数	|	类型	|	描述
--------|---------------|------------
blockNumber	|	int64	|	待查询的区块高度	|

> 响应数据

参数		|		 类型			|	描述	|
--------	|---------------|------------|
CloseTime	|	int64	|	区块关闭时间	|
Number	|	int64	|	区块高度	|
TxCount	|	int64	|	交易总量	|
Version	|	string	|	区块版本	|


> 错误码

异常	|	 错误码|描述|
-----------|-----------|--------
INVALID_BLOCKNUMBER_ERROR	|	11060	|	BlockNumber must bigger than 0	|
CONNECTNETWORK_ERROR	|	11007	|	Fail to Connect network	|
SYSTEM_ERROR	|	20000	|	System error	|

> 示例

```
var reqData model.BlockGetInfoRequest
var blockNumber int64 = 581283
reqData.SetBlockNumber(blockNumber)
resData := testSdk.Block.GetInfo(reqData)
if resData.ErrorCode == 0 {
	data, _ := json.Marshal(resData.Result.Header)
	fmt.Println("Header:", string(data))
}
```


#### GetLatest
> 接口说明

获取最新区块信息

> 调用方法

 GetLatest() model.BlockGetLatestResponse

> 响应数据

参数		|		 类型			|	描述	|
--------	|---------------|------------|
CloseTime	|	int64	|	区块关闭时间	|
Number	|	int64	|	区块高度	|
TxCount	|	int64	|	交易总量	|
Version	|	string	|	区块版本	|


> 错误码

异常	|	 错误码|描述|
-----------|-----------|--------
CONNECTNETWORK_ERROR	|	11007	|	Fail to Connect network	|
SYSTEM_ERROR	|	20000	|	System error	|

> 示例

```
resData := testSdk.Block.GetLatest()
if resData.ErrorCode == 0 {
	data, _ := json.Marshal(resData.Result.Header)
	fmt.Println("Header:", string(data))
}
```


#### GetValidators
> 接口说明

获取指定区块中所有验证节点数

> 调用方法

GetValidators(model.BlockGetValidatorsRequest) model.BlockGetValidatorsResponse

> 请求参数

参数	|	类型	|	描述
--------|---------------|------------
blockNumber	|	int64	|	待查询的区块高度	|

> 响应数据


参数	|	类型	|	描述
-----------|------------|----------------
validators|[] [ValidatorInfo](#validatorinfo)|验证节点列表

#### ValidatorInfo

参数|	 类型	|	描述	|
-----------|------------|----------------
Address	|	String	|	共识节点地址
PledgeCoinAmount	|	int64	|	验证节点押金

> 错误码

异常	|	 错误码|描述|
-----------|-----------|--------
INVALID_BLOCKNUMBER_ERROR	|	11060	|	BlockNumber must bigger than 0	|
CONNECTNETWORK_ERROR	|	11007	|	Fail to Connect network	|
SYSTEM_ERROR	|	20000	|	System error	|

> 示例

```
var reqData model.BlockGetValidatorsRequest
var blockNumber int64 = 581283
reqData.SetBlockNumber(blockNumber)
resData := testSdk.Block.GetValidators(reqData)
if resData.ErrorCode == 0 {
	data, _ := json.Marshal(resData.Result.Validators)
	fmt.Println("Validators:", string(data))
}
```


#### GetLatestValidators

> 接口说明

获取最新区块中所有验证节点数

> 调用方法

GetLatestValidators() model.BlockGetLatestValidatorsResponse


> 响应数据

参数	|	类型	|	描述
-----------|------------|----------------
validators|[] [ValidatorInfo](#validatorinfo)|验证节点列表

> 错误码

异常	|	 错误码|描述|
-----------|-----------|--------
INVALID_BLOCKNUMBER_ERROR	|	11060	|	BlockNumber must bigger than 0	|
CONNECTNETWORK_ERROR	|	11007	|	Fail to Connect network	|
SYSTEM_ERROR	|	20000	|	System error	|

> 示例

```
resData := testSdk.Block.GetLatestValidators()
if resData.ErrorCode == 0 {
	data, _ := json.Marshal(resData.Result.Validators)
	fmt.Println("Validators:", string(data))
}
```


#### GetReward
> 接口说明

获取指定区块中的区块奖励和验证节点奖励

> 调用方法

	GetReward(model.BlockGetRewardRequest) model.BlockGetRewardResponse

> 请求参数

参数	|	类型	|	描述
-----------|------------|----------------
blockNumber	|	int64	|	必填，待查询的区块高度

> 响应数据

参数	|	类型	|	描述
-----------|------------|----------------
BlockReward	|	int64	|	区块奖励数
ValidatorsReward	|	[] [ValidatorReward](#validatorreward)|	验证节点奖励情况

#### ValidatorReward

成员变量|	类型	|	描述
-----------|------------|----------------
Validator	|	String	|	验证节点地址
Reward	|	int64	|	验证节点奖励


> 错误码

异常	|	 错误码|描述|
-----------|-----------|--------
INVALID_BLOCKNUMBER_ERROR	|	11060	|	BlockNumber must bigger than 0
CONNECTNETWORK_ERROR	|	11007	|	Fail to Connect network
SYSTEM_ERROR	|	20000	|	System error
> 示例

```
var reqData model.BlockGetRewardRequest
var blockNumber int64 = 581283
reqData.SetBlockNumber(blockNumber)
resData := testSdk.Block.GetReward(reqData)
if resData.ErrorCode == 0 {
	fmt.Println("ValidatorsReward:", resData.Result.ValidatorsReward)
}
```


#### GetLatestReward
> 接口说明

获取最新区块中的区块奖励和验证节点奖励

> 调用方法

GetLatestReward() model.BlockGetLatestRewardResponse

> 响应数据

参数	|	类型	|	描述
-----------|------------|----------------
BlockReward	|	int64	|	区块奖励数
ValidatorsReward	|	[] [ValidatorReward](#validatorreward)|	验证节点奖励情况


> 错误码

异常	|	 错误码|描述|
-----------|-----------|--------
CONNECTNETWORK_ERROR	|	11007	|	Fail to Connect network
SYSTEM_ERROR	|	20000	|	System error

> 示例

```
resData := testSdk.Block.GetLatestReward()
if resData.ErrorCode == 0 {
	fmt.Println("ValidatorsReward:", resData.Result.ValidatorsReward)
}
```


#### GetFees
> 接口说明

获取指定区块中的账户最低资产限制和打包费用

> 调用方法

GetFees(model.BlockGetFeesRequest) model.BlockGetFeesResponse

> 请求参数

参数	|	类型	|	描述
-----------|------------|----------------
blockNumber	|	int64	|	必填，待查询的区块高度	|

> 响应数据

参数	|	类型	|	描述
-----------|------------|----------------
Fees	|	[Fees](#fees)	|	费用|

#### Fees

成员变量|	类型	|	描述
-----------|------------|----------------
BaseReserve	|	int64	|	账户最低资产限制|
GasPrice	|	int64	|	打包费用，单位MO，1 BU = 10^8 MO|

> 错误码

异常	|	 错误码|描述|
-----------|-----------|--------
INVALID_BLOCKNUMBER_ERROR	|	11060	|	BlockNumber must bigger than 0	|
CONNECTNETWORK_ERROR	|	11007	|	Fail to Connect network	|
SYSTEM_ERROR	|	20000	|	System error	|

> 示例

```
var reqData model.BlockGetFeesRequest
var blockNumber int64 = 581283
reqData.SetBlockNumber(blockNumber)
resData := testSdk.Block.GetFees(reqData)
if resData.ErrorCode == 0 {
	data, _ := json.Marshal(resData.Result.Fees)
	fmt.Println("Fees:", string(data))
}
```


#### GetLatestFees
> 接口说明

获取最新区块中的账户最低资产限制和打包费用

> 调用方法

GetLatestFees() model.BlockGetLatestFeesResponse

> 响应数据

参数	|	类型	|	描述
-----------|------------|----------------
Fees	|	[Fees](#fees)	|	费用

> 错误码

异常	|	 错误码|描述|
-----------|-----------|--------
CONNECTNETWORK_ERROR	|	11007	|	Fail to Connect network
SYSTEM_ERROR	|	20000	|	System error


> 示例

```
resData := testSdk.Block.GetLatestFees()
if resData.ErrorCode == 0 {
	data, _ := json.Marshal(resData.Result.Fees)
	fmt.Println("Fees:", string(data))
}
```





### 错误码
> 公共错误码信息

参数|描述
-----|-----
11001|Create account failed.
11002|Invalid sourceAddress.
11003|Invalid destAddress.
11004|InitBalance must be between 1 and max(int64).
11005|SourceAddress cannot be equal to destAddress.
11006|Invalid address.
11007|Fail to connect network.
11008|AssetAmount this will be issued mustbetween 1 and max(int64).
11009|The account does not have this asset
11010|The account does not have this metadata.
11011|The length of key must be between 1 and 1024.
11012|The length of value must be between 0 and 256k.
11013|The version must be bigger than and equal to 0.
11015|MasterWeight must be between 0 and max(uint32).
11016|Invalid signer address.
11017|Signer weight must be between 0 and max(uint32).
11018|TxThreshold must be between 0 and max(int64).
11019|Type of TypeThreshold is invalid.
11020|TypeThreshold must be between 0 and max(int64).
11023|The length of code must be between 1 and 64.
11024|AssetAmount must be between 0 and max(int64).
11026|BuAmount must be between 0 and max(int64).
11027|Invalid issuer address.
11030|The length of ctp must be between 1 and 64.
11031|The length of token name must be between 1 and 1024.
11032|The length of symbol must be between 1 and 1024.
11033|Decimals must be between 0 and 8.
11034|TotalSupply must be between 1 and max(int64).
11035|Invalid token owner.
11036|Fail to get allowance.
11037|Invalid contract address.
11038|contractAddress is not a contract account.
11039|Amount must be between 1 and max(int64).
11040|SourceAddress cannot be equal to contractAddress.
11041|Invalid fromAddress.
11042|FromAddress cannot be equal to destAddress.
11043|Invalid spender.
11045|The length of key must be between 1 and 128.
11046|The length of value must be between 1 and 1024.
11048|Nonce must be between 1 and max(int64).
11049|GasPrice must be between 1000 and max(int64).
11050|FeeLimit must be between 1 and max(int64).
11051|Operations cannot be empty.
11052|CeilLedgerSeq must be equal or bigger than 0.
11053|One of operations cannot be resolved.
11054|SignatureNumber must be between 1 and max(int32).
11055|Invalid transaction hash.
11056|Invalid blob.
11057|PrivateKeys cannot be empty.
11058|One of privateKeys is invalid.
11060|BlockNumber must be bigger than 0.
11062|Url cannot be empty.
11063|ContractAddress and code cannot be empty at the same time.
11064|OptType must be between 0 and 2.
11065|Fail to get allowance.
11067|The signatures cannot be empty.
11066|Fail to get token info.
20000|System error.


> Go错误码信息

参数|描述
-----|-----
14000|The function 'GetEncPublicKey' failed.
14001|The function 'Sign' failed.
14002|The parameter'payload' is invalid.
14003|The query failed.
14004|Query no results.

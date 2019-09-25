---
id: sdk_go
title: BUMO GO SDK
sidebar_label: GO
---

## 概述
本文档详细说明Bumo Go SDK常用接口文档, 使开发者更方便地操作和查询BU区块链。

## 包导入

go必须是1.10.1或更高版本。

项目所依赖的包在src文件夹中，获取包的方法如下：

```go
//获取包
go get github.com/bumoproject/bumo-sdk-go
```

## 请求参数与响应数据格式

本章节将详细介绍请求参数与响应数据的格式。

### 请求参数

接口的请求参数的类名，是**服务名** + **方法名** + **Request**，比如: 账户服务下的[getInfo](#getinfo)接口的请求参数格式是AccountGetInfoRequest。

请求参数的成员，是各个接口的入参的成员。例如：账户服务下的[getInfo](#getinfo)接口的入参成员是address，那么该接口的请求参数的完整结构如下：
```go
type AccountGetInfoRequest struct {
address string
}
```

### 响应数据

接口的响应数据的类名，是**服务名** + **方法名** + **Response**，比如：账户服务下的[getNonce](#getnonce)接口的响应数据格式是`AccountGetNonceResponse`。

响应数据的成员，包括错误码、错误描述和返回结果，比如账户服务下的[getInfo](#getinfo)接口的响应数据的成员如下：

```go
type AccountGetInfoResponse struct {
  ErrorCode int
  ErrorDesc string
  Result  AccountGetInfoResult
}
```

说明：
- errorCode: **错误码**。0表示无错误，大于0表示有错误

- errorDesc: 错误描述。

- result: 返回结果。一个结构体，其类名是**服务名** + **方法名** + **Result**，其成员是各个接口返回值的成员，例如：账户服务下的[getNonce](#getnonce)接口的结果类名是`AccountGetNonceResult`，成员有nonce, 完整结构如下：

```go
type AccountGetNonceResult struct {
  Nonce int64
}
```

## 使用方法

这里介绍SDK的使用流程，首先需要生成SDK实例，然后调用相应服务的接口，其中服务包括[账户服务](#账户服务)、[资产服务](#资产服务)、[合约服务](#合约服务)、[交易服务](#交易服务)、[区块服务](#区块服务)，接口按使用分类分为[生成公私钥地址](#生成公私钥地址)、[有效性校验](#有效性校验)、[查询](#查询)、[广播交易](#广播交易)相关接口

### 包引入

生成SDK实例之前导入使用的包：

```go 
import(
  "github.com/bumoproject/bumo-sdk-go/src/model"
  "github.com/bumoproject/bumo-sdk-go/src/sdk"

```

### 生成SDK实例

初始化SDK结构方法：

```go
var testSdk sdk.sdk
```

调用SDK的接口Init：

```go
url :="http://seed1.bumotest.io:26002"
var reqData model.SDKInitRequest
reqData.SetUrl(url)
resData := testSdk.Init(reqData)
```

### 生成公私钥地址

此接口生成BU区块链账户的公钥、私钥和地址，直接调用`Keypair.generator`接口即可，具体调用如下：
```go
resData :=testSdk.Account.Create()
```

### 有效性校验
此接口用于校验信息的有效性的，直接调用相应的接口即可，比如，校验账户地址有效性，具体调用如下：
```go
//初始化传入参数
var reqData model.AccountCheckValidRequest
var address string = "buQtfFxpQP9JCFgmu4WBojBbEnVyQGaJDgGn"
reqData.SetAddress(address)
//调用接口检查
resData := testSdk.Account.CheckValid(reqData)
```

### 查询
此接口用于查询BU区块链上的数据，直接调用相应的接口即可，比如，查询账户信息，具体调用如下：
```go
//初始化传入参数
var reqData model.AccountGetInfoRequest
var address string = "buQtfFxpQP9JCFgmu4WBojBbEnVyQGaJDgGn"
reqData.SetAddress(address)
//调用接口查询
resData := testSdk.Account.GetInfo(reqData)
```

### 广播交易
广播交易是指通过广播的方式发起交易。广播交易包括以下步骤：

1. [获取账户nonce值](#获取账户nonce值)
2. [构建操作](#构建操作)
3. [序列化交易](#序列化交易)
4. [签名交易](#签名交易)
5. [提交交易](#提交交易)

#### 获取账户nonce值

开发者可自己维护各个账户`nonce`，在提交完一个交易后，自动为nonce值递增1，这样可以在短时间内发送多笔交易，否则，必须等上一个交易执行完成后，账户的`nonce`值才会加1。接口详情请见[getNonce](#getnonce)，调用如下：
```go
//初始化请求参数
var reqData model.AccountGetNonceRequest
var address string = "buQtfFxpQP9JCFgmu4WBojBbEnVyQGaJDgGn"
reqData.SetAddress(address)
//调用GetNonce接口
resData := testSdk.Account.GetNonce(reqData)
```

#### 构建操作

这里的操作是指在交易中做的一些动作，便于序列化交易和评估费用。操作详情请见[操作](#操作)。例如，构建发送BU操作(`BUSendOperation`)，接口调用如下：
```go 
var buSendOperation model.BUSendOperation
buSendOperation.Init()
var amount int64 = 100
var address string = "buQVU86Jm4FeRW4JcQTD9Rx9NkUkHikYGp6z"
buSendOperation.SetAmount(amount)
buSendOperation.SetDestAddress(address)
```

#### 序列化交易

该接口用于序列化交易，并生成交易Blob串，便于网络传输。其中nonce和operation是上面接口得到的，接口详情请见[buildBlob](#buildblob)，调用如下：
```go 
//初始化传入参数
var reqDataBlob model.TransactionBuildBlobRequest
reqDataBlob.SetSourceAddress(sourceAddress)
reqDataBlob.SetFeeLimit(feeLimit)
reqDataBlob.SetGasPrice(gasPrice)
reqDataBlob.SetNonce(senderNonce)
reqDataBlob.SetOperation(buSendOperation)
//调用BuildBlob接口
resDataBlob := testSdk.Transaction.BuildBlob(reqDataBlob)
}
```

#### 签名交易

该接口用于交易发起者使用其账户私钥对交易进行签名。其中transactionBlob是上面接口得到的，接口详情请见[sign](#sign)，调用如下：
```go
//初始化传入参数
PrivateKey := []string{"privbUPxs6QGkJaNdgWS2hisny6ytx1g833cD7V9C3YET9mJ25wdcq6h"}
var reqData model.TransactionSignRequest
reqData.SetBlob(resDataBlob.Result.Blob)
reqData.SetPrivateKeys(PrivateKey)
//调用Sign接口
resDataSign := testSdk.Transaction.Sign(reqData)
}
```

#### 提交交易

该接口用于向BU区块链发送交易请求，触发交易的执行。其中transactionBlob和signResult是上面接口得到的，接口详情见[submit](#submit)调用如下：
```go
//初始化传入参数
var reqData model.TransactionSubmitRequest
reqData.SetBlob(resDataBlob.Result.Blob)
reqData.SetSignatures(resDataSign.Result.Signatures)
//调用Submit接口
resDataSubmit := testSdk.Transaction.Submit(reqData)
```

## 交易服务

交易服务提供交易相关的接口，目前有5个接口：`BuildBlob`、 `EvaluateFee`、`sign`、 `Submit`、 `GetInfo`。

### buildBlob

> **注意:** 调用**buildBlob**之前需要构建一些操作，详情见[操作](#操作)。

- **接口说明**

   该接口用于序列化交易，生成交易Blob串，便于网络传输

- **调用方法**

  `BuildBlob(model.TransactionBuildBlobRequest)model.TransactionBuildBlobResponse;`

- **请求参数**

   参数      |     类型     |        描述       
   ----------- | ------------ | ---------------- 
   sourceAddress|String|必填，发起该操作的源账户地址
   nonce|int64|必填，待发起的交易序列号，函数里+1，大小限制[1, max(int64)]
   gasPrice|int64|必填，交易燃料单价，单位MO，1 BU = 10^8 MO，大小限制[1000, max(int64)]
   feeLimit|int64|必填，交易要求的最低的手续费，单位MO，1 BU = 10^8 MO，大小限制[1, max(int64)]
   operation|`[]`BaseOperation|必填，待提交的操作列表，不能为空
   ceilLedgerSeq|int64|选填，距离当前区块高度指定差值的区块内执行的限制，当区块超出当时区块高度与所设差值的和后，交易执行失败。必须大于等于0，是0时不限制
   metadata|String|选填，备注

- **响应数据**

   参数      |     类型     |        描述       
   ----------- | ------------ | ---------------- 
   transactionBlob|String|Transaction序列化后的16进制字符串
   hash|String|交易hash

- **错误码**

   异常       |     错误码   |   描述  
   -----------  | ----------- | -------- 
   INVALID_SOURCEADDRESS_ERROR|11002|Invalid sourceAddress
   INVALID_NONCE_ERROR|11048|Nonce must be between 1 and max(int64)
   INVALID_DESTADDRESS_ERROR|11003|Invalid destAddress
   INVALID_INITBALANCE_ERROR|11004|InitBalance must be between 1 and max(int64) 
   SOURCEADDRESS_EQUAL_DESTADDRESS_ERROR|11005|SourceAddress cannot be equal to destAddress
   INVALID_ISSUE_AMMOUNT_ERROR|11008|AssetAmount this will be issued must be between 1 and max(int64)
   INVALID_DATAKEY_ERROR|11011|The length of key must be between 1 and 1024
   INVALID_DATAVALUE_ERROR|11012|The length of value must be between 0 and 256000
   INVALID_DATAVERSION_ERROR|11013|The version must be equal to or greater than 0 
   INVALID_MASTERWEIGHT _ERROR|11015|MasterWeight must be between 0 andmax(uint32)
   INVALID_SIGNER_ADDRESS_ERROR|11016|Invalid signer address
   INVALID_SIGNER_WEIGHT _ERROR|11017|Signer weight must be between 0 andmax(uint32)
   INVALID_TX_THRESHOLD_ERROR|11018|TxThreshold must be between 0 and max(int64)
   INVALID_OPERATION_TYPE_ERROR|11019|Operation type must be between 1 and 100
   INVALID_TYPE_THRESHOLD_ERROR|11020|TypeThreshold must be between 0 and max(int64)
   INVALID_ASSET_CODE _ERROR|11023|The length of asset code must be between 1 and 64
   INVALID_ASSET_AMOUNT_ERROR|11024|AssetAmount must be between 0 and max(int64)
   INVALID_BU_AMOUNT_ERROR|11026|BuAmount must be between 0 and max(int64)
   INVALID_ISSUER_ADDRESS_ERROR|11027|Invalid issuer address
   NO_SUCH_TOKEN_ERROR|11030|No such token
   INVALID_TOKEN_NAME_ERROR|11031|The length of token name must be between 1 and 1024
   INVALID_TOKEN_SYMBOL_ERROR|11032|The length of symbol must be between 1 and 1024
   INVALID_TOKEN_DECIMALS_ERROR|11033|Decimals must be between 0 and 8
   INVALID_TOKEN_TOTALSUPPLY_ERROR|11034|TotalSupply must be between 1 and max(int64)
   INVALID_TOKENOWNER_ERRPR|11035|Invalid token owner
   INVALID_CONTRACTADDRESS_ERROR|11037|Invalid contract address
   CONTRACTADDRESS_NOT_CONTRACTACCOUNT_ERROR|11038|ContractAddress is not a contract account
   INVALID_TOKEN_AMOUNT_ERROR|11039|Token amount must be between 1 and max(int64)
   SOURCEADDRESS_EQUAL_CONTRACTADDRESS_ERROR|11040|SourceAddress cannot be equal to contractAddress
   INVALID_FROMADDRESS_ERROR|11041|Invalid fromAddress
   FROMADDRESS_EQUAL_DESTADDRESS_ERROR|11042|FromAddress cannot be equal to destAddress
   INVALID_SPENDER_ERROR|11043|Invalid spender
   PAYLOAD_EMPTY_ERROR|11044|Payload cannot be empty
   INVALID_LOG_TOPIC_ERROR|11045|The length of a log topic must be between 1 and 128
   INVALID_LOG_DATA_ERROR|11046|The length of one piece of log data must be between 1 and 1024
   INVALID_CONTRACT_TYPE_ERROR|11047|Type must be equal or bigger than 0 
   INVALID_NONCE_ERROR|11048|Nonce must be between 1 and max(int64)
   INVALID_ GASPRICE_ERROR|11049|GasPrice must be between 0 and max(int64)
   INVALID_FEELIMIT_ERROR|11050|FeeLimit must be between 0 and max(int64)
   OPERATIONS_EMPTY_ERROR|11051|Operations cannot be empty
   INVALID_CEILLEDGERSEQ_ERROR|11052|CeilLedgerSeq must be equal to or greater than 0
   OPERATIONS_ONE_ERROR|11053|One of the operations cannot be resolved
   REQUEST_NULL_ERROR|12001|Request parameter cannot be null
   SYSTEM_ERROR|20000|System error

- **示例**

   ```go
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

### evaluateFee

- **接口说明**

   该接口实现交易的费用评估。

- **调用方法**

  `EvaluateFee(model.TransactionEvaluateFeeRequest)model.TransactionEvaluateFeeResponse;`

- **请求参数**

   参数      |     类型     |        描述       
   ----------- | ------------ | ---------------- 
   sourceAddress|String|必填，发起该操作的源账户地址
   nonce|int64|必填，待发起的交易序列号，大小限制[1, max(int64)]
   operation|`[]`BaseOperation|必填，待提交的操作列表，不能为空
   signtureNumber|String|选填，待签名者的数量，默认是1，大小限制[1, max(int32)]
   ceilLedgerSeq|int64|选填，距离当前区块高度指定差值的区块内执行的限制，当区块超出当时区块高度与所设差值的和后，交易执行失败。必须大于等于0，是0时不限制
   metadata|String|选填，备注

- **响应数据**

   参数      |     类型     |        描述       
   ----------- | ------------ | ---------------- 
   txs     |   `[]`[TestTx](#testtx)   |  评估交易集   

- **错误码**

   异常       |     错误码   |   描述  
   -----------  | ----------- | -------- 
   INVALID_SOURCEADDRESS_ERROR|11002|Invalid sourceAddress
   INVALID_NONCE_ERROR|11048|Nonce must be between 1 and max(int64)
   OPERATIONS_EMPTY_ERROR|11051|Operations cannot be empty
   OPERATIONS_ONE_ERROR|11053|One of the operations cannot be resolved
   INVALID_SIGNATURENUMBER_ERROR|11054|SignagureNumber must be between 1 and max(int32)
   SYSTEM_ERROR|20000|System error

- **示例**

   ```go
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

### sign

- **接口说明**

   该接口用于实现交易的签名

- **调用方法**

  `Sign(model.TransactionSignRequest) model.TransactionSignResponse;`

- **请求参数**

   参数      |     类型     |        描述       
   ----------- | ------------ | ---------------- 
   blob|String|必填，待签名的交易Blob
   privateKeys|`[]`String|必填，私钥列表


- **响应数据**

   参数      |     类型     |        描述       
   ----------- | ------------ | ---------------- 
   signatures|`[]`[Signature](#signature)|签名后的数据列表

- **错误码**

   异常       |     错误码   |   描述  
   -----------  | ----------- | -------- 
   INVALID_BLOB_ERROR|11056|Invalid blob
   PRIVATEKEY_NULL_ERROR|11057|PrivateKeys cannot be empty
   PRIVATEKEY_ONE_ERROR|11058|One of privateKeys is invalid
   GET_ENCPUBLICKEY_ERROR|14000|The function ‘GetEncPublicKey’ failed
   SIGN_ERROR|14001|The function ‘Sign’ failed
   SYSTEM_ERROR|20000|System error

- **示例**

   ```go
   PrivateKey := []string{"privbUPxs6QGkJaNdgWS2hisny6ytx1g833cD7V9C3YET9mJ25wdcq6h"}
   var reqData model.TransactionSignRequest
   reqData.SetBlob(resDataBlob.Result.Blob)
   reqData.SetPrivateKeys(PrivateKey)
   resDataSign := testSdk.Transaction.Sign(reqData)
   if resDataSign.ErrorCode == 0 {
      fmt.Println("Sign:", resDataSign.Result)
   }
   ```

### submit

- **接口说明**

   该接口用于实现交易的提交。

- **调用方法**

  `Submit(model.TransactionSubmitRequest) model.TransactionSubmitResponse;`

- **请求参数**

   参数      |     类型     |        描述       
   ----------- | ------------ | ---------------- 
    blob|String|必填，交易blob
    signature|`[]`[Signature](#signature)|必填，签名列表

- **响应数据**

   参数      |     类型     |        描述       
   ----------- | ------------ | ---------------- 
   hash|String|交易hash

- **错误码**

   异常       |     错误码   |   描述  
   -----------  | ----------- | -------- 
   INVALID_BLOB_ERROR|11056|Invalid blob
   SIGNATURE_EMPTY_ERROR|11067|The signatures cannot be empty
   SYSTEM_ERROR|20000|System error

- **示例**

   ```go
   var reqData model.TransactionSubmitRequest
   reqData.SetBlob(resDataBlob.Result.Blob)
   reqData.SetSignatures(resDataSign.Result.Signatures)
   resDataSubmit := testSdk.Transaction.Submit(reqData.Result)
   if resDataSubmit.ErrorCode == 0 {
      fmt.Println("Hash:", resDataSubmit.Result.Hash)
   }
   ```

### getInfo

- **接口说明**

   该接口用于实现根据交易hash查询交易。

- **调用方法**

  `GetInfo(model.TransactionGetInfoRequest)model.TransactionGetInfoResponse;`

- **请求参数**

   参数      |     类型     |        描述       
   ----------- | ------------ | ---------------- 
   hash|String|交易hash

- **响应数据**

   参数      |     类型     |        描述       
   ----------- | ------------ | ---------------- 
   totalCount|int64|返回的总交易数
   transactions|`[]`[TransactionHistory](#transactionhistory)|交易内容

- **错误码**

   异常       |     错误码   |   描述  
   -----------  | ----------- | -------- 
   INVALID_HASH_ERROR|11055|Invalid transaction hash
   CONNECTNETWORK_ERROR|11007|Failed to connect to the network
   SYSTEM_ERROR|20000|System error

- **示例**

   ```go 
   var reqData model.TransactionGetInfoRequest
   var hash string = "cd33ad1e033d6dfe3db3a1d29a55e190935d9d1ff40a138d777e9406ebe0fdb1"
   reqData.SetHash(hash)
   resData := testSdk.Transaction.GetInfo(reqData)
   if resData.ErrorCode == 0 {
      data, _ := json.Marshal(resData.Result)
      fmt.Println("info:", string(data)
   }
   ```


## 操作

操作是指在交易在要做的事情，在构建操作之前，需要构建操作。目前操作有10种，分别是 [AccountActivateOperation](#accountactivateoperation)、[AccountSetMetadataOperation](#accountsetmetadataoperation)、 [AccountSetPrivilegeOperation](#accountsetprivilegeoperation)、 [AssetIssueOperation](#assetissueoperation)、 [AssetSendOperation](#assetsendoperation)、 [BUSendOperation](#busendoperation)、 [ContractCreateOperation](#contractcreateoperation)、 [ContractInvokeByAssetOperation](#contractinvokebyassetoperation)、 [ContractInvokeByBUOperation](#contractinvokebybuoperation)、 [LogCreateOperation](#logcreateoperation)。

**BaseOperation**

BaseOperation是buildBlob接口中所有操作的基类。

成员变量    |     类型  |        描述                           
------------- | -------- | ----------------------------------   
sourceAddress |   String |  选填，操作源账户地址
metadata      |   String |  选填，备注

### AccountActivateOperation

- 功能

  该操作用于激活账户。AccountActivateOperation继承于BaseOperation。

- 费用

  FeeLimit目前(2018.07.26)固定是0.01 BU。

- 成员

   成员变量    |     类型  |        描述                           
   ------------- | -------- | ---------------------------------- 
   sourceAddress |   String |  选填，操作源账户地址 
   destAddress   |   String |  必填，目标账户地址                     
   initBalance   |   int64   |  必填，初始化资产，单位MO，1 BU = 10^8 MO, 大小(0, max(int64)] 
   metadata|String|选填，备注

### AccountSetMetadataOperation

- 功能

  该操作用于设置账户metadata。AccountSetMetadataOperation继承于BaseOperation。

- 费用

  FeeLimit目前(2018.07.26)固定是0.01 BU。

- 成员

   成员变量    |     类型   |        描述                         
   ------------- | --------- | ------------------------------- 
   sourceAddress |   String |  选填，操作源账户地址
   key           |   String  |  必填，metadata的关键词，长度限制[1, 1024]
   value         |   String  |  必填，metadata的内容，长度限制[0, 256000]
   version       |   int64   |  选填，metadata的版本
   deleteFlag    |   Boolean |  选填，是否删除metadata
   metadata|String|选填，备注           

### AccountSetPrivilegeOperation

- 功能

  该操作用于设置账户权限。AccountSetPrivilegeOperation继承于BaseOperation。

- 费用

  feeLimit目前(2018.07.26)固定是0.01 BU。

- 成员

   成员变量    |     类型   |        描述               
   ------------- | --------- | --------------------------
   sourceAddress |   String |  选填，操作源账户地址
   masterWeight|String|选填，账户自身权重，大小限制[0, max(uint32)]
   signers|[Signer](#signer)[]|选填，签名者权重列表
   txThreshold|String|选填，交易门限，大小限制[0, max(int64)]
   typeThreshold|[TypeThreshold](#typethreshold)[]|选填，指定类型交易门限
   metadata|String|选填，备注

### AssetIssueOperation

- 功能

  该操作用于发行资产。AssetIssueOperation继承于BaseOperation。

- 费用

  FeeLimit目前(2018.07.26)固定是50.01 BU。

- 成员

   成员变量    |     类型   |        描述             
   ------------- | --------- | ------------------------
   sourceAddress|String|选填，操作源账户地址
   code|String|必填，资产编码，长度限制[1, 64]
   assetAmount|int64|必填，资产发行数量，大小限制[0, max(int64)]
   metadata|String|选填，备注

### AssetSendOperation

**注意**：若目标账户未激活，必须先调用激活账户操作。

- 功能

  该操作用于转移资产。AssetSendOperation继承于BaseOperation。

- 费用

  FeeLimit目前(2018.07.26)固定是0.01 BU。

- 成员

   成员变量    |     类型   |        描述            
   ------------- | --------- | ----------------------
   sourceAddress|String|选填，操作源账户地址
   destAddress|String|必填，目标账户地址
   code|String|必填，资产编码，长度限制[1, 64]
   issuer|String|必填，资产发行账户地址
   assetAmount|int64|必填，资产数量，大小限制[0, max(int64)]
   metadata|String|选填，备注

### BUSendOperation

**注意**：若目标账户未激活，该操作也可使目标账户激活。

- 功能

  该操作用于转移BU。BUSendOperation继承于BaseOperation。

- 费用

  FeeLimit目前(2018.07.26)固定是0.01 BU。

- 成员

   成员变量    |     类型   |        描述          
   ------------- | --------- | ---------------------
   sourceAddress|String|选填，操作源账户地址
   destAddress|String|必填，目标账户地址
   buAmount|int64|必填，BU 数量，大小限制[0, max(int64)]
   metadata|String|选填，备注

### ContractCreateOperation

- 功能

  该操作用于创建合约。ContractCreateOperation继承于BaseOperation。

- 费用

  FeeLimit目前(2018.07.26)固定是10.01 BU。

- 成员

   成员变量    |     类型   |        描述          
   ------------- | --------- | ---------------------
   sourceAddress|String|选填，操作源账户地址
   initBalance|int64|必填，给合约账户的初始化资产，单位MO，1 BU = 10^8 MO, 大小限制[1, max(int64)]
   type|Integer|选填，合约的语种，默认是0
   payload|String|必填，对应语种的合约代码
   initInput|String|选填，合约代码中init方法的入参
   metadata|String|选填，备注

### ContractInvokeByAssetOperation

**注意**：若合约账户不存在，必须先创建合约账户。

- 功能

  该操作用于转移资产并触发合约。ContractInvokeByAssetOperation继承于BaseOperation。

- 费用

  FeeLimit要根据合约中执行交易来做添加手续费，首先发起交易手续费目前(2018.07.26)是0.01BU，然后合约中的交易也需要交易发起者添加相应交易的手续费。

- 成员

   成员变量    |     类型   |        描述          
   ------------- | --------- | ---------------------
   sourceAddress|String|选填，操作源账户地址
   contractAddress|String|必填，合约账户地址
   code|String|选填，资产编码，长度限制[0, 64];当为空时，仅触发合约;
   issuer|String|选填，资产发行账户地址，当null时，仅触发合约
   assetAmount|int64|选填，资产数量，大小限制[0, max(int64)]，当是0时，仅触发合约
   input|String|选填，待触发的合约的main()入参
   metadata|String|选填，备注

### ContractInvokeByBUOperation

**注意**：若目标账户非合约账户且未激活，该操作也可使目标账户激活。

- 功能

  该操作用于转移BU并触发合约。ContractInvokeByBUOperation继承于BaseOperation。

- 费用

  FeeLimit要根据合约中执行交易来做添加手续费，首先发起交易手续费目前(2018.07.26)是0.01BU，然后合约中的交易也需要交易发起者添加相应交易的手续费。

- 成员

   成员变量    |     类型   |        描述          
   ------------- | --------- | ---------------------
   sourceAddress|String|选填，操作源账户地址
   contractAddress|String|必填，合约账户地址
   buAmount|int64|选填，BU 数量，大小限制[0, max(int64)]，当0时仅触发合约
   input|String|选填，待触发的合约的main()入参
   metadata|String|选填，备注

### LogCreateOperation

- 功能

  该操作用于记录日志。LogCreateOperation继承于BaseOperation。

- 费用

  FeeLimit目前(2018.07.26)固定是0.01 BU。

- 成员

   成员变量    |     类型   |        描述          
   ------------- | --------- | ---------------------
   sourceAddress|String|选填，操作源账户地址
   topic|String|必填，日志主题，长度限制[1, 128]
   datas|List<String>|必填，日志内容，每个字符串长度限制[1, 1024]
   metadata|String|选填，备注

## 账户服务

账户服务提供账户相关的接口，包括6个接口：`CheckValid`、 `GetInfo`、 `GetNonce`、 `GetBalance`、`GetAssets`、 `GetMetadata`。

### CheckValid
- **接口说明**

   该接口用于检查区块链账户地址的有效性

- **调用方法**

  `CheckValid(model.AccountCheckValidRequest) model.AccountCheckValidResponse;`

- **请求参数**

   参数      |     类型     |        描述       
   ----------- | ------------ | ---------------- 
   address     |   String     |  必填，待检查的区块链账户地址   

- **响应数据**

   参数      |     类型     |        描述       
   ----------- | ------------ | ---------------- 
   isValid     | Boolean |  是否有效   

- **错误码**

   异常       |     错误码   |   描述  
   -----------  | ----------- | -------- 
   SYSTEM_ERROR |   20000     |  System error 

- **示例**

   ```go
   var reqData model.AccountCheckValidRequest
   address := "buQtfFxpQP9JCFgmu4WBojBbEnVyQGaJDgGn"
   reqData.SetAddress(address)
   resData := testSdk.Account.CheckValid(reqData)
   if resData.ErrorCode == 0 {
   fmt.Println(resData.Result.IsValid)
   }
   ```

### GetInfo

- **接口说明**

   该接口用于获取指定的账户信息

- **调用方法**

  `GetInfo(model.AccountGetInfoRequest) model.AccountGetInfoResponse;`

- **请求参数**

   参数      |     类型     |        描述       
   ----------- | ------------ | ---------------- 
   address     |   String     |  必填，待查询的区块链账户地址   

- **响应数据**

   参数    |     类型      |        描述       
   --------- | ------------- | ---------------- 
   address	  |    String     |    账户地址       
   balance	  |    int64       |    账户余额，单位MO，1 BU = 10^8 MO, 必须大于0
   nonce	  |    int64       |    账户交易序列号，必须大于0
   priv	  | [Priv](#priv) |    账户权限

- **错误码**

   异常       |     错误码   |   描述  
   -----------  | ----------- | -------- 
   INVALID_ADDRESS_ERROR| 11006 | Invalid address
   CONNECTNETWORK_ERROR| 11007| Failed to connect to the network
   SYSTEM_ERROR |   20000     |  System error 

- **示例**

   ```go 
   var reqData model.AccountGetInfoRequest
   var address string = "buQtfFxpQP9JCFgmu4WBojBbEnVyQGaJDgGn"
   reqData.SetAddress(address)
   resData := testSdk.Account.GetInfo(reqData)
   if resData.ErrorCode == 0 {
   data, _ := json.Marshal(resData.Result)
   fmt.Println("Info:", string(data))
   }
   ```

### GetNonce

- **接口说明**

   该接口用于获取指定账户的nonce值

- **调用方法**

  `GetNonce(model.AccountGetNonceRequest)model.AccountGetNonceResponse;`

- **请求参数**

   参数      |     类型     |        描述       
   ----------- | ------------ | ---------------- 
   address     |   String     |  必填，待查询的区块链账户地址   

- **响应数据**

   参数      |     类型     |        描述       
   ----------- | ------------ | ---------------- 
   nonce       |   int64       |  账户交易序列号   

- **错误码**

   异常       |     错误码   |   描述  
   -----------  | ----------- | -------- 
   INVALID_ADDRESS_ERROR| 11006 | Invalid address
   CONNECTNETWORK_ERROR| 11007| Failed to connect to the network
   SYSTEM_ERROR |   20000     |  System error 

- **示例**

   ```go
   var reqData model.AccountGetNonceRequest
   var address string = "buQtfFxpQP9JCFgmu4WBojBbEnVyQGaJDgGn"
   reqData.SetAddress(address)
   if resData.ErrorCode == 0 {
   fmt.Println(resData.Result.Nonce)
   }
   ```

### GetBalance

- **接口说明**

   该接口用于获取指定账户的BU的余额

- **调用方法**

  `GetBalance(model.AccountGetBalanceRequest)model.AccountGetBalanceResponse;`

- **请求参数**

   参数      |     类型     |        描述       
   ----------- | ------------ | ---------------- 
   address     |   String     |  必填，待查询的区块链账户地址   

- **响应数据**

   参数      |     类型     |        描述       
   ----------- | ------------ | ---------------- 
   balance     |   int64       | BU的余额，单位MO，1 BU = 10^8 MO 

- **错误码**

   异常       |     错误码   |   描述  
   -----------  | ----------- | -------- 
   INVALID_ADDRESS_ERROR| 11006 | Invalid address
   CONNECTNETWORK_ERROR| 11007| Failed to connect to the network
   SYSTEM_ERROR |   20000     |  System error 

- **示例**

   ```go
   var reqData model.AccountGetBalanceRequest
   var address string = "buQtfFxpQP9JCFgmu4WBojBbEnVyQGaJDgGn"
   reqData.SetAddress(address)
   resData := testSdk.Account.GetBalance(reqData)
   if resData.ErrorCode == 0 {
   fmt.Println("Balance", resData.Result.Balance)
   }
   ```

### GetAssets

- **接口说明**

   该接口用于获取指定账户的所有资产信息

- **调用方法**

  `GetAssets(model.AccountGetAssetsRequest)model.AccountGetAssetsResponse;`

- **请求参数**

   参数      |     类型     |        描述       
   ----------- | ------------ | ---------------- 
   address     |   String     |  必填，待查询的账户地址   

- **响应数据**

   参数      |     类型     |        描述       
   ----------- | ------------ | ---------------- 
   asset	    | `[]`[AssetInfo](#assetinfo) |账户资产

- **错误码**

   异常       |     错误码   |   描述  
   -----------  | ----------- | -------- 
   INVALID_ADDRESS_ERROR| 11006 | Invalid address
   CONNECTNETWORK_ERROR| 11007| Failed to connect to the network
   NO_ASSET_ERROR|11009|The account does not have the asset
   SYSTEM_ERROR|20000|System error

- **示例**

   ```go
   var reqData model.AccountGetAssetsRequest
   var address string = "buQtfFxpQP9JCFgmu4WBojBbEnVyQGaJDgGn"
   reqData.SetAddress(address)
   resData := testSdk.Account.GetAssets(reqData)
   if resData.ErrorCode == 0 {
   data, _ := json.Marshal(resData.Result.Assets)
   fmt.Println("Assets:", string(data))
   }
   ```

### GetMetadata

- **接口说明**

   该接口用于获取指定账户的metadata信息

- **调用方法**

  `GetMetadata(model.AccountGetMetadataRequest)model.AccountGetMetadataResponse;`

- **请求参数**

   参数   |   类型   |        描述       
   -------- | -------- | ---------------- 
   address  |  String  |  必填，待查询的账户地址  
   key      |  String  |  选填，metadata关键字，长度限制[1, 1024]

- **响应数据**

   参数      |     类型    |        描述       
   ----------- | ----------- | ---------------- 
   metadata    |`[]`[Metadata](#metadata)   |  账户

- **错误码**

   异常       |     错误码   |   描述  
   -----------  | ----------- | -------- 
   INVALID_ADDRESS_ERROR | 11006 | Invalid address
   CONNECTNETWORK_ERROR | 11007 | Failed to connect to the network
   NO_METADATA_ERROR|11010|The account does not have the metadata
   INVALID_DATAKEY_ERROR | 11011 | The length of key must be between 1 and 1024
   SYSTEM_ERROR | 20000| System error


- **示例**

   ```go
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

遵循ATP1.0协议，账户服务提供资产相关的接口，目前有1个接口：`GetInfo`

### GetInfo

- **接口说明**

   该接口用于获取指定账户的指定资产信息

- **调用方法**

  `GetInfo(model.AssetGetInfoRequest) model.AssetGetInfoResponse;`

- **请求参数**

   参数      |     类型     |        描述       
   ----------- | ------------ | ---------------- 
   address     |   String    |  必填，待查询的账户地址
   code        |   String    |  必填，资产编码，长度限制[1, 64]
   issuer      |   String    |  必填，资产发行账户地址

- **响应数据**

   参数      |     类型     |        描述       
   ----------- | ------------ | ---------------- 
   asset	    | `[]`[AssetInfo](#assetinfo) |账户资产   

- **错误码**

   异常       |     错误码   |   描述  
   -----------  | ----------- | -------- 
   INVALID_ADDRESS_ERROR|11006|Invalid address
   CONNECTNETWORK_ERROR|11007|Failed to connect to the network
   INVALID_ASSET_CODE_ERROR|11023|The length of asset code must be between 1 and 64
   INVALID_ISSUER_ADDRESS_ERROR|11027|Invalid issuer address
   SYSTEM_ERROR|20000|System error

- **示例**

   ```go
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

合约服务提供合约相关的接口，目前有4个接口：`CheckValid`、 `GetInfo`、 `GetAddress`、 `Call`

### CheckValid

- **接口说明**

   该接口用于检测合约账户的有效性

- **调用方法**

  `CheckValid(reqData model.ContractCheckValidRequest) model.ContractCheckValidResponse;`

- **请求参数**

   参数      |     类型     |        描述       
   ----------- | ------------ | ---------------- 
   contractAddress     |   String     |  待检测的合约账户地址   

- **响应数据**

   参数      |     类型     |        描述       
   ----------- | ------------ | ---------------- 
   isValid     |   Boolean     |  是否有效   

- **错误码**

   异常       |     错误码   |   描述  
   -----------  | ----------- | -------- 
   INVALID_CONTRACTADDRESS_ERROR|11037|Invalid contract address
   CONTRACTADDRESS_NOT_CONTRACTACCOUNT_ERROR|11038|ContractAddress is not a contract account
   SYSTEM_ERROR |   20000     |  System error 

- **示例**

   ```go
   var reqData model.ContractCheckValidRequest
   var address string = "buQXmYrmqt6ohcKtLFKgWFSZ5CjYKaSzaMjT"
   reqData.SetAddress(address)
   resData := testSdk.Contract.CheckValid(reqData)
   if resData.ErrorCode != 0 {
      t.Errorf(resData.ErrorDesc)
   } else {
      t.Log("Test_Contract_CheckValid succeed", resData.Result)
   }
   ```

### GetInfo

- **接口说明**

   该接口用于查询合约代码

- **调用方法**

  `GetInfo(model.ContractGetInfoRequest) model.ContractGetInfoResponse;`

- **请求参数**

   参数      |     类型     |        描述       
   ----------- | ------------ | ---------------- 
   contractAddress     |   String     |  待查询的合约账户地址   

- **响应数据**

   参数      |     类型     |        描述       
   ----------- | ------------ | ---------------- 
   contract|[ContractInfo](#contractinfo)|合约信息

- **错误码**

   异常       |     错误码   |   描述  
   -----------  | ----------- | -------- 
   INVALID_CONTRACTADDRESS_ERROR|11037|Invalid contract address
   CONTRACTADDRESS_NOT_CONTRACTACCOUNT_ERROR|11038|contractAddress is not a contract account
   NO_SUCH_TOKEN_ERROR|11030|No such token
   GET_TOKEN_INFO_ERROR|11066|Failed to get token info
   SYSTEM_ERROR|20000|System error

- **示例**

   ```go
   var reqData model.ContractGetInfoRequest
   var address string = "buQfnVYgXuMo3rvCEpKA6SfRrDpaz8D8A9Ea"
   reqData.SetAddress(address)
   resData := testSdk.Contract.GetInfo(reqData)
   if resData.ErrorCode == 0 {
   data, _ := json.Marshal(resData.Result.Contract)
   fmt.Println("Contract:", string(data))
   }
   ```

### getAddress

- **接口说明**

  该接口用于查询合约地址

- **调用方法**

  `GetAddress(reqData model.ContractGetAddressRequest) model.ContractGetAddressResponse;`

- **请求参数**

   参数      |     类型     |        描述       
   ----------- | ------------ | ---------------- 
   hash     |   String     |  创建合约交易的hash   

- **响应数据**

   参数      |     类型     |        描述       
   ----------- | ------------ | ---------------- 
   contractAddressList|List<[ContractAddressInfo](#contractaddressinfo)>|合约地址列表

- **错误码**

   异常       |     错误码   |   描述  
   -----------  | ----------- | -------- 
   INVALID_HASH_ERROR|11055|Invalid transaction hash
   CONNECTNETWORK_ERROR|11007|Failed to connect to the network
   SYSTEM_ERROR|20000|System error

- **示例**

   ```go
   // 初始化请求参数
   var reqData model.ContractGetAddressRequest();
   reqData.SetAddress("44246c5ba1b8b835a5cbc29bdc9454cdb9a9d049870e41227f2dcfbcf7a07689");

   resData := sdk.Contract.GetAddress(reqData);
   if resData.ErrorCode == 0 {
   fmt.Println("Address:", resData.Result.Address);
   }
   ```

### Call 

- **接口说明**

   该接口用于调试合约代码

- **调用方法**

  `Call(reqData model.ContractCallRequest) model.ContractCallResponse;`

- **请求参数**

   参数      |     类型     |        描述       
   ----------- | ------------ | ---------------- 
   sourceAddress|String|选填，合约触发账户地址
   contractAddress|String|选填，合约账户地址，与code不能同时为空
   code|String|选填，合约代码，与contractAddress不能同时为空
   input|String|选填，合约入参
   contractBalance|int64|选填，赋予合约的初始 BU 余额, 单位MO，1 BU = 10^8 MO, 大小限制[1, max(int64)]
   optType|Integer|必填，0: 调用合约的读写接口 init, 1: 调用合约的读写接口 main, 2 :调用只读接口 query
   feeLimit|int64|交易要求的最低手续费， 大小限制[1, max(int64)]
   gasPrice|int64|交易燃料单价，大小限制[1000, max(int64)]


- **响应数据**

   参数      |     类型     |        描述       
   ----------- | ------------ | ---------------- 
   logs|JSONObject|日志信息
   queryRets|JSONArray|查询结果集
   stat|[ContractStat](#contractstat)|合约资源占用信息
   txs|`[]`[TransactionEnvs](#transactionenvs)|交易集

- **错误码**

   异常       |     错误码   |   描述  
   -----------  | ----------- | -------- 
   INVALID_SOURCEADDRESS_ERROR|11002|Invalid sourceAddress
   INVALID_CONTRACTADDRESS_ERROR|11037|Invalid contract address
   CONTRACTADDRESS_CODE_BOTH_NULL_ERROR|11063|ContractAddress and code cannot be empty at the same time
   INVALID_OPTTYPE_ERROR|11064|OptType must be between 0 and 2
   CONNECTNETWORK_ERROR|11007|Failed to connect to the network
   SYSTEM_ERROR|20000|System error

- **示例**

   ```go 
   var reqData model.ContractCallRequest
   var contractAddress string = "buQXmYrmqt6ohcKtLFKgWFSZ5CjYKaSzaMjT"
   var feeLimit int64 = 1000000
   var gasPrice int64 = 1000
   var contractBalance string = "100000000000"
   var input string = "input"
   var optType int64 = 2
   var code string = "HNC"

   reqData.SetContractAddress(contractAddress)
   reqData.SetContractBalance(contractBalance)
   reqData.SetFeeLimit(feeLimit)
   reqData.SetGasPrice(gasPrice)
   reqData.SetInput(input)
   reqData.SetOptType(optType)
   reqData.SetCode(code)
   resData := testSdk.Contract.Call(reqData)

   if resData.ErrorCode != 0 {
      t.Errorf(resData.ErrorDesc)
   } else {
      t.Log("Test_Contract_Call succeed", resData.Result)
   }
   ```


## 区块服务

区块服务主要是区块相关的接口，目前有11个接口：`GetNumber`、 `CheckStatus`、 `GetTransactions`、 `GetInfo`、 `GetLatest`、 `GetValidators`、 `GetLatestValidators`、 `GetReward`、 `GetLatestReward`、 `GetFees`、 `GetLatestFees`。

### getNumber

- **接口说明**

   该接口用于查询最新的区块高度。

- **调用方法**

  `GetNumber() model.BlockGetNumberResponse;`

- **响应数据**

   参数      |     类型     |        描述       
   ----------- | ------------ | ---------------- 
   blockNumber|int64|最新的区块高度，对应底层字段seq

- **错误码**

   异常       |     错误码   |   描述  
   -----------  | ----------- | -------- 
   CONNECTNETWORK_ERROR|11007|Failed to connect to the network
   SYSTEM_ERROR|20000|System error

- **示例**

   ```go 
   resData := testSdk.Block.GetNumber()
   if resData.ErrorCode == 0 {
      fmt.Println("BlockNumber:", resData.Result.BlockNumber)
   }
   ```

### checkStatus

- **接口说明**

   该接口用于检查本地节点区块是否同步完成。

- **调用方法**

  `CheckStatus() model.BlockCheckStatusResponse;`

- **响应数据**

   参数      |     类型     |        描述       |
   ----------- | ------------ | ---------------- |
   isSynchronous    |   Boolean     |  区块是否同步   |

- **错误码**

   异常       |     错误码   |   描述  |
   -----------  | ----------- | -------- |
   CONNECTNETWORK_ERROR|11007|Failed to connect to the network
   SYSTEM_ERROR|20000|System error

- **示例**

   ```go
   resData := testSdk.Block.CheckStatus()
   if resData.ErrorCode == 0 {
      fmt.Println("IsSynchronous:", resData.Result.IsSynchronous)
   }
   ```

### getTransactions

- **接口说明**

   该接口用于查询指定区块高度下的所有交易。

- **调用方法**

   `GetTransactions(model.BlockGetTransactionRequest)model.BlockGetTransactionResponse;`

- **请求参数**

   参数      |     类型     |        描述       
   ----------- | ------------ | ---------------- 
   blockNumber|int64|必填，待查询的区块高度，必须大于0

- **响应数据**

   参数      |     类型     |        描述       
   ----------- | ------------ | ---------------- 
   totalCount|int64|返回的总交易数
   transactions|`[]`[TransactionHistory](#transactionhistory)|交易内容

- **错误码**

   异常       |     错误码   |   描述  
   -----------  | ----------- | -------- 
   INVALID_BLOCKNUMBER_ERROR|11060|BlockNumber must bigger than 0
   CONNECTNETWORK_ERROR|11007|Failed to connect to the network
   SYSTEM_ERROR|20000|System error

- **示例**

   ```go 
   var reqData model.BlockGetTransactionRequest
   var blockNumber int64 = 581283
   reqData.SetBlockNumber(blockNumber)
   resData := testSdk.Block.GetTransactions(reqData)
   if resData.ErrorCode == 0 {
      data, _ := json.Marshal(resData.Result.Transactions)
      fmt.Println("Transactions:", string(data))
   }
   ```

### getInfo

- **接口说明**

   该接口用于获取区块信息。

- **调用方法**

  `GetInfo(model.BlockGetInfoRequest) model.BlockGetInfoResponse;`

- **请求参数**

   参数      |     类型     |        描述       
   ----------- | ------------ | ---------------- 
   blockNumber|int64|必填，待查询的区块高度

- **响应数据**

   参数      |     类型     |        描述       
   ----------- | ------------ | ---------------- 
   closeTime|int64|区块关闭时间
   number|int64|区块高度
   txCount|int64|交易总量
   version|String|区块版本

- **错误码**

   异常       |     错误码   |   描述  
   -----------  | ----------- | -------- 
   INVALID_BLOCKNUMBER_ERROR|11060|BlockNumber must bigger than 0
   CONNECTNETWORK_ERROR|11007|Failed to connect to the network
   SYSTEM_ERROR|20000|System error

- **示例**

   ```go 
   var reqData model.BlockGetInfoRequest
   var blockNumber int64 = 581283
   reqData.SetBlockNumber(blockNumber)
   resData := testSdk.Block.GetInfo(reqData)
   if resData.ErrorCode == 0 {
      data, _ := json.Marshal(resData.Result.Header)
      fmt.Println("Header:", string(data))
   }
   ```

### getLatest

- **接口说明**

   该接口用于获取最新区块信息。

- **调用方法**

  `GetLatest() model.BlockGetLatestResponse;`

- **响应数据**

   参数      |     类型     |        描述       
   ----------- | ------------ | ---------------- 
   closeTime|int64|区块关闭时间
   number|int64|区块高度，对应底层字段seq
   txCount|int64|交易总量
   version|String|区块版本


- **错误码**

   异常       |     错误码   |   描述  
   -----------  | ----------- | -------- 
   CONNECTNETWORK_ERROR|11007|Failed to connect to the network
   SYSTEM_ERROR|20000|System error

- **示例**

   ```go
   resData := testSdk.Block.GetLatest()
   if resData.ErrorCode == 0 {
      data, _ := json.Marshal(resData.Result.Header)
      fmt.Println("Header:", string(data))
   }
   ```

### getValidators

- **接口说明**

   该接口用于获取指定区块中所有验证节点。

- **调用方法**

  `GetValidators(model.BlockGetValidatorsRequest)model.BlockGetValidatorsResponse;`

- **请求参数**

   参数      |     类型     |        描述       
   ----------- | ------------ | ---------------- 
   blockNumber|int64|必填，待查询的区块高度，必须大于0

- **响应数据**

   参数      |     类型     |        描述       
   ----------- | ------------ | ---------------- 
   validators|`[]`String|验证节点列表

- **错误码**

   异常       |     错误码   |   描述  
   -----------  | ----------- | -------- 
   INVALID_BLOCKNUMBER_ERROR|11060|BlockNumber must bigger than 0
   CONNECTNETWORK_ERROR|11007|Failed to connect to the network
   SYSTEM_ERROR|20000|System error

- **示例**

   ```go
   var reqData model.BlockGetValidatorsRequest
   var blockNumber int64 = 581283
   reqData.SetBlockNumber(blockNumber)
   resData := testSdk.Block.GetValidators(reqData)
   if resData.ErrorCode == 0 {
      data, _ := json.Marshal(resData.Result.Validators)
      fmt.Println("Validators:", string(data))
   }
   ```

### getLatestValidators

- **接口说明**

   该接口用于获取最新区块中所有验证节点。

- **调用方法**

  `GetLatestValidators() model.BlockGetLatestValidatorsResponse;`

- **响应数据**

   参数      |     类型     |        描述       
   ----------- | ------------ | ---------------- 
   validators|`[]`String|验证节点列表

- **错误码**

   异常       |     错误码   |   描述  
   -----------  | ----------- | -------- 
   CONNECTNETWORK_ERROR|11007|Failed to connect to the network
   SYSTEM_ERROR|20000|System error

- **示例**

   ```go
   resData := testSdk.Block.GetLatestValidators()
   if resData.ErrorCode == 0 {
      data, _ := json.Marshal(resData.Result.Validators)
      fmt.Println("Validators:", string(data))
   }
   ```

### getReward

- **接口说明**

   该接口用于获取指定区块中的区块奖励和验证节点奖励。

- **调用方法**

  `GetReward(model.BlockGetRewardRequest) model.BlockGetRewardResponse;`

- **请求参数**

   参数      |     类型     |        描述       
   ----------- | ------------ | ---------------- 
   blockNumber|int64|必填，待查询的区块高度，必须大于0

- **响应数据**

   参数      |     类型     |        描述       
   ----------- | ------------ | ---------------- 
    validators | [Rewards](#Rewards)[] | 验证节点奖励 
    kols       | [Rewards](#Rewards)[] | 生态节点奖励 


- **错误码**

   异常       |     错误码   |   描述  
   -----------  | ----------- | -------- 
   INVALID_BLOCKNUMBER_ERROR|11060|BlockNumber must bigger than 0
   CONNECTNETWORK_ERROR|11007|Failed to connect to the network
   SYSTEM_ERROR|20000|System error

- **示例**

   ```go
   var reqData model.BlockGetRewardRequest
   var blockNumber int64 = 581283
   reqData.SetBlockNumber(blockNumber)
   resData := testSdk.Block.GetReward(reqData)
   if resData.ErrorCode == 0 {
      fmt.Println("ValidatorsReward:", resData.Result.ValidatorsReward)
   }
   ```

### getLatestReward

- **接口说明**

   获取最新区块中的区块奖励和验证节点奖励。

- **调用方法**

  `GetLatestReward() model.BlockGetLatestRewardResponse;`

- **响应数据**

   参数      |     类型     |        描述       
   ----------- | ------------ | ---------------- 
    validators | [Rewards](#Rewards)[] | 验证节点奖励 
    kols       | [Rewards](#Rewards)[] | 生态节点奖励 

- **错误码**

   异常       |     错误码   |   描述  
   -----------  | ----------- | -------- 
   CONNECTNETWORK_ERROR|11007|Failed to connect to the network
   SYSTEM_ERROR|20000|System error

- **示例**

   ```go 
   resData := testSdk.Block.GetLatestReward()
   if resData.ErrorCode == 0 {
      fmt.Println("ValidatorsReward:", resData.Result.ValidatorsReward)
   }
   ```

### getFees

- **接口说明**

   获取指定区块中的账户最低资产限制和燃料单价。

- **调用方法**

  `GetFees(model.BlockGetFeesRequest) model.BlockGetFeesResponse;`

- **请求参数**

   参数      |     类型     |        描述       
   ----------- | ------------ | ---------------- 
   blockNumber|int64|必填，待查询的区块高度，必须大于0

- **响应数据**

   参数      |     类型     |        描述       
   ----------- | ------------ | ---------------- 
   fees|[Fees](#fees)|费用

- **错误码**

   异常       |     错误码   |   描述  
   -----------  | ----------- | -------- 
   INVALID_BLOCKNUMBER_ERROR|11060|BlockNumber must bigger than 0
   CONNECTNETWORK_ERROR|11007|Failed to connect to the network
   SYSTEM_ERROR|20000|System error

- **示例**

   ```go 
   var reqData model.BlockGetFeesRequest
   var blockNumber int64 = 581283
   reqData.SetBlockNumber(blockNumber)
   resData := testSdk.Block.GetFees(reqData)
   if resData.ErrorCode == 0 {
      data, _ := json.Marshal(resData.Result.Fees)
      fmt.Println("Fees:", string(data))
   }
   ```

### getLatestFees

- **接口说明**

   该接口用于获取最新区块中的账户最低资产限制和燃料单价。

- **调用方法**

  `GetLatestFees() model.BlockGetLatestFeesResponse;`

- **响应数据**

   参数      |     类型     |        描述       
   ----------- | ------------ | ---------------- 
   fees|[Fees](#fees)|费用

- **错误码**

   异常       |     错误码   |   描述  
   -----------  | ----------- | -------- 
   CONNECTNETWORK_ERROR|11007|Failed to connect to the network
   SYSTEM_ERROR|20000|System error

- **示例**

   ```go
   resData := testSdk.Block.GetLatestFees()
   if resData.ErrorCode == 0 {
      data, _ := json.Marshal(resData.Result.Fees)
      fmt.Println("Fees:", string(data))
   }
   ```



## 数据对象
#### Priv

| 成员         | 类型                    | 描述                                                    |
| ------------ | ----------------------- | ------------------------------------------------------- |
| masterWeight | int64                    | 账户自身权重，大小限制[0,max(uint32)] |
| signers      | `[]`[Signer](#signer) | 签名者权重列表                                          |
| threshold    | [Threshold](#threshold) | 门限                                                    |

#### Signer

| 成员    | 类型   | 描述                                                  |
| ------- | ------ | ----------------------------------------------------- |
| address | String | 签名者区块链账户地址                                  |
| weight  | int64   | 签名者权重，大小限制[0,max(uint32)] |

#### Threshold

| 成员           | 类型                              | 描述                                      |
| -------------- | --------------------------------- | ----------------------------------------- |
| txThreshold    | int64                              | 交易默认门限，大小限制[0, max(int64)] |
| typeThresholds | `[]`[TypeThreshold](#typethreshold) | 不同类型交易的门限                        |

#### TypeThreshold

| 成员      | 类型 | 描述                                |
| --------- | ---- | ----------------------------------- |
| type      | int64 | 操作类型，必须大于0                 |
| threshold | int64 | 门限值，大小限制[0, max(int64)] |

#### AssetInfo

| 成员        | 类型        | 描述         |
| ----------- | ----------- | ------------ |
| key         | [Key](#key) | 资产惟一标识 |
| assetAmount | int64        | 资产数量     |

#### Key

| 成员   | 类型   | 描述             |
| ------ | ------ | ---------------- |
| code   | String | 资产编码         |
| issuer | String | 资产发行账户地址 |

#### ContractInfo

成员      |     类型     |        描述       |
----------- | ------------ | ---------------- |
type|Integer|合约类型，默认0
payload|String|合约代码

#### Metadata
   成员      |     类型    |        描述       
----------- | ----------- | ---------------- 
key         |  String     |  metadata的关键词
value       |  String     |  metadata的内容
version     |  int64      |  metadata的版本

#### ContractAddressInfo

成员      |     类型     |        描述       
----------- | ------------ | ---------------- 
contractAddress|String|合约地址
operationIndex|Integer|所在操作的下标

#### ContractStat

成员      |     类型     |        描述       
----------- | ------------ | ---------------- 
applyTime|int64|接收时间
memoryUsage|int64|内存占用量
stackUsage|int64|堆栈占用量
step|int64|几步完成

#### TransactionEnvs

成员      |     类型     |        描述       
----------- | ------------ | ---------------- 
transactionEnv|[TransactionEnv](#transactionenv)|交易

#### TransactionEnv

成员      |     类型     |        描述       
----------- | ------------ | ---------------- 
transaction|[TransactionInfo](#transactioninfo)|交易内容
trigger|[ContractTrigger](#contracttrigger)|合约触发者

#### TransactionInfo

成员      |     类型     |        描述       
----------- | ------------ | ---------------- 
sourceAddress|String|交易发起的源账户地址
feeLimit|int64|交易要求的最低费用
gasPrice|int64|交易燃料单价
nonce|int64|交易序列号
operations|`[]`[Operation](#operation)|操作列表

#### ContractTrigger
成员      |     类型     |        描述       
----------- | ------------ | ---------------- 
transaction|[TriggerTransaction](#triggertransaction)|触发交易

#### Operation

成员      |     类型     |        描述       
----------- | ------------ | ---------------- 
type|Integer|操作类型
sourceAddress|String|操作发起源账户地址
metadata|String|备注
createAccount|[OperationCreateAccount](#operationcreateaccount)|创建账户操作
issueAsset|[OperationIssueAsset](#operationissueasset)|发行资产操作
payAsset|[OperationPayAsset](#operationpayasset)|转移资产操作
payCoin|[OperationPayCoin](#operationpaycoin)|发送BU操作
setMetadata|[OperationSetMetadata](#operationsetmetadata)|设置metadata操作
setPrivilege|[OperationSetPrivilege](#operationsetprivilege)|设置账户权限操作
log|[OperationLog](#operationlog)|记录日志

#### TriggerTransaction

成员      |     类型     |        描述       
----------- | ------------ | ---------------- 
hash|String|交易hash

#### OperationCreateAccount

成员      |     类型     |        描述       
----------- | ------------ | ---------------- 
destAddress|String|目标账户地址
contract|[Contract](#contract)|合约信息
priv|[Priv](#priv)|账户权限
metadata|`[]`[Metadata](#metadata)|账户
initBalance|int64|账户资产, 单位MO，1 BU = 10^8 MO, 
initInput|String|合约init函数的入参

#### Contract

成员      |     类型     |        描述       
----------- | ------------ | ---------------- 
type|Integer| 合约的语种，默认不赋值
payload|String|对应语种的合约代码

#### MetadataInfo

成员      |     类型     |        描述       
----------- | ------------ | ---------------- 
key|String|metadata的关键词
value|String|metadata的内容
version|int64|metadata的版本

#### OperationIssueAsset

成员      |     类型     |        描述       
----------- | ------------ | ---------------- 
code|String|资产编码
assetAmount|int64|资产数量

#### OperationPayAsset

成员      |     类型     |        描述       
----------- | ------------ | ---------------- 
destAddress|String|待转移的目标账户地址
asset|[AssetInfo](#assetinfo)|账户资产
input|String|合约main函数入参

#### OperationPayCoin

成员      |     类型     |        描述       
----------- | ------------ | ---------------- 
destAddress|String|待转移的目标账户地址
buAmount|int64|待转移的BU数量
input|String|合约main函数入参

#### OperationSetMetadata

成员      |     类型     |        描述       
----------- | ------------ | ---------------- 
key|String|metadata的关键词
value|String|metadata的内容
version|int64|metadata的版本
deleteFlag|boolean|是否删除metadata

#### OperationSetPrivilege

成员      |     类型     |        描述       
----------- | ------------ | ---------------- 
masterWeight|String|账户自身权重，大小限制[0,max(uint32)]
signers|`[]`[Signer](#signer)|签名者权重列表
txThreshold|String|交易门限，大小限制[0, max(int64)]
typeThreshold|[TypeThreshold](#typethreshold)|指定类型交易门限

#### OperationLog

成员      |     类型     |        描述       
----------- | ------------ | ---------------- 
topic|String|日志主题
data|String[]|日志内容



#### TestTx

| 成员变量       | 类型                                        | 描述         |
| -------------- | ------------------------------------------- | ------------ |
| transactionEnv | [TestTransactionFees](#testtransactionfees) | 评估交易费用 |

#### TestTransactionFees

| 成员变量        | 类型                                | 描述     |
| --------------- | ----------------------------------- | -------- |
| transactionFees | [TransactionFees](#transactionfees) | 交易费用 |

#### TransactionFees

| 成员变量 | 类型 | 描述               |
| -------- | ---- | ------------------ |
| feeLimit | int64 | 交易要求的最低费用 |
| gasPrice | int64 | 交易燃料单价       |

#### Signature

| 成员变量  | 类型 | 描述       |
| --------- | ---- | ---------- |
| signData  | int64 | 签名后数据 |
| publicKey | int64 | 公钥       |

#### TransactionHistory

| 成员变量    | 类型                                | 描述         |
| ----------- | ----------------------------------- | ------------ |
| actualFee   | String                              | 交易实际费用 |
| closeTime   | int64                                | 交易关闭时间 |
| errorCode   | int64                                | 交易错误码   |
| errorDesc   | String                              | 交易描述     |
| hash        | String                              | 交易hash     |
| ledgerSeq   | int64                                | 区块序列号   |
| transaction | [TransactionInfo](#transactioninfo) | 交易内容列表 |
| signatures  | `[]`[Signature](#signature)       | 签名列表     |
| txSize      | int64                                | 交易大小     |

#### ValidatorReward

| 成员变量  | 类型   | 描述         |
| --------- | ------ | ------------ |
| validator | String | 验证节点地址 |
| reward    | int64   | 验证节点奖励 |

#### Rewards

| 成员变量 | 类型   | 描述     |
| -------- | ------ | -------- |
| address  | String | 节点地址 |
| reward   | Array  | 节点奖励 |

#### ValidatorInfo

| 成员变量        | 类型   | 描述         |
| --------------- | ------ | ------------ |
| address         | String | 共识节点地址 |

#### Fees

| 成员变量    | 类型 | 描述                                 |
| ----------- | ---- | ------------------------------------ |
| baseReserve | int64 | 账户最低资产限制                     |
| gasPrice    | int64 | 交易燃料单价，单位MO，1 BU = 10^8 MO |



## **错误码**

   异常       |     错误码   |   描述  
-----------  | ----------- | -------- 
ACCOUNT_CREATE_ERROR|11001|Failed to create the account 
INVALID_SOURCEADDRESS_ERROR|11002|Invalid sourceAddress
INVALID_DESTADDRESS_ERROR|11003|Invalid destAddress
INVALID_INITBALANCE_ERROR|11004|InitBalance must be between 1 and max(int64) 
SOURCEADDRESS_EQUAL_DESTADDRESS_ERROR|11005|SourceAddress cannot be equal to destAddress
INVALID_ADDRESS_ERROR|11006|Invalid address
CONNECTNETWORK_ERROR|11007|Failed to connect to the network
INVALID_ISSUE_AMOUNT_ERROR|11008|Amount of the token to be issued must be between 1 and max(int64)
NO_ASSET_ERROR|11009|The account does not have the asset
NO_METADATA_ERROR|11010|The account does not have the metadata
INVALID_DATAKEY_ERROR|11011|The length of key must be between 1 and 1024
INVALID_DATAVALUE_ERROR|11012|The length of value must be between 0 and 256000
INVALID_DATAVERSION_ERROR|11013|The version must be equal to or greater than 0 
INVALID_MASTERWEIGHT_ERROR|11015|MasterWeight must be between 0 andmax(uint32)
INVALID_SIGNER_ADDRESS_ERROR|11016|Invalid signer address
INVALID_SIGNER_WEIGHT_ERROR|11017|Signer weight must be between 0 andmax(uint32)
INVALID_TX_THRESHOLD_ERROR|11018|TxThreshold must be between 0 and max(int64)
INVALID_OPERATION_TYPE_ERROR|11019|Operation type must be between 1 and 100
INVALID_TYPE_THRESHOLD_ERROR|11020|TypeThreshold must be between 0 and max(int64)
INVALID_ASSET_CODE_ERROR|11023|The length of asset code must be between 1 and 64
INVALID_ASSET_AMOUNT_ERROR|11024|AssetAmount must be between 0 and max(int64)
INVALID_BU_AMOUNT_ERROR|11026|BuAmount must be between 0 and max(int64)
INVALID_ISSUER_ADDRESS_ERROR|11027|Invalid issuer address
NO_SUCH_TOKEN_ERROR|11030|No such token
INVALID_TOKEN_NAME_ERROR|11031|The length of token name must be between 1 and 1024
INVALID_TOKEN_SIMBOL_ERROR|11032|The length of symbol must be between 1 and 1024
INVALID_TOKEN_DECIMALS_ERROR|11033|Decimals must be between 0 and 8
INVALID_TOKEN_TOTALSUPPLY_ERROR|11034|TotalSupply must be between 1 and max(int64)
INVALID_TOKENOWNER_ERRPR|11035|Invalid token owner
INVALID_CONTRACTADDRESS_ERROR|11037|Invalid contract address
CONTRACTADDRESS_NOT_CONTRACTACCOUNT_ERROR|11038|contractAddress is not a contract account
INVALID_TOKEN_AMOUNT_ERROR|11039|TokenAmount must be between 1 and max(int64)
SOURCEADDRESS_EQUAL_CONTRACTADDRESS_ERROR|11040|SourceAddress cannot be equal to contractAddress
INVALID_FROMADDRESS_ERROR|11041|Invalid fromAddress
FROMADDRESS_EQUAL_DESTADDRESS_ERROR|11042|FromAddress cannot be equal to destAddress
INVALID_SPENDER_ERROR|11043|Invalid spender
PAYLOAD_EMPTY_ERROR|11044|Payload cannot be empty
INVALID_LOG_TOPIC_ERROR|11045|The length of a log topic must be between 1 and 128
INVALID_LOG_DATA_ERROR|11046|The length of one piece of log data must be between 1 and 1024
INVALID_CONTRACT_TYPE_ERROR|11047|Invalid contract type
INVALID_NONCE_ERROR|11048|Nonce must be between 1 and max(int64)
INVALID_GASPRICE_ERROR|11049|GasPrice must be between 0 and max(int64)
INVALID_FEELIMIT_ERROR|11050|FeeLimit must be between 0 and max(int64)
OPERATIONS_EMPTY_ERROR|11051|Operations cannot be empty
INVALID_CEILLEDGERSEQ_ERROR|11052|CeilLedgerSeq must be equal to or greater than 0
OPERATIONS_ONE_ERROR|11053|One of the operations cannot be resolved
INVALID_SIGNATURENUMBER_ERROR|11054|SignagureNumber must be between 1 and max(int32)
INVALID_HASH_ERROR|11055|Invalid transaction hash
INVALID_BLOB_ERROR|11056|Invalid blob
PRIVATEKEY_NULL_ERROR|11057|PrivateKeys cannot be empty
PRIVATEKEY_ONE_ERROR|11058|One of privateKeys is invalid
SIGNDATA_NULL_ERROR|11059|SignData cannot be empty
INVALID_BLOCKNUMBER_ERROR|11060|BlockNumber must be bigger than 0
PUBLICKEY_NULL_ERROR|11061|PublicKey cannot be empty
URL_EMPTY_ERROR|11062|Url cannot be empty
CONTRACTADDRESS_CODE_BOTH_NULL_ERROR|11063|ContractAddress and code cannot be empty at the same time
INVALID_OPTTYPE_ERROR|11064|OptType must be between 0 and 2
GET_ALLOWANCE_ERROR|11065|Failed to get allowance
GET_TOKEN_INFO_ERROR|11066|Failed to get token info
SIGNATURE_EMPTY_ERROR|11067|The signatures cannot be empty
REQUEST_NULL_ERROR|12001|Request parameter cannot be null
CONNECTN_BLOCKCHAIN_ERROR|19999|Failed to connect to the blockchain 
SYSTEM_ERROR|20000|System error
GET_ENCPUBLICKEY_ERROR|14000|The function ‘GetEncPublicKey’ failed
SIGN_ERROR|14001|The function ‘Sign’ failed

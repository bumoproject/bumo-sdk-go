---
id: sdk_go
title: BUMO GO SDK
sidebar_label: GO
---

## Overview
This document details the common interfaces of the BUMO GO SDK, making it easier for developers to operate and query the BuChain.

## Installation

GO 1.10.1 or above is required.

The packages on which the projects depend are in the src folder, and you can get the packages as follows:

```go
// To get the packages.
go get github.com/bumoproject/bumo-sdk-go
```

## Format of Request Parameters and Response Data

This section details the format of the request parameters and response data.

### Request Parameters

The class name of the request parameter of the interface is composed of **Service Name** + **Method Name** + **Request**. For example, the request parameter format of the [getInfo](#getinfo) interface in Account Service is `AccountGetInfoRequest`.

The member of the request parameter is the member of the input parameter of each interface. For example, if the input parameter of the [getInfo](#getinfo) interface in Account Service is address, the complete structure of the request parameters of the interface is as follows:
```go
type AccountGetInfoRequest struct {
address string
}
```

### Response Data

The class name of the response data of the interface is composed of **Service Name** + **Method Name** + **Response**. For example, the response data format of the [getNonce](#getnonce) interface in Account Service is `AccountGetNonceResponse`.

The members of the response data include error codes, error descriptions, and return results. For example, the members of the response data of the [getInfo](#getinfo) interface in Account Services are as follows:
```go
type AccountGetInfoResponse struct {
  ErrorCode int
  ErrorDesc string
  Result  AccountGetInfoResult
}
```

**Note**: 
- errorCode: **error code**. 0 means no error, greater than 0 means there is an error.
- errorDesc: Error description.
- result: Return the result. A structure whose class name is Service **Service Name** + **Method Name** + **Result**, whose members are members of the return value of each interface. For example, the result class name of the [getNonce](#getnonce) interface in Account Service is `AccountGetNonceResult`, and the member has a nonce. The complete structure is as follows:

```go
type AccountGetNonceResult struct {
  Nonce int64
}
```

## Usage

This section describes the process of using the SDK. First you need to generate the SDK instance and then call the interface of the corresponding service. Services include [Account Service](#account-service), [Asset Service](#asset-service), [Contract Service](#contract-service), [Transaction Service](#transaction-service), and [Block Service](#block-service). Interfaces are classified into [Generating Public-Private Keys and Addresses](#generating-public-private-keys-and-addresses), [Checking Validity](#checking-validity), [Querying](#querying), and [Groadcasting Transaction](#broadcasting-transactions).

### Importing Packagings

Import the packages before generating the SDK instance.

```go 
import(
  "github.com/bumoproject/bumo-sdk-go/src/model"
  "github.com/bumoproject/bumo-sdk-go/src/sdk"

```

### Generating SDK Instances

This method is to initialize SDK structure:

```go
var testSdk sdk.sdk
```

Call the Init interface of SDK structure:

```go
url :="http://seed1.bumotest.io:26002"
var reqData model.SDKInitRequest
reqData.SetUrl(url)
resData := testSdk.Init(reqData)
```

### Generating Public-Private Keys and Addresses

The public-private key address interface is used to generate the public key, private key, and address for the account on the BuChain. This can be achieved by directly calling the `create` interface of account service. The specific call is as follows:
```go
resData :=testSdk.Account.Create()
```

### Checking Validity

The validity check interface is used to verify the validity of the information, and the information validity check can be achieved by directly invoking the corresponding interface. For example, to verify the validity of the account address, the specific call is as follows:
```go
//Initialize request parameters
var reqData model.AccountCheckValidRequest
var address string = "buQtfFxpQP9JCFgmu4WBojBbEnVyQGaJDgGn"
reqData.SetAddress(address)
//Call the checkValid interface
resData := testSdk.Account.CheckValid(reqData)
```

### Querying
The query interface is used to query data on the BuChain, and data query can be implemented by directly invoking the corresponding interface. For example, to query the account information, the specific call is as follows:
```go
//Initialize request parameters
var reqData model.AccountGetInfoRequest
var address string = "buQtfFxpQP9JCFgmu4WBojBbEnVyQGaJDgGn"
reqData.SetAddress(address)
//Call the getInfo interface
resData := testSdk.Account.GetInfo(reqData)
```

### Broadcasting Transactions
Broadcasting transactions refers to the initiation of a transaction by means of broadcasting. The broadcast transaction consists of the following steps:

1. [Obtaining the Nonce Value of the account](#obtaining-the-nonce-value-of-the-account)
2. [Building Operations](#building-operations)
3. [Serializing Transactions](#serializing-transactions)
4. [Signing Transactions](#signing-transactions)
5. [Submitting Transactions](#submitting-transactions)

#### Obtaining the nonce value of the account

The developer can maintain the `nonce` value of each account, and automatically increments by 1 for the `nounce` value after submitting a transaction, so that multiple transactions can be sent in a short time; otherwise, the `nonce` value of the account must be added 1 after the execution of the previous transaction is completed. For interface details, see [getNonce](#getnonce), which calls as follows:
```go
//Initialize request parameters
var reqData model.AccountGetNonceRequest
var address string = "buQtfFxpQP9JCFgmu4WBojBbEnVyQGaJDgGn"
reqData.SetAddress(address)
//Call GetNonce interface
resData := testSdk.Account.GetNonce(reqData)
```

#### Building operations

The operation refers to some of the actions that are done in the transaction to facilitate serialization of transactions and evaluation of fees. For more details, see [Operations](#operations). For example, to build an operation to send BU (`BUSendOperation`), the specific interface call is as follows:
```go 
var buSendOperation model.BUSendOperation
buSendOperation.Init()
var amount int64 = 100
var address string = "buQVU86Jm4FeRW4JcQTD9Rx9NkUkHikYGp6z"
buSendOperation.SetAmount(amount)
buSendOperation.SetDestAddress(address)
```

#### Serializing transactions

The `transaction serialization` interface is used to serialize transactions and generate transaction blob strings for network transmission. The nonce value and operation are obtained from the interface called. For interface details, see [buildBlob](#buildblob), which calls as follows:
```go 
//Initialize request parameters
var reqDataBlob model.TransactionBuildBlobRequest
reqDataBlob.SetSourceAddress(sourceAddress)
reqDataBlob.SetFeeLimit(feeLimit)
reqDataBlob.SetGasPrice(gasPrice)
reqDataBlob.SetNonce(senderNonce)
reqDataBlob.SetOperation(buSendOperation)
//Call BuildBlob interface
resDataBlob := testSdk.Transaction.BuildBlob(reqDataBlob)
}
```

#### Signing transactions

The `signature transaction` interface is used by the transaction initiator to sign the transaction using the private key of the account. The transactionBlob is obtained from the interface called. For interface details, see [sign](#sign), which calls as follows:
```go
//Initialize request parameters
PrivateKey := []string{"privbUPxs6QGkJaNdgWS2hisny6ytx1g833cD7V9C3YET9mJ25wdcq6h"}
var reqData model.TransactionSignRequest
reqData.SetBlob(resDataBlob.Result.Blob)
reqData.SetPrivateKeys(PrivateKey)
//Call Sign interface
resDataSign := testSdk.Transaction.Sign(reqData)
}
```

#### Submitting transactions

The `submit interface` is used to send a transaction request to the BU blockchain, triggering the execution of the transaction. transactionBlob and signResult are obtained from the interfaces called. For interface details, see [submit](#submit), which calls as follows:
```go
//Initialize request parameters
var reqData model.TransactionSubmitRequest
reqData.SetBlob(resDataBlob.Result.Blob)
reqData.SetSignatures(resDataSign.Result.Signatures)
//Call Submit interface
resDataSubmit := testSdk.Transaction.Submit(reqData)
```

## Transaction Service

Transaction Service provide transaction-related interfaces and currently have five interfaces: `BuildBlob`, `EvaluateFee`, `sign`, `Submit`, and `GetInfo`.

### buildBlob

> **Note:** Before you call **buildBlob**, you shold make some operations, details for [Operations](#operations).

- **Interface description**

   The `buildBlob` interface is used to serialize transactions and generate transaction blob strings for network transmission.

- **Calling method**

  `BuildBlob(model.TransactionBuildBlobRequest)model.TransactionBuildBlobResponse;`

- **Request parameters**

   Parameter      |     Type     |        Description       
   ----------- | ------------ | ---------------- 
   sourceAddress|String|Required, the source account address initiating the operation
   nonce|int64|Required, the transaction serial number to be initiated, add 1 in the function, size limit [1, max(int64)]
   gasPrice|int64|Required, transaction gas price, unit MO, 1 BU = 10^8 MO, size limit [1000, max(int64)]
   feeLimit|int64|Required, the minimum fees required for the transaction, unit MO, 1 BU = 10^8 MO, size limit [1, max(int64)]
   operation|`[]`BaseOperation|Required, list of operations to be committed which cannot be empty
   ceilLedgerSeq|int64|Optional, set a value which will be combined with the current block height to restrict transactions. If transactions do not complete within the set value plus the current block height, the transactions fail. The value you set must be greater than 0. If the value is set to 0, no limit is set.
   metadata|String|Optional, note

- **Response data**

   Parameter      |     Type     |        Description       
   ----------- | ------------ | ---------------- 
   transactionBlob|String|Serialized transaction hex string
   hash|String|Transaction hash

- **Error code**

   Error Message      |     Error Code     |        Description   
   -----------  | ----------- | -------- 
   INVALID_SOURCEADDRESS_ERROR|11002|Invalid sourceAddress
   INVALID_NONCE_ERROR|11048|Nonce must be between 1 and max(int64)
   INVALID_DESTADDRESS_ERROR|11003|Invalid destAddress
   INVALID_INITBALANCE_ERROR|11004|InitBalance must be between 1 and max(int64) 
   SOURCEADDRESS_EQUAL_DESTADDRESS_ERROR|11005|SourceAddress cannot be equal to destAddress
   INVALID_ISSUE_AMOUNT_ERROR|11008|AssetAmount this will be issued must be between 1 and max(int64)
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
   INVALID_ GASPRICE_ERROR|11049|GasPrice must be between 1000 and max(int64)
   INVALID_FEELIMIT_ERROR|11050|FeeLimit must be between 1 and max(int64)
   OPERATIONS_EMPTY_ERROR|11051|Operations cannot be empty
   INVALID_CEILLEDGERSEQ_ERROR|11052|CeilLedgerSeq must be equal to or greater than 0
   OPERATIONS_ONE_ERROR|11053|One of the operations cannot be resolved
   REQUEST_NULL_ERROR|12001|Request parameter cannot be null
   SYSTEM_ERROR|20000|System error

- **Example**

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

- **Interface description**

   The `evaluateFee` interface implements the cost estimate for the transaction.

- **Calling method**

  `EvaluateFee(model.TransactionEvaluateFeeRequest)model.TransactionEvaluateFeeResponse;`

- **Request parameters**

   Parameter      |     Type     |        Description       
   ----------- | ------------ | ---------------- 
   sourceAddress|String|Required, the source account address initiating the operation
   nonce|int64|Required, transaction serial number to be initiated, size limit [1, max(int64)]
   operation|`[]`BaseOperation|Required, list of operations to be committed which cannot be empty
   signtureNumber|String|Optional, the number of people to sign, the default is 1, size limit [1, max(int32)]
   ceilLedgerSeq|int64|Optional, set a value which will be combined with the current block height to restrict transactions. If transactions do not complete within the set value plus the current block height, the transactions fail. The value you set must be greater than 0. If the value is set to 0, no limit is set.
   metadata|String|Optional, note

- **Response data**

   Parameter      |     Type     |        Description       
   ----------- | ------------ | ---------------- 
   txs     |   `[]`[TestTx](#testtx)   |  Evaluation transaction set

- **Error code**

   Error Message      |     Error Code     |        Description   
   -----------  | ----------- | -------- 
   INVALID_SOURCEADDRESS_ERROR|11002|Invalid sourceAddress
   INVALID_NONCE_ERROR|11048|Nonce must be between 1 and max(int64)
   OPERATIONS_EMPTY_ERROR|11051|Operations cannot be empty
   OPERATIONS_ONE_ERROR|11053|One of the operations cannot be resolved
   INVALID_SIGNATURENUMBER_ERROR|11054|SignagureNumber must be between 1 and max(int32)
   SYSTEM_ERROR|20000|System error

- **Example**

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

- **Interface description**

   The `sign` interface is used to implement the signature of the transaction.

- **Calling method**

   `Sign(model.TransactionSignRequest) model.TransactionSignResponse;`

- **Request parameters**

   Parameter      |     Type     |        Description       
   ----------- | ------------ | ---------------- 
   blob|String|Required, pending transaction blob to be signed
   privateKeys|`[]`String|Required, private key list


- **Response data**

   Parameter      |     Type     |        Description       
   ----------- | ------------ | ---------------- 
   signatures|`[]`[Signature](#signature)|Signed data list

- **Error code**

   Error Message      |     Error Code     |        Description   
   -----------  | ----------- | -------- 
   INVALID_BLOB_ERROR|11056|Invalid blob
   PRIVATEKEY_NULL_ERROR|11057|PrivateKeys cannot be empty
   PRIVATEKEY_ONE_ERROR|11058|One of privateKeys is invalid
   GET_ENCPUBLICKEY_ERROR|14000|The function ‘GetEncPublicKey’ failed
   SIGN_ERROR|14001|The function ‘Sign’ failed
   SYSTEM_ERROR|20000|System error

- **Example**

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

- **Interface description**

   The `submit` interface is used to implement the submission of the transaction.

- **Calling method**

  `Submit(model.TransactionSubmitRequest) model.TransactionSubmitResponse;`

- **Request parameters**

   Parameter      |     Type     |        Description       
   ----------- | ------------ | ---------------- 
   blob|String|Required, transaction blob
   signature|`[]`[Signature](#signature)|Required, signature list

- **Response data**

   Parameter      |     Type     |        Description       
   ----------- | ------------ | ---------------- 
   hash|String|Transaction hash

- **Error code**

   Error Message      |     Error Code     |        Description   
   -----------  | ----------- | -------- 
   INVALID_BLOB_ERROR|11056|Invalid blob
   SIGNATURE_EMPTY_ERROR|11067|The signatures cannot be empty
   SYSTEM_ERROR|20000|System error

- **Example**

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

- **Interface description**

   The `getInfo` interface is used to implement query transactions based on transaction hashes.

- **Calling method**

  `GetInfo(model.TransactionGetInfoRequest)model.TransactionGetInfoResponse;`

- **Request parameters**

   Parameter      |     Type     |        Description       
   ----------- | ------------ | ---------------- 
   hash|String|Transaction hash

- **Response data**

   Parameter      |     Type     |        Description       
   ----------- | ------------ | ---------------- 
   totalCount|int64|Total number of transactions returned
   transactions|`[]`[TransactionHistory](#transactionhistory)|Transaction content

- **Error code**

   Error Message      |     Error Code     |        Description   
   -----------  | ----------- | -------- 
   INVALID_HASH_ERROR|11055|Invalid transaction hash
   CONNECTNETWORK_ERROR|11007|Failed to connect to the network
   SYSTEM_ERROR|20000|System error

- **Example**

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

## Operations

Operations refer to the things that are to be done in a transaction, and the operations that need to be built before the operations are to be built. At present, there are 10 kinds of operations, which include [AccountActivateOperation](#accountactivateoperation)、[AccountSetMetadataOperation](#accountsetmetadataoperation)、 [AccountSetPrivilegeOperation](#accountsetprivilegeoperation)、 [AssetIssueOperation](#assetissueoperation)、 [AssetSendOperation](#assetsendoperation)、 [BUSendOperation](#busendoperation)、 [ContractCreateOperation](#contractcreateoperation)、 [ContractInvokeByAssetOperation](#contractinvokebyassetoperation)、 [ContractInvokeByBUOperation](#contractinvokebybuoperation)、 [LogCreateOperation](#logcreateoperation).


**BaseOperation**

BaseOperation is the base class for all operations in the buildBlob interface. The following table describes BaseOperation:

Member    |     Type  |        Description                           
------------- | -------- | ----------------------------------   
sourceAddress |   String |  Optional, source account address of the operation
metadata      |   String |  Optional, note

### AccountActivateOperation

- Function

  This operation is used to activate an account. AccountActivateOperation inherits from BaseOperation.

- Fee

  FeeLimit is currently fixed at 0.01 BU (2018.07.26).

- Member

   Member    |     Type  |        Description                           
   ------------- | -------- | ---------------------------------- 
   sourceAddress |   String |  Optional, source account address of the operation 
   destAddress   |   String |  Required, target account address
   initBalance   |   int64   |  Required, initialize the asset, unit MO, 1 BU = 10^8 MO, size (0, max(int64)] 
   metadata|String|Optional, note

### AccountSetMetadataOperation

- Function

  This operation is used to set the metadata of an account. AccountSetMetadataOperation inherits from BaseOperation.

- Fee

  FeeLimit is currently fixed at 0.01 BU (2018.07.26).

- Member

   Member    |     Type  |        Description                         
   ------------- | --------- | ------------------------------- 
   sourceAddress |   String |  Optional, source account address of the operation
   key           |   String  |  Required, metadata keyword, length limit [1, 1024]
   value         |   String  |  Required, metadata content, length limit [0, 256000]
   version       |   int64    |  Optional, metadata version
   deleteFlag    |   Boolean |  Optional, whether to delete metadata
   metadata|String|Optional, note           

### AccountSetPrivilegeOperation

- Function

  This operation is used to set the privilege of an account. AccountSetPrivilegeOperation inherits from BaseOperation.

- Fee

  FeeLimit is currently fixed at 0.01 BU (2018.07.26).

- Member

   Member    |     Type  |        Description               
   ------------- | --------- | --------------------------
   sourceAddress |   String |  Optional, source account address of the operation
   masterWeight|String|	Optional, account weight, size limit [0, max(uint32)]
   signers|[Signer](#signer)[]|Optional, signer weight list
   txThreshold|String|Optional, transaction threshold, size limit [0, max(int64)]
   typeThreshold|[TypeThreshold](#typethreshold)[]|Optional, specify transaction threshold
   metadata|String|Optional, note

### AssetIssueOperation

- Function

  This operation is used to issue assets. AssetIssueOperation inherits from BaseOperation.

- Fee

  FeeLimit is currently fixed at 50.01 BU (2018.07.26).

- Member

   Member    |     Type  |        Description             
   ------------- | --------- | ------------------------
   sourceAddress|String|Optional, source account address of the operation
   code|String|Required, asset code, length limit [1, 64]
   assetAmount|int64|Required, asset code, length limit [0, max(int64)]
   metadata|String|Optional, note

### AssetSendOperation

> **Note**: If the destination account is not activated, the activation account operation must be invoked first.

- Function

  This operation is used to send assets. AssetSendOperation inherits from BaseOperation.

- Fee

  FeeLimit is currently fixed at 0.01 BU (2018.07.26).

- Member

   Member    |     Type  |        Description            
   ------------- | --------- | ----------------------
   sourceAddress|String|Optional, source account address of the operation
   destAddress|String|Required, target account address
   code|String|Required, asset code, length limit [1, 64]
   issuer|String|Required, the account address for issuing assets
   assetAmount|int64|Required, asset amount, size limit [0, max(int64)]
   metadata|String|Optional, note

### BUSendOperation

> **Note**: If the destination account is not activated, this operation will activate this account.

- Function

  This operation is used to send BUs. BUSendOperation inherits from BaseOperation.

- Fee

  FeeLimit is currently fixed at 0.01 BU (2018.07.26).

- Member

   Member    |     Type  |        Description          
   ------------- | --------- | ---------------------
   sourceAddress|String|Optional, source account address of the operation
   destAddress|String|Required, target account address
   buAmount|int64|Required, the amount of BU to be transferred, length limit [0, max(int64)]
   metadata|String|Optional, note

### ContractCreateOperation

- Function

  This operation is used to create a contract. ContractCreateOperation inherits from BaseOperation.

- Fee

  FeeLimit is currently fixed at 10.01 BU (2018.07.26).

- Member

   Member    |     Type  |        Description          
   ------------- | --------- | ---------------------
   sourceAddress|String|Optional, source account address of the operation
   initBalance|int64|Required, initial asset for contract account, unit MO, 1 BU = 10^8 MO, size limit [1, max(int64)]
   type|Integer|Optional, the language of the contract, the default is 
   payload|String|Required, contract code for the corresponding language
   initInput|String|Optional, the input parameters of the init method in the contract code
   metadata|String|Optional, note

### ContractInvokeByAssetOperation

> **Note**: If the destination account is not activated, the activation account operation must be invoked first.

- Function

  This operation is used to send assets and invoke a contract. ContractInvokeByAssetOperation inherits from BaseOperation.

- Fee

  FeeLimit requires to add fees according to the execution of the transaction in the contract. First, the transaction fee is initiated. At present the fee (2018.07.26) is 0.01BU, and then the transaction in the contract also requires the transaction initiator to add the transaction fees.

- Member

   Member    |     Type  |        Description          
   ------------- | --------- | ---------------------
   sourceAddress|String|Optional, source account address of the operation
   contractAddress|String|Required, contract account address
   code|String|Optional, asset code, length limit [0, 1024]; when it is empty, only the contract is triggered
   issuer|String|Optional, the account address issuing assets; when it is null, only trigger the contract
   assetAmount|int64|Optional, asset amount, size limit[0, max(int64)]when it is 0, only trigger the contract
   input|String|Optional, the input parameter of the main() method for the contract to be triggered
   metadata|String|Optional, note

### ContractInvokeByBUOperation

> **Note**: If the destination account is not a contract and it is not activated, this operation will activate this account.

- Function

  This operation is used to send BUs and invoke an contract. ContractInvokeByBUOperation inherits from BaseOperation.

- Fee

  FeeLimit requires to add fees according to the execution of the transaction in the contract. First, the transaction fee is initiated. At present the fee (2018.07.26) is 0.01BU, and then the transaction in the contract also requires the transaction initiator to add the transaction fees.

- Member

   Member    |     Type  |        Description          
   ------------- | --------- | ---------------------
   sourceAddress|String|Optional, source account address of the operation
   contractAddress|String|Required, contract account address
   buAmount|int64|Optional, the amount of BU to be transferred, size limit [0, max(int64)], when it is 0 only triggers the contract
   input|String|Optional, the input parameter of the main() method for the contract to be triggered
   metadata|String|Optional, note

### LogCreateOperation

- Function

  This operation is used to record a log. LogCreateOperation inherits from BaseOperation.

- Fee

  FeeLimit is currently fixed at 0.01 BU (2018.07.26).

- Member

   Member    |     Type  |        Description          
   ------------- | --------- | ---------------------
   sourceAddress|String|Optional, source account address of the operation
   topic|String|Required, Log theme，length limit [1, 128]
   datas|List<String>|Required, Log content，length limit of each string [1, 1024]
   metadata|String|Optional, note



## Account Service

Account Service provide account-related interfaces, which include six interfaces: `CheckValid`, `GetInfo`, `GetNonce`, `GetBalance`, `GetAssets`, and `GetMetadata`.

### CheckValid

- **Interface description**

   The `create` interface in account service can generate private key, public key and address of an new account.

- **Calling method**

  `CheckValid(model.AccountCheckValidRequest) model.AccountCheckValidResponse;`

- **Request parameters**

   Parameter      |     Type     |        Description       
   ----------- | ------------ | ---------------- 
   address     |   String     |  Required, the account address to be checked on the blockchain

- **Response data**

   Parameter      |     Type     |        Description       
   ----------- | ------------ | ---------------- 
   isValid     | Boolean |  Whether the response data is valid   

- **Error code**

   Error Message      |     Error Code     |        Description   
   -----------  | ----------- | -------- 
   SYSTEM_ERROR |   20000     |  System error 

- **Example**

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

- **Interface description**

   The getInfo interface is used to obtain the specified account information.

- **Calling method**

  `GetInfo(model.AccountGetInfoRequest) model.AccountGetInfoResponse;`

- **Request parameters**

   Parameter      |     Type     |        Description       
   ----------- | ------------ | ---------------- 
   address     |   String     |  Required, the account address to be queried on the blockchain 

- **Response data**

   Parameter      |     Type     |        Description       
   --------- | ------------- | ---------------- 
   address	  |    String     |    Account address       
   balance	  |    int64       |    Account balance, unit is MO, 1 BU = 10^8 MO, the account balance must be > 0
   nonce	  |    int64       |    Account transaction serial number must be greater than 0
   priv	  | [Priv](#priv) |    Account privilege

- **Error code**

   Error Message      |     Error Code     |        Description   
   -----------  | ----------- | -------- 
   INVALID_ADDRESS_ERROR| 11006 | Invalid address
   CONNECTNETWORK_ERROR| 11007| Failed to connect to the network
   SYSTEM_ERROR |   20000     |  System error 

- **Example**

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

- **Interface description**

   The `getNonce` interface is used to obtain the nonce value of the specified account.

- **Calling method**

  `GetNonce(model.AccountGetNonceRequest)model.AccountGetNonceResponse;`

- **Request parameters**

   Parameter      |     Type     |        Description       
   ----------- | ------------ | ---------------- 
   address     |   String     |  Required, the account address to be queried on the blockchain 

- **Response data**

   Parameter      |     Type     |        Description       
   ----------- | ------------ | ---------------- 
   nonce       |   int64       |  Account transaction serial number

- **Error code**

   Error Message      |     Error Code     |        Description   
   -----------  | ----------- | -------- 
   INVALID_ADDRESS_ERROR| 11006 | Invalid address
   CONNECTNETWORK_ERROR| 11007| Failed to connect to the network
   SYSTEM_ERROR |   20000     |  System error 

- **Example**

   ```go
   var reqData model.AccountGetNonceRequest
   var address string = "buQtfFxpQP9JCFgmu4WBojBbEnVyQGaJDgGn"
   reqData.SetAddress(address)
   if resData.ErrorCode == 0 {
   fmt.Println(resData.Result.Nonce)
   }
   ```

### GetBalance

- **Interface description**

   The `getBalance` interface is used to obtain the BU balance of the specified account.

- **Calling method**

  `GetBalance(model.AccountGetBalanceRequest)model.AccountGetBalanceResponse;`

- **Request parameters**

   Parameter      |     Type     |        Description       
   ----------- | ------------ | ---------------- 
   address     |   String     |  Required, the account address to be queried on the blockchain 

- **Response data**

   Parameter      |     Type     |        Description       
   ----------- | ------------ | ---------------- 
   balance     |   int64       | BU balance, unit MO, 1 BU = 10^8 MO

- **Error code**

   Error Message      |     Error Code     |        Description   
   -----------  | ----------- | -------- 
   INVALID_ADDRESS_ERROR| 11006 | Invalid address
   CONNECTNETWORK_ERROR| 11007| Failed to connect to the network
   SYSTEM_ERROR |   20000     |  System error 

- **Example**

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

- **Interface description**

   The `getAssets` interface is used to get all the asset information of the specified account.

- **Calling method**

  `GetAssets(model.AccountGetAssetsRequest)model.AccountGetAssetsResponse;`

- **Request parameters**

   Parameter      |     Type     |        Description       
   ----------- | ------------ | ---------------- 
   address     |   String     |  Required, the account address to be queried   

- **Response data**

   Parameter      |     Type     |        Description       
   ----------- | ------------ | ---------------- 
   asset	    | `[]`[AssetInfo](#assetinfo) |Account asset

- **Error code**

   Error Message      |     Error Code     |        Description   
   -----------  | ----------- | -------- 
   INVALID_ADDRESS_ERROR| 11006 | Invalid address
   CONNECTNETWORK_ERROR| 11007| Failed to connect to the network
   NO_ASSET_ERROR|11009|The account does not have the asset
   SYSTEM_ERROR|20000|System error

- **Example**

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

- **Interface description**

   The `getMetadata` interface is used to obtain the metadata information of the specified account.

- **Calling method**

  `GetMetadata(model.AccountGetMetadataRequest)model.AccountGetMetadataResponse;`

- **Request parameters**

   Parameter      |     Type     |        Description       
   -------- | -------- | ---------------- 
   address  |  String  |  Required, the account address to be queried  
   key      |  String  |  Optional, metadata keyword, length limit [1, 1024]

- **Response data**

   Parameter      |     Type     |        Description       
   ----------- | ----------- | ---------------- 
   metadata    |`[]`[Metadata](#metadata)   | Account metadata

- **Error code**

   Error Message      |     Error Code     |        Description   
   -----------  | ----------- | -------- 
   INVALID_ADDRESS_ERROR | 11006 | Invalid address
   CONNECTNETWORK_ERROR | 11007 | Failed to connect to the network
   NO_METADATA_ERROR|11010|The account does not have the metadata
   INVALID_DATAKEY_ERROR | 11011 | The length of key must be between 1 and 1024
   SYSTEM_ERROR | 20000| System error


- **Example**

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

## Asset Service

Asset Service follow the ATP 1.0 protocol, and Account Service provide an asset-related interface. Currently there is one interface: `GetInfo`.

### getInfo

- **Interface description**

   The `getInfo` interface is used to obtain the specified asset information of the specified account.

- **Calling method**

  `GetInfo(model.AssetGetInfoRequest) model.AssetGetInfoResponse;`

- **Request parameters**

   Parameter      |     Type     |        Description       
   ----------- | ------------ | ---------------- 
   address     |   String    |  Required, the account address to be queried
   code        |   String    |  Required, asset code, length limit [1, 64]
   issuer      |   String    |  Required, the account address for issuing assets

- **Response data**

   Parameter      |     Type     |        Description       
   ----------- | ------------ | ---------------- 
   asset	    | `[]`[AssetInfo](#assetinfo) |Account asset   

- **Error code**

   Error Message      |     Error Code     |        Description   
   -----------  | ----------- | -------- 
   INVALID_ADDRESS_ERROR|11006|Invalid address
   CONNECTNETWORK_ERROR|11007|Failed to connect to the network
   INVALID_ASSET_CODE_ERROR|11023|The length of asset code must be between 1 and 64
   INVALID_ISSUER_ADDRESS_ERROR|11027|Invalid issuer address
   SYSTEM_ERROR|20000|System error

- **Example**

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

## Contract Service

Contract Service provide contract-related interfaces and currently have four interfaces: `CheckValid`, `GetInfo`, `GetAddress`, and `Call`.

### checkValid

- **Interface description**

   The `checkValid` interface is used to check the validity of the contract account.

- **Calling method**

  `CheckValid(reqData model.ContractCheckValidRequest) model.ContractCheckValidResponse;`

- **Request parameters**

   Parameter      |     Type     |        Description       
   ----------- | ------------ | ---------------- 
   contractAddress     |   String     |  Contract account address to be tested

- **Response data**

   Parameter      |     Type     |        Description       
   ----------- | ------------ | ---------------- 
   isValid     |   Boolean     |  Whether the response data is valid   

- **Error code**

   Error Message      |     Error Code     |        Description   
   -----------  | ----------- | -------- 
   INVALID_CONTRACTADDRESS_ERROR|11037|Invalid contract address
   CONTRACTADDRESS_NOT_CONTRACTACCOUNT_ERROR|11038|ContractAddress is not a contract account
   SYSTEM_ERROR |   20000     |  System error 

- **Example**

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

### getInfo

- **Interface description**

   The `getInfo` interface is used to query the contract code.

- **Calling method**

  `GetInfo(model.ContractGetInfoRequest) model.ContractGetInfoResponse;`

- **Request parameters**

   Parameter      |     Type     |        Description       
   ----------- | ------------ | ---------------- 
   contractAddress     |   String     |  Contract account address to be queried   

- **Response data**

   Parameter      |     Type     |        Description       
   ----------- | ------------ | ---------------- 
   contract|[ContractInfo](#contractinfo)|Contract info

- **Error code**

   Error Message      |     Error Code     |        Description   
   -----------  | ----------- | -------- 
   INVALID_CONTRACTADDRESS_ERROR|11037|Invalid contract address
   CONTRACTADDRESS_NOT_CONTRACTACCOUNT_ERROR|11038|contractAddress is not a contract account
   NO_SUCH_TOKEN_ERROR|11030|No such token
   GET_TOKEN_INFO_ERROR|11066|Failed to get token info
   SYSTEM_ERROR|20000|System error

- **Example**

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

- **Interface description**

  The `getAddress` interface is used to query the contract address.

- **Calling method**

  `GetAddress(reqData model.ContractGetAddressRequest) model.ContractGetAddressResponse;`

- **Request parameters**

   Parameter      |     Type     |        Description       
   ----------- | ------------ | ---------------- 
   hash     |   String     |  The hash used to create a contract transaction   

- **Response data**

   Parameter      |     Type     |        Description       
   ----------- | ------------ | ---------------- 
   contractAddressList|List<[ContractAddressInfo](#contractaddressinfo)>|Contract address list

- **Error code**

   Error Message      |     Error Code     |        Description   
   -----------  | ----------- | -------- 
   INVALID_HASH_ERROR|11055|Invalid transaction hash
   CONNECTNETWORK_ERROR|11007|Failed to connect to the network
   SYSTEM_ERROR|20000|System error

- **Example**

   ```go
   // Initialize request parameters
   var reqData model.ContractGetAddressRequest();
   reqData.SetAddress("44246c5ba1b8b835a5cbc29bdc9454cdb9a9d049870e41227f2dcfbcf7a07689");

   resData := sdk.Contract.GetAddress(reqData);
   if resData.ErrorCode == 0 {
   fmt.Println("Address:", resData.Result.Address);
   }
   ```

### Call 

- **Interface description**

   The `call` interface is used to debug the contract code.

- **Calling method**

  `Call(reqData model.ContractCallRequest) model.ContractCallResponse;`

- **Request parameters**

   Parameter      |     Type     |        Description       
   ----------- | ------------ | ---------------- 
   sourceAddress|String|Optional, the account address to trigger the contract
   contractAddress|String|Optional, the contract account address and code cannot be empty at the same time
   code|String|Optional, the contract code and contractAddress cannot be empty at the same time
   input|String|Optional, input parameter for the contract
   contractBalance|int64|Optional, the initial BU balance given to the contract, unit MO, 1 BU = 10^8 MO, size limit [1, max(int64)]
   optType|Integer|Required, 0: Call the read/write interface of the contract init, 1: Call the read/write interface of the contract main, 2: Call the read-only interface query
   feeLimit|int64|Minimum fee required for the transaction, size limit [1, max(int64)]
   gasPrice|int64|Transaction fuel price, size limit [1000, max(int64)]


- **Response data**

   Parameter      |     Type     |        Description       
   ----------- | ------------ | ---------------- 
   logs|JSONObject|Log information
   queryRets|JSONArray|Query the result set
   stat|[ContractStat](#contractstat)|Contract resource occupancy
   txs|`[]`[TransactionEnvs](#transactionenvs)|Transaction set

- **Error code**

   Error Message      |     Error Code     |        Description   
   -----------  | ----------- | -------- 
   INVALID_SOURCEADDRESS_ERROR|11002|Invalid sourceAddress
   INVALID_CONTRACTADDRESS_ERROR|11037|Invalid contract address
   CONTRACTADDRESS_CODE_BOTH_NULL_ERROR|11063|ContractAddress and code cannot be empty at the same time
   INVALID_OPTTYPE_ERROR|11064|OptType must be between 0 and 2
   CONNECTNETWORK_ERROR|11007|Failed to connect to the network
   SYSTEM_ERROR|20000|System error

- **Example**

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

## Block service

Block service provide block-related interfaces. There are currently 11 interfaces: `GetNumber`, `CheckStatus`, `GetTransactions`, `GetInfo`, `GetLatest`, `GetValidators`, `GetLatestValidators`, `GetReward`, `GetLatestReward`, `GetFees`, and `GetLatestFees`.

### getNumber

- **Interface description**

   The `getNumber` interface is used to query the latest block height.

- **Calling method**

  `GetNumber() model.BlockGetNumberResponse;`

- **Response data**

   Parameter      |     Type     |        Description       
   ----------- | ------------ | ---------------- 
   header|BlockHeader|Block head
   blockNumber|int64|The latest block height,corresponding to the underlying field sequence

- **Error code**

   Error Message      |     Error Code     |        Description   
   -----------  | ----------- | -------- 
   CONNECTNETWORK_ERROR|11007|Failed to connect to the network
   SYSTEM_ERROR|20000|System error

- **Example**

   ```go 
   resData := testSdk.Block.GetNumber()
   if resData.ErrorCode == 0 {
      fmt.Println("BlockNumber:", resData.Result.BlockNumber)
   }
   ```

### checkStatus

- **Interface description**

   The `checkStatus` interface is used to check if the local node block is synchronized.

- **Calling method**

  `CheckStatus() model.BlockCheckStatusResponse;`

- **Response data**

   Parameter      |     Type     |        Description       |
   ----------- | ------------ | ---------------- |
   isSynchronous    |   Boolean     |  Whether the block is synchronized  |

- **Error code**

   Error Message      |     Error Code     |        Description   |
   -----------  | ----------- | -------- |
   CONNECTNETWORK_ERROR|11007|Failed to connect to the network
   SYSTEM_ERROR|20000|System error

- **Example**

   ```go
   resData := testSdk.Block.CheckStatus()
   if resData.ErrorCode == 0 {
      fmt.Println("IsSynchronous:", resData.Result.IsSynchronous)
   }
   ```

### getTransactions

- **Interface description**

   The `getTransactions` interface is used to query all transactions at the specified block height.

- **Calling method**

   `GetTransactions(model.BlockGetTransactionRequest)model.BlockGetTransactionResponse;`

- **Request parameters**

   Parameter      |     Type     |        Description       
   ----------- | ------------ | ---------------- 
   blockNumber|int64|Required, the height of the block to be queried must be greater than 0

- **Response data**

   Parameter      |     Type     |        Description       
   ----------- | ------------ | ---------------- 
   totalCount|int64|Total number of transactions returned
   transactions|`[]`[TransactionHistory](#transactionhistory)|Transaction content

- **Error code**

   Error Message      |     Error Code     |        Description   
   -----------  | ----------- | -------- 
   INVALID_BLOCKNUMBER_ERROR|11060|BlockNumber must bigger than 0
   CONNECTNETWORK_ERROR|11007|Failed to connect to the network
   SYSTEM_ERROR|20000|System error

- **Example**

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

- **Interface description**

   The `getInfo` interface is used to obtain block information.

- **Calling method**

  `GetInfo(model.BlockGetInfoRequest) model.BlockGetInfoResponse;`

- **Request parameters**

   Parameter      |     Type     |        Description       
   ----------- | ------------ | ---------------- 
   blockNumber|int64|Required, the height of the block to be queried

- **Response data**

   Parameter      |     Type     |        Description       
   ----------- | ------------ | ---------------- 
   closeTime|int64|Block closure time
   number|int64|Block height
   txCount|int64|Total transactions amount
   version|String|Block version

- **Error code**

   Error Message      |     Error Code     |        Description   
   -----------  | ----------- | -------- 
   INVALID_BLOCKNUMBER_ERROR|11060|BlockNumber must bigger than 0
   CONNECTNETWORK_ERROR|11007|Failed to connect to the network
   SYSTEM_ERROR|20000|System error

- **Example**

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

- **Interface description**

   The `getLatest` interface is used to get the latest block information.

- **Calling method**

  `GetLatest() model.BlockGetLatestResponse;`

- **Response data**

   Parameter      |     Type     |        Description       
   ----------- | ------------ | ---------------- 
   closeTime|int64|Block closure time
   number|int64|Block height,corresponding to the underlying field seq
   txCount|int64|Total transactions amount
   version|String|Block version


- **Error code**

   Error Message      |     Error Code     |        Description   
   -----------  | ----------- | -------- 
   CONNECTNETWORK_ERROR|11007|Failed to connect to the network
   SYSTEM_ERROR|20000|System error

- **Example**

   ```go
   resData := testSdk.Block.GetLatest()
   if resData.ErrorCode == 0 {
      data, _ := json.Marshal(resData.Result.Header)
      fmt.Println("Header:", string(data))
   }
   ```

### getValidators

- **Interface description**

   The `getValidators` interface is used to get the number of all the authentication nodes in the specified block.

- **Calling method**

  `GetValidators(model.BlockGetValidatorsRequest)model.BlockGetValidatorsResponse;`

- **Request parameters**

   Parameter      |     Type     |        Description       
   ----------- | ------------ | ---------------- 
   blockNumber|int64|Required, the height of the block to be queried must be greater than 0

- **Response data**

   Parameter      |     Type     |        Description       
   ----------- | ------------ | ---------------- 
   validators|`[]`String|Validators list

- **Error code**

   Error Message      |     Error Code     |        Description   
   -----------  | ----------- | -------- 
   INVALID_BLOCKNUMBER_ERROR|11060|BlockNumber must bigger than 0
   CONNECTNETWORK_ERROR|11007|Failed to connect to the network
   SYSTEM_ERROR|20000|System error

- **Example**

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

- **Interface description**

   `The getLatestValidators` interface is used to get the number of all validators in the latest block.

- **Calling method**

  `GetLatestValidators() model.BlockGetLatestValidatorsResponse;`

- **Response data**

   Parameter      |     Type     |        Description       
   ----------- | ------------ | ---------------- 
   validators|`[]`String|Validators list

- **Error code**

   Error Message      |     Error Code     |        Description   
   -----------  | ----------- | -------- 
   CONNECTNETWORK_ERROR|11007|Failed to connect to the network
   SYSTEM_ERROR|20000|System error

- **Example**

   ```go
   resData := testSdk.Block.GetLatestValidators()
   if resData.ErrorCode == 0 {
      data, _ := json.Marshal(resData.Result.Validators)
      fmt.Println("Validators:", string(data))
   }
   ```

### getReward

- **Interface description**

   The `getReward` interface is used to retrieve the block reward and valicator node rewards in the specified block.

- **Calling method**

  `GetReward(model.BlockGetRewardRequest) model.BlockGetRewardResponse;`

- **Request parameters**

   Parameter      |     Type     |        Description       
   ----------- | ------------ | ---------------- 
   blockNumber|int64|Required, the height of the block to be queried must be greater than 0

- **Response data**

   Parameter      |     Type     |        Description       
   ----------- | ------------ | ---------------- 
    validators | [Rewards](#Rewards)[] | Validators rewards 
    kols       | [Rewards](#Rewards)[] | Kols rewards       


- **Error code**

   Error Message      |     Error Code     |        Description   
   -----------  | ----------- | -------- 
   INVALID_BLOCKNUMBER_ERROR|11060|BlockNumber must bigger than 0
   CONNECTNETWORK_ERROR|11007|Failed to connect to the network
   SYSTEM_ERROR|20000|System error

- **Example**

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

- **Interface description**

   The `getLatestReward` interface gets the block rewards and validator rewards in the latest block. The method call is as follows:

- **Calling method**

  `GetLatestReward() model.BlockGetLatestRewardResponse;`

- **Response data**

   Parameter      |     Type     |        Description       
   ----------- | ------------ | ---------------- 
    validators | [Rewards](#Rewards)[] | Validators rewards 
    kols       | [Rewards](#Rewards)[] | Kols rewards       

- **Error code**

   Error Message      |     Error Code     |        Description   
   -----------  | ----------- | -------- 
   CONNECTNETWORK_ERROR|11007|Failed to connect to the network
   SYSTEM_ERROR|20000|System error

- **Example**

   ```go 
   resData := testSdk.Block.GetLatestReward()
   if resData.ErrorCode == 0 {
      fmt.Println("ValidatorsReward:", resData.Result.ValidatorsReward)
   }
   ```

### getFees

- **Interface description**

   The `getFees` interface gets the minimum asset limit and fuel price of the account in the specified block.

- **Calling method**

  `GetFees(model.BlockGetFeesRequest) model.BlockGetFeesResponse;`

- **Request parameters**

   Parameter      |     Type     |        Description       
   ----------- | ------------ | ---------------- 
   blockNumber|int64|Required, the height of the block to be queried must be greater than 0

- **Response data**

   Parameter      |     Type     |        Description       
   ----------- | ------------ | ---------------- 
   fees|[Fees](#fees)|Fees

- **Error code**

   Error Message      |     Error Code     |        Description   
   -----------  | ----------- | -------- 
   INVALID_BLOCKNUMBER_ERROR|11060|BlockNumber must bigger than 0
   CONNECTNETWORK_ERROR|11007|Failed to connect to the network
   SYSTEM_ERROR|20000|System error

- **Example**

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

- **Interface description**

   The `getLatestFees` interface is used to obtain the minimum asset limit and fuel price of the account in the latest block.

- **Calling method**

  `GetLatestFees() model.BlockGetLatestFeesResponse;`

- **Response data**

   Parameter      |     Type     |        Description       
   ----------- | ------------ | ---------------- 
   fees|[Fees](#fees)|Fees

- **Error code**

   Error Message      |     Error Code     |        Description   
   -----------  | ----------- | -------- 
   CONNECTNETWORK_ERROR|11007|Failed to connect to the network
   SYSTEM_ERROR|20000|System error

- **Example**

   ```go
   resData := testSdk.Block.GetLatestFees()
   if resData.ErrorCode == 0 {
      data, _ := json.Marshal(resData.Result.Fees)
      fmt.Println("Fees:", string(data))
   }
   ```



## Data Object

#### Priv

| Member       |     Type     |       Description                                                    |
| ------------ | ----------------------- | ------------------------------------------------------- |
| masterWeight | int64                    | Account weight, size limit [0,max(uint32)] |
| signers      | `[]`[Signer](#signer) | Signer weight list                                          |
| threshold    | [Threshold](#threshold) | Threshold                                                    |

#### Signer

| Member       |     Type     |       Description                                                  |
| ------- | ------ | ----------------------------------------------------- |
| address | String | The account address of the signer on the blockchain                                  |
| weight  | int64   | Signer weight, size limit [0,max(uint32)] |

#### Threshold

| Member       |     Type     |       Description                                     |
| -------------- | --------------------------------- | ----------------------------------------- |
| txThreshold    | int64                              | Transaction default threshold, size limit [0, max(int64)] |
| typeThresholds | `[]`[TypeThreshold](#typethreshold) | Thresholds for different types of transactions                        |

#### TypeThreshold

| Member       |     Type     |       Description                               |
| --------- | ---- | ----------------------------------- |
| type      | int64 | The operation type must be greater than 0                 |
| threshold | int64 | Threshold, size limit [0, max(int64)] |

#### AssetInfo

| Member       |     Type     |       Description         |
| ----------- | ----------- | ------------ |
| key         | [Key](#key) | Unique identifier for asset |
| assetAmount | int64        | Amount of assets     |

#### Key

| Member       |     Type     |       Description             |
| ------ | ------ | ---------------- |
| code   | String | Asset code         |
| issuer | String | The account address for issuing assets |

#### ContractInfo

Member       |     Type     |       Description      |
----------- | ------------ | ---------------- |
type|Integer|Contract type, default is 0
payload|String|Contract code

#### Metadata
Member       |     Type     |       Description       
----------- | ----------- | ---------------- 
key         |  String     |  Metadata keyword
value       |  String     |  Metadata content
version     |  int64      |  Metadata version

#### ContractAddressInfo

Member       |     Type     |       Description      
----------- | ------------ | ---------------- 
contractAddress|String|Contract address
operationIndex|Integer|The subscript of the operation

#### ContractStat

Member       |     Type     |       Description      
----------- | ------------ | ---------------- 
applyTime|int64|Receipt time
memoryUsage|int64|Memory footprint
stackUsage|int64|Stack occupancy
step|int64|Steps needed

#### TransactionEnvs

Member       |     Type     |       Description      
----------- | ------------ | ---------------- 
transactionEnv|[TransactionEnv](#transactionenv)|Transaction

#### TransactionEnv

Member       |     Type     |       Description      
----------- | ------------ | ---------------- 
transaction|[TransactionInfo](#transactioninfo)|Transaction content
trigger|[ContractTrigger](#contracttrigger)|Contract trigger

#### TransactionInfo

Member       |     Type     |       Description      
----------- | ------------ | ---------------- 
sourceAddress|String|The source account address initiating the transaction
feeLimit|int64|Minimum fees required for the transaction
gasPrice|int64|Transaction fuel price
nonce|int64|Transaction serial number
operations|`[]`[Operation](#operation)|Operations list

#### ContractTrigger
Member       |     Type     |       Description      
----------- | ------------ | ---------------- 
transaction|[TriggerTransaction](#triggertransaction)|Trigger transactions

#### Operation

Member       |     Type     |       Description      
----------- | ------------ | ---------------- 
type|Integer|Operation type
sourceAddress|String|The source account address initiating operations
metadata|String|Metadata
createAccount|[OperationCreateAccount](#operationcreateaccount)|Operation of creating accounts
issueAsset|[OperationIssueAsset](#operationissueasset)|Operation of issuing assets
payAsset|[OperationPayAsset](#operationpayasset)|Operation of transferring assets
payCoin|[OperationPayCoin](#operationpaycoin)|Operation of sending BU
setMetadata|[OperationSetMetadata](#operationsetmetadata)|Operation of setting metadata
setPrivilege|[OperationSetPrivilege](#operationsetprivilege)|Operation of setting account privilege
log|[OperationLog](#operationlog)|Record log

#### TriggerTransaction

Member       |     Type     |       Description      
----------- | ------------ | ---------------- 
hash|String|Transaction hash

#### OperationCreateAccount

Member       |     Type     |       Description      
----------- | ------------ | ---------------- 
destAddress|String|Target account address
contract|[Contract](#contract)|Contract info
priv|[Priv](#priv)|Account privilege
metadata|`[]`[Metadata](#metadata)|Account
initBalance|int64|Account assets, unit MO, 1 BU = 10^8 MO,
initInput|String|The input parameter for the init function of the contract

#### Contract

Member       |     Type     |       Description      
----------- | ------------ | ---------------- 
type|Integer| The contract language is not assigned value by default
payload|String|The contract code for the corresponding language

#### MetadataInfo

Member       |     Type     |       Description      
----------- | ------------ | ---------------- 
key|String|Metadata keyword
value|String|Metadata content
version|int64|Metadata version

#### OperationIssueAsset

Member       |     Type     |       Description      
----------- | ------------ | ---------------- 
code|String|Asset code
assetAmount|int64|Amount of assets

#### OperationPayAsset

Member       |     Type     |       Description      
----------- | ------------ | ---------------- 
destAddress|String|The target account address to which the asset is transferred
asset|[AssetInfo](#assetinfo)|Account asset
input|String|Input parameters for the main function of the contract

#### OperationPayCoin

Member       |     Type     |       Description      
----------- | ------------ | ---------------- 
destAddress|String|The target account address to which the asset is transferred
buAmount|int64|The amount of BU to be transferred
input|String|Input parameters for the main function of the contract

#### OperationSetMetadata

Member       |     Type     |       Description      
----------- | ------------ | ---------------- 
key|String|Metadata keyword
value|String|Metadata content
version|int64|Metadata version
deleteFlag|boolean|Whether to delete metadata

#### OperationSetPrivilege

Member       |     Type     |       Description      
----------- | ------------ | ---------------- 
masterWeight|String|Account weight, size limit [0,max(uint32)]
signers|`[]`[Signer](#signer)|Signer weight list
txThreshold|String|Transaction threshold, size limit [0, max(int64)]
typeThreshold|[TypeThreshold](#typethreshold)|Threshold for specified transaction type

#### OperationLog

Member       |     Type     |       Description      
----------- | ------------ | ---------------- 
topic|String|Log theme
data|String[]|Log content



#### TestTx

| Member       |     Type     |       Description         |
| -------------- | ------------------------------------------- | ------------ |
| transactionEnv | [TestTransactionFees](#testtransactionfees) | Assess transaction costs |

#### TestTransactionFees

| Member       |     Type     |       Description     |
| --------------- | ----------------------------------- | -------- |
| transactionFees | [TransactionFees](#transactionfees) | Transaction fees |

#### TransactionFees

| Member       |     Type     |       Description               |
| -------- | ---- | ------------------ |
| feeLimit | int64 | Minimum fees required for the transaction |
| gasPrice | int64 | Transaction fuel price       |

#### Signature

| Member       |     Type     |       Description       |
| --------- | ---- | ---------- |
| signData  | int64 | Signed data |
| publicKey | int64 | Public key       |

#### TransactionHistory

| Member       |     Type     |       Description         |
| ----------- | ----------------------------------- | ------------ |
| actualFee   | String                              | Actual transaction cost |
| closeTime   | int64                                | Transaction closure time |
| errorCode   | int64                                | Transaction error code   |
| errorDesc   | String                              | Transaction description     |
| hash        | String                              | Transaction hash     |
| ledgerSeq   | int64                                | Block serial number   |
| transaction | [TransactionInfo](#transactioninfo) | List of transaction contents |
| signatures  | `[]`[Signature](#signature)       | Signature list     |
| txSize      | int64                                | Transaction size     |

#### ValidatorReward

| Member       |     Type     |       Description        |
| --------- | ------ | ------------ |
| validator | String | Validator address |
| reward    | int64   | Validator reward |

#### Rewards

| Member  | Type   | Description  |
| ------- | ------ | ------------ |
| address | String | Node address |
| reward  | Array  | Node reward  |

#### ValidatorInfo

| Parameter      |     Type     |        Description         |
| --------------- | ------ | ------------ |
| address         | String | Consensus node address |

#### Fees

| Member       |     Type     |       Description                                 |
| ----------- | ---- | ------------------------------------ |
| baseReserve | int64 | Minimum asset limit for the account                     |
| gasPrice    | int64 | Transaction fuel price, unit MO, 1 BU = 10^8 MO |



## Error Code

Error Message      |     Error Code     |        Description   
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
INVALID_GASPRICE_ERROR|11049|GasPrice must be between 1000 and max(int64)
INVALID_FEELIMIT_ERROR|11050|FeeLimit must be between 1 and max(int64)
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

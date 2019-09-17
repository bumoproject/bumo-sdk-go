// response
package model

import (
	"github.com/bumoproject/bumo-sdk-go/src/crypto/protocol"
)

type deal struct {
	Items []Items `json:"items"`
}
type Items struct {
	TransactionBlob string      `json:"transaction_blob"`
	Signatures      []Signature `json:"signatures"`
}
type Signature struct {
	SignData  string `json:"sign_data"`
	PublicKey string `json:"public_key"`
}

// account
//CheckValid
type AccountCheckValidResponse struct {
	ErrorCode int              `json:"error_code"`
	ErrorDesc string           `json:"error_desc"`
	Result    CheckValidResult `json:"result"`
}
type CheckValidResult struct {
	IsValid bool
}

//CheckActivated
type AccountCheckActivatedResponse struct {
	ErrorCode int                  `json:"error_code"`
	ErrorDesc string               `json:"error_desc"`
	Result    CheckActivatedResult `json:"result"`
}
type CheckActivatedResult struct {
	IsActivated bool
}

//Create
type AccountCreateResponse struct {
	ErrorCode int                 `json:"error_code"`
	ErrorDesc string              `json:"error_desc"`
	Result    AccountCreateResult `json:"result"`
}
type AccountCreateResult struct {
	PrivateKey string `json:"private_key"`
	PublicKey  string `json:"public_key"`
	Address    string `json:"address"`
}
type AccountActivateResponse struct {
	ErrorCode int                   `json:"error_code"`
	ErrorDesc string                `json:"error_desc"`
	Result    AccountActivateResult `json:"result"`
}
type AccountActivateResult struct {
	Operation protocol.Operation `json:"operation"`
}
type AccountGetInfoResponse struct {
	ErrorCode int                  `json:"error_code"`
	ErrorDesc string               `json:"error_desc"`
	Result    AccountGetInfoResult `json:"result"`
}
type AccountGetInfoResult struct {
	Address  string   `json:"address"`
	Balance  int64    `json:"balance"`
	Nonce    int64    `json:"nonce"`
	Priv     Priv     `json:"priv"`
	Contract Contract `json:"contract"`
}
type Asset struct {
	Amount int64 `json:"amount"`
	Key    Key   `json:"key"`
}
type Key struct {
	Code   string `json:"code"`
	Issuer string `json:"issuer"`
}
type Priv struct {
	MasterWeight int64     `json:"master_weight"`
	Signers      []Signer  `json:"signers"`
	Thresholds   Threshold `json:"thresholds"`
}
type Signer struct {
	Address string `json:"address"`
	Weight  int64  `json:"weight"`
}
type Threshold struct {
	TxThreshold    int64           `json:"tx_threshold"`
	TypeThresholds []TypeThreshold `json:"type_thresholds"`
}
type TypeThreshold struct {
	Type      int64 `json:"type"`
	Threshold int64 `json:"threshold"`
}
type AccountGetNonceResponse struct {
	ErrorCode int                   `json:"error_code"`
	ErrorDesc string                `json:"error_desc"`
	Result    AccountGetNonceResult `json:"result"`
}
type AccountGetNonceResult struct {
	Nonce int64 `json:"nonce"`
}
type AccountGetBalanceResponse struct {
	ErrorCode int                     `json:"error_code"`
	ErrorDesc string                  `json:"error_desc"`
	Result    AccountGetBalanceResult `json:"result"`
}
type AccountGetBalanceResult struct {
	Balance int64 `json:"balance"`
}
type AccountSetMetadataResponse struct {
	ErrorCode int                      `json:"error_code"`
	ErrorDesc string                   `json:"error_desc"`
	Result    AccountSetMetadataResult `json:"result"`
}
type AccountSetMetadataResult struct {
	Operation protocol.Operation `json:"operation"`
}
type AccountSetPrivilegeResponse struct {
	ErrorCode int                       `json:"error_code"`
	ErrorDesc string                    `json:"error_desc"`
	Result    AccountSetPrivilegeResult `json:"result"`
}
type AccountSetPrivilegeResult struct {
	Operation protocol.Operation `json:"operation"`
}
type AccountGetAssetsResponse struct {
	ErrorCode int                    `json:"error_code"`
	ErrorDesc string                 `json:"error_desc"`
	Result    AccountGetAssetsResult `json:"result"`
}
type AccountGetAssetsResult struct {
	Assets []Asset `json:"assets"`
}
type AccountGetMetadataResponse struct {
	ErrorCode int                      `json:"error_code"`
	ErrorDesc string                   `json:"error_desc"`
	Result    AccountGetMetadataResult `json:"result"`
}
type AccountGetMetadataResult struct {
	Metadatas []Metadata `json:"metadatas"`
}
type Metadata struct {
	Key     string `json:"key"`
	Value   string `json:"value"`
	Version int64  `json:"version"`
}

// asset

type AssetIssueResponse struct {
	ErrorCode int                `json:"error_code"`
	ErrorDesc string             `json:"error_desc"`
	Result    AccountIssueResult `json:"result"`
}
type AccountIssueResult struct {
	Operation protocol.Operation `json:"operation"`
}
type AssetSendResponse struct {
	ErrorCode int             `json:"error_code"`
	ErrorDesc string          `json:"error_desc"`
	Result    AssetSendResult `json:"result"`
}
type AssetSendResult struct {
	Operation protocol.Operation `json:"operation"`
}
type AssetGetInfoResponse struct {
	ErrorCode int                `json:"error_code"`
	ErrorDesc string             `json:"error_desc"`
	Result    AssetGetInfoResult `json:"result"`
}
type AssetGetInfoResult struct {
	Assets []Asset `json:"assets"`
}

//BU
type BUSendResponse struct {
	ErrorCode int          `json:"error_code"`
	ErrorDesc string       `json:"error_desc"`
	Result    BUSendResult `json:"result"`
}
type BUSendResult struct {
	Operation protocol.Operation `json:"operation"`
}

//Contract
type ContractCreateResponse struct {
	ErrorCode int                  `json:"error_code"`
	ErrorDesc string               `json:"error_desc"`
	Result    ContractCreateResult `json:"result"`
}
type ContractCreateResult struct {
	Operation protocol.Operation `json:"operation"`
}

type ContractCheckValidResponse struct {
	ErrorCode int              `json:"error_code"`
	ErrorDesc string           `json:"error_desc"`
	Result    CheckValidResult `json:"result"`
}
type ContractGetInfoResponse struct {
	ErrorCode int          `json:"error_code"`
	ErrorDesc string       `json:"error_desc"`
	Result    GetPayResult `json:"result"`
}
type GetPayResult struct {
	Contract Contract `json:"contract"`
}
type Contract struct {
	Payload string `json:"payload"`
	Type    int64  `json:"type"`
}

type ContractCallResponse struct {
	ErrorCode int                `json:"error_code"`
	ErrorDesc string             `json:"error_desc"`
	Result    ContractCallResult `json:"result"`
}
type ContractCallResult struct {
	Logs      map[string]interface{} `json:"logs"`
	QueryRets []QueryRet             `json:"query_rets"`
	Stat      Stat                   `json:"stat"`
	Txs       []Tx                   `json:"txs"`
}
type Stat struct {
	ApplyTime   int64 `json:"apply_time"`
	MemoryUsage int64 `json:"memory_usage"`
	StackUsage  int64 `json:"stack_usage"`
	Step        int64 `json:"step"`
}
type ContractGetAddressResponse struct {
	ErrorCode int                      `json:"error_code"`
	ErrorDesc string                   `json:"error_desc"`
	Result    ContractGetAddressResult `json:"result"`
}
type ContractGetAddressResult struct {
	ContractAddresInfos []ContractAddresInfo
}
type ContractAddresInfo struct {
	ContractAddres string `json:"contract_address"`
	OperationIndex int    `json:"operation_index"`
}
type ContractInvokeByAssetResponse struct {
	ErrorCode int                 `json:"error_code"`
	ErrorDesc string              `json:"error_desc"`
	Result    InvokeByAssetResult `json:"result"`
}

type InvokeByAssetResult struct {
	Operation protocol.Operation `json:"operation"`
}
type ContractInvokeByBUResponse struct {
	ErrorCode int              `json:"error_code"`
	ErrorDesc string           `json:"error_desc"`
	Result    InvokeByBUResult `json:"result"`
}
type InvokeByBUResult struct {
	Operation protocol.Operation `json:"operation"`
}

// Transaction
type TransactionBuildBlobResponse struct {
	ErrorCode int             `json:"error_code"`
	ErrorDesc string          `json:"error_desc"`
	Result    BuildBlobResult `json:"result"`
}
type BuildBlobResult struct {
	Blob string `json:"transaction_blob"`
}
type WebTransactionEvaluateFeeResponse struct {
	Items []Item `json:"items"`
}
type Item struct {
	TransactionJson TransactionJson `json:"transaction_json"`
	SignatureNumber int64           `json:"signature_number"`
}
type TransactionJson struct {
	SourceAddress string             `json:"source_address"`
	Metadata      string             `json:"metadata"`
	Nonce         int64              `json:"nonce"`
	CeilLedgerSeq int64              `json:"ceil_ledger_seq"`
	Operations    []OperationEvaluat `json:"operations"`
}
type OperationEvaluat struct {
	Type          protocol.Operation_Type `protobuf:"varint,1,opt,name=type,enum=protocol.Operation_Type" json:"type,omitempty"`
	SourceAddress string                  `protobuf:"bytes,2,opt,name=source_address,json=sourceAddress" json:"source_address,omitempty"`
	Metadata      string                  `protobuf:"bytes,3,opt,name=metadata,proto3" json:"metadata,omitempty"`
	//
	CreateAccount        *protocol.OperationCreateAccount   `protobuf:"bytes,4,opt,name=create_account,json=createAccount" json:"create_account,omitempty"`
	IssueAsset           *protocol.OperationIssueAsset      `protobuf:"bytes,5,opt,name=issue_asset,json=issueAsset" json:"issue_asset,omitempty"`
	PayAsset             *protocol.OperationPayAsset        `protobuf:"bytes,6,opt,name=pay_asset,json=payAsset" json:"pay_asset,omitempty"`
	SetMetadata          *protocol.OperationSetMetadata     `protobuf:"bytes,7,opt,name=set_metadata,json=setMetadata" json:"set_metadata,omitempty"`
	SetSignerWeight      *protocol.OperationSetSignerWeight `protobuf:"bytes,8,opt,name=set_signer_weight,json=setSignerWeight" json:"set_signer_weight,omitempty"`
	SetThreshold         *protocol.OperationSetThreshold    `protobuf:"bytes,9,opt,name=set_threshold,json=setThreshold" json:"set_threshold,omitempty"`
	PayCoin              *protocol.OperationPayCoin         `protobuf:"bytes,10,opt,name=pay_coin,json=payCoin" json:"pay_coin,omitempty"`
	Log                  *protocol.OperationLog             `protobuf:"bytes,11,opt,name=log" json:"log,omitempty"`
	SetPrivilege         *protocol.OperationSetPrivilege    `protobuf:"bytes,12,opt,name=set_privilege,json=setPrivilege" json:"set_privilege,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                           `json:"-"`
	XXX_unrecognized     []byte                             `json:"-"`
	XXX_sizecache        int32                              `json:"-"`
}
type TransactionEvaluateFeeResponse struct {
	ErrorCode int               `json:"error_code"`
	ErrorDesc string            `json:"error_desc"`
	Result    EvaluateFeeResult `json:"result"`
}
type EvaluateFeeResult struct {
	GasPrice int64 `json:"gas_price"`
	FeeLimit int64 `json:"fee_limit"`
}
type TransactionEvaluateFeeData struct {
	ErrorCode int         `json:"error_code"`
	ErrorDesc string      `json:"error_desc"`
	Result    EvFeeResult `json:"result"`
}
type EvFeeResult struct {
	Txs []Tx `json:"txs"`
}
type Tx struct {
	TransactionEnv TransactionEnv `json:"transaction_env"`
}
type TransactionEnv struct {
	Transaction Transaction `json:"transaction"`
}
type TransactionSignResponse struct {
	ErrorCode int        `json:"error_code"`
	ErrorDesc string     `json:"error_desc"`
	Result    SignResult `json:"result"`
}
type SignResult struct {
	Signatures []Signature `json:"signatures"`
}
type TransactionSubmitResponse struct {
	ErrorCode int          `json:"error_code"`
	ErrorDesc string       `json:"error_desc"`
	Result    SubmitResult `json:"result"`
}
type SubmitResult struct {
	Hash string `json:"hash"`
}
type SubmitResults struct {
	Hash string `json:"hash"`
}
type TransactionSubmitData struct {
	Results      []SubmitsResult `json:"results"`
	SuccessCount int64           `json:"success_count"`
}
type SubmitsResult struct {
	ErrorCode int    `json:"error_code"`
	ErrorDesc string `json:"error_desc"`
	Hash      string `json:"hash"`
}
type TransactionGetInfoResponse struct {
	ErrorCode int            `json:"error_code"`
	ErrorDesc string         `json:"error_desc"`
	Result    GetInfoResults `json:"result"`
}
type GetInfoResults struct {
	TotalCount   int               `json:"total_count"`
	Transactions []Transactioninfo `json:"transactions"`
}

//Block
type BlockGetTransactionResponse struct {
	ErrorCode int                      `json:"error_code"`
	ErrorDesc string                   `json:"error_desc"`
	Result    GetTransactionInfoResult `json:"result"`
}
type GetTransactionInfoResult struct {
	TotalCount   int64             `json:"total_count"`
	Transactions []Transactioninfo `json:"Transactions"`
}
type Transactioninfo struct {
	ActualFee   int64       `json:"actual_fee"`
	CloseTime   int64       `json:"close_time"`
	ContractTxHashes []string `json:"contract_tx_hashes"`
	ErrorCode   int64       `json:"error_code"`
	ErrorDesc   string      `json:"error_desc"`
	Hash        string      `json:"hash"`
	LedgerSeq   int64       `json:"ledger_seq"`
	Signatures  []Signature `json:"signatures"`
	Transaction Transaction `json:"transaction"`
	TxSize      int64       `json:"tx_size"`
}
type Transaction struct {
	SourceAddress string      `json:"source_address"`
	Nonce         int64       `json:"nonce"`
	GasPrice      int64       `json:"gas_price"`
	FeeLimit      int64       `json:"fee_limit"`
	Metadata      string      `json:"metadata"`
	Operations    []Operation `json:"operations"`
}
type Operation struct {
	SourceAddress string        `json:"source_address"`
	Type          int64         `json:"type"`
	Metadata      string        `json:"metadata"`
	CreateAccount CreateAccount `json:"create_account"`
	IssueAsset    IssueAsset    `json:"issue_asset"`
	PayAsset      PayAsset      `json:"pay_asset"`
	PayCoin       PayCoin       `json:"pay_coin"`
	SetMetadata   SetMetadata   `json:"set_metadata"`
	SetPrivilege  SetPrivilege  `json:"set_privilege"`
	Log           Log           `json:"set_privilege"`
}
type CreateAccount struct {
	DestAddress string        `json:"dest_address"`
	Contract    Contract      `json:"contract"`
	Priv        Priv          `json:"priv"`
	Metadatas   []SetMetadata `json:"metadatas"`
	InitBalance int64         `json:"init_balance"`
	InitInput   string        `json:"init_input"`
}
type SetMetadata struct {
	Key        string `json:"key"`
	Value      string `json:"value"`
	Version    int64  `json:"version"`
	DeleteFlag bool   `json:"delete_flag"`
}

type IssueAsset struct {
	Code   string `json:"code"`
	Amount int64  `json:"amount"`
}
type PayAsset struct {
	DestAddress string `json:"dest_address"`
	Asset       Asset  `json:"asset"`
	Input       string `json:"input"`
}
type PayCoin struct {
	DestAddress string `json:"dest_address"`
	Amount      int64  `json:"amount"`
	Input       string `json:"input"`
}
type SetPrivilege struct {
	MasterWeight   string          `json:"master_weight"`
	Signers        []Signer        `json:"signers"`
	TxThreshold    string          `json:"tx_threshold"`
	TypeThresholds []TypeThreshold `json:"type_thresholds"`
}
type Log struct {
	Topic string   `json:"topic"`
	Datas []string `json:"datas"`
}
type BlockGetInfoResponse struct {
	ErrorCode int                `json:"error_code"`
	ErrorDesc string             `json:"error_desc"`
	Result    BlockGetInfoResult `json:"result"`
}
type BlockGetInfoResult struct {
	Header GetInfoHeader `json:"header"`
}
type GetInfoHeader struct {
	CloseTime int64 `json:"close_time"`
	Number    int64 `json:"seq"`
	TxCount   int64 `json:"tx_count"`
	Version   int64 `json:"version"`
}

//GetLatest
type BlockGetLatestResponse struct {
	ErrorCode int             `json:"error_code"`
	ErrorDesc string          `json:"error_desc"`
	Result    GetLatestResult `json:"result"`
}
type GetLatestResult struct {
	Header GetLatestHeader `json:"header"`
}
type GetLatestHeader struct {
	CloseTime int64 `json:"close_time"`
	Number    int64 `json:"seq"`
	TxCount   int64 `json:"tx_count"`
	Version   int64 `json:"version"`
}

//GetNumber
type BlockGetNumberResponse struct {
	ErrorCode int             `json:"error_code"`
	ErrorDesc string          `json:"error_desc"`
	Result    GetNumberResult `json:"result"`
}
type GetNumberResult struct {
	Header GetNumberHeader `json:"header"`
}
type GetNumberHeader struct {
	BlockNumber int64 `json:"seq"`
}

//CheckStatus
type BlockCheckStatusResponse struct {
	ErrorCode int               `json:"error_code"`
	ErrorDesc string            `json:"error_desc"`
	Result    CheckStatusResult `json:"result"`
}
type CheckStatusResult struct {
	IsSynchronous bool `json:"check_status"`
}
type BlockGetValidatorsResponse struct {
	ErrorCode int                 `json:"error_code"`
	ErrorDesc string              `json:"error_desc"`
	Result    GetValidatorsResult `json:"result"`
}
type GetValidatorsResult struct {
	Validators []string `json:"validators"`
}
type Validator struct {
	Address          string `json:"address"`
	PledgeCoinAmount int64  `json:"pledge_coin_amount"`
}

//GetLatestValidators
type BlockGetLatestValidatorsResponse struct {
	ErrorCode int                 `json:"error_code"`
	ErrorDesc string              `json:"error_desc"`
	Result    GetValidatorsResult `json:"result"`
}
type GetLatestValidatorsResult struct {
	Validators []Validator `json:"validators"`
}
type WebBlockGetRewardResponse struct {
	ErrorCode int                `json:"error_code"`
	ErrorDesc string             `json:"error_desc"`
	Result    WebGetRewardResult `json:"result"`
}
type WebGetRewardResult struct {
	ValidatorsReward map[string]int64 `json:"validators_reward"`
}
type BlockGetRewardResponse struct {
	ErrorCode int             `json:"error_code"`
	ErrorDesc string          `json:"error_desc"`
	Result    BlockGetRewardResult `json:"result"`
}
type BlockGetRewardResult struct {
	Validators      []Rewards  `json:"validators"`
	Kols []Rewards `json:"kols"`
}

type Rewards struct {
	Address string `json:"address"`
	Reward []interface{} `json:"reward"`
}

type GetRewardResult struct {
	Validators map[string][]interface{} `json:"validators"`
	Kols map[string][]interface{} `json:"kols"`
}

//GetLatestReward
type WebBlockGetLatestRewardResponse struct {
	ErrorCode int                `json:"error_code"`
	ErrorDesc string             `json:"error_desc"`
	Result    WebGetRewardResult `json:"result"`
}
type WebGetLatestRewardResult struct {
	ValidatorsReward map[string]int64 `json:"validators_reward"`
}
type BlockGetLatestRewardResponse struct {
	ErrorCode int                   `json:"error_code"`
	ErrorDesc string                `json:"error_desc"`
	Result    GetLatestRewardResult `json:"result"`
}
type GetLatestRewardResult struct {
	BlockReward      int64
	ValidatorsReward []ValidatorReward
}
type ValidatorReward struct {
	Validator string
	Reward    int64
}

type BlockGetFeesResponse struct {
	ErrorCode int           `json:"error_code"`
	ErrorDesc string        `json:"error_desc"`
	Result    GetFeesResult `json:"result"`
}
type GetFeesResult struct {
	Fees Fees `json:"fees"`
}
type Fees struct {
	BaseReserve int64 `json:"base_reserve"`
	GasPrice    int64 `json:"gas_price"`
}

//GetLatestFees
type BlockGetLatestFeesResponse struct {
	ErrorCode int                 `json:"error_code"`
	ErrorDesc string              `json:"error_desc"`
	Result    GetLatestFeesResult `json:"result"`
}
type GetLatestFeesResult struct {
	Fees Fees `json:"fees"`
}

// sdk
//Init
type SDKInitResponse struct {
	ErrorCode int    `json:"error_code"`
	ErrorDesc string `json:"error_desc"`
}
type LogCreateResponse struct {
	ErrorCode int             `json:"error_code"`
	ErrorDesc string          `json:"error_desc"`
	Result    LogCreateResult `json:"result"`
}
type LogCreateResult struct {
	Operation protocol.Operation `json:"operation"`
}

//Ctp10Token
type Ctp10TokenIssueResponse struct {
	ErrorCode int                   `json:"error_code"`
	ErrorDesc string                `json:"error_desc"`
	Result    Ctp10TokenIssueResult `json:"result"`
}
type Ctp10TokenIssueResult struct {
	Operation protocol.Operation `json:"operation"`
}

//Atp10TokenIssue
type Atp10TokenIssueResponse struct {
	ErrorCode int                   `json:"error_code"`
	ErrorDesc string                `json:"error_desc"`
	Result    Atp10TokenIssueResult `json:"result"`
}
type Atp10TokenIssueResult struct {
	Operations []protocol.Operation `json:"operations"`
}

//Atp10TokenAppendToIssue
type Atp10TokenAppendToIssueResponse struct {
	ErrorCode int                           `json:"error_code"`
	ErrorDesc string                        `json:"error_desc"`
	Result    Atp10TokenAppendToIssueResult `json:"result"`
}
type Atp10TokenAppendToIssueResult struct {
	Operations []protocol.Operation `json:"operations"`
}
type Ctp10TokenTransferResponse struct {
	ErrorCode int                      `json:"error_code"`
	ErrorDesc string                   `json:"error_desc"`
	Result    Ctp10TokenTransferResult `json:"result"`
}
type Ctp10TokenTransferResult struct {
	Operation protocol.Operation `json:"operation"`
}
type Ctp10TokenTransferFromResponse struct {
	ErrorCode int                          `json:"error_code"`
	ErrorDesc string                       `json:"error_desc"`
	Result    Ctp10TokenTransferFromResult `json:"result"`
}
type Ctp10TokenTransferFromResult struct {
	Operation protocol.Operation `json:"operation"`
}
type Ctp10TokenApproveResponse struct {
	ErrorCode int                     `json:"error_code"`
	ErrorDesc string                  `json:"error_desc"`
	Result    Ctp10TokenApproveResult `json:"result"`
}
type Ctp10TokenApproveResult struct {
	Operation protocol.Operation `json:"operation"`
}
type Ctp10TokenAssignResponse struct {
	ErrorCode int                    `json:"error_code"`
	ErrorDesc string                 `json:"error_desc"`
	Result    Ctp10TokenAssignResult `json:"result"`
}
type Ctp10TokenAssignResult struct {
	Operation protocol.Operation `json:"operation"`
}
type Ctp10TokenChangeOwnerResponse struct {
	ErrorCode int                         `json:"error_code"`
	ErrorDesc string                      `json:"error_desc"`
	Result    Ctp10TokenChangeOwnerResult `json:"result"`
}
type Ctp10TokenChangeOwnerResult struct {
	Operation protocol.Operation `json:"operation"`
}
type Input struct {
	Params Params `json:"params"`
	Method string `json:"method"`
}
type Params struct {
	Ctp             string `json:"ctp"`
	Name            string `json:"name"`
	Symbol          string `json:"symbol"`
	Decimals        int64  `json:"decimals"`
	TotalSupply     string `json:"totalSupply"`
	Supply          string `json:"supply"`
	Ctp10TokenOwner string `json:"contractOwner"`
	To              string `json:"to"`
	Value           string `json:"value"`
	From            string `json:"from"`
	Spender         string `json:"spender"`
	Address         string `json:"address"`
	Owner           string `json:"owner"`
	Balance         string `json:"balance"`
}
type Ctp10TokenCallResponse struct {
	ErrorCode int                  `json:"error_code"`
	ErrorDesc string               `json:"error_desc"`
	Result    Ctp10TokenCallResult `json:"result"`
}
type Ctp10TokenCallResult struct {
	QueryRets []QueryRet `json:"query_rets"`
}
type QueryRet struct {
	Result Ctp10TokenResult `json:"result"`
	Error  Error            `json:"error"`
}
type Ctp10TokenResult struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}
type Error struct {
	Data Data `json:"data"`
}
type Data struct {
	Exception string `json:"exception"`
}
type Value struct {
	ContractInfo ContractInfo `json:"contractInfo"`
}
type ContractInfo struct {
	Ctp           string `json:"ctp"`
	Name          string `json:"name"`
	Symbol        string `json:"symbol"`
	Decimals      int64  `json:"decimals"`
	TotalSupply   string `json:"totalSupply"`
	contractOwner string `json:"contractOwner"`
	Balance       string `json:"balance"`
}
type Ctp10TokenCheckValidResponse struct {
	ErrorCode int              `json:"error_code"`
	ErrorDesc string           `json:"error_desc"`
	Result    CheckValidResult `json:"result"`
}
type Ctp10TokenAllowanceResponse struct {
	ErrorCode int                       `json:"error_code"`
	ErrorDesc string                    `json:"error_desc"`
	Result    Ctp10TokenAllowanceResult `json:"result"`
}
type Ctp10TokenAllowanceResult struct {
	Allowance int64 `json:"Allowance"`
}
type CallGetInfoResponse struct {
	ErrorCode int               `json:"error_code"`
	ErrorDesc string            `json:"error_desc"`
	Result    CallGetInfoResult `json:"result"`
}
type CallGetInfoResult struct {
	QueryRets []Ctp10TokenGetInfoResponse `json:"query_rets"`
}
type Ctp10TokenGetInfoResponse struct {
	ErrorCode int                     `json:"error_code"`
	ErrorDesc string                  `json:"error_desc"`
	Result    Ctp10TokenGetInfoResult `json:"result"`
}
type Ctp10TokenGetInfoResult struct {
	Symbol      string `json:"symbol"`
	Decimals    int64  `json:"decimals"`
	TotalSupply int64  `json:"totalSupply"`
	Name        string `json:"name"`
}
type CallGetNameResponse struct {
	ErrorCode int               `json:"error_code"`
	ErrorDesc string            `json:"error_desc"`
	Result    CallGetNameResult `json:"result"`
}
type CallGetNameResult struct {
	QueryRets []Ctp10TokenGetNameResponse `json:"query_rets"`
}
type Ctp10TokenGetNameResponse struct {
	ErrorCode int                     `json:"error_code"`
	ErrorDesc string                  `json:"error_desc"`
	Result    Ctp10TokenGetNameResult `json:"result"`
}
type Ctp10TokenGetNameResult struct {
	Name string `json:"name"`
}
type CallGetSymbolResponse struct {
	ErrorCode int                 `json:"error_code"`
	ErrorDesc string              `json:"error_desc"`
	Result    CallGetSymbolResult `json:"result"`
}
type CallGetSymbolResult struct {
	QueryRets []Ctp10TokenGetSymbolResponse `json:"query_rets"`
}
type Ctp10TokenGetSymbolResponse struct {
	ErrorCode int                       `json:"error_code"`
	ErrorDesc string                    `json:"error_desc"`
	Result    Ctp10TokenGetSymbolResult `json:"result"`
}
type Ctp10TokenGetSymbolResult struct {
	Symbol string `json:"symbol"`
}
type CallGetDecimalsResponse struct {
	ErrorCode int                   `json:"error_code"`
	ErrorDesc string                `json:"error_desc"`
	Result    CallGetDecimalsResult `json:"result"`
}
type CallGetDecimalsResult struct {
	QueryRets []Ctp10TokenGetDecimalsResponse `json:"query_rets"`
}
type Ctp10TokenGetDecimalsResponse struct {
	ErrorCode int                         `json:"error_code"`
	ErrorDesc string                      `json:"error_desc"`
	Result    Ctp10TokenGetDecimalsResult `json:"result"`
}
type Ctp10TokenGetDecimalsResult struct {
	Decimals int64 `json:"decimals"`
}
type CallGetTotalSupplyResponse struct {
	ErrorCode int                      `json:"error_code"`
	ErrorDesc string                   `json:"error_desc"`
	Result    CallGetTotalSupplyResult `json:"result"`
}
type CallGetTotalSupplyResult struct {
	QueryRets []Ctp10TokenGetTotalSupplyResponse `json:"query_rets"`
}
type Ctp10TokenGetTotalSupplyResponse struct {
	ErrorCode int                            `json:"error_code"`
	ErrorDesc string                         `json:"error_desc"`
	Result    Ctp10TokenGetTotalSupplyResult `json:"result"`
}
type Ctp10TokenGetTotalSupplyResult struct {
	TotalSupply int64 `json:"totalSupply"`
}
type CallGetBalanceResponse struct {
	ErrorCode int                  `json:"error_code"`
	ErrorDesc string               `json:"error_desc"`
	Result    CallGetBalanceResult `json:"result"`
}
type CallGetBalanceResult struct {
	QueryRets []Ctp10TokenGetBalanceResponse `json:"query_rets"`
}
type Ctp10TokenGetBalanceResponse struct {
	ErrorCode int                        `json:"error_code"`
	ErrorDesc string                     `json:"error_desc"`
	Result    Ctp10TokenGetBalanceResult `json:"result"`
}
type Ctp10TokenGetBalanceResult struct {
	Balance int64 `json:"balanceOf"`
}
type CallContractRequest struct {
	ContractAddress string `json:"contract_address"`
	Code            string `json:"code"`
	Input           string `json:"input"`
	ContractBalance string `json:"contract_balance"`
	FeeLimit        int64  `json:"fee_limit"`
	GasPrice        int64  `json:"gas_price"`
	OptType         int64  `json:"opt_type"`
	SourceAddress   string `json:"source_address"`
}

type RewardsGetInput struct {
	Method        string `json:"method"`
}

type RewardsResult struct {
	Rewards    GetRewardResult `json:"rewards"`
}

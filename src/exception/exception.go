// exception
package exception

type SDKResponse struct {
	ErrorCode int
	ErrorDesc string
	Result    string
}

const (
	SUCCESS                                   int = 0
	ACCOUNT_CREATE_ERROR                      int = 11001
	INVALID_SOURCEADDRESS_ERROR               int = 11002
	INVALID_DESTADDRESS_ERROR                 int = 11003
	INVALID_INITBALANCE_ERROR                 int = 11004
	SOURCEADDRESS_EQUAL_DESTADDRESS_ERROR     int = 11005
	INVALID_ADDRESS_ERROR                     int = 11006
	CONNECTNETWORK_ERROR                      int = 11007
	METADATA_NOT_HEX_STRING_ERROR             int = 11008
	NO_ASSET_ERROR                            int = 11009
	NO_METADATA_ERROR                         int = 11010
	INVALID_DATAKEY_ERROR                     int = 11011
	INVALID_DATAVALUE_ERROR                   int = 11012
	INVALID_DATAVERSION_ERROR                 int = 11013
	INVALID_MASTERWEIGHT_ERROR                int = 11015
	INVALID_SIGNER_ADDRESS_ERROR              int = 11016
	INVALID_SIGNER_WEIGHT_ERROR               int = 11017
	INVALID_TX_THRESHOLD_ERROR                int = 11018
	INVALID_OPERATION_TYPE_ERROR              int = 11019
	INVALID_TYPE_THRESHOLD_ERROR              int = 11020
	INVALID_ASSET_CODE_ERROR                  int = 11023
	INVALID_ASSET_AMOUNT_ERROR                int = 11024
	INVALID_BU_AMOUNT_ERROR                   int = 11026
	INVALID_ISSUER_ADDRESS_ERROR              int = 11027
	NO_SUCH_TOKEN_ERROR                       int = 11030
	INVALID_TOKEN_NAME_ERROR                  int = 11031
	INVALID_TOKEN_SIMBOL_ERROR                int = 11032
	INVALID_TOKEN_DECIMALS_ERROR              int = 11033
	INVALID_TOKEN_TOTALSUPPLY_ERROR           int = 11034
	INVALID_TOKENOWNER_ERROR                  int = 11035
	INVALID_CONTRACTADDRESS_ERROR             int = 11037
	CONTRACTADDRESS_NOT_CONTRACTACCOUNT_ERROR int = 11038
	INVALID_TOKEN_AMOUNT_ERROR                int = 11039
	INVALID_FROMADDRESS_ERROR                 int = 11041
	INVALID_SPENDER_ERROR                     int = 11043
	INVALID_LOG_TOPIC_ERROR                   int = 11045
	INVALID_LOG_DATA_ERROR                    int = 11046
	INVALID_NONCE_ERROR                       int = 11048
	INVALID_GASPRICE_ERROR                    int = 11049
	INVALID_FEELIMIT_ERROR                    int = 11050
	INVALID_OPERATIONS_ERROR                  int = 11051
	INVALID_CEILLEDGERSEQ_ERROR               int = 11052
	OPERATIONS_ONE_ERROR                      int = 11053
	INVALID_SIGNATURENUMBER_ERROR             int = 11054
	INVALID_HASH_ERROR                        int = 11055
	INVALID_BLOB_ERROR                        int = 11056
	PRIVATEKEY_NULL_ERROR                     int = 11057
	PRIVATEKEY_ONE_ERROR                      int = 11058
	INVALID_BLOCKNUMBER_ERROR                 int = 11060
	URL_EMPTY_ERROR                           int = 11062
	CONTRACTADDRESS_CODE_BOTH_NULL_ERROR      int = 11063
	INVALID_OPTTYPE_ERROR                     int = 11064
	SYSTEM_ERROR                              int = 20000
)
const (
	GET_ENCPUBLICKEY_ERROR int = iota + 17000
	SIGN_ERROR
	INVALID_PAYLOAD_ERROR
	THE_QUERY_FAILED
)

var SDKRes SDKResponse

var errm = map[int]string{
	SUCCESS:                                   "",
	ACCOUNT_CREATE_ERROR:                      "Create account failed.",
	INVALID_SOURCEADDRESS_ERROR:               "Invalid sourceAddress.",
	INVALID_DESTADDRESS_ERROR:                 "Invalid destAddress.",
	INVALID_INITBALANCE_ERROR:                 "InitBalance must between 1 and max(int64).",
	SOURCEADDRESS_EQUAL_DESTADDRESS_ERROR:     "SourceAddress cannot be equal to destAddress.",
	INVALID_ADDRESS_ERROR:                     "Invalid address.",
	CONNECTNETWORK_ERROR:                      "Connect network failed.",
	METADATA_NOT_HEX_STRING_ERROR:             "Metadata must be a hex string.",
	NO_ASSET_ERROR:                            "The account does not have this asset",
	NO_METADATA_ERROR:                         "The account does not have this metadata.",
	INVALID_DATAKEY_ERROR:                     "The length of key must between 1 and 1024.",
	INVALID_DATAVALUE_ERROR:                   "The length of value must between 0 and 256000.",
	INVALID_DATAVERSION_ERROR:                 "The version must be bigger than and equal to 0.",
	INVALID_MASTERWEIGHT_ERROR:                "MasterWeight must between 0 and max(uint32).",
	INVALID_SIGNER_ADDRESS_ERROR:              "Invalid signer address.",
	INVALID_SIGNER_WEIGHT_ERROR:               "Signer weight must between 0 and max(uint32).",
	INVALID_TX_THRESHOLD_ERROR:                "TxThreshold must between 0 and max(int64).",
	INVALID_OPERATION_TYPE_ERROR:              "Operation type must between 1 and 100.",
	INVALID_TYPE_THRESHOLD_ERROR:              "TypeThreshold must between 0 and max(int64).",
	INVALID_ASSET_CODE_ERROR:                  "The length of code must between 1 and 64.",
	INVALID_ASSET_AMOUNT_ERROR:                "AssetAmount must between 0 and max(int64).",
	INVALID_BU_AMOUNT_ERROR:                   "BuAmount must between 0 and max(int64).",
	INVALID_ISSUER_ADDRESS_ERROR:              "Invalid issuer address.",
	NO_SUCH_TOKEN_ERROR:                       "The length of ctp must between 1 and 64.",
	INVALID_TOKEN_NAME_ERROR:                  "The length of token name must between 1 and 1024.",
	INVALID_TOKEN_SIMBOL_ERROR:                "The length of symbol must between 1 and 1024.",
	INVALID_TOKEN_DECIMALS_ERROR:              "Decimals must less than 8.",
	INVALID_TOKEN_TOTALSUPPLY_ERROR:           "TotalSupply must between 1 and max(int64).",
	INVALID_TOKENOWNER_ERROR:                  "Invalid token owner.",
	INVALID_CONTRACTADDRESS_ERROR:             "Invalid contract address.",
	CONTRACTADDRESS_NOT_CONTRACTACCOUNT_ERROR: "contractAddress is not a contract account.",
	INVALID_TOKEN_AMOUNT_ERROR:                "Amount must between 1 and max(int64).",
	INVALID_FROMADDRESS_ERROR:                 "Invalid fromAddress.",
	INVALID_SPENDER_ERROR:                     "Invalid spender.",
	INVALID_LOG_TOPIC_ERROR:                   "The length of key must between 1 and 128.",
	INVALID_LOG_DATA_ERROR:                    "The length of value must between 1 and 1024.",
	INVALID_NONCE_ERROR:                       "Nonce must between 1 and max(int64).",
	INVALID_GASPRICE_ERROR:                    "Amount must between 0 and max(int64).",
	INVALID_FEELIMIT_ERROR:                    "FeeLimit must between 0 and max(int64).",
	INVALID_OPERATIONS_ERROR:                  "Operations cannot be resolved",
	INVALID_CEILLEDGERSEQ_ERROR:               "CeilLedgerSeq must be equal or bigger than 0.",
	OPERATIONS_ONE_ERROR:                      "One of operations cannot be resolved.",
	INVALID_SIGNATURENUMBER_ERROR:             "SignagureNumber must between 1 and max(int32).",
	INVALID_HASH_ERROR:                        "Invalid transaction hash.",
	INVALID_BLOB_ERROR:                        "Invalid blob.",
	PRIVATEKEY_NULL_ERROR:                     "PrivateKeys cannot be empty.",
	PRIVATEKEY_ONE_ERROR:                      "One of privateKeys is invalid.",
	URL_EMPTY_ERROR:                           "Url cannot be empty.",
	CONTRACTADDRESS_CODE_BOTH_NULL_ERROR:      "ContractAddress and code cannot be empty at the same time.",
	SYSTEM_ERROR:                              "System error.",
	INVALID_BLOCKNUMBER_ERROR:                 "BlockNumber must bigger than 0",
	INVALID_OPTTYPE_ERROR:                     "OptType must between 0 and 2",

	GET_ENCPUBLICKEY_ERROR: "The function 'GetEncPublicKey' failed.",
	SIGN_ERROR:             "The function 'Sign' failed.",
	INVALID_PAYLOAD_ERROR:  "The parameter 'payload' is invalid.",
	THE_QUERY_FAILED:       "The query failed.",
}

//GetSDKRes
func GetSDKRes(code int) SDKResponse {

	v, _ := errm[code]
	SDKRes.ErrorCode = code
	SDKRes.ErrorDesc = v
	return SDKRes
}

//GetErrDesc
func GetErrDesc(code int) string {
	v, _ := errm[code]
	return v
}

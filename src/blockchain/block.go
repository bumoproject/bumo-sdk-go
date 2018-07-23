// block
package blockchain

import (
	"bytes"
	"encoding/json"
	"strconv"

	"github.com/bumoproject/bumo-sdk-go/src/common"
	"github.com/bumoproject/bumo-sdk-go/src/exception"
	"github.com/bumoproject/bumo-sdk-go/src/model"
)

type BlockOperation struct {
	Url string
}

//获取区块高度 GetNumber
func (block *BlockOperation) GetNumber() model.BlockGetNumberResponse {
	var resData model.BlockGetNumberResponse
	get := "/getLedger"
	response, SDKRes := common.GetRequest(block.Url, get, "")
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
			return resData
		} else {
			if resData.ErrorCode == 4 {
				resData.ErrorDesc = "Get block failed"
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

//检查区块同步 CheckStatus
func (block *BlockOperation) CheckStatus() model.BlockCheckStatusResponse {
	var resData model.BlockCheckStatusResponse
	resData.Result.IsSynchronous = false
	get := "/getModulesStatus"
	response, SDKRes := common.GetRequest(block.Url, get, "")
	if SDKRes.ErrorCode != 0 {
		resData.ErrorCode = SDKRes.ErrorCode
		resData.ErrorDesc = SDKRes.ErrorDesc
		return resData
	}
	defer response.Body.Close()
	if response.StatusCode == 200 {
		data := make(map[string]interface{})
		decoder := json.NewDecoder(response.Body)
		decoder.UseNumber()
		err := decoder.Decode(&data)
		if err != nil {
			SDKRes := exception.GetSDKRes(exception.SYSTEM_ERROR)
			resData.ErrorCode = SDKRes.ErrorCode
			resData.ErrorDesc = SDKRes.ErrorDesc
			return resData
		}
		ledger_manager := data["ledger_manager"].(map[string]interface{})
		if ledger_manager["chain_max_ledger_seq"] == ledger_manager["ledger_sequence"] {
			resData.Result.IsSynchronous = true
		} else {
			resData.Result.IsSynchronous = false
		}
		resData.ErrorCode = exception.SUCCESS
		return resData
	} else {
		SDKRes := exception.GetSDKRes(exception.CONNECTNETWORK_ERROR)
		resData.ErrorCode = SDKRes.ErrorCode
		resData.ErrorDesc = SDKRes.ErrorDesc
		return resData
	}

}

//根据高度查询交易 GetTransactions
func (block *BlockOperation) GetTransactions(reqData model.BlockGetTransactionRequest) model.BlockGetTransactionResponse {
	var resData model.BlockGetTransactionResponse
	if reqData.GetBlockNumber() <= 0 {
		SDKRes := exception.GetSDKRes(exception.INVALID_BLOCKNUMBER_ERROR)
		resData.ErrorCode = SDKRes.ErrorCode
		resData.ErrorDesc = SDKRes.ErrorDesc
		return resData
	}
	bnstr := strconv.FormatInt(reqData.GetBlockNumber(), 10)
	get := "/getTransactionHistory?ledger_seq="
	response, SDKRes := common.GetRequest(block.Url, get, bnstr)
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
				resData.ErrorDesc = "Get transactions failed"
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

//获取区块信息 GetInfo
func (block *BlockOperation) GetInfo(reqData model.BlockGetInfoRequest) model.BlockGetInfoResponse {
	var resData model.BlockGetInfoResponse
	if reqData.GetBlockNumber() <= 0 {
		SDKRes := exception.GetSDKRes(exception.INVALID_BLOCKNUMBER_ERROR)
		resData.ErrorCode = SDKRes.ErrorCode
		resData.ErrorDesc = SDKRes.ErrorDesc
		return resData
	}
	str := strconv.FormatInt(reqData.GetBlockNumber(), 10)
	get := "/getLedger?seq="
	response, SDKRes := common.GetRequest(block.Url, get, str)
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
				resData.ErrorDesc = "Get block failed"
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

//获取最新区块信息 GetLatest
func (block *BlockOperation) GetLatest() model.BlockGetLatestResponse {
	var resData model.BlockGetLatestResponse
	get := "/getLedger"
	response, SDKRes := common.GetRequest(block.Url, get, "")
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
				resData.ErrorDesc = "Get block failed"
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

//获取指定区块中所有验证节点数 GetValidators
func (block *BlockOperation) GetValidators(reqData model.BlockGetValidatorsRequest) model.BlockGetValidatorsResponse {
	var resData model.BlockGetValidatorsResponse
	if reqData.GetBlockNumber() <= 0 {
		SDKRes := exception.GetSDKRes(exception.INVALID_BLOCKNUMBER_ERROR)
		resData.ErrorCode = SDKRes.ErrorCode
		resData.ErrorDesc = SDKRes.ErrorDesc
		return resData
	}
	get := "/getLedger?seq="
	bnstr := strconv.FormatInt(reqData.GetBlockNumber(), 10)
	var buf bytes.Buffer
	buf.WriteString(bnstr)
	buf.WriteString("&with_validator=true")
	str := buf.String()
	response, SDKRes := common.GetRequest(block.Url, get, str)
	if SDKRes.ErrorCode != 0 {
		resData.ErrorDesc = SDKRes.ErrorDesc
		resData.ErrorCode = SDKRes.ErrorCode
		return resData
	}
	defer response.Body.Close()
	if response.StatusCode == 200 {
		decoder := json.NewDecoder(response.Body)
		decoder.UseNumber()
		err := decoder.Decode(&resData)
		if err != nil {
			SDKRes := exception.GetSDKRes(exception.SYSTEM_ERROR)
			resData.ErrorDesc = SDKRes.ErrorDesc
			resData.ErrorCode = SDKRes.ErrorCode
			return resData
		}
		if resData.ErrorCode == 0 {
			SDKRes := exception.GetSDKRes(exception.SUCCESS)
			resData.ErrorDesc = SDKRes.ErrorDesc
			resData.ErrorCode = exception.SUCCESS
			return resData
		} else {
			if resData.ErrorCode == 4 {
				resData.ErrorDesc = "Get block failed"
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

//获取最新区块中所有验证节点数 GetLatestValidators
func (block *BlockOperation) GetLatestValidators() model.BlockGetLatestValidatorsResponse {
	var resData model.BlockGetLatestValidatorsResponse
	get := "/getLedger?"
	var buf bytes.Buffer
	buf.WriteString("with_validator=true")
	str := buf.String()
	response, SDKRes := common.GetRequest(block.Url, get, str)
	if SDKRes.ErrorCode != 0 {
		resData.ErrorDesc = SDKRes.ErrorDesc
		resData.ErrorCode = SDKRes.ErrorCode
		return resData
	}
	defer response.Body.Close()
	if response.StatusCode == 200 {
		decoder := json.NewDecoder(response.Body)
		decoder.UseNumber()
		err := decoder.Decode(&resData)
		if err != nil {
			SDKRes := exception.GetSDKRes(exception.SYSTEM_ERROR)
			resData.ErrorDesc = SDKRes.ErrorDesc
			resData.ErrorCode = SDKRes.ErrorCode
			return resData
		}
		if resData.ErrorCode == 0 {
			SDKRes := exception.GetSDKRes(exception.SUCCESS)
			resData.ErrorDesc = SDKRes.ErrorDesc
			resData.ErrorCode = exception.SUCCESS
			return resData
		} else {
			if resData.ErrorCode == 4 {
				resData.ErrorDesc = "Get block failed"
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

//获取指定区块中的区块奖励和验证节点奖励 GetReward
func (block *BlockOperation) GetReward(reqData model.BlockGetRewardRequest) model.BlockGetRewardResponse {
	var resData model.BlockGetRewardResponse
	var resDataWeb model.WebBlockGetRewardResponse
	if reqData.GetBlockNumber() <= 0 {
		SDKRes := exception.GetSDKRes(exception.INVALID_BLOCKNUMBER_ERROR)
		resData.ErrorCode = SDKRes.ErrorCode
		resData.ErrorDesc = SDKRes.ErrorDesc
		return resData
	}
	get := "/getLedger?seq="
	bnstr := strconv.FormatInt(reqData.GetBlockNumber(), 10)
	var buf bytes.Buffer
	buf.WriteString(bnstr)
	buf.WriteString("&with_block_reward=true")
	str := buf.String()
	response, SDKRes := common.GetRequest(block.Url, get, str)
	if SDKRes.ErrorCode != 0 {
		resData.ErrorDesc = SDKRes.ErrorDesc
		resData.ErrorCode = SDKRes.ErrorCode
		return resData
	}
	defer response.Body.Close()
	if response.StatusCode == 200 {
		decoder := json.NewDecoder(response.Body)
		decoder.UseNumber()
		err := decoder.Decode(&resDataWeb)
		if err != nil {
			SDKRes := exception.GetSDKRes(exception.SYSTEM_ERROR)
			resData.ErrorDesc = SDKRes.ErrorDesc
			resData.ErrorCode = SDKRes.ErrorCode
			return resData
		}
		if resDataWeb.ErrorCode == 0 {
			resData.Result.ValidatorsReward = make([]model.ValidatorReward, len(resDataWeb.Result.ValidatorsReward))
			var i int64 = 0
			for key, value := range resDataWeb.Result.ValidatorsReward {
				resData.Result.ValidatorsReward[i].Validator = key
				resData.Result.ValidatorsReward[i].Reward = value
				i++
			}
			resData.Result.BlockReward = i
			return resData
		} else {
			if resDataWeb.ErrorCode == 4 {
				resData.ErrorCode = resDataWeb.ErrorCode
				resData.ErrorDesc = "Get block failed"
				return resData
			}
			resData.ErrorCode = resDataWeb.ErrorCode
			resData.ErrorDesc = resDataWeb.ErrorDesc
			return resData
		}
	} else {
		SDKRes := exception.GetSDKRes(exception.CONNECTNETWORK_ERROR)
		resData.ErrorCode = SDKRes.ErrorCode
		resData.ErrorDesc = SDKRes.ErrorDesc
		return resData
	}
}

//获取最新区块中的区块奖励和验证节点奖励 GetLatestReward
func (block *BlockOperation) GetLatestReward() model.BlockGetLatestRewardResponse {
	var resData model.BlockGetLatestRewardResponse
	var resDataWeb model.WebBlockGetLatestRewardResponse
	get := "/getLedger?"
	var buf bytes.Buffer
	buf.WriteString("with_block_reward=true")
	str := buf.String()
	response, SDKRes := common.GetRequest(block.Url, get, str)
	if SDKRes.ErrorCode != 0 {
		resData.ErrorDesc = SDKRes.ErrorDesc
		resData.ErrorCode = SDKRes.ErrorCode
		return resData
	}
	defer response.Body.Close()
	if response.StatusCode == 200 {
		decoder := json.NewDecoder(response.Body)
		decoder.UseNumber()
		err := decoder.Decode(&resDataWeb)
		if err != nil {
			SDKRes := exception.GetSDKRes(exception.SYSTEM_ERROR)
			resData.ErrorDesc = SDKRes.ErrorDesc
			resData.ErrorCode = SDKRes.ErrorCode
			return resData
		}
		if resDataWeb.ErrorCode == 0 {
			resData.Result.ValidatorsReward = make([]model.ValidatorReward, len(resDataWeb.Result.ValidatorsReward))
			var i int64 = 0
			for key, value := range resDataWeb.Result.ValidatorsReward {
				resData.Result.ValidatorsReward[i].Validator = key
				resData.Result.ValidatorsReward[i].Reward = value
				i++
			}
			resData.Result.BlockReward = i
			return resData
		} else {
			if resDataWeb.ErrorCode == 4 {
				resData.ErrorCode = resDataWeb.ErrorCode
				resData.ErrorDesc = "Get block failed"
				return resData
			}
			resData.ErrorCode = resDataWeb.ErrorCode
			resData.ErrorDesc = resDataWeb.ErrorDesc
			return resData
		}
	} else {
		SDKRes := exception.GetSDKRes(exception.CONNECTNETWORK_ERROR)
		resData.ErrorCode = SDKRes.ErrorCode
		resData.ErrorDesc = SDKRes.ErrorDesc
		return resData
	}
}

//获取指定区块中的账户最低资产限制和打包费用 GetFees
func (block *BlockOperation) GetFees(reqData model.BlockGetFeesRequest) model.BlockGetFeesResponse {
	var resData model.BlockGetFeesResponse
	if reqData.GetBlockNumber() <= 0 {
		SDKRes := exception.GetSDKRes(exception.INVALID_BLOCKNUMBER_ERROR)
		resData.ErrorCode = SDKRes.ErrorCode
		resData.ErrorDesc = SDKRes.ErrorDesc
		return resData
	}
	get := "/getLedger?seq="
	bnstr := strconv.FormatInt(reqData.GetBlockNumber(), 10)
	var buf bytes.Buffer
	buf.WriteString(bnstr)
	buf.WriteString("&with_fee=true")
	str := buf.String()
	response, SDKRes := common.GetRequest(block.Url, get, str)
	if SDKRes.ErrorCode != 0 {
		resData.ErrorDesc = SDKRes.ErrorDesc
		resData.ErrorCode = SDKRes.ErrorCode
		return resData
	}
	defer response.Body.Close()
	if response.StatusCode == 200 {
		decoder := json.NewDecoder(response.Body)
		decoder.UseNumber()
		err := decoder.Decode(&resData)
		if err != nil {
			SDKRes := exception.GetSDKRes(exception.SYSTEM_ERROR)
			resData.ErrorDesc = SDKRes.ErrorDesc
			resData.ErrorCode = SDKRes.ErrorCode
			return resData
		}
		if resData.ErrorCode == 0 {
			SDKRes := exception.GetSDKRes(exception.SUCCESS)
			resData.ErrorDesc = SDKRes.ErrorDesc
			resData.ErrorCode = SDKRes.ErrorCode
			return resData
		} else {
			if resData.ErrorCode == 4 {
				resData.ErrorDesc = "Get block failed"
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

//获取最新区块中的账户最低资产限制和打包费用 GetLatestFees
func (block *BlockOperation) GetLatestFees() model.BlockGetLatestFeesResponse {
	var resData model.BlockGetLatestFeesResponse
	get := "/getLedger?"
	var buf bytes.Buffer
	buf.WriteString("with_fee=true")
	str := buf.String()
	response, SDKRes := common.GetRequest(block.Url, get, str)
	if SDKRes.ErrorCode != 0 {
		resData.ErrorDesc = SDKRes.ErrorDesc
		resData.ErrorCode = SDKRes.ErrorCode
		return resData
	}
	defer response.Body.Close()
	if response.StatusCode == 200 {
		decoder := json.NewDecoder(response.Body)
		decoder.UseNumber()
		err := decoder.Decode(&resData)
		if err != nil {
			SDKRes := exception.GetSDKRes(exception.SYSTEM_ERROR)
			resData.ErrorDesc = SDKRes.ErrorDesc
			resData.ErrorCode = SDKRes.ErrorCode
			return resData
		}
		if resData.ErrorCode == 0 {
			resData.ErrorCode = exception.SUCCESS
			return resData
		} else {
			if resData.ErrorCode == 4 {
				resData.ErrorDesc = "Get block failed"
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

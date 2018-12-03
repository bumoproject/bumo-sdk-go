// common
package common

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"net"
	"net/http"
	"net/url"
	"strconv"
	"sync"
	"time"

	"github.com/bumoproject/bumo-sdk-go/src/crypto/keypair"
	"github.com/bumoproject/bumo-sdk-go/src/exception"
	"github.com/bumoproject/bumo-sdk-go/src/model"
)

type common struct {
	Url              string
	ConnectTimeout   int64
	ReadWriteTimeout int64
	ChainId          int64
}

var ins *common
var once sync.Once

func GetIns() *common {
	once.Do(func() {
		ins = &common{}
	})
	return ins
}
func GetChainId() int64 {
	return ins.ChainId
}
func TimeoutDialer(cTimeout time.Duration, rwTimeout time.Duration) func(net, addr string) (c net.Conn, err error) {
	return func(netw, addr string) (net.Conn, error) {
		conn, err := net.DialTimeout(netw, addr, cTimeout)
		if err != nil {
			return nil, err
		}
		conn.SetDeadline(time.Now().Add(rwTimeout))
		return conn, nil
	}
}

//http get
func GetRequest(get string, str string) (*http.Response, exception.SDKResponse) {
	var buf bytes.Buffer
	connectTimeout := time.Duration(ins.ConnectTimeout) * time.Second
	readWriteTimeout := time.Duration(ins.ReadWriteTimeout) * time.Second
	buf.WriteString(ins.Url)
	buf.WriteString(get)
	buf.WriteString(url.PathEscape(str))
	strUrl := buf.String()
	if connectTimeout != 0 || readWriteTimeout != 0 {
		client := &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
				Dial:            TimeoutDialer(connectTimeout, readWriteTimeout),
			},
		}
		newRequest, err := http.NewRequest("GET", strUrl, nil)
		if err != nil {
			return nil, exception.GetSDKRes(exception.CONNECTNETWORK_ERROR)
		}
		response, err := client.Do(newRequest)
		if err != nil {
			return response, exception.GetSDKRes(exception.CONNECTNETWORK_ERROR)
		}
		return response, exception.GetSDKRes(exception.SUCCESS)
	} else {
		client := &http.Client{}
		newRequest, err := http.NewRequest("GET", strUrl, nil)
		if err != nil {
			return nil, exception.GetSDKRes(exception.CONNECTNETWORK_ERROR)
		}
		response, err := client.Do(newRequest)
		if err != nil {
			return response, exception.GetSDKRes(exception.CONNECTNETWORK_ERROR)
		}
		return response, exception.GetSDKRes(exception.SUCCESS)
	}

}

//http post
func PostRequest(post string, data []byte) (*http.Response, exception.SDKResponse) {
	var buf bytes.Buffer
	buf.WriteString(ins.Url)
	buf.WriteString(post)
	strUrl := buf.String()
	client := &http.Client{}
	newRequest, err := http.NewRequest("POST", strUrl, bytes.NewReader(data))
	if err != nil {
		return nil, exception.GetSDKRes(exception.CONNECTNETWORK_ERROR)
	}
	response, err := client.Do(newRequest)
	if err != nil {
		return nil, exception.GetSDKRes(exception.CONNECTNETWORK_ERROR)
	}
	return response, exception.GetSDKRes(exception.SUCCESS)
}

//Json
func GetRequestJson(reqData model.TransactionSubmitRequests) ([]byte, exception.SDKResponse) {
	request := make(map[string]interface{})
	items := make([]map[string]interface{}, len(reqData.Items))
	for i := range reqData.Items {
		items[i] = make(map[string]interface{})
		items[i]["transaction_blob"] = reqData.Items[i].GetBlob()
		items[i]["signatures"] = reqData.Items[i].GetSignatures()
	}
	request["items"] = items
	requestJson, err := json.Marshal(request)
	if err != nil {
		return nil, exception.GetSDKRes(exception.SYSTEM_ERROR)
	}
	return requestJson, exception.GetSDKRes(exception.SUCCESS)
}

//获取最新fees
func GetLatestFees() (int64, int64, exception.SDKResponse) {
	get := "/getLedger?with_fee=true"
	response, SDKRes := GetRequest(get, "")
	if SDKRes.ErrorCode != 0 {
		return 0, 0, SDKRes
	}
	defer response.Body.Close()
	if response.StatusCode == 200 {
		data := make(map[string]interface{})
		decoder := json.NewDecoder(response.Body)
		decoder.UseNumber()
		err := decoder.Decode(&data)
		if err != nil {
			return 0, 0, exception.GetSDKRes(exception.SYSTEM_ERROR)
		}
		if data["error_code"].(json.Number) == "0" {
			result := data["result"].(map[string]interface{})
			fees := result["fees"].(map[string]interface{})
			gasPriceStr, ok := fees["gas_price"].(json.Number)
			if ok != true {
				return 0, 0, exception.GetSDKRes(exception.SUCCESS)
			}
			baseReserveStr, ok := fees["base_reserve"].(json.Number)
			if ok != true {
				return 0, 0, exception.GetSDKRes(exception.SUCCESS)
			}
			gasPrice, err := strconv.ParseInt(string(gasPriceStr), 10, 64)
			if err != nil {
				return 0, 0, exception.GetSDKRes(exception.SYSTEM_ERROR)
			}
			baseReserve, err := strconv.ParseInt(string(baseReserveStr), 10, 64)
			if err != nil {
				return 0, 0, exception.GetSDKRes(exception.SYSTEM_ERROR)
			}
			return gasPrice, baseReserve, exception.GetSDKRes(exception.SUCCESS)
		} else {
			errorCodeStr := data["error_code"].(json.Number)
			errorCode, err := strconv.ParseInt(string(errorCodeStr), 10, 64)
			if err != nil {
				return 0, 0, exception.GetSDKRes(exception.SYSTEM_ERROR)
			}
			SDKRes.ErrorCode = int(float64(errorCode))
			SDKRes.ErrorDesc = data["error_desc"].(string)
			return 0, 0, SDKRes
		}
	} else {
		return 0, 0, exception.GetSDKRes(exception.CONNECTNETWORK_ERROR)
	}
}

func GetCallDataStr(funcStr string, ContractAddress string, TokenOwner string) (string, exception.SDKResponse) {
	if !keypair.CheckAddress(ContractAddress) {
		return "", exception.GetSDKRes(exception.INVALID_CONTRACTADDRESS_ERROR)
	}
	if TokenOwner != "" {
		if !keypair.CheckAddress(TokenOwner) {
			return "", exception.GetSDKRes(exception.INVALID_TOKENOWNER_ERROR)
		}
	}
	var Input model.Input
	Input.Method = funcStr
	Input.Params.Address = TokenOwner
	InputStr, err := json.Marshal(Input)
	if err != nil {
		return "", exception.GetSDKRes(exception.SYSTEM_ERROR)
	}
	callData := model.CallContractRequest{
		ContractAddress: ContractAddress,
		Input:           string(InputStr),
	}
	callDataStr, err := json.Marshal(callData)
	if err != nil {
		return "", exception.GetSDKRes(exception.SYSTEM_ERROR)
	}
	return string(callDataStr), exception.GetSDKRes(exception.SUCCESS)

}

func CheckActivated(address string) (bool, exception.SDKResponse) {
	var resData model.AccountGetInfoResponse
	if !keypair.CheckAddress(address) {
		return false, exception.GetSDKRes(exception.INVALID_ADDRESS_ERROR)
	}
	response, SDKRes := GetRequest("/getAccount?address=", address)
	if SDKRes.ErrorCode != 0 {
		return false, SDKRes
	}
	defer response.Body.Close()
	if response.StatusCode == 200 {
		decoder := json.NewDecoder(response.Body)
		decoder.UseNumber()
		err := decoder.Decode(&resData)
		if err != nil {
			return false, exception.GetSDKRes(exception.SYSTEM_ERROR)
		}
		if resData.ErrorCode == 0 {
			return true, exception.GetSDKRes(exception.SUCCESS)
		} else if resData.ErrorCode == 4 {
			return false, exception.GetSDKRes(exception.SUCCESS)
		} else {
			SDKRes.ErrorCode = resData.ErrorCode
			SDKRes.ErrorDesc = resData.ErrorDesc
			return false, SDKRes
		}
	} else {
		return false, exception.GetSDKRes(exception.CONNECTNETWORK_ERROR)
	}
}

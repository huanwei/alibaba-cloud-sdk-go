package cloudwf

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

func (client *Client) SaveApRadioConfig(request *SaveApRadioConfigRequest) (response *SaveApRadioConfigResponse, err error) {
	response = CreateSaveApRadioConfigResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) SaveApRadioConfigWithChan(request *SaveApRadioConfigRequest) (<-chan *SaveApRadioConfigResponse, <-chan error) {
	responseChan := make(chan *SaveApRadioConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SaveApRadioConfig(request)
		responseChan <- response
		errChan <- err
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

func (client *Client) SaveApRadioConfigWithCallback(request *SaveApRadioConfigRequest, callback func(response *SaveApRadioConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SaveApRadioConfigResponse
		var err error
		defer close(result)
		response, err = client.SaveApRadioConfig(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

type SaveApRadioConfigRequest struct {
	*requests.RpcRequest
	RequireMode        string `position:"Query" name:"RequireMode"`
	Minrate            string `position:"Query" name:"Minrate"`
	Uapsd              string `position:"Query" name:"Uapsd"`
	Mac                string `position:"Query" name:"Mac"`
	RadioIndex         string `position:"Query" name:"RadioIndex"`
	BeaconInt          string `position:"Query" name:"BeaconInt"`
	BcastRate          string `position:"Query" name:"BcastRate"`
	Shortgi            string `position:"Query" name:"Shortgi"`
	InstantlyEffective string `position:"Query" name:"InstantlyEffective"`
	Id                 string `position:"Query" name:"Id"`
	Frag               string `position:"Query" name:"Frag"`
	Txpower            string `position:"Query" name:"Txpower"`
	Rts                string `position:"Query" name:"Rts"`
	Htmode             string `position:"Query" name:"Htmode"`
	Probereq           string `position:"Query" name:"Probereq"`
	Channel            string `position:"Query" name:"Channel"`
	Disabled           string `position:"Query" name:"Disabled"`
	Hwmode             string `position:"Query" name:"Hwmode"`
	Noscan             string `position:"Query" name:"Noscan"`
}

type SaveApRadioConfigResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Message   string `json:"Message" xml:"Message"`
	Data      string `json:"Data" xml:"Data"`
	ErrorCode int    `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMsg  string `json:"ErrorMsg" xml:"ErrorMsg"`
}

func CreateSaveApRadioConfigRequest() (request *SaveApRadioConfigRequest) {
	request = &SaveApRadioConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cloudwf", "2017-03-28", "SaveApRadioConfig", "", "")
	return
}

func CreateSaveApRadioConfigResponse() (response *SaveApRadioConfigResponse) {
	response = &SaveApRadioConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
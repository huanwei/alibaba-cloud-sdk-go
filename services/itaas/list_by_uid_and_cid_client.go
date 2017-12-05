package itaas

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

func (client *Client) ListByUidAndCidClient(request *ListByUidAndCidClientRequest) (response *ListByUidAndCidClientResponse, err error) {
	response = CreateListByUidAndCidClientResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) ListByUidAndCidClientWithChan(request *ListByUidAndCidClientRequest) (<-chan *ListByUidAndCidClientResponse, <-chan error) {
	responseChan := make(chan *ListByUidAndCidClientResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListByUidAndCidClient(request)
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

func (client *Client) ListByUidAndCidClientWithCallback(request *ListByUidAndCidClientRequest, callback func(response *ListByUidAndCidClientResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListByUidAndCidClientResponse
		var err error
		defer close(result)
		response, err = client.ListByUidAndCidClient(request)
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

type ListByUidAndCidClientRequest struct {
	*requests.RpcRequest
	Uid           string `position:"Query" name:"Uid"`
	AuthType      string `position:"Query" name:"AuthType"`
	Sid           string `position:"Query" name:"Sid"`
	OsType        string `position:"Query" name:"OsType"`
	AppName       string `position:"Query" name:"AppName"`
	AppVersion    string `position:"Query" name:"AppVersion"`
	Language      string `position:"Query" name:"Language"`
	Cid           string `position:"Query" name:"Cid"`
	Umid          string `position:"Query" name:"Umid"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
}

type ListByUidAndCidClientResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Success      bool   `json:"Success" xml:"Success"`
	Code         string `json:"Code" xml:"Code"`
	BusinessCode string `json:"BusinessCode" xml:"BusinessCode"`
	Message      string `json:"Message" xml:"Message"`
	ResultData   []struct {
		Id           int64  `json:"Id" xml:"Id"`
		Uid          int64  `json:"Uid" xml:"Uid"`
		Cid          string `json:"Cid" xml:"Cid"`
		Umid         string `json:"Umid" xml:"Umid"`
		OsType       string `json:"OsType" xml:"OsType"`
		DeviceName   string `json:"DeviceName" xml:"DeviceName"`
		Status       int    `json:"Status" xml:"Status"`
		RegisterTime int64  `json:"RegisterTime" xml:"RegisterTime"`
	} `json:"ResultData" xml:"ResultData"`
}

func CreateListByUidAndCidClientRequest() (request *ListByUidAndCidClientRequest) {
	request = &ListByUidAndCidClientRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ITaaS", "2017-05-12", "ListByUidAndCidClient", "", "")
	return
}

func CreateListByUidAndCidClientResponse() (response *ListByUidAndCidClientResponse) {
	response = &ListByUidAndCidClientResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
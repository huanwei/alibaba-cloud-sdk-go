package jaq

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

func (client *Client) CheckAccountAndPasswordRisk(request *CheckAccountAndPasswordRiskRequest) (response *CheckAccountAndPasswordRiskResponse, err error) {
	response = CreateCheckAccountAndPasswordRiskResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) CheckAccountAndPasswordRiskWithChan(request *CheckAccountAndPasswordRiskRequest) (<-chan *CheckAccountAndPasswordRiskResponse, <-chan error) {
	responseChan := make(chan *CheckAccountAndPasswordRiskResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CheckAccountAndPasswordRisk(request)
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

func (client *Client) CheckAccountAndPasswordRiskWithCallback(request *CheckAccountAndPasswordRiskRequest, callback func(response *CheckAccountAndPasswordRiskResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CheckAccountAndPasswordRiskResponse
		var err error
		defer close(result)
		response, err = client.CheckAccountAndPasswordRisk(request)
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

type CheckAccountAndPasswordRiskRequest struct {
	*requests.RpcRequest
	AccountName  string `position:"Query" name:"AccountName"`
	CallerName   string `position:"Query" name:"CallerName"`
	PasswordHash string `position:"Query" name:"PasswordHash"`
}

type CheckAccountAndPasswordRiskResponse struct {
	*responses.BaseResponse
	ErrorCode int    `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMsg  string `json:"ErrorMsg" xml:"ErrorMsg"`
	Data      struct {
		FnalDecision     int    `json:"FnalDecision" xml:"FnalDecision"`
		EventId          string `json:"EventId" xml:"EventId"`
		UserId           string `json:"UserId" xml:"UserId"`
		FinalScore       int    `json:"FinalScore" xml:"FinalScore"`
		FinalDesc        string `json:"FinalDesc" xml:"FinalDesc"`
		Detail           string `json:"Detail" xml:"Detail"`
		CaptchaCheckData string `json:"CaptchaCheckData" xml:"CaptchaCheckData"`
	} `json:"Data" xml:"Data"`
}

func CreateCheckAccountAndPasswordRiskRequest() (request *CheckAccountAndPasswordRiskRequest) {
	request = &CheckAccountAndPasswordRiskRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("jaq", "2017-08-25", "CheckAccountAndPasswordRisk", "", "")
	return
}

func CreateCheckAccountAndPasswordRiskResponse() (response *CheckAccountAndPasswordRiskResponse) {
	response = &CheckAccountAndPasswordRiskResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
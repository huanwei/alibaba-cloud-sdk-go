package vod

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

func (client *Client) RefreshUploadMaterialToken(request *RefreshUploadMaterialTokenRequest) (response *RefreshUploadMaterialTokenResponse, err error) {
	response = CreateRefreshUploadMaterialTokenResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) RefreshUploadMaterialTokenWithChan(request *RefreshUploadMaterialTokenRequest) (<-chan *RefreshUploadMaterialTokenResponse, <-chan error) {
	responseChan := make(chan *RefreshUploadMaterialTokenResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RefreshUploadMaterialToken(request)
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

func (client *Client) RefreshUploadMaterialTokenWithCallback(request *RefreshUploadMaterialTokenRequest, callback func(response *RefreshUploadMaterialTokenResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RefreshUploadMaterialTokenResponse
		var err error
		defer close(result)
		response, err = client.RefreshUploadMaterialToken(request)
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

type RefreshUploadMaterialTokenRequest struct {
	*requests.RpcRequest
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	OwnerId              string `position:"Query" name:"OwnerId"`
	MaterialId           string `position:"Query" name:"MaterialId"`
}

type RefreshUploadMaterialTokenResponse struct {
	*responses.BaseResponse
	RequestId  string `json:"RequestId" xml:"RequestId"`
	UploadAuth string `json:"UploadAuth" xml:"UploadAuth"`
}

func CreateRefreshUploadMaterialTokenRequest() (request *RefreshUploadMaterialTokenRequest) {
	request = &RefreshUploadMaterialTokenRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("vod", "2017-07-13", "RefreshUploadMaterialToken", "", "")
	return
}

func CreateRefreshUploadMaterialTokenResponse() (response *RefreshUploadMaterialTokenResponse) {
	response = &RefreshUploadMaterialTokenResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
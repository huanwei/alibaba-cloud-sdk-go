package cloudapi

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

func (client *Client) DescribeSignaturesByApi(request *DescribeSignaturesByApiRequest) (response *DescribeSignaturesByApiResponse, err error) {
	response = CreateDescribeSignaturesByApiResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DescribeSignaturesByApiWithChan(request *DescribeSignaturesByApiRequest) (<-chan *DescribeSignaturesByApiResponse, <-chan error) {
	responseChan := make(chan *DescribeSignaturesByApiResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeSignaturesByApi(request)
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

func (client *Client) DescribeSignaturesByApiWithCallback(request *DescribeSignaturesByApiRequest, callback func(response *DescribeSignaturesByApiResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeSignaturesByApiResponse
		var err error
		defer close(result)
		response, err = client.DescribeSignaturesByApi(request)
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

type DescribeSignaturesByApiRequest struct {
	*requests.RpcRequest
	StageName string `position:"Query" name:"StageName"`
	GroupId   string `position:"Query" name:"GroupId"`
	ApiId     string `position:"Query" name:"ApiId"`
}

type DescribeSignaturesByApiResponse struct {
	*responses.BaseResponse
	RequestId  string `json:"RequestId" xml:"RequestId"`
	Signatures []struct {
		SignatureId   string `json:"SignatureId" xml:"SignatureId"`
		SignatureName string `json:"SignatureName" xml:"SignatureName"`
		BoundTime     string `json:"BoundTime" xml:"BoundTime"`
	} `json:"Signatures" xml:"Signatures"`
}

func CreateDescribeSignaturesByApiRequest() (request *DescribeSignaturesByApiRequest) {
	request = &DescribeSignaturesByApiRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudAPI", "2016-07-14", "DescribeSignaturesByApi", "", "")
	return
}

func CreateDescribeSignaturesByApiResponse() (response *DescribeSignaturesByApiResponse) {
	response = &DescribeSignaturesByApiResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
package emr

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

func (client *Client) ListClusterScripts(request *ListClusterScriptsRequest) (response *ListClusterScriptsResponse, err error) {
	response = CreateListClusterScriptsResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) ListClusterScriptsWithChan(request *ListClusterScriptsRequest) (<-chan *ListClusterScriptsResponse, <-chan error) {
	responseChan := make(chan *ListClusterScriptsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListClusterScripts(request)
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

func (client *Client) ListClusterScriptsWithCallback(request *ListClusterScriptsRequest, callback func(response *ListClusterScriptsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListClusterScriptsResponse
		var err error
		defer close(result)
		response, err = client.ListClusterScripts(request)
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

type ListClusterScriptsRequest struct {
	*requests.RpcRequest
	ClusterId       string `position:"Query" name:"ClusterId"`
	ResourceOwnerId string `position:"Query" name:"ResourceOwnerId"`
}

type ListClusterScriptsResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	ClusterScripts []struct {
		Id        string `json:"Id" xml:"Id"`
		Name      string `json:"Name" xml:"Name"`
		StartTime int64  `json:"StartTime" xml:"StartTime"`
		EndTime   int64  `json:"EndTime" xml:"EndTime"`
		Path      string `json:"Path" xml:"Path"`
		Args      string `json:"Args" xml:"Args"`
		Status    string `json:"Status" xml:"Status"`
	} `json:"ClusterScripts" xml:"ClusterScripts"`
}

func CreateListClusterScriptsRequest() (request *ListClusterScriptsRequest) {
	request = &ListClusterScriptsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Emr", "2016-04-08", "ListClusterScripts", "", "")
	return
}

func CreateListClusterScriptsResponse() (response *ListClusterScriptsResponse) {
	response = &ListClusterScriptsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
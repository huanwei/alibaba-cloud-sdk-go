package integration

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ram"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/sts"

	"fmt"
	"os"
	"strings"
)

var role_doc = `{
		"Statement": [{
		    "Action": "sts:AssumeRole",
		    "Effect": "Allow",
		    "Principal": {
		     	"RAM": [
				      "acs:ram::%s:root"
		        ]
            }
	    }],
	   "Version": "1"
	}`

var (
	travisValue = strings.Split(os.Getenv("TRAVIS_JOB_NUMBER"), ".")
	username    = "testuser" + travisValue[len(travisValue)-1]
	rolename    = "testrole" + travisValue[len(travisValue)-1]
)

func createRole(userid string) (string, string, error) {
	listRequest := ram.CreateListRolesRequest()
	listRequest.Scheme = "HTTPS"
	client, err := ram.NewClientWithAccessKey("cn-hangzhou", os.Getenv("ACCESS_KEY_ID"), os.Getenv("ACCESS_KEY_SECRET"))
	if err != nil {
		return "", "", err
	}
	listResponse, err := client.ListRoles(listRequest)
	if err != nil {
		return "", "", err
	}
	for _, role := range listResponse.Roles.Role {
		if strings.ToLower(role.RoleName) == rolename {
			return role.RoleName, role.Arn, nil
		}
	}
	createRequest := ram.CreateCreateRoleRequest()
	createRequest.Scheme = "HTTPS"
	createRequest.RoleName = rolename
	createRequest.AssumeRolePolicyDocument = fmt.Sprintf(role_doc, userid)
	res, err := client.CreateRole(createRequest)
	if err != nil {
		return "", "", err
	}
	return res.Role.RoleName, res.Role.Arn, nil
}

func createUser() error {
	listRequest := ram.CreateListUsersRequest()
	listRequest.Scheme = "HTTPS"
	client, err := ram.NewClientWithAccessKey("cn-hangzhou", os.Getenv("ACCESS_KEY_ID"), os.Getenv("ACCESS_KEY_SECRET"))
	if err != nil {
		return err
	}
	listResponse, err := client.ListUsers(listRequest)
	if err != nil {
		return err
	}
	for _, user := range listResponse.Users.User {
		if user.UserName == username {
			return nil
		}
	}
	createRequest := ram.CreateCreateUserRequest()
	createRequest.Scheme = "HTTPS"
	createRequest.UserName = username
	_, err = client.CreateUser(createRequest)
	if err != nil {
		return err
	}
	return nil
}

func createAttachPolicyToUser() error {
	listRequest := ram.CreateListPoliciesForUserRequest()
	listRequest.UserName = username
	listRequest.Scheme = "HTTPS"
	client, err := ram.NewClientWithAccessKey("cn-hangzhou", os.Getenv("ACCESS_KEY_ID"), os.Getenv("ACCESS_KEY_SECRET"))
	if err != nil {
		return err
	}
	listResponse, err := client.ListPoliciesForUser(listRequest)
	if err != nil {
		return err
	}
	for _, policy := range listResponse.Policies.Policy {
		if policy.PolicyName == "AliyunSTSAssumeRoleAccess" {
			return nil
		}
	}
	createRequest := ram.CreateAttachPolicyToUserRequest()
	createRequest.Scheme = "HTTPS"
	createRequest.PolicyName = "AliyunSTSAssumeRoleAccess"
	createRequest.UserName = username
	createRequest.PolicyType = "System"
	_, err = client.AttachPolicyToUser(createRequest)
	if err != nil {
		return err
	}
	return nil
}

func createAccessKey() (string, string, error) {
	client, err := ram.NewClientWithAccessKey("cn-hangzhou", os.Getenv("ACCESS_KEY_ID"), os.Getenv("ACCESS_KEY_SECRET"))
	if err != nil {
		return "", "", err
	}
	listrequest := ram.CreateListAccessKeysRequest()
	listrequest.UserName = username
	listrequest.Scheme = "HTTPS"
	listresponse, err := client.ListAccessKeys(listrequest)
	if err != nil {
		return "", "", err
	}
	if listresponse.AccessKeys.AccessKey != nil {
		if len(listresponse.AccessKeys.AccessKey) >= 2 {
			accesskey := listresponse.AccessKeys.AccessKey[0]
			deleterequest := ram.CreateDeleteAccessKeyRequest()
			deleterequest.UserAccessKeyId = accesskey.AccessKeyId
			deleterequest.UserName = username
			deleterequest.Scheme = "HTTPS"
			_, err := client.DeleteAccessKey(deleterequest)
			if err != nil {
				return "", "", err
			}
		}
	}
	request := ram.CreateCreateAccessKeyRequest()
	request.Scheme = "HTTPS"
	request.UserName = username
	response, err := client.CreateAccessKey(request)
	if err != nil {
		return "", "", err
	}

	return response.AccessKey.AccessKeyId, response.AccessKey.AccessKeySecret, nil
}

func createAssumeRole() (*sts.AssumeRoleResponse, error) {
	err := createUser()
	if err != nil {
		return nil, err
	}
	_, _, err = createRole(os.Getenv("USER_ID"))
	if err != nil {
		return nil, err
	}
	err = createAttachPolicyToUser()
	if err != nil {
		return nil, err
	}
	subaccesskeyid, subaccesskeysecret, err := createAccessKey()
	if err != nil {
		return nil, err
	}
	request := sts.CreateAssumeRoleRequest()
	request.RoleArn = fmt.Sprintf("acs:ram::%s:role/testrole", os.Getenv("USER_ID"))
	request.RoleSessionName = "alice_test"
	request.Scheme = "HTTPS"
	client, err := sts.NewClientWithAccessKey("cn-hangzhou", subaccesskeyid, subaccesskeysecret)
	response, err := client.AssumeRole(request)
	if err != nil {
		return nil, err
	}
	return response, nil
}
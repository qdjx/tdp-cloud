package lighthouse

import (
	lighthouse "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/lighthouse/v20200324"

	"tdp-cloud/core/midware"
	"tdp-cloud/core/qcloud"
)

// 创建客户端

func NewClient(ud *midware.Userdata) (*lighthouse.Client, error) {

	credential, cpf := qcloud.NewCredentialProfile(ud)

	if ud.Region != "" {
		cpf.HttpProfile.Endpoint = "lighthouse." + ud.Region + ".tencentcloudapi.com"
	}

	return lighthouse.NewClient(credential, ud.Region, cpf)

}

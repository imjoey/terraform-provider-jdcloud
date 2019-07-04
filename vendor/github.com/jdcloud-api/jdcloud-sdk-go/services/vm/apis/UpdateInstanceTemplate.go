// Copyright 2018 JDCLOUD.COM
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// NOTE: This class is auto generated by the jdcloud code generator program.

package apis

import (
	"github.com/jdcloud-api/jdcloud-sdk-go/core"
)

type UpdateInstanceTemplateRequest struct {
	core.JDCloudRequest

	/* 地域ID  */
	RegionId string `json:"regionId"`

	/* 启动模板ID  */
	InstanceTemplateId string `json:"instanceTemplateId"`

	/* 模板描述，<a href="http://docs.jdcloud.com/virtual-machines/api/general_parameters">参考公共参数规范</a>。 (Optional) */
	Description *string `json:"description"`

	/* 模板名称，<a href="http://docs.jdcloud.com/virtual-machines/api/general_parameters">参考公共参数规范</a>。 (Optional) */
	Name *string `json:"name"`
}

/*
 * param regionId: 地域ID (Required)
 * param instanceTemplateId: 启动模板ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewUpdateInstanceTemplateRequest(
	regionId string,
	instanceTemplateId string,
) *UpdateInstanceTemplateRequest {

	return &UpdateInstanceTemplateRequest{
		JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/instanceTemplates/{instanceTemplateId}",
			Method:  "PATCH",
			Header:  nil,
			Version: "v1",
		},
		RegionId:           regionId,
		InstanceTemplateId: instanceTemplateId,
	}
}

/*
 * param regionId: 地域ID (Required)
 * param instanceTemplateId: 启动模板ID (Required)
 * param description: 模板描述，<a href="http://docs.jdcloud.com/virtual-machines/api/general_parameters">参考公共参数规范</a>。 (Optional)
 * param name: 模板名称，<a href="http://docs.jdcloud.com/virtual-machines/api/general_parameters">参考公共参数规范</a>。 (Optional)
 */
func NewUpdateInstanceTemplateRequestWithAllParams(
	regionId string,
	instanceTemplateId string,
	description *string,
	name *string,
) *UpdateInstanceTemplateRequest {

	return &UpdateInstanceTemplateRequest{
		JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/instanceTemplates/{instanceTemplateId}",
			Method:  "PATCH",
			Header:  nil,
			Version: "v1",
		},
		RegionId:           regionId,
		InstanceTemplateId: instanceTemplateId,
		Description:        description,
		Name:               name,
	}
}

/* This constructor has better compatible ability when API parameters changed */
func NewUpdateInstanceTemplateRequestWithoutParam() *UpdateInstanceTemplateRequest {

	return &UpdateInstanceTemplateRequest{
		JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/instanceTemplates/{instanceTemplateId}",
			Method:  "PATCH",
			Header:  nil,
			Version: "v1",
		},
	}
}

/* param regionId: 地域ID(Required) */
func (r *UpdateInstanceTemplateRequest) SetRegionId(regionId string) {
	r.RegionId = regionId
}

/* param instanceTemplateId: 启动模板ID(Required) */
func (r *UpdateInstanceTemplateRequest) SetInstanceTemplateId(instanceTemplateId string) {
	r.InstanceTemplateId = instanceTemplateId
}

/* param description: 模板描述，<a href="http://docs.jdcloud.com/virtual-machines/api/general_parameters">参考公共参数规范</a>。(Optional) */
func (r *UpdateInstanceTemplateRequest) SetDescription(description string) {
	r.Description = &description
}

/* param name: 模板名称，<a href="http://docs.jdcloud.com/virtual-machines/api/general_parameters">参考公共参数规范</a>。(Optional) */
func (r *UpdateInstanceTemplateRequest) SetName(name string) {
	r.Name = &name
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r UpdateInstanceTemplateRequest) GetRegionId() string {
	return r.RegionId
}

type UpdateInstanceTemplateResponse struct {
	RequestID string                       `json:"requestId"`
	Error     core.ErrorResponse           `json:"error"`
	Result    UpdateInstanceTemplateResult `json:"result"`
}

type UpdateInstanceTemplateResult struct {
}

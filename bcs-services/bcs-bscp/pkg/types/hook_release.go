/*
Tencent is pleased to support the open source community by making Basic Service Configuration Platform available.
Copyright (C) 2019 THL A29 Limited, a Tencent company. All rights reserved.
Licensed under the MIT License (the "License"); you may not use this file except
in compliance with the License. You may obtain a copy of the License at
http://opensource.org/licenses/MIT
Unless required by applicable law or agreed to in writing, software distributed under
the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
either express or implied. See the License for the specific language governing permissions and
limitations under the License.
*/

package types

import (
	"errors"

	"bscp.io/pkg/dal/table"
)

// ListHookReleasesOption defines the response details of requested ListHookReleasesOption.
type ListHookReleasesOption struct {
	BizID     uint32    `json:"biz_id"`
	HookID    uint32    `json:"hook_id"`
	Page      *BasePage `json:"page"`
	SearchKey string    `json:"search_key"`
}

// ListHookReleaseDetails defines the response details of requested ListHooksReleaseOption.
type ListHookReleaseDetails struct {
	Count   uint32               `json:"count"`
	Details []*table.HookRelease `json:"details"`
}

// Validate the list release options
func (opt *ListHookReleasesOption) Validate(po *PageOption) error {
	if opt.BizID <= 0 {
		return errors.New("invalid biz id, should >= 1")
	}

	if opt.HookID <= 0 {
		return errors.New("invalid hook id id, should >= 1")
	}

	if opt.Page == nil {
		return errors.New("page is null")
	}

	if err := opt.Page.Validate(po); err != nil {
		return err
	}

	return nil
}

// GetByPubStateOption defines options to get hr by PubState
type GetByPubStateOption struct {
	BizID  uint32
	HookID uint32
	State  table.ReleaseStatus
}

// Validate the get ByPubState option
func (opt *GetByPubStateOption) Validate() error {
	if opt.BizID <= 0 {
		return errors.New("invalid biz id, should >= 1")
	}

	if opt.HookID <= 0 {
		return errors.New("invalid hook id id, should >= 1")
	}

	if err := opt.State.Validate(); err != nil {
		return err
	}

	return nil
}

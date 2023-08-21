/*
Tencent is pleased to support the open source community by making Basic Service Configuration Platform available.
Copyright (C) 2019 THL A29 Limited, a Tencent company. All rights reserved.
Licensed under the MIT License (the "License"); you may not use this file except
in compliance with the License. You may obtain a copy of the License at
http://opensource.org/licenses/MIT
Unless required by applicable law or agreed to in writing, software distributed under
the License is distributed on an "as IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
either express or implied. See the License for the specific language governing permissions and
limitations under the License.
*/

package service

import (
	"context"
	"encoding/json"
	"fmt"

	"bscp.io/pkg/dal/table"
	"bscp.io/pkg/kit"
	"bscp.io/pkg/logs"
	pbatb "bscp.io/pkg/protocol/core/app-template-binding"
	pbbase "bscp.io/pkg/protocol/core/base"
	pbds "bscp.io/pkg/protocol/data-service"
	"bscp.io/pkg/tools"
	"bscp.io/pkg/types"
)

// CreateAppTemplateBinding create app template binding.
func (s *Service) CreateAppTemplateBinding(ctx context.Context, req *pbds.CreateAppTemplateBindingReq) (
	*pbds.CreateResp, error) {
	kt := kit.FromGrpcContext(ctx)

	appTemplateBinding := &table.AppTemplateBinding{
		Spec:       req.Spec.AppTemplateBindingSpec(),
		Attachment: req.Attachment.AppTemplateBindingAttachment(),
		Revision: &table.Revision{
			Creator: kt.User,
			Reviser: kt.User,
		},
	}

	pbs := parseBindings(appTemplateBinding.Spec.Bindings)

	if err := s.validateATBUpsert(kt, pbs); err != nil {
		return nil, err
	}

	if err := s.fillUnspecifiedTemplates(kt, appTemplateBinding.Spec.Bindings, pbs); err != nil {
		return nil, err
	}

	if err := s.validateATBUniqueKey(kt, pbs, req.Attachment.BizId, req.Attachment.AppId); err != nil {
		return nil, err
	}

	if err := s.fillATBModel(kt, appTemplateBinding, pbs); err != nil {
		return nil, err
	}

	id, err := s.dao.AppTemplateBinding().Create(kt, appTemplateBinding)
	if err != nil {
		logs.Errorf("create app template binding failed, err: %v, rid: %s", err, kt.Rid)
		return nil, err
	}

	resp := &pbds.CreateResp{Id: id}
	return resp, nil
}

// ListAppTemplateBindings list app template binding.
func (s *Service) ListAppTemplateBindings(ctx context.Context, req *pbds.ListAppTemplateBindingsReq) (*pbds.
	ListAppTemplateBindingsResp,
	error) {
	kt := kit.FromGrpcContext(ctx)

	opt := &types.BasePage{Start: req.Start, Limit: uint(req.Limit)}
	if err := opt.Validate(types.DefaultPageOption); err != nil {
		return nil, err
	}

	details, count, err := s.dao.AppTemplateBinding().List(kt, req.BizId, req.AppId, opt)
	if err != nil {
		logs.Errorf("list app template bindings failed, err: %v, rid: %s", err, kt.Rid)
		return nil, err
	}

	resp := &pbds.ListAppTemplateBindingsResp{
		Count:   uint32(count),
		Details: pbatb.PbAppTemplateBindings(details),
	}
	return resp, nil
}

// UpdateAppTemplateBinding update app template binding.
func (s *Service) UpdateAppTemplateBinding(ctx context.Context, req *pbds.UpdateAppTemplateBindingReq) (*pbbase.
	EmptyResp,
	error) {
	kt := kit.FromGrpcContext(ctx)

	appTemplateBinding := &table.AppTemplateBinding{
		ID:         req.Id,
		Spec:       req.Spec.AppTemplateBindingSpec(),
		Attachment: req.Attachment.AppTemplateBindingAttachment(),
		Revision: &table.Revision{
			Reviser: kt.User,
		},
	}

	pbs := parseBindings(appTemplateBinding.Spec.Bindings)

	if err := s.validateATBUpsert(kt, pbs); err != nil {
		return nil, err
	}

	if err := s.fillUnspecifiedTemplates(kt, appTemplateBinding.Spec.Bindings, pbs); err != nil {
		return nil, err
	}

	if err := s.validateATBUniqueKey(kt, pbs, req.Attachment.BizId, req.Attachment.AppId); err != nil {
		return nil, err
	}

	if err := s.fillATBModel(kt, appTemplateBinding, pbs); err != nil {
		return nil, err
	}

	if err := s.dao.AppTemplateBinding().Update(kt, appTemplateBinding); err != nil {
		logs.Errorf("update app template binding failed, err: %v, rid: %s", err, kt.Rid)
		return nil, err
	}

	return new(pbbase.EmptyResp), nil
}

// DeleteAppTemplateBinding delete app template binding.
func (s *Service) DeleteAppTemplateBinding(ctx context.Context, req *pbds.DeleteAppTemplateBindingReq) (
	*pbbase.EmptyResp, error) {
	kt := kit.FromGrpcContext(ctx)

	appTemplateBinding := &table.AppTemplateBinding{
		ID:         req.Id,
		Attachment: req.Attachment.AppTemplateBindingAttachment(),
	}
	if err := s.dao.AppTemplateBinding().Delete(kt, appTemplateBinding); err != nil {
		logs.Errorf("delete app template binding failed, err: %v, rid: %s", err, kt.Rid)
		return nil, err
	}

	return new(pbbase.EmptyResp), nil
}

// ValidateAppTemplateBindingUniqueKey validate the unique key name+path for an app.
// if the unique key name+path exists in table app_template_binding for the app, return error.
func (s *Service) ValidateAppTemplateBindingUniqueKey(kt *kit.Kit, bizID, appID uint32, name,
	path string) error {
	opt := &types.BasePage{All: true}
	details, _, err := s.dao.AppTemplateBinding().List(kt, bizID, appID, opt)
	if err != nil {
		logs.Errorf("validate app template binding unique key failed, err: %v, rid: %s", err, kt.Rid)
		return err
	}
	// so far, no any template config item exists for the app
	if len(details) == 0 {
		return nil
	}

	templateRevisions, err := s.dao.TemplateRevision().ListByIDs(kt, details[0].Spec.TemplateRevisionIDs)
	if err != nil {
		logs.Errorf("validate app template binding unique key failed, err: %v, rid: %s", err, kt.Rid)
		return err
	}
	for _, tr := range templateRevisions {
		if name == tr.Spec.Name && path == tr.Spec.Path {
			return fmt.Errorf("config item's same name %s and path %s already exists", name, path)
		}
	}

	return nil
}

// fillATBModel fill model AppTemplateBinding's fields
// including TemplateSetIDs,TemplateRevisionIDs,TemplateSpaceIDs,TemplateIDs
func (s *Service) fillATBModel(kit *kit.Kit, g *table.AppTemplateBinding, pbs *parsedBindings) error {
	g.Spec.TemplateSetIDs = pbs.TemplateSetIDs
	g.Spec.TemplateRevisionIDs = pbs.TemplateRevisionIDs
	g.Spec.LatestTemplateRevisionIDs = pbs.LatestTemplateRevisionIDs
	g.Spec.TemplateIDs = pbs.TemplateIDs

	templateRevisions, err := s.dao.TemplateRevision().ListByIDs(kit, pbs.TemplateRevisionIDs)
	if err != nil {
		return err
	}

	templateSpaceIDs := make(map[uint32]struct{})
	for _, tr := range templateRevisions {
		templateSpaceIDs[tr.Attachment.TemplateSpaceID] = struct{}{}
	}
	g.Spec.TemplateSpaceIDs = convertToSlice(templateSpaceIDs)

	return nil
}

// no need to validate the bindings here, it is already done in config server
func parseBindings(bindings []*table.TemplateBinding) *parsedBindings {
	pbs := new(parsedBindings)
	for _, b := range bindings {
		pbs.TemplateSetIDs = append(pbs.TemplateSetIDs, b.TemplateSetID)

		for _, r := range b.TemplateRevisions {
			pbs.TemplateRevisionIDs = append(pbs.TemplateRevisionIDs, r.TemplateRevisionID)
			pbs.TemplateIDs = append(pbs.TemplateIDs, r.TemplateID)
			if r.IsLatest {
				pbs.LatestTemplateRevisionIDs = append(pbs.LatestTemplateRevisionIDs, r.TemplateRevisionID)
			}
		}
	}

	return pbs
}

// fillUnspecifiedTemplates update the pbs's unspecified templates and revisions
func (s *Service) fillUnspecifiedTemplates(kit *kit.Kit, bindings []*table.TemplateBinding, pbs *parsedBindings) error {
	for _, b := range bindings {
		var templateIDs []uint32
		for _, r := range b.TemplateRevisions {
			templateIDs = append(templateIDs, r.TemplateID)
		}

		// get all the templates belong to the template set
		templateSets, err := s.dao.TemplateSet().ListByIDs(kit, []uint32{b.TemplateSetID})
		if err != nil {
			logs.Errorf("parse bindings failed, err: %v, rid: %s", err, kit.Rid)
			return err
		}

		unspecified := tools.SliceDiff(templateSets[0].Spec.TemplateIDs, templateIDs)
		// Note: Use list_template_revision_by_id to get all latest revisions and update pbs's unspecified templates
		if len(unspecified) > 0 {

		}
	}

	return nil
}

type parsedBindings struct {
	TemplateIDs               []uint32
	TemplateSetIDs            []uint32
	TemplateRevisionIDs       []uint32
	LatestTemplateRevisionIDs []uint32
	TemplateBindings          []*table.TemplateBinding
}

func convertToSlice(m map[uint32]struct{}) []uint32 {
	var keys []uint32
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

// validateUpsert validate for create or update operation of app template binding
func (s *Service) validateATBUpsert(kit *kit.Kit, b *parsedBindings) error {

	if err := s.dao.Validator().ValidateTemplateSetsExist(kit, b.TemplateSetIDs); err != nil {
		return err
	}

	if err := s.dao.Validator().ValidateTemplatesExist(kit, b.TemplateIDs); err != nil {
		return err
	}

	if err := s.dao.Validator().ValidateTemplateRevisionsExist(kit, b.TemplateRevisionIDs); err != nil {
		return err
	}

	return nil
}

// validateATBUniqueKey validate unique key for app template binding
func (s *Service) validateATBUniqueKey(kit *kit.Kit, b *parsedBindings, bizID, appID uint32) error {
	templateRevisions, err := s.dao.TemplateRevision().ListByIDs(kit, b.TemplateRevisionIDs)
	if err != nil {
		return err
	}

	// validates unique key name+path both in table app_template_bindings and config_items
	// validate the input is equivalent to validate in table app_template_bindings
	if err := validateUniqueKeyOfInput(templateRevisions); err != nil {
		return err
	}
	// validate in table config_items
	for _, tr := range templateRevisions {
		if _, err := s.dao.ConfigItem().GetByUniqueKey(kit, bizID, appID, tr.Spec.Name, tr.Spec.Path); err == nil {
			return fmt.Errorf("config item's same name %s and path %s already exists", tr.Spec.Name, tr.Spec.Path)
		}
	}

	return nil
}

// validateUniqueKeyOfInput validates unique key which is name+path of input only
func validateUniqueKeyOfInput(templateRevisions []*table.TemplateRevision) error {
	var uids []uid
	for _, tr := range templateRevisions {
		uids = append(uids, uid{
			Name: tr.Spec.Name,
			Path: tr.Spec.Path,
		})
	}
	repeated := findRepeatedElements(uids)
	if len(repeated) > 0 {
		js, _ := json.Marshal(repeated)
		return fmt.Errorf("config item's name and path must be unique, these are repeated: %s", js)
	}

	return nil
}

// validateUniqueKeyForApp validates unique key which is name+path for an app
func validateUniqueKeyForApp(templateRevisions []*table.TemplateRevision, name, path string) error {
	for _, tr := range templateRevisions {
		if name == tr.Spec.Name && path == tr.Spec.Path {
			return fmt.Errorf("config item's same name %s and path %s already exists", name, path)
		}
	}

	return nil
}

type uid struct {
	Name string
	Path string
}

func findRepeatedElements(slice []uid) []uid {
	frequencyMap := make(map[uid]int)
	var repeatedElements []uid

	// Count the frequency of each uID in the slice
	for _, key := range slice {
		frequencyMap[key]++
	}

	// Check if any uID appears more than once
	for key, count := range frequencyMap {
		if count > 1 {
			repeatedElements = append(repeatedElements, key)
		}
	}

	return repeatedElements
}
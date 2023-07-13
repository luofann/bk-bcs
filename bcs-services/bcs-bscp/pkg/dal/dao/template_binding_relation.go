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

package dao

import (
	"bscp.io/pkg/dal/gen"
	"bscp.io/pkg/kit"
	"bscp.io/pkg/types"

	"gorm.io/datatypes"
	rawgen "gorm.io/gen"
)

// TemplateBindingRelation supplies all the template binding relation query operations.
type TemplateBindingRelation interface {
	// GetTemplateBoundUnnamedAppCount get bound unnamed app count of the target template.
	GetTemplateBoundUnnamedAppCount(kit *kit.Kit, bizID, templateID uint32) (uint32, error)
	// GetTemplateBoundNamedAppCount get bound named app count of the target template.
	GetTemplateBoundNamedAppCount(kit *kit.Kit, bizID, templateID uint32) (uint32, error)
	// GetTemplateBoundTemplateSetCount get bound template set count of the target template.
	GetTemplateBoundTemplateSetCount(kit *kit.Kit, bizID, templateID uint32) (uint32, error)
	// GetTemplateReleaseBoundUnnamedAppCount get bound unnamed app count of the target template release.
	GetTemplateReleaseBoundUnnamedAppCount(kit *kit.Kit, bizID, templateReleaseID uint32) (uint32, error)
	// GetTemplateReleaseBoundNamedAppCount get bound named app count of the target template release.
	GetTemplateReleaseBoundNamedAppCount(kit *kit.Kit, bizID, templateReleaseID uint32) (uint32, error)
	// GetTemplateSetBoundUnnamedAppCount get bound unnamed app count of the target template set.
	GetTemplateSetBoundUnnamedAppCount(kit *kit.Kit, bizID, templateSetID uint32) (uint32, error)
	// GetTemplateSetBoundNamedAppCount get bound named app count of the target template set.
	GetTemplateSetBoundNamedAppCount(kit *kit.Kit, bizID, templateSetID uint32) (uint32, error)
	// ListTemplateBoundUnnamedAppDetails list bound unnamed app details of the target template.
	ListTemplateBoundUnnamedAppDetails(kit *kit.Kit, bizID, templateID uint32) ([]*types.TmplBoundUnnamedAppDetail, error)
	// ListTemplateBoundNamedAppDetails list bound named app details of the target template.
	ListTemplateBoundNamedAppDetails(kit *kit.Kit, bizID, templateID uint32) ([]*types.TmplBoundNamedAppDetail, error)
	// ListTemplateBoundTemplateSetDetails list bound template set details of the target template.
	ListTemplateBoundTemplateSetDetails(kit *kit.Kit, bizID, templateID uint32) ([]uint32, error)
	// ListTemplateReleaseBoundUnnamedAppDetails list bound unnamed app details of the target template release.
	ListTemplateReleaseBoundUnnamedAppDetails(kit *kit.Kit, bizID, templateReleaseID uint32) ([]uint32, error)
	// ListTemplateReleaseBoundNamedAppDetails list bound named app details of the target template release.
	ListTemplateReleaseBoundNamedAppDetails(kit *kit.Kit, bizID, templateReleaseID uint32) ([]*types.TmplReleaseBoundNamedAppDetail, error)
	// ListTemplateSetBoundUnnamedAppDetails list bound unnamed app details of the target template set.
	ListTemplateSetBoundUnnamedAppDetails(kit *kit.Kit, bizID, templateSetID uint32) ([]uint32, error)
	// ListTemplateSetBoundNamedAppDetails list bound named app details of the target template set.
	ListTemplateSetBoundNamedAppDetails(kit *kit.Kit, bizID, templateSetID uint32) ([]*types.TmplSetBoundNamedAppDetail, error)
}

var _ TemplateBindingRelation = new(templateBindingRelationDao)

type templateBindingRelationDao struct {
	genQ     *gen.Query
	idGen    IDGenInterface
	auditDao AuditDao
}

type countResult struct {
	Counts uint32 `json:"counts"`
}

// GetTemplateBoundUnnamedAppCount get bound unnamed app count of the target template.
func (dao *templateBindingRelationDao) GetTemplateBoundUnnamedAppCount(kit *kit.Kit, bizID, templateID uint32) (
	uint32, error) {
	m := dao.genQ.AppTemplateBinding
	q := dao.genQ.AppTemplateBinding.WithContext(kit.Ctx)
	var rs countResult
	if err := q.Select(m.AppID.Distinct().Count().As("counts")).
		Where(m.BizID.Eq(bizID)).
		Where(rawgen.Cond(datatypes.JSONArrayQuery("template_ids").Contains(templateID))...).
		Scan(&rs); err != nil {
		return 0, err
	}

	return rs.Counts, nil
}

// GetTemplateBoundNamedAppCount get bound named app count of the target template.
func (dao *templateBindingRelationDao) GetTemplateBoundNamedAppCount(kit *kit.Kit, bizID, templateID uint32) (
	uint32, error) {
	m := dao.genQ.ReleasedAppTemplateBinding
	q := dao.genQ.ReleasedAppTemplateBinding.WithContext(kit.Ctx)
	var rs countResult
	if err := q.Select(m.AppID.Distinct().Count().As("counts")).
		Where(m.BizID.Eq(bizID)).
		Where(rawgen.Cond(datatypes.JSONArrayQuery("template_ids").Contains(templateID))...).
		Scan(&rs); err != nil {
		return 0, err
	}

	return rs.Counts, nil
}

// GetTemplateBoundTemplateSetCount get bound template set count of the target template.
func (dao *templateBindingRelationDao) GetTemplateBoundTemplateSetCount(kit *kit.Kit, bizID, templateID uint32) (
	uint32, error) {
	m := dao.genQ.TemplateSet
	q := dao.genQ.TemplateSet.WithContext(kit.Ctx)
	var rs countResult
	if err := q.Select(m.ID.Distinct().Count().As("counts")).
		Where(m.BizID.Eq(bizID)).
		Where(rawgen.Cond(datatypes.JSONArrayQuery("template_ids").Contains(templateID))...).
		Scan(&rs); err != nil {
		return 0, err
	}

	return rs.Counts, nil
}

// GetTemplateReleaseBoundUnnamedAppCount get bound unnamed app count of the target template release.
func (dao *templateBindingRelationDao) GetTemplateReleaseBoundUnnamedAppCount(kit *kit.Kit, bizID,
	templateReleaseID uint32) (uint32, error) {
	m := dao.genQ.AppTemplateBinding
	q := dao.genQ.AppTemplateBinding.WithContext(kit.Ctx)
	var rs countResult
	if err := q.Select(m.AppID.Distinct().Count().As("counts")).
		Where(m.BizID.Eq(bizID)).
		Where(rawgen.Cond(datatypes.JSONArrayQuery("template_release_ids").Contains(templateReleaseID))...).
		Scan(&rs); err != nil {
		return 0, err
	}

	return rs.Counts, nil
}

// GetTemplateReleaseBoundNamedAppCount get bound named app count of the target template release.
func (dao *templateBindingRelationDao) GetTemplateReleaseBoundNamedAppCount(kit *kit.Kit, bizID,
	templateReleaseID uint32) (uint32, error) {
	m := dao.genQ.ReleasedAppTemplateBinding
	q := dao.genQ.ReleasedAppTemplateBinding.WithContext(kit.Ctx)
	var rs countResult
	if err := q.Select(m.AppID.Distinct().Count().As("counts")).
		Where(m.BizID.Eq(bizID)).
		Where(rawgen.Cond(datatypes.JSONArrayQuery("template_release_ids").Contains(templateReleaseID))...).
		Scan(&rs); err != nil {
		return 0, err
	}

	return rs.Counts, nil
}

// GetTemplateSetBoundUnnamedAppCount get bound unnamed app count of the target template set.
func (dao *templateBindingRelationDao) GetTemplateSetBoundUnnamedAppCount(kit *kit.Kit, bizID,
	templateSetID uint32) (uint32, error) {
	m := dao.genQ.AppTemplateBinding
	q := dao.genQ.AppTemplateBinding.WithContext(kit.Ctx)
	var rs countResult
	if err := q.Select(m.AppID.Distinct().Count().As("counts")).
		Where(m.BizID.Eq(bizID)).
		Where(rawgen.Cond(datatypes.JSONArrayQuery("template_set_ids").Contains(templateSetID))...).
		Scan(&rs); err != nil {
		return 0, err
	}

	return rs.Counts, nil
}

// GetTemplateSetBoundNamedAppCount get bound named app count of the target template set.
func (dao *templateBindingRelationDao) GetTemplateSetBoundNamedAppCount(kit *kit.Kit, bizID, templateSetID uint32) (
	uint32, error) {
	m := dao.genQ.ReleasedAppTemplateBinding
	q := dao.genQ.ReleasedAppTemplateBinding.WithContext(kit.Ctx)
	var rs countResult
	if err := q.Select(m.AppID.Distinct().Count().As("counts")).
		Where(m.BizID.Eq(bizID)).
		Where(rawgen.Cond(datatypes.JSONArrayQuery("template_set_ids").Contains(templateSetID))...).
		Scan(&rs); err != nil {
		return 0, err
	}

	return rs.Counts, nil
}

// ListTemplateBoundUnnamedAppDetails list bound unnamed app details of the target template.
func (dao *templateBindingRelationDao) ListTemplateBoundUnnamedAppDetails(kit *kit.Kit, bizID, templateID uint32) (
	[]*types.TmplBoundUnnamedAppDetail, error) {
	m := dao.genQ.AppTemplateBinding
	q := dao.genQ.AppTemplateBinding.WithContext(kit.Ctx)
	var rs []*types.TmplBoundUnnamedAppDetail
	if err := q.Select(m.AppID, m.TemplateReleaseIDs).
		Where(m.BizID.Eq(bizID)).
		Where(rawgen.Cond(datatypes.JSONArrayQuery("template_ids").Contains(templateID))...).
		Scan(&rs); err != nil {
		return nil, err
	}

	return rs, nil
}

// ListTemplateBoundNamedAppDetails list bound named app details of the target template.
func (dao *templateBindingRelationDao) ListTemplateBoundNamedAppDetails(kit *kit.Kit, bizID, templateID uint32) (
	[]*types.TmplBoundNamedAppDetail, error) {
	m := dao.genQ.ReleasedAppTemplateBinding
	q := dao.genQ.ReleasedAppTemplateBinding.WithContext(kit.Ctx)
	var rs []*types.TmplBoundNamedAppDetail
	if err := q.Select(m.AppID, m.ReleaseID, m.TemplateReleaseIDs).
		Where(m.BizID.Eq(bizID)).
		Where(rawgen.Cond(datatypes.JSONArrayQuery("template_ids").Contains(templateID))...).
		Scan(&rs); err != nil {
		return nil, err
	}

	return rs, nil
}

// ListTemplateBoundTemplateSetDetails list bound template set details of the target template.
func (dao *templateBindingRelationDao) ListTemplateBoundTemplateSetDetails(kit *kit.Kit, bizID, templateID uint32) (
	[]uint32, error) {
	m := dao.genQ.TemplateSet
	q := dao.genQ.TemplateSet.WithContext(kit.Ctx)
	var templateSetIDs []uint32
	if err := q.Select(m.ID).
		Where(m.BizID.Eq(bizID)).
		Where(rawgen.Cond(datatypes.JSONArrayQuery("template_ids").Contains(templateID))...).
		Pluck(m.ID, &templateSetIDs); err != nil {
		return nil, err
	}

	return templateSetIDs, nil
}

// ListTemplateReleaseBoundUnnamedAppDetails list bound unnamed app details of the target template release.
func (dao *templateBindingRelationDao) ListTemplateReleaseBoundUnnamedAppDetails(kit *kit.Kit, bizID,
	templateReleaseID uint32) ([]uint32, error) {
	m := dao.genQ.AppTemplateBinding
	q := dao.genQ.AppTemplateBinding.WithContext(kit.Ctx)
	var appIDs []uint32
	if err := q.Select(m.AppID).
		Where(m.BizID.Eq(bizID)).
		Where(rawgen.Cond(datatypes.JSONArrayQuery("template_release_ids").Contains(templateReleaseID))...).
		Pluck(m.AppID, &appIDs); err != nil {
		return nil, err
	}

	return appIDs, nil
}

// ListTemplateReleaseBoundNamedAppDetails list bound named app details of the target template release.
func (dao *templateBindingRelationDao) ListTemplateReleaseBoundNamedAppDetails(kit *kit.Kit, bizID,
	templateReleaseID uint32) ([]*types.TmplReleaseBoundNamedAppDetail, error) {
	m := dao.genQ.ReleasedAppTemplateBinding
	q := dao.genQ.ReleasedAppTemplateBinding.WithContext(kit.Ctx)
	var rs []*types.TmplReleaseBoundNamedAppDetail
	if err := q.Select(m.AppID, m.ReleaseID).
		Where(m.BizID.Eq(bizID)).
		Where(rawgen.Cond(datatypes.JSONArrayQuery("template_release_ids").Contains(templateReleaseID))...).
		Scan(&rs); err != nil {
		return nil, err
	}

	return rs, nil
}

// ListTemplateSetBoundUnnamedAppDetails list bound unnamed app details of the target template set.
func (dao *templateBindingRelationDao) ListTemplateSetBoundUnnamedAppDetails(kit *kit.Kit, bizID,
	templateSetID uint32) ([]uint32, error) {
	m := dao.genQ.AppTemplateBinding
	q := dao.genQ.AppTemplateBinding.WithContext(kit.Ctx)
	var appIDs []uint32
	if err := q.Select(m.AppID).
		Where(m.BizID.Eq(bizID)).
		Where(rawgen.Cond(datatypes.JSONArrayQuery("template_set_ids").Contains(templateSetID))...).
		Pluck(m.AppID, &appIDs); err != nil {
		return nil, err
	}

	return appIDs, nil
}

// ListTemplateSetBoundNamedAppDetails list bound named app details of the target template set.
func (dao *templateBindingRelationDao) ListTemplateSetBoundNamedAppDetails(kit *kit.Kit, bizID, templateSetID uint32) (
	[]*types.TmplSetBoundNamedAppDetail, error) {
	m := dao.genQ.ReleasedAppTemplateBinding
	q := dao.genQ.ReleasedAppTemplateBinding.WithContext(kit.Ctx)
	var rs []*types.TmplSetBoundNamedAppDetail
	if err := q.Select(m.AppID, m.ReleaseID).
		Where(m.BizID.Eq(bizID)).
		Where(rawgen.Cond(datatypes.JSONArrayQuery("template_set_ids").Contains(templateSetID))...).
		Scan(&rs); err != nil {
		return nil, err
	}

	return rs, nil
}

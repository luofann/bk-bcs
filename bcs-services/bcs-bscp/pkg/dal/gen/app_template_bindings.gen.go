// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package gen

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"bscp.io/pkg/dal/table"
)

func newAppTemplateBinding(db *gorm.DB, opts ...gen.DOOption) appTemplateBinding {
	_appTemplateBinding := appTemplateBinding{}

	_appTemplateBinding.appTemplateBindingDo.UseDB(db, opts...)
	_appTemplateBinding.appTemplateBindingDo.UseModel(&table.AppTemplateBinding{})

	tableName := _appTemplateBinding.appTemplateBindingDo.TableName()
	_appTemplateBinding.ALL = field.NewAsterisk(tableName)
	_appTemplateBinding.ID = field.NewUint32(tableName, "id")
	_appTemplateBinding.TemplateSpaceIDs = field.NewField(tableName, "template_space_ids")
	_appTemplateBinding.TemplateSetIDs = field.NewField(tableName, "template_set_ids")
	_appTemplateBinding.TemplateIDs = field.NewField(tableName, "template_ids")
	_appTemplateBinding.TemplateRevisionIDs = field.NewField(tableName, "template_revision_ids")
	_appTemplateBinding.Bindings = field.NewField(tableName, "bindings")
	_appTemplateBinding.BizID = field.NewUint32(tableName, "biz_id")
	_appTemplateBinding.AppID = field.NewUint32(tableName, "app_id")
	_appTemplateBinding.Creator = field.NewString(tableName, "creator")
	_appTemplateBinding.Reviser = field.NewString(tableName, "reviser")
	_appTemplateBinding.CreatedAt = field.NewTime(tableName, "created_at")
	_appTemplateBinding.UpdatedAt = field.NewTime(tableName, "updated_at")

	_appTemplateBinding.fillFieldMap()

	return _appTemplateBinding
}

type appTemplateBinding struct {
	appTemplateBindingDo appTemplateBindingDo

	ALL                 field.Asterisk
	ID                  field.Uint32
	TemplateSpaceIDs    field.Field
	TemplateSetIDs      field.Field
	TemplateIDs         field.Field
	TemplateRevisionIDs field.Field
	Bindings            field.Field
	BizID               field.Uint32
	AppID               field.Uint32
	Creator             field.String
	Reviser             field.String
	CreatedAt           field.Time
	UpdatedAt           field.Time

	fieldMap map[string]field.Expr
}

func (a appTemplateBinding) Table(newTableName string) *appTemplateBinding {
	a.appTemplateBindingDo.UseTable(newTableName)
	return a.updateTableName(newTableName)
}

func (a appTemplateBinding) As(alias string) *appTemplateBinding {
	a.appTemplateBindingDo.DO = *(a.appTemplateBindingDo.As(alias).(*gen.DO))
	return a.updateTableName(alias)
}

func (a *appTemplateBinding) updateTableName(table string) *appTemplateBinding {
	a.ALL = field.NewAsterisk(table)
	a.ID = field.NewUint32(table, "id")
	a.TemplateSpaceIDs = field.NewField(table, "template_space_ids")
	a.TemplateSetIDs = field.NewField(table, "template_set_ids")
	a.TemplateIDs = field.NewField(table, "template_ids")
	a.TemplateRevisionIDs = field.NewField(table, "template_revision_ids")
	a.Bindings = field.NewField(table, "bindings")
	a.BizID = field.NewUint32(table, "biz_id")
	a.AppID = field.NewUint32(table, "app_id")
	a.Creator = field.NewString(table, "creator")
	a.Reviser = field.NewString(table, "reviser")
	a.CreatedAt = field.NewTime(table, "created_at")
	a.UpdatedAt = field.NewTime(table, "updated_at")

	a.fillFieldMap()

	return a
}

func (a *appTemplateBinding) WithContext(ctx context.Context) IAppTemplateBindingDo {
	return a.appTemplateBindingDo.WithContext(ctx)
}

func (a appTemplateBinding) TableName() string { return a.appTemplateBindingDo.TableName() }

func (a appTemplateBinding) Alias() string { return a.appTemplateBindingDo.Alias() }

func (a *appTemplateBinding) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := a.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (a *appTemplateBinding) fillFieldMap() {
	a.fieldMap = make(map[string]field.Expr, 12)
	a.fieldMap["id"] = a.ID
	a.fieldMap["template_space_ids"] = a.TemplateSpaceIDs
	a.fieldMap["template_set_ids"] = a.TemplateSetIDs
	a.fieldMap["template_ids"] = a.TemplateIDs
	a.fieldMap["template_revision_ids"] = a.TemplateRevisionIDs
	a.fieldMap["bindings"] = a.Bindings
	a.fieldMap["biz_id"] = a.BizID
	a.fieldMap["app_id"] = a.AppID
	a.fieldMap["creator"] = a.Creator
	a.fieldMap["reviser"] = a.Reviser
	a.fieldMap["created_at"] = a.CreatedAt
	a.fieldMap["updated_at"] = a.UpdatedAt
}

func (a appTemplateBinding) clone(db *gorm.DB) appTemplateBinding {
	a.appTemplateBindingDo.ReplaceConnPool(db.Statement.ConnPool)
	return a
}

func (a appTemplateBinding) replaceDB(db *gorm.DB) appTemplateBinding {
	a.appTemplateBindingDo.ReplaceDB(db)
	return a
}

type appTemplateBindingDo struct{ gen.DO }

type IAppTemplateBindingDo interface {
	gen.SubQuery
	Debug() IAppTemplateBindingDo
	WithContext(ctx context.Context) IAppTemplateBindingDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IAppTemplateBindingDo
	WriteDB() IAppTemplateBindingDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IAppTemplateBindingDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IAppTemplateBindingDo
	Not(conds ...gen.Condition) IAppTemplateBindingDo
	Or(conds ...gen.Condition) IAppTemplateBindingDo
	Select(conds ...field.Expr) IAppTemplateBindingDo
	Where(conds ...gen.Condition) IAppTemplateBindingDo
	Order(conds ...field.Expr) IAppTemplateBindingDo
	Distinct(cols ...field.Expr) IAppTemplateBindingDo
	Omit(cols ...field.Expr) IAppTemplateBindingDo
	Join(table schema.Tabler, on ...field.Expr) IAppTemplateBindingDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IAppTemplateBindingDo
	RightJoin(table schema.Tabler, on ...field.Expr) IAppTemplateBindingDo
	Group(cols ...field.Expr) IAppTemplateBindingDo
	Having(conds ...gen.Condition) IAppTemplateBindingDo
	Limit(limit int) IAppTemplateBindingDo
	Offset(offset int) IAppTemplateBindingDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IAppTemplateBindingDo
	Unscoped() IAppTemplateBindingDo
	Create(values ...*table.AppTemplateBinding) error
	CreateInBatches(values []*table.AppTemplateBinding, batchSize int) error
	Save(values ...*table.AppTemplateBinding) error
	First() (*table.AppTemplateBinding, error)
	Take() (*table.AppTemplateBinding, error)
	Last() (*table.AppTemplateBinding, error)
	Find() ([]*table.AppTemplateBinding, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*table.AppTemplateBinding, err error)
	FindInBatches(result *[]*table.AppTemplateBinding, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*table.AppTemplateBinding) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IAppTemplateBindingDo
	Assign(attrs ...field.AssignExpr) IAppTemplateBindingDo
	Joins(fields ...field.RelationField) IAppTemplateBindingDo
	Preload(fields ...field.RelationField) IAppTemplateBindingDo
	FirstOrInit() (*table.AppTemplateBinding, error)
	FirstOrCreate() (*table.AppTemplateBinding, error)
	FindByPage(offset int, limit int) (result []*table.AppTemplateBinding, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IAppTemplateBindingDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (a appTemplateBindingDo) Debug() IAppTemplateBindingDo {
	return a.withDO(a.DO.Debug())
}

func (a appTemplateBindingDo) WithContext(ctx context.Context) IAppTemplateBindingDo {
	return a.withDO(a.DO.WithContext(ctx))
}

func (a appTemplateBindingDo) ReadDB() IAppTemplateBindingDo {
	return a.Clauses(dbresolver.Read)
}

func (a appTemplateBindingDo) WriteDB() IAppTemplateBindingDo {
	return a.Clauses(dbresolver.Write)
}

func (a appTemplateBindingDo) Session(config *gorm.Session) IAppTemplateBindingDo {
	return a.withDO(a.DO.Session(config))
}

func (a appTemplateBindingDo) Clauses(conds ...clause.Expression) IAppTemplateBindingDo {
	return a.withDO(a.DO.Clauses(conds...))
}

func (a appTemplateBindingDo) Returning(value interface{}, columns ...string) IAppTemplateBindingDo {
	return a.withDO(a.DO.Returning(value, columns...))
}

func (a appTemplateBindingDo) Not(conds ...gen.Condition) IAppTemplateBindingDo {
	return a.withDO(a.DO.Not(conds...))
}

func (a appTemplateBindingDo) Or(conds ...gen.Condition) IAppTemplateBindingDo {
	return a.withDO(a.DO.Or(conds...))
}

func (a appTemplateBindingDo) Select(conds ...field.Expr) IAppTemplateBindingDo {
	return a.withDO(a.DO.Select(conds...))
}

func (a appTemplateBindingDo) Where(conds ...gen.Condition) IAppTemplateBindingDo {
	return a.withDO(a.DO.Where(conds...))
}

func (a appTemplateBindingDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IAppTemplateBindingDo {
	return a.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (a appTemplateBindingDo) Order(conds ...field.Expr) IAppTemplateBindingDo {
	return a.withDO(a.DO.Order(conds...))
}

func (a appTemplateBindingDo) Distinct(cols ...field.Expr) IAppTemplateBindingDo {
	return a.withDO(a.DO.Distinct(cols...))
}

func (a appTemplateBindingDo) Omit(cols ...field.Expr) IAppTemplateBindingDo {
	return a.withDO(a.DO.Omit(cols...))
}

func (a appTemplateBindingDo) Join(table schema.Tabler, on ...field.Expr) IAppTemplateBindingDo {
	return a.withDO(a.DO.Join(table, on...))
}

func (a appTemplateBindingDo) LeftJoin(table schema.Tabler, on ...field.Expr) IAppTemplateBindingDo {
	return a.withDO(a.DO.LeftJoin(table, on...))
}

func (a appTemplateBindingDo) RightJoin(table schema.Tabler, on ...field.Expr) IAppTemplateBindingDo {
	return a.withDO(a.DO.RightJoin(table, on...))
}

func (a appTemplateBindingDo) Group(cols ...field.Expr) IAppTemplateBindingDo {
	return a.withDO(a.DO.Group(cols...))
}

func (a appTemplateBindingDo) Having(conds ...gen.Condition) IAppTemplateBindingDo {
	return a.withDO(a.DO.Having(conds...))
}

func (a appTemplateBindingDo) Limit(limit int) IAppTemplateBindingDo {
	return a.withDO(a.DO.Limit(limit))
}

func (a appTemplateBindingDo) Offset(offset int) IAppTemplateBindingDo {
	return a.withDO(a.DO.Offset(offset))
}

func (a appTemplateBindingDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IAppTemplateBindingDo {
	return a.withDO(a.DO.Scopes(funcs...))
}

func (a appTemplateBindingDo) Unscoped() IAppTemplateBindingDo {
	return a.withDO(a.DO.Unscoped())
}

func (a appTemplateBindingDo) Create(values ...*table.AppTemplateBinding) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Create(values)
}

func (a appTemplateBindingDo) CreateInBatches(values []*table.AppTemplateBinding, batchSize int) error {
	return a.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (a appTemplateBindingDo) Save(values ...*table.AppTemplateBinding) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Save(values)
}

func (a appTemplateBindingDo) First() (*table.AppTemplateBinding, error) {
	if result, err := a.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*table.AppTemplateBinding), nil
	}
}

func (a appTemplateBindingDo) Take() (*table.AppTemplateBinding, error) {
	if result, err := a.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*table.AppTemplateBinding), nil
	}
}

func (a appTemplateBindingDo) Last() (*table.AppTemplateBinding, error) {
	if result, err := a.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*table.AppTemplateBinding), nil
	}
}

func (a appTemplateBindingDo) Find() ([]*table.AppTemplateBinding, error) {
	result, err := a.DO.Find()
	return result.([]*table.AppTemplateBinding), err
}

func (a appTemplateBindingDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*table.AppTemplateBinding, err error) {
	buf := make([]*table.AppTemplateBinding, 0, batchSize)
	err = a.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (a appTemplateBindingDo) FindInBatches(result *[]*table.AppTemplateBinding, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return a.DO.FindInBatches(result, batchSize, fc)
}

func (a appTemplateBindingDo) Attrs(attrs ...field.AssignExpr) IAppTemplateBindingDo {
	return a.withDO(a.DO.Attrs(attrs...))
}

func (a appTemplateBindingDo) Assign(attrs ...field.AssignExpr) IAppTemplateBindingDo {
	return a.withDO(a.DO.Assign(attrs...))
}

func (a appTemplateBindingDo) Joins(fields ...field.RelationField) IAppTemplateBindingDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Joins(_f))
	}
	return &a
}

func (a appTemplateBindingDo) Preload(fields ...field.RelationField) IAppTemplateBindingDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Preload(_f))
	}
	return &a
}

func (a appTemplateBindingDo) FirstOrInit() (*table.AppTemplateBinding, error) {
	if result, err := a.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*table.AppTemplateBinding), nil
	}
}

func (a appTemplateBindingDo) FirstOrCreate() (*table.AppTemplateBinding, error) {
	if result, err := a.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*table.AppTemplateBinding), nil
	}
}

func (a appTemplateBindingDo) FindByPage(offset int, limit int) (result []*table.AppTemplateBinding, count int64, err error) {
	result, err = a.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = a.Offset(-1).Limit(-1).Count()
	return
}

func (a appTemplateBindingDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = a.Count()
	if err != nil {
		return
	}

	err = a.Offset(offset).Limit(limit).Scan(result)
	return
}

func (a appTemplateBindingDo) Scan(result interface{}) (err error) {
	return a.DO.Scan(result)
}

func (a appTemplateBindingDo) Delete(models ...*table.AppTemplateBinding) (result gen.ResultInfo, err error) {
	return a.DO.Delete(models)
}

func (a *appTemplateBindingDo) withDO(do gen.Dao) *appTemplateBindingDo {
	a.DO = *do.(*gen.DO)
	return a
}

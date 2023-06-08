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

func newHook(db *gorm.DB, opts ...gen.DOOption) hook {
	_hook := hook{}

	_hook.hookDo.UseDB(db, opts...)
	_hook.hookDo.UseModel(&table.Hook{})

	tableName := _hook.hookDo.TableName()
	_hook.ALL = field.NewAsterisk(tableName)
	_hook.ID = field.NewUint32(tableName, "id")
	_hook.Name = field.NewString(tableName, "name")
	_hook.Type = field.NewString(tableName, "type")
	_hook.PublishNum = field.NewUint32(tableName, "publish_num")
	_hook.PubState = field.NewString(tableName, "pub_state")
	_hook.Tag = field.NewString(tableName, "tag")
	_hook.Memo = field.NewString(tableName, "memo")
	_hook.BizID = field.NewUint32(tableName, "biz_id")
	_hook.Creator = field.NewString(tableName, "creator")
	_hook.Reviser = field.NewString(tableName, "reviser")
	_hook.CreatedAt = field.NewTime(tableName, "created_at")
	_hook.UpdatedAt = field.NewTime(tableName, "updated_at")

	_hook.fillFieldMap()

	return _hook
}

type hook struct {
	hookDo hookDo

	ALL        field.Asterisk
	ID         field.Uint32
	Name       field.String
	Type       field.String
	PublishNum field.Uint32
	PubState   field.String
	Tag        field.String
	Memo       field.String
	BizID      field.Uint32
	Creator    field.String
	Reviser    field.String
	CreatedAt  field.Time
	UpdatedAt  field.Time

	fieldMap map[string]field.Expr
}

func (h hook) Table(newTableName string) *hook {
	h.hookDo.UseTable(newTableName)
	return h.updateTableName(newTableName)
}

func (h hook) As(alias string) *hook {
	h.hookDo.DO = *(h.hookDo.As(alias).(*gen.DO))
	return h.updateTableName(alias)
}

func (h *hook) updateTableName(table string) *hook {
	h.ALL = field.NewAsterisk(table)
	h.ID = field.NewUint32(table, "id")
	h.Name = field.NewString(table, "name")
	h.Type = field.NewString(table, "type")
	h.PublishNum = field.NewUint32(table, "publish_num")
	h.PubState = field.NewString(table, "pub_state")
	h.Tag = field.NewString(table, "tag")
	h.Memo = field.NewString(table, "memo")
	h.BizID = field.NewUint32(table, "biz_id")
	h.Creator = field.NewString(table, "creator")
	h.Reviser = field.NewString(table, "reviser")
	h.CreatedAt = field.NewTime(table, "created_at")
	h.UpdatedAt = field.NewTime(table, "updated_at")

	h.fillFieldMap()

	return h
}

func (h *hook) WithContext(ctx context.Context) IHookDo { return h.hookDo.WithContext(ctx) }

func (h hook) TableName() string { return h.hookDo.TableName() }

func (h hook) Alias() string { return h.hookDo.Alias() }

func (h *hook) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := h.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (h *hook) fillFieldMap() {
	h.fieldMap = make(map[string]field.Expr, 12)
	h.fieldMap["id"] = h.ID
	h.fieldMap["name"] = h.Name
	h.fieldMap["type"] = h.Type
	h.fieldMap["publish_num"] = h.PublishNum
	h.fieldMap["pub_state"] = h.PubState
	h.fieldMap["tag"] = h.Tag
	h.fieldMap["memo"] = h.Memo
	h.fieldMap["biz_id"] = h.BizID
	h.fieldMap["creator"] = h.Creator
	h.fieldMap["reviser"] = h.Reviser
	h.fieldMap["created_at"] = h.CreatedAt
	h.fieldMap["updated_at"] = h.UpdatedAt
}

func (h hook) clone(db *gorm.DB) hook {
	h.hookDo.ReplaceConnPool(db.Statement.ConnPool)
	return h
}

func (h hook) replaceDB(db *gorm.DB) hook {
	h.hookDo.ReplaceDB(db)
	return h
}

type hookDo struct{ gen.DO }

type IHookDo interface {
	gen.SubQuery
	Debug() IHookDo
	WithContext(ctx context.Context) IHookDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IHookDo
	WriteDB() IHookDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IHookDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IHookDo
	Not(conds ...gen.Condition) IHookDo
	Or(conds ...gen.Condition) IHookDo
	Select(conds ...field.Expr) IHookDo
	Where(conds ...gen.Condition) IHookDo
	Order(conds ...field.Expr) IHookDo
	Distinct(cols ...field.Expr) IHookDo
	Omit(cols ...field.Expr) IHookDo
	Join(table schema.Tabler, on ...field.Expr) IHookDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IHookDo
	RightJoin(table schema.Tabler, on ...field.Expr) IHookDo
	Group(cols ...field.Expr) IHookDo
	Having(conds ...gen.Condition) IHookDo
	Limit(limit int) IHookDo
	Offset(offset int) IHookDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IHookDo
	Unscoped() IHookDo
	Create(values ...*table.Hook) error
	CreateInBatches(values []*table.Hook, batchSize int) error
	Save(values ...*table.Hook) error
	First() (*table.Hook, error)
	Take() (*table.Hook, error)
	Last() (*table.Hook, error)
	Find() ([]*table.Hook, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*table.Hook, err error)
	FindInBatches(result *[]*table.Hook, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*table.Hook) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IHookDo
	Assign(attrs ...field.AssignExpr) IHookDo
	Joins(fields ...field.RelationField) IHookDo
	Preload(fields ...field.RelationField) IHookDo
	FirstOrInit() (*table.Hook, error)
	FirstOrCreate() (*table.Hook, error)
	FindByPage(offset int, limit int) (result []*table.Hook, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IHookDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (h hookDo) Debug() IHookDo {
	return h.withDO(h.DO.Debug())
}

func (h hookDo) WithContext(ctx context.Context) IHookDo {
	return h.withDO(h.DO.WithContext(ctx))
}

func (h hookDo) ReadDB() IHookDo {
	return h.Clauses(dbresolver.Read)
}

func (h hookDo) WriteDB() IHookDo {
	return h.Clauses(dbresolver.Write)
}

func (h hookDo) Session(config *gorm.Session) IHookDo {
	return h.withDO(h.DO.Session(config))
}

func (h hookDo) Clauses(conds ...clause.Expression) IHookDo {
	return h.withDO(h.DO.Clauses(conds...))
}

func (h hookDo) Returning(value interface{}, columns ...string) IHookDo {
	return h.withDO(h.DO.Returning(value, columns...))
}

func (h hookDo) Not(conds ...gen.Condition) IHookDo {
	return h.withDO(h.DO.Not(conds...))
}

func (h hookDo) Or(conds ...gen.Condition) IHookDo {
	return h.withDO(h.DO.Or(conds...))
}

func (h hookDo) Select(conds ...field.Expr) IHookDo {
	return h.withDO(h.DO.Select(conds...))
}

func (h hookDo) Where(conds ...gen.Condition) IHookDo {
	return h.withDO(h.DO.Where(conds...))
}

func (h hookDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IHookDo {
	return h.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (h hookDo) Order(conds ...field.Expr) IHookDo {
	return h.withDO(h.DO.Order(conds...))
}

func (h hookDo) Distinct(cols ...field.Expr) IHookDo {
	return h.withDO(h.DO.Distinct(cols...))
}

func (h hookDo) Omit(cols ...field.Expr) IHookDo {
	return h.withDO(h.DO.Omit(cols...))
}

func (h hookDo) Join(table schema.Tabler, on ...field.Expr) IHookDo {
	return h.withDO(h.DO.Join(table, on...))
}

func (h hookDo) LeftJoin(table schema.Tabler, on ...field.Expr) IHookDo {
	return h.withDO(h.DO.LeftJoin(table, on...))
}

func (h hookDo) RightJoin(table schema.Tabler, on ...field.Expr) IHookDo {
	return h.withDO(h.DO.RightJoin(table, on...))
}

func (h hookDo) Group(cols ...field.Expr) IHookDo {
	return h.withDO(h.DO.Group(cols...))
}

func (h hookDo) Having(conds ...gen.Condition) IHookDo {
	return h.withDO(h.DO.Having(conds...))
}

func (h hookDo) Limit(limit int) IHookDo {
	return h.withDO(h.DO.Limit(limit))
}

func (h hookDo) Offset(offset int) IHookDo {
	return h.withDO(h.DO.Offset(offset))
}

func (h hookDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IHookDo {
	return h.withDO(h.DO.Scopes(funcs...))
}

func (h hookDo) Unscoped() IHookDo {
	return h.withDO(h.DO.Unscoped())
}

func (h hookDo) Create(values ...*table.Hook) error {
	if len(values) == 0 {
		return nil
	}
	return h.DO.Create(values)
}

func (h hookDo) CreateInBatches(values []*table.Hook, batchSize int) error {
	return h.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (h hookDo) Save(values ...*table.Hook) error {
	if len(values) == 0 {
		return nil
	}
	return h.DO.Save(values)
}

func (h hookDo) First() (*table.Hook, error) {
	if result, err := h.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*table.Hook), nil
	}
}

func (h hookDo) Take() (*table.Hook, error) {
	if result, err := h.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*table.Hook), nil
	}
}

func (h hookDo) Last() (*table.Hook, error) {
	if result, err := h.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*table.Hook), nil
	}
}

func (h hookDo) Find() ([]*table.Hook, error) {
	result, err := h.DO.Find()
	return result.([]*table.Hook), err
}

func (h hookDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*table.Hook, err error) {
	buf := make([]*table.Hook, 0, batchSize)
	err = h.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (h hookDo) FindInBatches(result *[]*table.Hook, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return h.DO.FindInBatches(result, batchSize, fc)
}

func (h hookDo) Attrs(attrs ...field.AssignExpr) IHookDo {
	return h.withDO(h.DO.Attrs(attrs...))
}

func (h hookDo) Assign(attrs ...field.AssignExpr) IHookDo {
	return h.withDO(h.DO.Assign(attrs...))
}

func (h hookDo) Joins(fields ...field.RelationField) IHookDo {
	for _, _f := range fields {
		h = *h.withDO(h.DO.Joins(_f))
	}
	return &h
}

func (h hookDo) Preload(fields ...field.RelationField) IHookDo {
	for _, _f := range fields {
		h = *h.withDO(h.DO.Preload(_f))
	}
	return &h
}

func (h hookDo) FirstOrInit() (*table.Hook, error) {
	if result, err := h.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*table.Hook), nil
	}
}

func (h hookDo) FirstOrCreate() (*table.Hook, error) {
	if result, err := h.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*table.Hook), nil
	}
}

func (h hookDo) FindByPage(offset int, limit int) (result []*table.Hook, count int64, err error) {
	result, err = h.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = h.Offset(-1).Limit(-1).Count()
	return
}

func (h hookDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = h.Count()
	if err != nil {
		return
	}

	err = h.Offset(offset).Limit(limit).Scan(result)
	return
}

func (h hookDo) Scan(result interface{}) (err error) {
	return h.DO.Scan(result)
}

func (h hookDo) Delete(models ...*table.Hook) (result gen.ResultInfo, err error) {
	return h.DO.Delete(models)
}

func (h *hookDo) withDO(do gen.Dao) *hookDo {
	h.DO = *do.(*gen.DO)
	return h
}

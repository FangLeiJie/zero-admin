// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"
	"database/sql"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/feihua/zero-admin/rpc/sms/gen/model"
)

func newSmsCouponScope(db *gorm.DB, opts ...gen.DOOption) smsCouponScope {
	_smsCouponScope := smsCouponScope{}

	_smsCouponScope.smsCouponScopeDo.UseDB(db, opts...)
	_smsCouponScope.smsCouponScopeDo.UseModel(&model.SmsCouponScope{})

	tableName := _smsCouponScope.smsCouponScopeDo.TableName()
	_smsCouponScope.ALL = field.NewAsterisk(tableName)
	_smsCouponScope.ID = field.NewInt64(tableName, "id")
	_smsCouponScope.CouponID = field.NewInt64(tableName, "coupon_id")
	_smsCouponScope.ScopeType = field.NewInt32(tableName, "scope_type")
	_smsCouponScope.ScopeID = field.NewInt64(tableName, "scope_id")
	_smsCouponScope.CreateTime = field.NewTime(tableName, "create_time")

	_smsCouponScope.fillFieldMap()

	return _smsCouponScope
}

// smsCouponScope 优惠券使用范围表
type smsCouponScope struct {
	smsCouponScopeDo smsCouponScopeDo

	ALL        field.Asterisk
	ID         field.Int64 // 主键ID
	CouponID   field.Int64 // 优惠券ID
	ScopeType  field.Int32 // 范围类型：0-全场，1-分类，2-商品
	ScopeID    field.Int64 // 范围ID（分类ID或商品ID）
	CreateTime field.Time  // 创建时间

	fieldMap map[string]field.Expr
}

func (s smsCouponScope) Table(newTableName string) *smsCouponScope {
	s.smsCouponScopeDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s smsCouponScope) As(alias string) *smsCouponScope {
	s.smsCouponScopeDo.DO = *(s.smsCouponScopeDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *smsCouponScope) updateTableName(table string) *smsCouponScope {
	s.ALL = field.NewAsterisk(table)
	s.ID = field.NewInt64(table, "id")
	s.CouponID = field.NewInt64(table, "coupon_id")
	s.ScopeType = field.NewInt32(table, "scope_type")
	s.ScopeID = field.NewInt64(table, "scope_id")
	s.CreateTime = field.NewTime(table, "create_time")

	s.fillFieldMap()

	return s
}

func (s *smsCouponScope) WithContext(ctx context.Context) ISmsCouponScopeDo {
	return s.smsCouponScopeDo.WithContext(ctx)
}

func (s smsCouponScope) TableName() string { return s.smsCouponScopeDo.TableName() }

func (s smsCouponScope) Alias() string { return s.smsCouponScopeDo.Alias() }

func (s smsCouponScope) Columns(cols ...field.Expr) gen.Columns {
	return s.smsCouponScopeDo.Columns(cols...)
}

func (s *smsCouponScope) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *smsCouponScope) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 5)
	s.fieldMap["id"] = s.ID
	s.fieldMap["coupon_id"] = s.CouponID
	s.fieldMap["scope_type"] = s.ScopeType
	s.fieldMap["scope_id"] = s.ScopeID
	s.fieldMap["create_time"] = s.CreateTime
}

func (s smsCouponScope) clone(db *gorm.DB) smsCouponScope {
	s.smsCouponScopeDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s smsCouponScope) replaceDB(db *gorm.DB) smsCouponScope {
	s.smsCouponScopeDo.ReplaceDB(db)
	return s
}

type smsCouponScopeDo struct{ gen.DO }

type ISmsCouponScopeDo interface {
	gen.SubQuery
	Debug() ISmsCouponScopeDo
	WithContext(ctx context.Context) ISmsCouponScopeDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ISmsCouponScopeDo
	WriteDB() ISmsCouponScopeDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ISmsCouponScopeDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ISmsCouponScopeDo
	Not(conds ...gen.Condition) ISmsCouponScopeDo
	Or(conds ...gen.Condition) ISmsCouponScopeDo
	Select(conds ...field.Expr) ISmsCouponScopeDo
	Where(conds ...gen.Condition) ISmsCouponScopeDo
	Order(conds ...field.Expr) ISmsCouponScopeDo
	Distinct(cols ...field.Expr) ISmsCouponScopeDo
	Omit(cols ...field.Expr) ISmsCouponScopeDo
	Join(table schema.Tabler, on ...field.Expr) ISmsCouponScopeDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ISmsCouponScopeDo
	RightJoin(table schema.Tabler, on ...field.Expr) ISmsCouponScopeDo
	Group(cols ...field.Expr) ISmsCouponScopeDo
	Having(conds ...gen.Condition) ISmsCouponScopeDo
	Limit(limit int) ISmsCouponScopeDo
	Offset(offset int) ISmsCouponScopeDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ISmsCouponScopeDo
	Unscoped() ISmsCouponScopeDo
	Create(values ...*model.SmsCouponScope) error
	CreateInBatches(values []*model.SmsCouponScope, batchSize int) error
	Save(values ...*model.SmsCouponScope) error
	First() (*model.SmsCouponScope, error)
	Take() (*model.SmsCouponScope, error)
	Last() (*model.SmsCouponScope, error)
	Find() ([]*model.SmsCouponScope, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SmsCouponScope, err error)
	FindInBatches(result *[]*model.SmsCouponScope, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.SmsCouponScope) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ISmsCouponScopeDo
	Assign(attrs ...field.AssignExpr) ISmsCouponScopeDo
	Joins(fields ...field.RelationField) ISmsCouponScopeDo
	Preload(fields ...field.RelationField) ISmsCouponScopeDo
	FirstOrInit() (*model.SmsCouponScope, error)
	FirstOrCreate() (*model.SmsCouponScope, error)
	FindByPage(offset int, limit int) (result []*model.SmsCouponScope, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Rows() (*sql.Rows, error)
	Row() *sql.Row
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ISmsCouponScopeDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (s smsCouponScopeDo) Debug() ISmsCouponScopeDo {
	return s.withDO(s.DO.Debug())
}

func (s smsCouponScopeDo) WithContext(ctx context.Context) ISmsCouponScopeDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s smsCouponScopeDo) ReadDB() ISmsCouponScopeDo {
	return s.Clauses(dbresolver.Read)
}

func (s smsCouponScopeDo) WriteDB() ISmsCouponScopeDo {
	return s.Clauses(dbresolver.Write)
}

func (s smsCouponScopeDo) Session(config *gorm.Session) ISmsCouponScopeDo {
	return s.withDO(s.DO.Session(config))
}

func (s smsCouponScopeDo) Clauses(conds ...clause.Expression) ISmsCouponScopeDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s smsCouponScopeDo) Returning(value interface{}, columns ...string) ISmsCouponScopeDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s smsCouponScopeDo) Not(conds ...gen.Condition) ISmsCouponScopeDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s smsCouponScopeDo) Or(conds ...gen.Condition) ISmsCouponScopeDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s smsCouponScopeDo) Select(conds ...field.Expr) ISmsCouponScopeDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s smsCouponScopeDo) Where(conds ...gen.Condition) ISmsCouponScopeDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s smsCouponScopeDo) Order(conds ...field.Expr) ISmsCouponScopeDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s smsCouponScopeDo) Distinct(cols ...field.Expr) ISmsCouponScopeDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s smsCouponScopeDo) Omit(cols ...field.Expr) ISmsCouponScopeDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s smsCouponScopeDo) Join(table schema.Tabler, on ...field.Expr) ISmsCouponScopeDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s smsCouponScopeDo) LeftJoin(table schema.Tabler, on ...field.Expr) ISmsCouponScopeDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s smsCouponScopeDo) RightJoin(table schema.Tabler, on ...field.Expr) ISmsCouponScopeDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s smsCouponScopeDo) Group(cols ...field.Expr) ISmsCouponScopeDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s smsCouponScopeDo) Having(conds ...gen.Condition) ISmsCouponScopeDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s smsCouponScopeDo) Limit(limit int) ISmsCouponScopeDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s smsCouponScopeDo) Offset(offset int) ISmsCouponScopeDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s smsCouponScopeDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ISmsCouponScopeDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s smsCouponScopeDo) Unscoped() ISmsCouponScopeDo {
	return s.withDO(s.DO.Unscoped())
}

func (s smsCouponScopeDo) Create(values ...*model.SmsCouponScope) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s smsCouponScopeDo) CreateInBatches(values []*model.SmsCouponScope, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s smsCouponScopeDo) Save(values ...*model.SmsCouponScope) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s smsCouponScopeDo) First() (*model.SmsCouponScope, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.SmsCouponScope), nil
	}
}

func (s smsCouponScopeDo) Take() (*model.SmsCouponScope, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.SmsCouponScope), nil
	}
}

func (s smsCouponScopeDo) Last() (*model.SmsCouponScope, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.SmsCouponScope), nil
	}
}

func (s smsCouponScopeDo) Find() ([]*model.SmsCouponScope, error) {
	result, err := s.DO.Find()
	return result.([]*model.SmsCouponScope), err
}

func (s smsCouponScopeDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SmsCouponScope, err error) {
	buf := make([]*model.SmsCouponScope, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s smsCouponScopeDo) FindInBatches(result *[]*model.SmsCouponScope, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s smsCouponScopeDo) Attrs(attrs ...field.AssignExpr) ISmsCouponScopeDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s smsCouponScopeDo) Assign(attrs ...field.AssignExpr) ISmsCouponScopeDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s smsCouponScopeDo) Joins(fields ...field.RelationField) ISmsCouponScopeDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s smsCouponScopeDo) Preload(fields ...field.RelationField) ISmsCouponScopeDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s smsCouponScopeDo) FirstOrInit() (*model.SmsCouponScope, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.SmsCouponScope), nil
	}
}

func (s smsCouponScopeDo) FirstOrCreate() (*model.SmsCouponScope, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.SmsCouponScope), nil
	}
}

func (s smsCouponScopeDo) FindByPage(offset int, limit int) (result []*model.SmsCouponScope, count int64, err error) {
	result, err = s.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = s.Offset(-1).Limit(-1).Count()
	return
}

func (s smsCouponScopeDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s smsCouponScopeDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s smsCouponScopeDo) Delete(models ...*model.SmsCouponScope) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *smsCouponScopeDo) withDO(do gen.Dao) *smsCouponScopeDo {
	s.DO = *do.(*gen.DO)
	return s
}

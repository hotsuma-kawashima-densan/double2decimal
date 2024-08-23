// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"project+test/pkg/db/model"
)

func newTemp(db *gorm.DB, opts ...gen.DOOption) temp {
	_temp := temp{}

	_temp.tempDo.UseDB(db, opts...)
	_temp.tempDo.UseModel(&model.Temp{})

	tableName := _temp.tempDo.TableName()
	_temp.ALL = field.NewAsterisk(tableName)
	_temp.ID = field.NewInt32(tableName, "id")
	_temp.Input = field.NewString(tableName, "input")
	_temp.A = field.NewFloat64(tableName, "A")
	_temp.B = field.NewField(tableName, "B")
	_temp.B0 = field.NewField(tableName, "B0")
	_temp.B1 = field.NewField(tableName, "B1")
	_temp.B2 = field.NewField(tableName, "B2")
	_temp.B3 = field.NewField(tableName, "B3")
	_temp.B4 = field.NewField(tableName, "B4")

	_temp.fillFieldMap()

	return _temp
}

type temp struct {
	tempDo tempDo

	ALL   field.Asterisk
	ID    field.Int32
	Input field.String
	A     field.Float64
	B     field.Field
	B0    field.Field
	B1    field.Field
	B2    field.Field
	B3    field.Field
	B4    field.Field

	fieldMap map[string]field.Expr
}

func (t temp) Table(newTableName string) *temp {
	t.tempDo.UseTable(newTableName)
	return t.updateTableName(newTableName)
}

func (t temp) As(alias string) *temp {
	t.tempDo.DO = *(t.tempDo.As(alias).(*gen.DO))
	return t.updateTableName(alias)
}

func (t *temp) updateTableName(table string) *temp {
	t.ALL = field.NewAsterisk(table)
	t.ID = field.NewInt32(table, "id")
	t.Input = field.NewString(table, "input")
	t.A = field.NewFloat64(table, "A")
	t.B = field.NewField(table, "B")
	t.B0 = field.NewField(table, "B0")
	t.B1 = field.NewField(table, "B1")
	t.B2 = field.NewField(table, "B2")
	t.B3 = field.NewField(table, "B3")
	t.B4 = field.NewField(table, "B4")

	t.fillFieldMap()

	return t
}

func (t *temp) WithContext(ctx context.Context) ITempDo { return t.tempDo.WithContext(ctx) }

func (t temp) TableName() string { return t.tempDo.TableName() }

func (t temp) Alias() string { return t.tempDo.Alias() }

func (t temp) Columns(cols ...field.Expr) gen.Columns { return t.tempDo.Columns(cols...) }

func (t *temp) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := t.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (t *temp) fillFieldMap() {
	t.fieldMap = make(map[string]field.Expr, 9)
	t.fieldMap["id"] = t.ID
	t.fieldMap["input"] = t.Input
	t.fieldMap["A"] = t.A
	t.fieldMap["B"] = t.B
	t.fieldMap["B0"] = t.B0
	t.fieldMap["B1"] = t.B1
	t.fieldMap["B2"] = t.B2
	t.fieldMap["B3"] = t.B3
	t.fieldMap["B4"] = t.B4
}

func (t temp) clone(db *gorm.DB) temp {
	t.tempDo.ReplaceConnPool(db.Statement.ConnPool)
	return t
}

func (t temp) replaceDB(db *gorm.DB) temp {
	t.tempDo.ReplaceDB(db)
	return t
}

type tempDo struct{ gen.DO }

type ITempDo interface {
	gen.SubQuery
	Debug() ITempDo
	WithContext(ctx context.Context) ITempDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ITempDo
	WriteDB() ITempDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ITempDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ITempDo
	Not(conds ...gen.Condition) ITempDo
	Or(conds ...gen.Condition) ITempDo
	Select(conds ...field.Expr) ITempDo
	Where(conds ...gen.Condition) ITempDo
	Order(conds ...field.Expr) ITempDo
	Distinct(cols ...field.Expr) ITempDo
	Omit(cols ...field.Expr) ITempDo
	Join(table schema.Tabler, on ...field.Expr) ITempDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ITempDo
	RightJoin(table schema.Tabler, on ...field.Expr) ITempDo
	Group(cols ...field.Expr) ITempDo
	Having(conds ...gen.Condition) ITempDo
	Limit(limit int) ITempDo
	Offset(offset int) ITempDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ITempDo
	Unscoped() ITempDo
	Create(values ...*model.Temp) error
	CreateInBatches(values []*model.Temp, batchSize int) error
	Save(values ...*model.Temp) error
	First() (*model.Temp, error)
	Take() (*model.Temp, error)
	Last() (*model.Temp, error)
	Find() ([]*model.Temp, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Temp, err error)
	FindInBatches(result *[]*model.Temp, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Temp) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ITempDo
	Assign(attrs ...field.AssignExpr) ITempDo
	Joins(fields ...field.RelationField) ITempDo
	Preload(fields ...field.RelationField) ITempDo
	FirstOrInit() (*model.Temp, error)
	FirstOrCreate() (*model.Temp, error)
	FindByPage(offset int, limit int) (result []*model.Temp, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ITempDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (t tempDo) Debug() ITempDo {
	return t.withDO(t.DO.Debug())
}

func (t tempDo) WithContext(ctx context.Context) ITempDo {
	return t.withDO(t.DO.WithContext(ctx))
}

func (t tempDo) ReadDB() ITempDo {
	return t.Clauses(dbresolver.Read)
}

func (t tempDo) WriteDB() ITempDo {
	return t.Clauses(dbresolver.Write)
}

func (t tempDo) Session(config *gorm.Session) ITempDo {
	return t.withDO(t.DO.Session(config))
}

func (t tempDo) Clauses(conds ...clause.Expression) ITempDo {
	return t.withDO(t.DO.Clauses(conds...))
}

func (t tempDo) Returning(value interface{}, columns ...string) ITempDo {
	return t.withDO(t.DO.Returning(value, columns...))
}

func (t tempDo) Not(conds ...gen.Condition) ITempDo {
	return t.withDO(t.DO.Not(conds...))
}

func (t tempDo) Or(conds ...gen.Condition) ITempDo {
	return t.withDO(t.DO.Or(conds...))
}

func (t tempDo) Select(conds ...field.Expr) ITempDo {
	return t.withDO(t.DO.Select(conds...))
}

func (t tempDo) Where(conds ...gen.Condition) ITempDo {
	return t.withDO(t.DO.Where(conds...))
}

func (t tempDo) Order(conds ...field.Expr) ITempDo {
	return t.withDO(t.DO.Order(conds...))
}

func (t tempDo) Distinct(cols ...field.Expr) ITempDo {
	return t.withDO(t.DO.Distinct(cols...))
}

func (t tempDo) Omit(cols ...field.Expr) ITempDo {
	return t.withDO(t.DO.Omit(cols...))
}

func (t tempDo) Join(table schema.Tabler, on ...field.Expr) ITempDo {
	return t.withDO(t.DO.Join(table, on...))
}

func (t tempDo) LeftJoin(table schema.Tabler, on ...field.Expr) ITempDo {
	return t.withDO(t.DO.LeftJoin(table, on...))
}

func (t tempDo) RightJoin(table schema.Tabler, on ...field.Expr) ITempDo {
	return t.withDO(t.DO.RightJoin(table, on...))
}

func (t tempDo) Group(cols ...field.Expr) ITempDo {
	return t.withDO(t.DO.Group(cols...))
}

func (t tempDo) Having(conds ...gen.Condition) ITempDo {
	return t.withDO(t.DO.Having(conds...))
}

func (t tempDo) Limit(limit int) ITempDo {
	return t.withDO(t.DO.Limit(limit))
}

func (t tempDo) Offset(offset int) ITempDo {
	return t.withDO(t.DO.Offset(offset))
}

func (t tempDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ITempDo {
	return t.withDO(t.DO.Scopes(funcs...))
}

func (t tempDo) Unscoped() ITempDo {
	return t.withDO(t.DO.Unscoped())
}

func (t tempDo) Create(values ...*model.Temp) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Create(values)
}

func (t tempDo) CreateInBatches(values []*model.Temp, batchSize int) error {
	return t.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (t tempDo) Save(values ...*model.Temp) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Save(values)
}

func (t tempDo) First() (*model.Temp, error) {
	if result, err := t.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Temp), nil
	}
}

func (t tempDo) Take() (*model.Temp, error) {
	if result, err := t.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Temp), nil
	}
}

func (t tempDo) Last() (*model.Temp, error) {
	if result, err := t.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Temp), nil
	}
}

func (t tempDo) Find() ([]*model.Temp, error) {
	result, err := t.DO.Find()
	return result.([]*model.Temp), err
}

func (t tempDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Temp, err error) {
	buf := make([]*model.Temp, 0, batchSize)
	err = t.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (t tempDo) FindInBatches(result *[]*model.Temp, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return t.DO.FindInBatches(result, batchSize, fc)
}

func (t tempDo) Attrs(attrs ...field.AssignExpr) ITempDo {
	return t.withDO(t.DO.Attrs(attrs...))
}

func (t tempDo) Assign(attrs ...field.AssignExpr) ITempDo {
	return t.withDO(t.DO.Assign(attrs...))
}

func (t tempDo) Joins(fields ...field.RelationField) ITempDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Joins(_f))
	}
	return &t
}

func (t tempDo) Preload(fields ...field.RelationField) ITempDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Preload(_f))
	}
	return &t
}

func (t tempDo) FirstOrInit() (*model.Temp, error) {
	if result, err := t.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Temp), nil
	}
}

func (t tempDo) FirstOrCreate() (*model.Temp, error) {
	if result, err := t.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Temp), nil
	}
}

func (t tempDo) FindByPage(offset int, limit int) (result []*model.Temp, count int64, err error) {
	result, err = t.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = t.Offset(-1).Limit(-1).Count()
	return
}

func (t tempDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = t.Count()
	if err != nil {
		return
	}

	err = t.Offset(offset).Limit(limit).Scan(result)
	return
}

func (t tempDo) Scan(result interface{}) (err error) {
	return t.DO.Scan(result)
}

func (t tempDo) Delete(models ...*model.Temp) (result gen.ResultInfo, err error) {
	return t.DO.Delete(models)
}

func (t *tempDo) withDO(do gen.Dao) *tempDo {
	t.DO = *do.(*gen.DO)
	return t
}

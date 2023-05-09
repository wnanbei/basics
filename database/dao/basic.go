package dao

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type MODEL interface {
	schema.Tabler
}

// IDao 数据访问层
type IDao[M MODEL] interface {
	GetOneByID(id int64) (*M, error)
	GetManyByIDs(ids []int64) ([]*M, error)
	InsertOne(i *M) error
	InsertMany(ips []*M) error
	UpdateByID(id int64, values map[string]any) (int64, error)
	DeleteManyByIDs(ids []int64) (int64, error)
	Find(page, pageSize int) ([]*M, int64, error)

	Clone(db *gorm.DB) *Dao[M]
}

// NewDao 创建数据访问层
func NewDao[M MODEL](ctx context.Context, db *gorm.DB) IDao[M] {
	return &Dao[M]{
		Ctx: ctx,
		DB:  db,
	}
}

type Dao[M MODEL] struct {
	Ctx context.Context
	DB  *gorm.DB
}

// GetOneByID 根据 ID 单条查询
func (d *Dao[M]) GetOneByID(id int64) (*M, error) {
	var data *M

	result := d.DB.Where("id = ?", id).First(&data)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, result.Error
	}

	return data, nil
}

// GetManyByIDs 根据 IDs 批量查询
func (d *Dao[M]) GetManyByIDs(ids []int64) ([]*M, error) {
	var data []*M
	result := d.DB.Model(new(M)).Where("id IN ?", ids).Find(&data)
	return data, result.Error
}

// InsertOne 插入一条数据
func (d *Dao[M]) InsertOne(data *M) error {
	result := d.DB.Create(data)
	return result.Error
}

// InsertMany 插入多条数据
func (d *Dao[M]) InsertMany(data []*M) error {
	result := d.DB.Create(data)
	return result.Error
}

// UpdateByID 根据 ID 更新数据
func (d *Dao[M]) UpdateByID(id int64, values map[string]any) (int64, error) {
	result := d.DB.Model(new(M)).Where("id = ?", id).UpdateColumns(values)
	return result.RowsAffected, result.Error
}

// DeleteManyByIDs 根据 IDs 批量删除
func (d *Dao[M]) DeleteManyByIDs(ids []int64) (int64, error) {
	result := d.DB.Where("id IN ?", ids).Model(new(M)).Delete(new(M))
	return result.RowsAffected, result.Error
}

// Find 查询
func (d *Dao[M]) Find(page, pageSize int) ([]*M, int64, error) {
	data := make([]*M, 0)
	var count int64

	result := d.DB.Model(new(M)).Count(&count).Order("id DESC")
	if result.Error != nil || count == 0 {
		return data, 0, result.Error
	}

	result = d.DB.Offset((page - 1) * pageSize).Limit(pageSize).Find(&data)
	return data, count, result.Error
}

func (d *Dao[M]) Clone(db *gorm.DB) *Dao[M] {
	d.DB = d.DB.Session(&gorm.Session{Initialized: true}).Session(&gorm.Session{})
	d.DB.Statement.ConnPool = db.ConnPool
	return d
}

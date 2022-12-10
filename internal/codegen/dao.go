package codegen

import (
	"context"
	"errors"

	"github.com/go-bamboo/pkg/store/gormx"
)

// Dao dao interface
type Dao interface {
	Close()
	Ping(ctx context.Context) (err error)

	GetDBTablesPage(dbname string, e *DBTables, pageSize int, pageIndex int) ([]DBTables, int64, error)
	GetDBColumnsList(dbname string, e *DBColumns) ([]DBColumns, error)
}

// dao dao.
type dao struct {
	orm *gormx.DB
}

// New new a dao and return.
func New() (d Dao, err error) {
	return newDao()
}

func newDao() (d *dao, err error) {
	db := gormx.MustNew(&gormx.Conf{
		Driver:   "mysql",
		Source:   dsn,
		LogLevel: 3,
	})
	d = &dao{
		orm: db,
	}
	return
}

// Close close the resource.
func (d *dao) Close() {
}

// Ping ping the resource.
func (d *dao) Ping(ctx context.Context) (err error) {
	return nil
}

func (d *dao) GetDBColumnsList(dbname string, e *DBColumns) (doc []DBColumns, err error) {
	if e.TableName == "" {
		return nil, errors.New("table name cannot be empty！")
	}
	table := d.orm.Table("information_schema.columns")
	table = table.Where("table_schema= ? ", dbname)
	table = table.Where("TABLE_NAME = ?", e.TableName)

	if err := table.Find(&doc).Error; err != nil {
		return doc, err
	}
	return doc, nil
}

func (d *dao) GetDBTablesPage(dbname string, e *DBTables, pageSize int, pageIndex int) (doc []DBTables, count int64, err error) {

	table := d.orm.Table("tables")
	// table = table.Where("TABLE_NAME not in (select table_name from " + dbname + ".sys_tables) ")
	table = table.Where("table_schema= ? ", dbname)

	if e.TableName != "" {
		table = table.Where("TABLE_NAME = ?", e.TableName)
	}
	if err = table.Count(&count).Error; err != nil {
		return
	}
	if err = table.Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&doc).Error; err != nil {
		return
	}
	return
}

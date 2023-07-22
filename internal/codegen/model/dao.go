package model

import (
	"context"
	"errors"
	"strings"

	"github.com/go-bamboo/pkg/store/gormx"
)

// dao dao.
type Dao struct {
	orm *gormx.DB
}

func New(dsn string) (d *Dao, err error) {
	db := gormx.MustNew(&gormx.Conf{
		Driver: gormx.DBType_mysql,
		Source: dsn,
	})
	d = &Dao{
		orm: db,
	}
	return
}

func (d *Dao) GetDBColumnsList(dbname string, e *DBColumns) (doc []DBColumns, err error) {
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

func (d *Dao) GetDBTablesPage(dbname string, e *DBTables, pageSize int, pageIndex int) (doc []DBTables, count int64, err error) {
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

// GetDBTableList 根据数据库，表面获取列表
func (uc *Dao) getDBTableList(ctx context.Context, dbname, tbname string, pageSize, pageIndex int) (page []DBTables, total int64, err error) {
	var filter = DBTables{
		TableName: tbname,
	}
	page, total, err = uc.GetDBTablesPage(dbname, &filter, pageSize, pageIndex)
	if err != nil {
		return
	}
	return
}

func (uc *Dao) initTable(ctx context.Context, dbname string, dbtable *DBTables) (reply *SysTables, err error) {

	var data SysTables
	data.TBName = dbtable.TableName
	// data.CreateBy = fmt.Sprint(dp.UserId)
	tablenamelist := strings.Split(dbtable.TableName, "_")
	for i := 0; i < len(tablenamelist); i++ {
		strStart := string([]byte(tablenamelist[i])[:1])
		strend := string([]byte(tablenamelist[i])[1:])
		data.ClassName += strings.ToUpper(strStart) + strend
		data.PackageName += strings.ToLower(strStart) + strings.ToLower(strend)
		data.ModuleName += strings.ToLower(strStart) + strings.ToLower(strend)
	}
	data.TplCategory = "crud"
	data.Crud = true

	// data.CreateBy = fmt.Sprint(dp.UserId)
	data.TableComment = dbtable.TableComment
	if dbtable.TableComment == "" {
		data.TableComment = data.ClassName
	}

	data.FunctionName = data.TableComment
	data.BusinessName = data.ModuleName
	data.IsLogicalDelete = "1"
	data.LogicalDelete = true
	data.LogicalDeleteColumn = "is_del"
	data.FunctionAuthor = "bing"

	// 获取列信息
	data.Columns = make([]SysColumns, 0)
	dbcolumns, err := uc.GetDBColumnsList(dbname, &DBColumns{TableName: dbtable.TableName})
	for i := 0; i < len(dbcolumns); i++ {
		dbcolumn := dbcolumns[i]
		var column SysColumns
		column.ColumnComment = dbcolumn.ColumnComment
		column.ColumnName = dbcolumn.ColumnName
		column.ColumnType = dbcolumn.ColumnType
		column.Sort = int32(i + 1)
		column.Insert = true
		column.IsInsert = "1"
		column.QueryType = "EQ"
		column.IsPk = "0"

		namelist := strings.Split(dbcolumn.ColumnName, "_")
		for i := 0; i < len(namelist); i++ {
			strStart := string([]byte(namelist[i])[:1])
			strend := string([]byte(namelist[i])[1:])
			column.GoField += strings.ToUpper(strStart) + strend
			if i == 0 {
				column.JsonField = strings.ToLower(strStart) + strend
			} else {
				column.JsonField += strings.ToUpper(strStart) + strend
			}
		}
		if strings.Contains(dbcolumn.ColumnKey, "PR") {
			column.IsPk = "1"
			column.Pk = true
			data.PkColumn = dbcolumn.ColumnName
			data.PkGoField = column.GoField
			data.PkJsonField = column.JsonField
		}
		column.IsRequired = "0"
		if strings.Contains(dbcolumn.IsNullable, "NO") {
			column.IsRequired = "1"
			column.Required = true
		}

		// 类型系统转换
		// log.Printf("ctype: %v", dbcolumn.ColumnType)
		if strings.Compare(dbcolumn.ColumnType, "bigint unsigned") == 0 {
			column.GoType = "uint64"
			column.HtmlType = "input"
		} else if strings.Compare(dbcolumn.ColumnType, "bigint") == 0 {
			column.GoType = "int64"
			column.HtmlType = "input"
		} else if strings.Compare(dbcolumn.ColumnType, "int unsigned") == 0 {
			column.GoType = "uint32"
			column.HtmlType = "input"
		} else if strings.Compare(dbcolumn.ColumnType, "int") == 0 {
			column.GoType = "int32"
			column.HtmlType = "input"
		} else if strings.Contains(dbcolumn.ColumnType, "timestamp") {
			column.GoType = "time.Time"
			column.HtmlType = "datetime"
		} else {
			column.GoType = "string"
			column.HtmlType = "input"
		}

		data.Columns = append(data.Columns, column)
	}
	reply = &data
	return
}

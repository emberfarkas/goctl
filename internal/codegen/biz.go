package codegen

import (
	"bytes"
	"context"
	"path"
	"strings"
	"text/template"

	xtem "github.com/emberfarkas/goctl/internal/codegen/template"
	"github.com/go-bamboo/pkg/filex"
	"github.com/go-bamboo/pkg/log"
)

type biz struct {
	d Dao
}

func NewBiz(d Dao) *biz {
	return &biz{
		d: d,
	}
}

// GetDBTableList 根据数据库，表面获取列表
func (uc *biz) getDBTableList(ctx context.Context, dbname, tbname string, pageSize, pageIndex int) (page []DBTables, total int64, err error) {
	var filter DBTables
	filter.TableName = tbname
	page, total, err = uc.d.GetDBTablesPage(dbname, &filter, pageSize, pageIndex)
	if err != nil {
		return
	}
	return
}

func (uc *biz) initTable(ctx context.Context, dbname string, dbtable *DBTables) (reply *SysTables, err error) {

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
	dbcolumns, err := uc.d.GetDBColumnsList(dbname, &DBColumns{TableName: dbtable.TableName})
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

// GenCode 生成代码工具
func (uc *biz) genCode(ctx context.Context, tab *SysTables) (err error) {
	if err = uc.genCodeModelGo(ctx, tab); err != nil {
		log.Fatalf("codegen model: %v", err)
		return
	}
	if err = uc.genCodeDaoGo(ctx, tab); err != nil {
		log.Fatalf("codegen dao: %v", err)
		return
	}
	if err = uc.genCodeApiJs(ctx, tab); err != nil {
		log.Fatalf("codegen api: %v", err)
		return
	}
	if err = uc.genCodeViewVue(ctx, tab); err != nil {
		log.Fatalf("codegen view: %v", err)
		return
	}
	return
}

func (uc *biz) genCodeModelGo(ctx context.Context, tab *SysTables) (err error) {
	tpl := "model.go.tpl"
	tt1, err := xtem.ReadFile(tpl)
	if err != nil {
		log.Fatal(err)
		return
	}
	t1 := template.New(tpl)
	t1, err = t1.Parse(string(tt1))
	if err != nil {
		log.Fatal(err)
		return
	}
	var b1 bytes.Buffer
	if err = t1.Execute(&b1, tab); err != nil {
		log.Fatalf("err: %v", err)
		return
	}
	var prefixPath string
	if path.IsAbs(backpath) {
		prefixPath = path.Join(backpath, tab.Module, "internal/model")
	} else {
		wd := filex.GetCurrentPath()
		prefixPath = path.Join(wd, backpath, tab.Module, "internal/model")
	}
	if err = filex.CreateFile(path.Join(prefixPath, tab.PackageName+".go")); err != nil {
		log.Fatal(err)
		return
	}
	return
}

func (uc *biz) genCodeDaoGo(ctx context.Context, tab *SysTables) (err error) {
	tpl := "dao.go.tpl"
	tt2, err := xtem.ReadFile(tpl)
	if err != nil {
		return
	}
	t2 := template.New(tpl)
	t2, err = t2.Parse(string(tt2))
	if err != nil {
		return
	}
	var b2 bytes.Buffer
	if err = t2.Execute(&b2, tab); err != nil {
		log.Fatalf("err: %v", err)
		return
	}

	var prefixPath string
	if path.IsAbs(backpath) {
		prefixPath = path.Join(backpath, tab.Module, "internal/dao")
	} else {
		wd := filex.GetCurrentPath()
		prefixPath = path.Join(wd, backpath, tab.Module, "internal/dao")
	}
	if err = filex.CreateFile(path.Join(prefixPath, tab.PackageName+".go")); err != nil {
		return
	}
	return
}

func (uc *biz) genCodeApiJs(ctx context.Context, tab *SysTables) (err error) {
	tpl := "api.js.go.tpl"
	tt3, err := xtem.ReadFile(tpl)
	if err != nil {
		return
	}
	t3 := template.New(tpl)
	t3, err = t3.Parse(string(tt3))
	if err != nil {
		return
	}
	var b3 bytes.Buffer
	if err = t3.Execute(&b3, tab); err != nil {
		log.Fatalf("err: %v", err)
		return
	}

	var prefixPath string
	if path.IsAbs(backpath) {
		prefixPath = path.Join(adminpath, "api", tab.Module)
	} else {
		wd := filex.GetCurrentPath()
		prefixPath = path.Join(wd, adminpath, "api", tab.Module)
	}
	if err = filex.CreateFile(path.Join(prefixPath, tab.PackageName+".js")); err != nil {
		return
	}
	return
}

func (uc *biz) genCodeViewVue(ctx context.Context, tab *SysTables) (err error) {
	tpl := "view.vue.go.tpl"
	tt4, err := xtem.ReadFile(tpl)
	if err != nil {
		return
	}
	t4 := template.New(tpl)
	t4, err = t4.Parse(string(tt4))
	if err != nil {
		return
	}
	var b4 bytes.Buffer
	if err = t4.Execute(&b4, tab); err != nil {
		log.Fatalf("err: %v", err)
		return
	}

	var prefixPath string
	if path.IsAbs(backpath) {
		prefixPath = path.Join(adminpath, "views", tab.Module, tab.PackageName)
	} else {
		wd := filex.GetCurrentPath()
		prefixPath = path.Join(wd, adminpath, "views", tab.Module, tab.PackageName)
	}
	if err = filex.CreateFile(path.Join(prefixPath, "index.vue")); err != nil {
		return
	}
	return
}

package codegen

import (
	"bytes"
	"context"
	"log"
	"path"
	"strings"
	"text/template"

	xtem "github.com/emberfarkas/goctl/internal/codegen/template"
	"github.com/go-bamboo/pkg/filex"
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
	if err = filex.PathCreate(prefixPath); err != nil {
		log.Fatal(err)
		return
	}
	if err = filex.FileCreate(b1, path.Join(prefixPath, tab.PackageName+".go")); err != nil {
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
	if err = filex.PathCreate(prefixPath); err != nil {
		return
	}
	if err = filex.FileCreate(b2, path.Join(prefixPath, tab.PackageName+".go")); err != nil {
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
	if err = filex.PathCreate(prefixPath); err != nil {
		return
	}
	if err = filex.FileCreate(b3, path.Join(prefixPath, tab.PackageName+".js")); err != nil {
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
	if err = filex.PathCreate(prefixPath); err != nil {
		return
	}
	if err = filex.FileCreate(b4, path.Join(prefixPath, "index.vue")); err != nil {
		return
	}
	return
}

// func (uc *biz) GenMenuAndApi(ctx context.Context, req *pb.GenMenuAndApiRequest) (reply *pb.GenMenuAndApiReply, err error) {

// 	table := model.SysTables{}
// 	timeNow := time.Now()
// 	table.TableId = (req.TableId)
// 	tab, _ := uc.d.GetSysTables(&table)
// 	Mmenu := model.Menu{}
// 	Mmenu.MenuName = tab.TBName + "管理"
// 	Mmenu.Title = tab.TableComment
// 	Mmenu.Icon = "pass"
// 	Mmenu.Path = "/" + tab.TBName
// 	Mmenu.MenuType = "M"
// 	Mmenu.Action = "无"
// 	Mmenu.ParentId = 0
// 	Mmenu.NoCache = false
// 	Mmenu.Component = "Layout"
// 	Mmenu.Sort = 0
// 	Mmenu.Visible = true
// 	Mmenu.IsFrame = false
// 	Mmenu.CreateBy = "1"
// 	Mmenu.UpdateBy = "1"
// 	Mmenu.CreatedAt = timeNow
// 	Mmenu.UpdatedAt = timeNow
// 	Mmenu.MenuId, err = uc.d.CreateMenu(&Mmenu)

// 	Cmenu := model.Menu{}
// 	Cmenu.MenuName = tab.TBName + "管理"
// 	Cmenu.Title = tab.TableComment
// 	Cmenu.Icon = "pass"
// 	Cmenu.Path = tab.TBName
// 	Cmenu.MenuType = "C"
// 	Cmenu.Action = "无"
// 	Cmenu.Permission = tab.PackageName + ":" + tab.ModuleName + ":list"
// 	Cmenu.ParentId = Mmenu.MenuId
// 	Cmenu.NoCache = false
// 	Cmenu.Component = "/" + tab.ModuleName + "/index"
// 	Cmenu.Sort = 0
// 	Cmenu.Visible = true
// 	Cmenu.IsFrame = false
// 	Cmenu.CreateBy = "1"
// 	Cmenu.UpdateBy = "1"
// 	Cmenu.CreatedAt = timeNow
// 	Cmenu.UpdatedAt = timeNow
// 	Cmenu.MenuId, err = uc.d.CreateMenu(&Cmenu)

// 	MList := model.Menu{}
// 	MList.MenuName = tab.TBName
// 	MList.Title = "分页获取" + tab.TableComment
// 	MList.Icon = "pass"
// 	MList.Path = tab.TBName
// 	MList.MenuType = "F"
// 	MList.Action = "无"
// 	MList.Permission = tab.PackageName + ":" + tab.ModuleName + ":query"
// 	MList.ParentId = Cmenu.MenuId
// 	MList.NoCache = false
// 	MList.Sort = 0
// 	MList.Visible = true
// 	MList.IsFrame = false
// 	MList.CreateBy = "1"
// 	MList.UpdateBy = "1"
// 	MList.CreatedAt = timeNow
// 	MList.UpdatedAt = timeNow
// 	MList.MenuId, err = uc.d.CreateMenu(&MList)

// 	MCreate := model.Menu{}
// 	MCreate.MenuName = tab.TBName
// 	MCreate.Title = "创建" + tab.TableComment
// 	MCreate.Icon = "pass"
// 	MCreate.Path = tab.TBName
// 	MCreate.MenuType = "F"
// 	MCreate.Action = "无"
// 	MCreate.Permission = tab.PackageName + ":" + tab.ModuleName + ":add"
// 	MCreate.ParentId = Cmenu.MenuId
// 	MCreate.NoCache = false
// 	MCreate.Sort = 0
// 	MCreate.Visible = true
// 	MCreate.IsFrame = false
// 	MCreate.CreateBy = "1"
// 	MCreate.UpdateBy = "1"
// 	MCreate.CreatedAt = timeNow
// 	MCreate.UpdatedAt = timeNow
// 	MCreate.MenuId, err = uc.d.CreateMenu(&MCreate)

// 	MUpdate := model.Menu{}
// 	MUpdate.MenuName = tab.TBName
// 	MUpdate.Title = "修改" + tab.TableComment
// 	MUpdate.Icon = "pass"
// 	MUpdate.Path = tab.TBName
// 	MUpdate.MenuType = "F"
// 	MUpdate.Action = "无"
// 	MUpdate.Permission = tab.PackageName + ":" + tab.ModuleName + ":edit"
// 	MUpdate.ParentId = Cmenu.MenuId
// 	MUpdate.NoCache = false
// 	MUpdate.Sort = 0
// 	MUpdate.Visible = true
// 	MUpdate.IsFrame = false
// 	MUpdate.CreateBy = "1"
// 	MUpdate.UpdateBy = "1"
// 	MUpdate.CreatedAt = timeNow
// 	MUpdate.UpdatedAt = timeNow
// 	MUpdate.MenuId, err = uc.d.CreateMenu(&MUpdate)

// 	MDelete := model.Menu{}
// 	MDelete.MenuName = tab.TBName
// 	MDelete.Title = "删除" + tab.TableComment
// 	MDelete.Icon = "pass"
// 	MDelete.Path = tab.TBName
// 	MDelete.MenuType = "F"
// 	MDelete.Action = "无"
// 	MDelete.Permission = tab.PackageName + ":" + tab.ModuleName + ":remove"
// 	MDelete.ParentId = Cmenu.MenuId
// 	MDelete.NoCache = false
// 	MDelete.Sort = 0
// 	MDelete.Visible = true
// 	MDelete.IsFrame = false
// 	MDelete.CreateBy = "1"
// 	MDelete.UpdateBy = "1"
// 	MDelete.CreatedAt = timeNow
// 	MDelete.UpdatedAt = timeNow
// 	MDelete.MenuId, err = uc.d.CreateMenu(&MDelete)

// 	var InterfaceId = int32(63)
// 	Amenu := model.Menu{}
// 	Amenu.MenuName = tab.TBName
// 	Amenu.Title = tab.TableComment
// 	Amenu.Icon = "bug"
// 	Amenu.Path = tab.TBName
// 	Amenu.MenuType = "M"
// 	Amenu.Action = "无"
// 	Amenu.ParentId = InterfaceId
// 	Amenu.NoCache = false
// 	Amenu.Sort = 0
// 	Amenu.Visible = true
// 	Amenu.IsFrame = false
// 	Amenu.CreateBy = "1"
// 	Amenu.UpdateBy = "1"
// 	Amenu.CreatedAt = timeNow
// 	Amenu.UpdatedAt = timeNow
// 	Amenu.MenuId, err = uc.d.CreateMenu(&Amenu)

// 	AList := model.Menu{}
// 	AList.MenuName = tab.TBName
// 	AList.Title = "分页获取" + tab.TableComment
// 	AList.Icon = "bug"
// 	AList.Path = "/api/v1/" + tab.ModuleName + "List"
// 	AList.MenuType = "A"
// 	AList.Action = "GET"
// 	AList.ParentId = Amenu.MenuId
// 	AList.NoCache = false
// 	AList.Sort = 0
// 	AList.Visible = true
// 	AList.IsFrame = false
// 	AList.CreateBy = "1"
// 	AList.UpdateBy = "1"
// 	AList.CreatedAt = timeNow
// 	AList.UpdatedAt = timeNow
// 	AList.MenuId, err = uc.d.CreateMenu(&AList)

// 	AGet := model.Menu{}
// 	AGet.MenuName = tab.TBName
// 	AGet.Title = "根据id获取" + tab.TableComment
// 	AGet.Icon = "bug"
// 	AGet.Path = "/api/v1/" + tab.ModuleName + "/:id"
// 	AGet.MenuType = "A"
// 	AGet.Action = "GET"
// 	AGet.ParentId = Amenu.MenuId
// 	AGet.NoCache = false
// 	AGet.Sort = 0
// 	AGet.Visible = true
// 	AGet.IsFrame = false
// 	AGet.CreateBy = "1"
// 	AGet.UpdateBy = "1"
// 	AGet.CreatedAt = timeNow
// 	AGet.UpdatedAt = timeNow
// 	AGet.MenuId, err = uc.d.CreateMenu(&AGet)

// 	ACreate := model.Menu{}
// 	ACreate.MenuName = tab.TBName
// 	ACreate.Title = "创建" + tab.TableComment
// 	ACreate.Icon = "bug"
// 	ACreate.Path = "/api/v1/" + tab.ModuleName
// 	ACreate.MenuType = "A"
// 	ACreate.Action = "POST"
// 	ACreate.ParentId = Amenu.MenuId
// 	ACreate.NoCache = false
// 	ACreate.Sort = 0
// 	ACreate.Visible = true
// 	ACreate.IsFrame = false
// 	ACreate.CreateBy = "1"
// 	ACreate.UpdateBy = "1"
// 	ACreate.CreatedAt = timeNow
// 	ACreate.UpdatedAt = timeNow
// 	ACreate.MenuId, err = uc.d.CreateMenu(&ACreate)

// 	AUpdate := model.Menu{}
// 	AUpdate.MenuName = tab.TBName
// 	AUpdate.Title = "修改" + tab.TableComment
// 	AUpdate.Icon = "bug"
// 	AUpdate.Path = "/api/v1/" + tab.ModuleName
// 	AUpdate.MenuType = "A"
// 	AUpdate.Action = "PUT"
// 	AUpdate.ParentId = Amenu.MenuId
// 	AUpdate.NoCache = false
// 	AUpdate.Sort = 0
// 	AUpdate.Visible = true
// 	AUpdate.IsFrame = false
// 	AUpdate.CreateBy = "1"
// 	AUpdate.UpdateBy = "1"
// 	AUpdate.CreatedAt = timeNow
// 	AUpdate.UpdatedAt = timeNow
// 	AUpdate.MenuId, err = uc.d.CreateMenu(&AUpdate)

// 	ADelete := model.Menu{}
// 	ADelete.MenuName = tab.TBName
// 	ADelete.Title = "删除" + tab.TableComment
// 	ADelete.Icon = "bug"
// 	ADelete.Path = "/api/v1/" + tab.ModuleName + "/:id"
// 	ADelete.MenuType = "A"
// 	ADelete.Action = "DELETE"
// 	ADelete.ParentId = Amenu.MenuId
// 	ADelete.NoCache = false
// 	ADelete.Sort = 0
// 	ADelete.Visible = true
// 	ADelete.IsFrame = false
// 	ADelete.CreateBy = "1"
// 	ADelete.UpdateBy = "1"
// 	ADelete.CreatedAt = timeNow
// 	ADelete.UpdatedAt = timeNow
// 	ADelete.MenuId, err = uc.d.CreateMenu(&ADelete)
// 	if err != nil {
// 		return
// 	}
// 	return
// }

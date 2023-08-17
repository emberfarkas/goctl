package template_test

import (
	"testing"
)

func TestGoModelTemplate(t *testing.T) {
	//b, err := ReadFile("view.vue.tpl")
	//if err != nil {
	//	t.Errorf("%v", err)
	//	return
	//}
	//t1 := template.New(string(b))

	//table := model.SysTables{
	//	TBName:      "hello",
	//	PackageName: "hello",
	//}
	//
	//file, err := os.Create("models/" + table.PackageName + ".go")
	//if err != nil {
	//	t.Errorf("err1: %v", err)
	//	return
	//}
	//defer file.Close()

	//if err := t1.Execute(file, table); err != nil {
	//	t.Errorf("err : %v", err)
	//}
}

// func TestGoApiTemplate(t *testing.T) {
// 	c := config.New(config.WithSource(
// 		file.NewSource("../configs/admin.yaml"),
// 	))
// 	if err := c.Load(); err != nil {
// 		panic(err)
// 	}

// 	logger := log.NewStdLogger(os.Stdout)
// 	var bc conf.Bootstrap
// 	if err := c.Scan(&bc); err != nil {
// 		panic(err)
// 	}

// 	d, err := dao.New(bc.Data, logger)
// 	if err != nil {
// 		t.Errorf("err = %v", err)
// 		return
// 	}
// 	defer d.Close()
// 	t1, err := template.ParseFiles("api.go.template")
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	table := model.SysTables{}
// 	tab, _ := d.GetSysTables(&table)
// 	file, err := os.Create("apis/" + table.PackageName + ".go")
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	defer file.Close()

// 	_ = t1.Execute(file, tab)
// 	t.Log("")
// }

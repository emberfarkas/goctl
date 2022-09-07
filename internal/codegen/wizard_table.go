package codegen

import (
	"context"
	"errors"
	"fmt"
)

// deployPage creates a new node configuration based on some user input.
func (w *wizard) deployPage(ctx context.Context) error {

	dbname := w.read("dbname:")

	// 获取表的信息
	list, total, err := w.uc.getDBTableList(ctx, dbname, "", 10, 1)
	if err != nil {
		return err
	}
	if total <= 0 {
		return errors.New("没有表")
	}
	fmt.Printf("开始查看表:%v\n", dbname)

	// 显示所有
	for i := 0; i < len(list); i++ {
		it := list[i]
		fmt.Println()
		choice := w.read(fmt.Sprintf("gen table %v (yes or no):", it.TableName))
		switch choice {
		case "yes":
			// 获取表的信息
			data, err := w.uc.initTable(ctx, dbname, &it)
			if err != nil {
				return err
			}
			// log.Printf("[%v]", data)
			module := w.read("service name:")
			data.Module = module

			prompt := `What would you like to gen template code? (default = all)
	1. model
	2. dao
	3. js
	4. view
	5. all
`
			fmt.Println("")
			fmt.Print(prompt)
			fmt.Println("")

			choice := w.read("choice:")
			switch {
			case choice == "1":
				if err := w.uc.genCodeModelGo(ctx, data); err != nil {
					fmt.Printf("err : %v", err)
				}
			case choice == "2":
				if err := w.uc.genCodeModelGo(ctx, data); err != nil {
					fmt.Printf("err : %v", err)
				}
			case choice == "3":
				if err := w.uc.genCodeApiJs(ctx, data); err != nil {
					fmt.Printf("err : %v", err)
				}
			case choice == "4":
				if err := w.uc.genCodeViewVue(ctx, data); err != nil {
					fmt.Printf("err : %v", err)
				}
			default:
				// 生成代码
				if err = w.uc.genCode(ctx, data); err != nil {
					return err
				}
			}
		default:
		}
	}
	return nil
}

// deployMenu 部署menu
func (w *wizard) deployMenu(ctx context.Context) error {
	// w.biz.GenCode(context.Background())
	// 获取表的信息
	// data, err := w.uc.initTable(ctx, dbname, tbname)
	// if err != nil {
	// 	return
	// }

	// 生成菜单
	// if err = uc.genCode(ctx, data); err != nil {
	// 	return
	// }
	return nil
}

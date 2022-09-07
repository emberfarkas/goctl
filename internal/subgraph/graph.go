package subgraph

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/machinebox/graphql"
	"github.com/spf13/cobra"
	"github.com/xuri/excelize/v2"
)

// Cmd represents the config command
var (
	Cmd = &cobra.Command{
		Use:   "stat",
		Short: "统计辅助工具",
		Long:  `一些批处理proxy的工具`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return run(cmd.Context())
		},
	}
	contract string // 合约
	out      string // 输出
)

const FreeToken string = "freetoken"
const STT = "stt"

func init() {

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// configCmd.PersistentFlags().String("foo", "", "A help for foo")

	Cmd.Flags().StringVarP(&contract, "token", "t", "freetoken", "freetoke, stt")
	Cmd.Flags().StringVarP(&out, "out", "o", "./Book1.xlsx", "freetoke, stt")
}

func run(ctx context.Context) error {
	f := excelize.NewFile()
	// Create a new sheet.
	index := f.NewSheet("Sheet2")
	// Set value of a cell.
	f.SetCellValue("Sheet2", "A1", "账号")
	f.SetCellValue("Sheet2", "B1", "余额")
	f.SetCellValue("Sheet2", "C1", "余额浮点数")

	// Set active sheet of the workbook.
	f.SetActiveSheet(index)

	for i := int64(0); ; i = i + 1000 {
		if err := getData(ctx, i, f); err != nil {
			log.Println(err)
			break
		}
	}

	// Save spreadsheet by the given path.
	if err := f.SaveAs(out); err != nil {
		return err
	}
	return nil
}

func getClient(ctx context.Context) (*graphql.Client, error) {
	if contract == FreeToken {
		return graphql.NewClient("http://54.179.35.101:8000/subgraphs/name/dccswap/freetoken"), nil
	} else if contract == STT {
		return graphql.NewClient("http://54.179.35.101:8000/subgraphs/name/dccswap/stt"), nil
	}
	return nil, errors.New("not support")
}

func getData(ctx context.Context, id int64, f *excelize.File) (err error) {
	client, err := getClient(ctx)
	if err != nil {
		return
	}
	// make a request
	greq := graphql.NewRequest(`
		query ($key: Int!) {
			users (first:1000,skip:$key,where: {balance_gt:0}) {
				id
				balance
				balancef
			}
		}
	`)
	greq.Var("key", id)

	// set header fields
	greq.Header.Set("Cache-Control", "no-cache")

	// run it and capture the response
	var respData struct {
		Users []User
	}
	{
	}

	if err = client.Run(ctx, greq, &respData); err != nil {
		return err
	}
	offset := id + 2
	for i := 0; i < len(respData.Users); i++ {
		u := respData.Users[i]
		log.Println("%v", u.Id)
		f.SetCellValue("Sheet2", fmt.Sprintf("A%d", offset), u.Id)
		f.SetCellValue("Sheet2", fmt.Sprintf("B%d", offset), u.Balance)
		f.SetCellValue("Sheet2", fmt.Sprintf("C%d", offset), u.Balancef)
		offset++
	}
	if len(respData.Users) <= 0 {
		return errors.New("users is empty")
	}
	return nil
}

package dao

import (
	"bytes"
	"context"
	"path"
	"text/template"

	"github.com/emberfarkas/goctl/internal/codegen/model"
	xtem "github.com/emberfarkas/goctl/internal/codegen/template"
	"github.com/go-bamboo/pkg/filex"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "dao",
	Short: "生dao模版",
	Long:  `生dao模版文件`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return genCodeDaoGo(cmd.Context(), nil)
	},
}
var backpath string

func init() {
	Cmd.Flags().StringVar(&backpath, "path", "app/service", "后端路径")
}

func genCodeDaoGo(ctx context.Context, tab *model.SysTables) (err error) {
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
		return
	}
	if err = filex.CreateFile(&b2, path.Join(tab.PackageName+".go")); err != nil {
		return
	}
	return
}

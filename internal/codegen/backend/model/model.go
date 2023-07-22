package model

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
	Use:   "model",
	Short: "生model模版",
	Long:  `生model模版文件`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return genCodeModelGo(cmd.Context(), nil)
	},
}
var backpath string

func init() {
	Cmd.Flags().StringVar(&backpath, "backpath", "app/service", "后端路径")
}

func genCodeModelGo(ctx context.Context, tab *model.SysTables) (err error) {
	tpl := "model.go.tpl"
	tt1, err := xtem.ReadFile(tpl)
	if err != nil {
		return
	}
	t1 := template.New(tpl)
	t1, err = t1.Parse(string(tt1))
	if err != nil {
		return
	}
	var b1 bytes.Buffer
	if err = t1.Execute(&b1, tab); err != nil {
		return
	}
	var prefixPath string
	if path.IsAbs(backpath) {
		prefixPath = path.Join(backpath, tab.Module, "internal/model")
	} else {
		wd := filex.GetCurrentPath()
		prefixPath = path.Join(wd, backpath, tab.Module, "internal/model")
	}
	if err = filex.CreateFile(&b1, path.Join(prefixPath, tab.PackageName+".go")); err != nil {
		return
	}
	return
}

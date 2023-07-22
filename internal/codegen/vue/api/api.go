package api

import (
	"bytes"
	"context"
	"github.com/emberfarkas/goctl/internal/codegen/model"
	xtem "github.com/emberfarkas/goctl/internal/codegen/template"
	"github.com/go-bamboo/pkg/filex"
	"github.com/spf13/cobra"
	"path"
	"text/template"
)

var Cmd = &cobra.Command{
	Use:   "api",
	Short: "生api模版",
	Long:  `生api模版文件`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return genCodeApiJs(cmd.Context(), nil)
	},
}
var out string

func init() {
	Cmd.Flags().StringVar(&out, "out", "app/service", "后端路径")
}

func genCodeApiJs(ctx context.Context, tab *model.SysTables) (err error) {
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
		return
	}
	if err = filex.CreateFile(&b3, path.Join(tab.PackageName+".js")); err != nil {
		return
	}
	return
}

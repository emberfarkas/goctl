package view

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
	Use:   "view",
	Short: "生view模版",
	Long:  `生view模版文件`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return genCodeViewVue(cmd.Context(), nil)
	},
}
var out string

func init() {
	Cmd.Flags().StringVar(&out, "out", "app/service", "后端路径")
}

func genCodeViewVue(ctx context.Context, tab *model.SysTables) (err error) {
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
		return
	}
	if err = filex.CreateFile(&b4, path.Join("index.vue")); err != nil {
		return
	}
	return
}

package pdfcov

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
)

// Cmd represents the config command
var Cmd = &cobra.Command{
	Use:   "pdfcov",
	Short: "pdfcov相关辅助工具",
	Long:  `一些批处理pdfcov的工具`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return run(cmd.Context())
	},
}
var src string
var out string

func init() {

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// configCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// configCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	Cmd.Flags().StringVar(&src, "src", "example.html", "html原文件")
	Cmd.Flags().StringVar(&src, "out", "example.pdf", "pdf输出文件")
}

func run(ctx context.Context) error {
	//html template data
	templateData := struct {
		Title       string
		Description string
		Company     string
		Contact     string
		Country     string
	}{
		Title:       "HTML to PDF generator",
		Description: "This is the simple HTML to PDF file.",
		Company:     "Jhon Lewis",
		Contact:     "Maria Anders",
		Country:     "Germany",
	}

	r := NewRequestPdf("")
	if err := r.ParseTemplate(src, templateData); err == nil {
		ok, _ := r.GeneratePDF(out)
		fmt.Println(ok, "pdf generated successfully")
	} else {
		fmt.Println(err)
	}
	return nil
}

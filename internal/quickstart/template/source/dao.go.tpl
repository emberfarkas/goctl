package dao

import (
    "context"
    {{$ii := 1}}
{{ range .Columns }}
    {{- if eq .GoType "time.Time" -}}{{- if eq $ii 1 -}} "time" {{$ii = 2}}{{- end}}{{- end}}
{{- end }}

    "bls/pkg/tools"
    "bls/pkg/ecode"
    "bls/service/{{.Module}}/internal/model"
)

//Page{{.ClassName}} 获取{{.ClassName}}带分页
func (d *dao) Page{{.ClassName}}(ctx context.Context, filter *model.{{.ClassName}},pageSize int, pageIndex int) (docs []model.{{.ClassName}}, total int64, err error) {

    table := d.orm.WithContext(ctx).Table(filter.TableName())
{{ range .Columns }}
{{- if .IsQuery }}
    if e.{{.GoField}} != {{if eq .GoType "string" -}} "" {{ else if eq .GoType "int" -}} 0 {{- end}} {
        table = table.Where("{{.ColumnName}}{{if eq .QueryType "EQ"}} = {{else if eq .QueryType "NE"}} != {{else if eq .QueryType "GT"}} >  {{else if eq .QueryType "GTE"}} >=  {{else if eq .QueryType "LT"}} < {{else if eq .QueryType "LTE"}} <= {{else if eq .QueryType "LIKE"}} like {{end}}?", {{ if eq .QueryType "LIKE"}}"%"+e.{{.GoField}}+"%"{{else}}e.{{.GoField}}{{end}})
    }
{{ end -}}
{{- end }}

    if err = table.Where("`deleted_at` IS NULL").Count(&total).Error; err != nil {
        return
    }
    if err = table.Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&docs).Error; err != nil {
        return
    }
    return
}

//BatchGet{{.ClassName}} 获取{{.ClassName}}列表
func (d *dao) BatchGet{{.ClassName}}(ctx context.Context, e *model.{{.ClassName}}) (docs []model.{{.ClassName}}, err error) {

    table := d.orm.WithContext(ctx).Table(e.TableName())
{{ range .Columns }}
{{- if .IsQuery }}
    if e.{{.GoField}} != {{if eq .GoType "string" -}} "" {{ else if eq .GoType "int" -}} 0 {{- end}} {
        table = table.Where("{{.ColumnName}}{{if eq .QueryType "EQ"}} = {{else if eq .QueryType "NE"}} != {{else if eq .QueryType "GT"}} >  {{else if eq .QueryType "GTE"}} >=  {{else if eq .QueryType "LT"}} < {{else if eq .QueryType "LTE"}} <= {{else if eq .QueryType "LIKE"}} like {{end}}?", {{ if eq .QueryType "LIKE"}}"%"+e.{{.GoField}}+"%"{{else}}e.{{.GoField}}{{end}})
    }
{{ end -}}
{{- end }}

    if err = table.Find(&docs).Error; err != nil {
        return
    }
    return
}

//Get{{.ClassName}} 获取{{.ClassName}}
func (d *dao) Get{{.ClassName}}(ctx context.Context, e *model.{{.ClassName}}) (ret *model.{{.ClassName}}, err error) {
    table := d.orm.WithContext(ctx).Table(e.TableName())
{{ range .Columns -}}
{{- $x := .Pk}}
{{- if ($x) }}
    if e.{{.GoField}} != {{if eq .GoType "string" -}} "" {{ else if eq .GoType "uint64" -}} 0 {{- end}} {
        table = table.Where("{{.ColumnName}}{{if eq .QueryType "EQ"}} = {{else if eq .QueryType "NE"}} != {{else if eq .QueryType "GT"}} >  {{else if eq .QueryType "GTE"}} >=  {{else if eq .QueryType "LT"}} < {{else if eq .QueryType "LTE"}} <= {{else if eq .QueryType "LIKE"}} like {{end}}?", {{ if eq .QueryType "LIKE"}}"%"+e.{{.GoField}}+"%"{{else}}e.{{.GoField}}{{end}})
    }
{{- else if .IsQuery }}
    if e.{{.GoField}} != {{if eq .GoType "string" -}} "" {{ else if eq .GoType "int" -}} 0 {{- end}} {
        table = table.Where("{{.ColumnName}}{{if eq .QueryType "EQ"}} = {{else if eq .QueryType "NE"}} != {{else if eq .QueryType "GT"}} >  {{else if eq .QueryType "GTE"}} >=  {{else if eq .QueryType "LT"}} < {{else if eq .QueryType "LTE"}} <= {{else if eq .QueryType "LIKE"}} like {{end}}?", {{ if eq .QueryType "LIKE"}}"%"+e.{{.GoField}}+"%"{{else}}e.{{.GoField}}{{end}})
    }
{{- end -}}
{{- end }}

    var doc model.{{.ClassName}}
    if err = table.First(&doc).Error; err != nil {
        return
    }
    ret = &doc
    return
}

//Create{{.ClassName}} 创建{{.ClassName}}
func (d *dao) Create{{.ClassName}}(ctx context.Context, e *model.{{.ClassName}}) (doc model.{{.ClassName}}, err error) {
    table := d.orm.WithContext(ctx).Table(e.TableName())
    if err = table.Create(&e).Error; err != nil {
        return
    }
    doc = *e
    return
}

//Update{{.ClassName}} 更新{{.ClassName}}
func (d *dao) Update{{.ClassName}}(ctx context.Context, e *model.{{.ClassName}},id uint64) (update model.{{.ClassName}}, err error) {
    table := d.orm.WithContext(ctx).Table(e.TableName())
    if err = table.Where("{{.PkColumn}} = ?", id).First(&update).Error; err != nil {
        return
    }

    //参数1:是要修改的数据
    //参数2:是修改的数据
    if err = table.Model(&update).Updates(&e).Error; err != nil {
        return
    }
    return
}

//Delete{{.ClassName}} 删除{{.ClassName}}
func (d *dao) Delete{{.ClassName}}(ctx context.Context, e *model.{{.ClassName}}, id uint64) (err error) {
    table := d.orm.WithContext(ctx).Table(e.TableName())
    if err = table.Where("{{.PkColumn}} = ?", id).Delete(&model.{{.ClassName}}{}).Error; err != nil {
        return
    }
    return
}

//BatchDelete{{.ClassName}} 批量删除
func (d *dao) BatchDelete{{.ClassName}}(ctx context.Context, e *model.{{.ClassName}}, id []uint64) (err error) {
    table := d.orm.WithContext(ctx).Table(e.TableName())
    if err = table.Where("{{.PkColumn}} in (?)", id).Delete(&model.{{.ClassName}}{}).Error; err != nil {
        return
    }
    return
}

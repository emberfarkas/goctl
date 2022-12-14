package model

import (
{{$ii := 1 -}}
{{- range .Columns -}}
{{- if eq .GoType "time.Time" -}}
{{- if eq $ii 1 -}} 
    "time" 
    {{- $ii = 2}}
{{- end }}
{{- end }}
{{- end }}
)

type {{.ClassName}} struct {
{{ range .Columns -}} 
{{- $x := .Pk}}
{{- if ($x) }}
    {{.GoField}} {{.GoType}} `json:"{{.JsonField}}" gorm:"type:{{.ColumnType}};primary_key"` // {{.ColumnComment}}
{{- else if eq .GoField "CreatedAt" -}}
{{- else if eq .GoField "UpdatedAt" -}}
{{- else if eq .GoField "DeletedAt" -}}
{{- else }} 
    {{.GoField}} {{.GoType}} `json:"{{.JsonField}}" gorm:"type:{{.ColumnType}};"` // {{.ColumnComment}}
{{- end -}}
{{- end }}
    BaseModel
}

func ({{.ClassName}}) TableName() string {
    return "{{.TBName}}"
}


package codegen

import (
	"time"
)

type BaseModel struct {
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt"`
}

type DBTables struct {
	TableName      string    `gorm:"column:TABLE_NAME" json:"tableName"`
	Engine         string    `gorm:"column:ENGINE" json:"engine"`
	TableRows      string    `gorm:"column:TABLE_ROWS" json:"tableRows"`
	TableCollation string    `gorm:"column:TABLE_COLLATION" json:"tableCollation"`
	CreateTime     time.Time `gorm:"column:CREATE_TIME" json:"createTime"`
	UpdateTime     time.Time `gorm:"column:UPDATE_TIME" json:"updateTime"`
	TableComment   string    `gorm:"column:TABLE_COMMENT" json:"tableComment"`
}

type DBColumns struct {
	TableCatalog           string `gorm:"column:TABLE_CATALOG" json:"tableCatalog"`
	TableSchema            string `gorm:"column:TABLE_SCHEMA" json:"tableSchema"`
	TableName              string `gorm:"column:TABLE_NAME" json:"tableName"`
	ColumnName             string `gorm:"column:COLUMN_NAME" json:"columnName"`
	ColumnDefault          string `gorm:"column:COLUMN_DEFAULT" json:"columnDefault"`
	IsNullable             string `gorm:"column:IS_NULLABLE" json:"isNullable"`
	DataType               string `gorm:"column:DATA_TYPE" json:"dataType"`
	CharacterMaximumLength string `gorm:"column:CHARACTER_MAXIMUM_LENGTH" json:"characterMaximumLength"`
	CharacterSetName       string `gorm:"column:CHARACTER_SET_NAME" json:"characterSetName"`
	ColumnType             string `gorm:"column:COLUMN_TYPE" json:"columnType"`
	ColumnKey              string `gorm:"column:COLUMN_KEY" json:"columnKey"`
	Extra                  string `gorm:"column:EXTRA" json:"extra"`
	ColumnComment          string `gorm:"column:COLUMN_COMMENT" json:"columnComment"`
}

type SysTables struct {
	TableId             int32        `gorm:"primary_key;auto_increment" json:"tableId"`    //表编码
	TBName              string       `gorm:"column:table_name;size:255;" json:"tableName"` //表名称
	TableComment        string       `gorm:"size:255;" json:"tableComment"`                //表备注
	ClassName           string       `gorm:"size:255;" json:"className"`                   //类名
	TplCategory         string       `gorm:"size:255;" json:"tplCategory"`                 //模板分类
	PackageName         string       `gorm:"size:255;" json:"packageName"`                 //包名
	ModuleName          string       `gorm:"size:255;" json:"moduleName"`                  //模块名
	BusinessName        string       `gorm:"size:255;" json:"businessName"`                //业务模块
	FunctionName        string       `gorm:"size:255;" json:"functionName"`                //功能名称
	FunctionAuthor      string       `gorm:"size:255;" json:"functionAuthor"`              //功能作者
	PkColumn            string       `gorm:"size:255;" json:"pkColumn"`
	PkGoField           string       `gorm:"size:255;" json:"pkGoField"`
	PkJsonField         string       `gorm:"size:255;" json:"pkJsonField"`
	Options             string       `gorm:"size:255;" json:"options"`
	TreeCode            string       `gorm:"size:255;" json:"treeCode"`
	TreeParentCode      string       `gorm:"size:255;" json:"treeParentCode"`
	TreeName            string       `gorm:"size:255;" json:"treeName"`
	Tree                bool         `gorm:"size:1;" json:"tree"`
	Crud                bool         `gorm:"size:1;" json:"crud"`
	Remark              string       `gorm:"size:255;" json:"remark"`
	IsLogicalDelete     string       `gorm:"size:1;" json:"isLogicalDelete"`
	LogicalDelete       bool         `gorm:"size:1;" json:"logicalDelete"`
	LogicalDeleteColumn string       `gorm:"size:128;" json:"logicalDeleteColumn"`
	CreateBy            string       `gorm:"size:128;" json:"createBy"`
	UpdateBy            string       `gorm:"size:128;" json:"updateBy"`
	Columns             []SysColumns `gorm:"-" json:"columns"`
	Module              string       `gorm:"-" json:"module"`
	BaseModel
}

func (SysTables) TableName() string {
	return "sys_tables"
}

type Params struct {
	TreeCode       string `gorm:"-" json:"treeCode"`
	TreeParentCode string `gorm:"-" json:"treeParentCode"`
	TreeName       string `gorm:"-" json:"treeName"`
}

type SysColumns struct {
	ColumnId      int32  `gorm:"primary_key;auto_increment" json:"columnId"`
	TableId       int32  `gorm:"" json:"tableId"`
	ColumnName    string `gorm:"size:128;" json:"columnName"`
	ColumnComment string `gorm:"column:column_comment;size:128;" json:"columnComment"`
	ColumnType    string `gorm:"column:column_type;size:128;" json:"columnType"`
	GoType        string `gorm:"column:go_type;size:128;" json:"goType"`
	GoField       string `gorm:"column:go_field;size:128;" json:"goField"`
	JsonField     string `gorm:"column:json_field;size:128;" json:"jsonField"`
	IsPk          string `gorm:"column:is_pk;size:4;" json:"isPk"`
	IsIncrement   string `gorm:"column:is_increment;size:4;" json:"isIncrement"`
	IsRequired    string `gorm:"column:is_required;size:4;" json:"isRequired"`
	IsInsert      string `gorm:"column:is_insert;size:4;" json:"isInsert"`
	IsEdit        string `gorm:"column:is_edit;size:4;" json:"isEdit"`
	IsList        string `gorm:"column:is_list;size:4;" json:"isList"`
	IsQuery       string `gorm:"column:is_query;size:4;" json:"isQuery"`
	QueryType     string `gorm:"column:query_type;size:128;" json:"queryType"`
	HtmlType      string `gorm:"column:html_type;size:128;" json:"htmlType"`
	DictType      string `gorm:"column:dict_type;size:128;" json:"dictType"`
	Sort          int32  `gorm:"column:sort;" json:"sort"`
	List          string `gorm:"column:list;size:1;" json:"list"`
	Pk            bool   `gorm:"column:pk;size:1;" json:"pk"`
	Required      bool   `gorm:"column:required;size:1;" json:"required"`
	SuperColumn   bool   `gorm:"column:super_column;size:1;" json:"superColumn"`
	UsableColumn  bool   `gorm:"column:usable_column;size:1;" json:"usableColumn"`
	Increment     bool   `gorm:"column:increment;size:1;" json:"increment"`
	Insert        bool   `gorm:"column:insert;size:1;" json:"insert"`
	Edit          bool   `gorm:"column:edit;size:1;" json:"edit"`
	Query         bool   `gorm:"column:query;size:1;" json:"query"`
	Remark        string `gorm:"column:remark;size:255;" json:"remark"`
	CreateBy      string `gorm:"column:create_by;size:128;" json:"createBy"`
	UpdateBy      string `gorm:"column:update_By;size:128;" json:"updateBy"`
	BaseModel
}

func (SysColumns) TableName() string {
	return "sys_columns"
}

{{$tableComment:=.TableComment}}
<template>
    <BasicLayout>
        <template #wrapper>
            <el-card class="box-card">
                <el-form ref="queryForm" :model="queryParams" :inline="true" label-width="68px">
                    {{range .Columns}}
                        {{- $x := .IsQuery -}}
                        {{- if ($x) -}}
                            <el-form-item label="{{.ColumnComment}}" prop="{{.JsonField}}">
                                {{if eq .DictType "" -}}
                                    <el-input v-model="queryParams.{{.JsonField}}" placeholder="请输入{{.ColumnComment}}" clearable
                                              size="small" @keyup.enter.native="handleQuery"/>
                                {{- else -}}
                                    <el-select v-model="queryParams.{{.JsonField}}"
                                               placeholder="{{$tableComment}}{{.ColumnComment}}" clearable size="small">
                                        <el-option
                                                v-for="dict in {{.JsonField}}Options"
                                                :key="dict.dictValue"
                                                :label="dict.dictLabel"
                                                :value="dict.dictValue"
                                        />
                                    </el-select>
                                {{- end}}
                            </el-form-item>
                        {{end}}
                    {{- end }}
                    <el-form-item>
                        <el-button type="primary" icon="el-icon-search" size="mini" @click="handleQuery">搜索</el-button>
                        <el-button icon="el-icon-refresh" size="mini" @click="resetQuery">重置</el-button>
                    </el-form-item>
                </el-form>

                <el-row :gutter="10" class="mb8">
                    <el-col :span="1.5">
                        <el-button
                                v-permisaction="['{{.PackageName}}:{{.ModuleName}}:add']"
                                type="primary"
                                icon="el-icon-plus"
                                size="mini"
                                @click="handleAdd"
                        >新增
                        </el-button>
                    </el-col>
                    <el-col :span="1.5">
                        <el-button
                                v-permisaction="['{{.PackageName}}:{{.ModuleName}}:edit']"
                                type="success"
                                icon="el-icon-edit"
                                size="mini"
                                :disabled="single"
                                @click="handleUpdate"
                        >修改
                        </el-button>
                    </el-col>
                    <el-col :span="1.5">
                        <el-button
                                v-permisaction="['{{.PackageName}}:{{.ModuleName}}:remove']"
                                type="danger"
                                icon="el-icon-delete"
                                size="mini"
                                :disabled="multiple"
                                @click="handleDelete"
                        >删除
                        </el-button>
                    </el-col>
                </el-row>

                <el-table v-loading="loading" border :data="{{.ModuleName}}List" @selection-change="handleSelectionChange">
                    <el-table-column type="selection" width="55" align="center"/>
                    {{- range .Columns -}}
                        {{- $x := .IsList -}}
                        {{- if ($x) }}
                            {{- if ne .DictType "" -}}
                                <el-table-column label="{{.ColumnComment}}" align="center" prop="{{.JsonField}}"
                                                 :formatter="{{.JsonField}}Format" width="100">
                                    <template slot-scope="scope">
                                        {{ "{{" }} {{.JsonField}}Format(scope.row) {{"}}"}}
                                    </template>
                                </el-table-column>
                            {{- end -}}
                            {{- if eq .DictType "" -}}
                                <el-table-column label="{{.ColumnComment}}" align="center" prop="{{.JsonField}}"
                                                 :show-overflow-tooltip="true"/>
                            {{- end -}}

                        {{- end }}
                    {{- end }}
                    <el-table-column label="操作" align="center" class-name="small-padding fixed-width">
                        <template slot-scope="scope">
                            <el-button
                                v-permisaction="['pmsbrand:pmsbrand:info']"
                                size="mini"
                                type="text"
                                icon="el-icon-info"
                                @click="handleDetail(scope.row)"
                            >详情
                            </el-button>
                            <el-button
                                    v-permisaction="['{{.PackageName}}:{{.ModuleName}}:edit']"
                                    size="mini"
                                    type="text"
                                    icon="el-icon-edit"
                                    @click="handleUpdate(scope.row)"
                            >修改
                            </el-button>
                            <el-button
                                    v-permisaction="['{{.PackageName}}:{{.ModuleName}}:remove']"
                                    size="mini"
                                    type="text"
                                    icon="el-icon-delete"
                                    @click="handleDelete(scope.row)"
                            >删除
                            </el-button>
                        </template>
                    </el-table-column>
                </el-table>

                <pagination
                        v-show="total>0"
                        :total="total"
                        :page.sync="queryParams.pageIndex"
                        :limit.sync="queryParams.pageSize"
                        @pagination="getList"
                />

                <!-- 添加或修改对话框 -->
                <el-dialog :title="title" :visible.sync="open" width="500px" modal="true" @close="cancel">
                    <el-form ref="form" :model="form" :rules="rules" label-width="80px">
                        {{ range .Columns }}
                            {{- $x := .IsInsert -}}
                            {{- if ($x) -}}
                                {{- if (.Pk) }}
                                {{- else if eq .GoField "CreatedAt" -}}
                                {{- else if eq .GoField "UpdatedAt" -}}
                                {{- else if eq .GoField "DeletedAt" -}}
                                {{- else if eq .GoField "UpdateBy" -}}
                                {{- else if eq .GoField "CreateBy" -}}
                                {{- else }}
                                    <el-form-item label="{{.ColumnComment}}" prop="{{.JsonField}}">
                                        {{ if eq "input" .HtmlType -}}
                                            <el-input v-model="form.{{.JsonField}}" placeholder="{{.ColumnComment}}"
                                                      {{if eq .IsEdit "false" -}}:disabled="isEdit" {{- end}}/>
                                        {{- else if eq "select" .HtmlType -}}
                                            <el-select v-model="form.{{.JsonField}}"
                                                       placeholder="请选择" {{if eq .IsEdit "false" -}} :disabled="isEdit" {{- end }}>
                                                <el-option
                                                        v-for="dict in {{.JsonField}}Options"
                                                        :key="dict.dictValue"
                                                        :label="dict.dictLabel"
                                                        :value="dict.dictValue"
                                                />
                                            </el-select>
                                        {{- else if eq "radio" .HtmlType -}}
                                            <el-radio-group v-model="form.{{.JsonField}}">
                                                <el-radio
                                                        v-for="dict in {{.JsonField}}Options"
                                                        :key="dict.dictValue"
                                                        :label="dict.dictValue"
                                                >{{"{{"}} dict.dictLabel {{"}}"}}</el-radio>
                                            </el-radio-group>
                                        {{- else if eq "datetime" .HtmlType -}}
                                            <el-date-picker
                                                    v-model="form.{{.JsonField}}"
                                                    type="datetime"
                                                    placeholder="选择日期">
                                            </el-date-picker>
                                        {{- else if eq "textarea" .HtmlType -}}
                                            <el-input
                                                    v-model="form.{{.JsonField}}"
                                                    type="textarea"
                                                    :rows="2"
                                                    placeholder="请输入内容">
                                            </el-input>
                                        {{- end }}
                                    </el-form-item>
                                {{- end }}
                            {{- end }}
                        {{- end }}
                    </el-form>
                    <div slot="footer" class="dialog-footer">
                        <el-button type="primary" @click="submitForm">确 定</el-button>
                        <el-button @click="cancel">取 消</el-button>
                    </div>
                </el-dialog>
            </el-card>
        </template>
    </BasicLayout>
</template>

<script>
    import {add{{.ClassName}}, del{{.ClassName}}, get{{.ClassName}}, page{{.ClassName}}, update{{.ClassName}}} from '@/api/{{.Module}}/{{.PackageName}}'
    export default {
        name: '{{.ClassName}}',
        data() {
            return {
                // 遮罩层
                loading: true,
                // 选中数组
                ids: [],
                // 非单个禁用
                single: true,
                // 非多个禁用
                multiple: true,
                // 总条数
                total: 0,
                // 弹出层标题
                title: '',
                // 是否显示弹出层
                open: false,
                isEdit: false,
                isDetail: false,

                // 类型数据字典
                typeOptions: [],
                {{.ModuleName}}List: [],
                {{range .Columns}}
                {{- if ne .DictType "" -}}
                {{.JsonField}}Options: [],
                {{- end -}}
                {{- end }}
                // 查询参数
                queryParams: {
                    pageIndex: 1,
                    pageSize: 10,
            {{ range .Columns }}
            {{- if (.IsQuery) -}}
            {{.JsonField}}:
            undefined,
            {{ end -}}
            {{- end }}
        },
            // 表单参数
            form: {
            }
        ,
            // 表单校验
            rules: {
                {{- range .Columns -}}
                {{- $x := .IsQuery -}}
                {{- if ($x) -}}
                {{.JsonField}}:
                [
                    {required: true, message: '{{.ColumnComment}}不能为空', trigger: 'blur'}
                ],
                {{ end }}
                {{- end -}}
            }
        }
        },
        created() {
            this.getList()
            {{range .Columns}}
            {{- if ne .DictType "" -}}

            this.getDicts('{{.DictType}}').then(response => {
                this.{{.JsonField}}Options = response.data
            })

            {{ end -}}
            {{- end -}}
        },
        methods: {
            /** 查询参数列表 */
            getList() {
                this.loading = true
                page{{.ClassName}}(this.addDateRange(this.queryParams, this.dateRange)).then(response => {
                        this.{{.ModuleName}}List = response.data
                        this.total = response.count
                        this.loading = false
                    }
                )
            },
            // 取消按钮
            cancel() {
                this.open = false
                this.isDetail = false
                this.reset()
            },
            // 表单重置
            reset() {
                this.form = {
                {{ range .Columns}}
                {{- $x := .IsInsert -}}
                {{- if ($x) -}}
                {{- if eq .GoField "CreatedAt" -}}
                {{- else if eq .GoField "UpdatedAt" -}}
                {{- else if eq .GoField "DeletedAt" -}}
                {{- else if eq .GoField "UpdateBy" -}}
                {{- else if eq .GoField "CreateBy" -}}
                {{- else }}
                {{.JsonField}}: undefined,
                {{- end }}
                {{- end -}}
                {{- end }}
            }
                this.resetForm('form')
            },
            {{range .Columns}}
            {{- if ne .DictType "" -}}
            {{.JsonField}}Format(row) {
                return this.selectDictLabel(this.{{.JsonField}}Options, row.{{.JsonField}})
            },
            {{ end -}}
            {{- end }}

            /** 搜索按钮操作 */
            handleQuery() {
                this.queryParams.pageIndex = 1
                this.getList()
            },
            /** 重置按钮操作 */
            resetQuery() {
                this.dateRange = []
                this.resetForm('queryForm')
                this.handleQuery()
            },
            /** 详情 */
            handleDetail(row) {
                this.reset()
                const id = row.id || this.ids
                get{{.ClassName}}(id).then((response) => {
                    this.form = response.data[0]
                    this.open = true
                    this.title = '{{.TableComment}}详情'
                    this.isEdit = false
                    this.isDetail = true
                })
            },
            /** 新增按钮操作 */
            handleAdd() {
                this.reset()
                this.open = true
                this.title = '添加{{.TableComment}}'
                this.isEdit = false
            },
            // 多选框选中数据
            handleSelectionChange(selection) {
                this.ids = selection.map(item => item.{{.PkJsonField}})
                this.single = selection.length !== 1
                this.multiple = !selection.length
            },
            /** 修改按钮操作 */
            handleUpdate(row) {
                this.reset()
                const {{.PkJsonField}} =
                row.{{.PkJsonField}} || this.ids
                get{{.ClassName}}({{.PkJsonField}}).then(response => {
                    this.form = response.data[0]
                    this.open = true
                    this.title = '修改{{.TableComment}}'
                    this.isEdit = true
                })
            },
            /** 提交按钮 */
            submitForm: function () {
                this.$refs['form'].validate(valid => {
                    if (valid) {
                        if (this.form.{{.PkJsonField}} !== undefined) {
                            update{{.ClassName}}(this.form).then(response => {
                                if (response.code === 0) {
                                    this.msgSuccess('修改成功')
                                    this.open = false
                                    this.getList()
                                } else {
                                    this.msgError(response.msg)
                                }
                            })
                        } else {
                            add{{.ClassName}}(this.form).then(response => {
                                if (response.code === 0) {
                                    this.msgSuccess('新增成功')
                                    this.open = false
                                    this.getList()
                                } else {
                                    this.msgError(response.msg)
                                }
                            })
                        }
                    }
                })
            },
            /** 删除按钮操作 */
            handleDelete(row) {
                const Ids = row.{{.PkJsonField}} || this.ids
                this.$confirm('是否确认删除编号为"' + Ids + '"的数据项?', '警告', {
                    confirmButtonText: '确定',
                    cancelButtonText: '取消',
                    type: 'warning'
                }).then(function () {
                    return del{{.ClassName}}(Ids)
                }).then(() => {
                    this.getList()
                    this.msgSuccess('删除成功')
                }).catch(function () {
                })
            }
        }
    }
</script>

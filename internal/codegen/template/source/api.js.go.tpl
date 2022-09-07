import request from '@/utils/request'
import qs from 'qs'

// 查询{{.ClassName}}分页
export async function list{{.ClassName}}(query) {
  return await request({
    url: '/admin/v1/{{.Module}}/{{.ModuleName}}List',
    method: 'get',
    params: query
  })
}

// 查询{{.ClassName}}列表
export async function batchGet{{.ClassName}}(query) {
  return await request({
    url: '/admin/v1/{{.Module}}/{{.ModuleName}}:batchGet',
    method: 'get',
    params: query
  })
}

// 查询{{.ClassName}}详细
export async function get{{.ClassName}} ({{.PkJsonField}}) {
  return await request({
    url: '/admin/v1/{{.Module}}/{{.ModuleName}}/' + {{.PkJsonField}},
    method: 'get'
  })
}


// 新增{{.ClassName}}
export async function add{{.ClassName}}(data) {
  return await request({
    url: '/admin/v1/{{.Module}}/{{.ModuleName}}',
    method: 'post',
    data: data
  })
}

// 修改{{.ClassName}}
export async function update{{.ClassName}}(data) {
  return await request({
    url: '/admin/v1/{{.Module}}/{{.ModuleName}}',
    method: 'put',
    data: data
  })
}

// 删除{{.ClassName}}
export async function del{{.ClassName}}({{.PkJsonField}}) {
  return await request({
    url: '/admin/v1/{{.Module}}/{{.ModuleName}}/' + {{.PkJsonField}},
    method: 'delete'
  })
}

import request from '@/utils/request'
import qs from 'qs'

// 查询{{.ClassName}}分页
export function page{{.ClassName}}(query) {
return request({
    url: '/admin/v1/{{.Module}}/{{.ModuleName}}Page',
    method: 'get',
    params: query
  })
}

// 查询{{.ClassName}}列表
export function batchget{{.ClassName}}(query) {
return request({
    url: '/admin/v1/{{.Module}}/{{.ModuleName}}List',
    method: 'get',
    params: query
  })
}

// 查询{{.ClassName}}详细
export function get{{.ClassName}} ({{.PkJsonField}}) {
  return request({
    url: '/admin/v1/{{.Module}}/{{.ModuleName}}/' + {{.PkJsonField}},
    method: 'get'
  })
}


// 新增{{.ClassName}}
export function add{{.ClassName}}(data) {
  return request({
    url: '/admin/v1/{{.Module}}/{{.ModuleName}}',
    method: 'post',
    data: data
  })
}

// 修改{{.ClassName}}
export function update{{.ClassName}}(data) {
  return request({
    url: '/admin/v1/{{.Module}}/{{.ModuleName}}',
    method: 'put',
    data: data
  })
}

// 删除{{.ClassName}}
export function del{{.ClassName}}({{.PkJsonField}}) {
  return request({
    url: '/admin/v1/{{.Module}}/{{.ModuleName}}/' + {{.PkJsonField}},
    method: 'delete'
  })
}

